package process

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andreylm/process-manager/router/routes/v1/process/forms"
	"github.com/andreylm/process-manager/utils"
	uuid "github.com/satori/go.uuid"

	"github.com/gorilla/mux"

	"github.com/andreylm/process-manager/models/task"
)

// Create - create new process
func Create(collection *task.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		taskName := mux.Vars(r)["name"]
		if len(taskName) == 0 || !collection.TaskExist(taskName) {
			utils.Respond(
				w,
				utils.PrepareData(
					http.StatusBadRequest,
					"Please provide correct task name",
					nil,
				),
			)
			return
		}

		form := &forms.CreateForm{}
		err := json.NewDecoder(r.Body).Decode(form)

		if err != nil {
			log.Println(err)
			utils.Respond(
				w,
				utils.PrepareData(
					http.StatusBadRequest,
					"Error parsing request",
					err,
				),
			)
			return
		}

		currentTask := collection.GetTask(taskName)
		currentTask.RemoveExpiredProcesses()
		newProcess, err := task.CreateProcess(form.Duration, form.Data)
		if err != nil {
			log.Println(err)
			utils.Respond(
				w,
				utils.PrepareData(
					http.StatusBadRequest,
					err.Error(),
					nil,
				),
			)
			return
		}

		err = currentTask.AddProcess(newProcess)
		if err != nil {
			log.Println(err)
			utils.Respond(
				w,
				utils.PrepareData(
					http.StatusBadRequest,
					err.Error(),
					nil,
				),
			)
			return
		}

		utils.Respond(
			w,
			utils.PrepareData(
				http.StatusOK,
				"Process",
				struct {
					UUID uuid.UUID `json:"uuid"`
				}{newProcess.UUID},
			),
		)
		return

	}
}
