package process

import (
	"net/http"

	"github.com/andreylm/process-manager/utils"

	"github.com/gorilla/mux"

	"github.com/andreylm/process-manager/models/task"
)

// Delete - create new process
func Delete(collection *task.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		taskName := mux.Vars(r)["name"]
		processUUID := mux.Vars(r)["uuid"]

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

		t := collection.GetTask(taskName)
		message := ""
		if err := t.DeleteProcess(processUUID); err != nil {
			message = err.Error()
		} else {
			message = "Process was deleted"
		}

		utils.Respond(
			w,
			utils.PrepareData(
				http.StatusOK,
				message,
				nil,
			),
		)
		return
	}
}
