package task

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andreylm/process-manager/utils"

	"github.com/andreylm/process-manager/models/task"
)

// Create - create new process
func Create(collection *task.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		task := &task.Task{}
		err := json.NewDecoder(r.Body).Decode(task)

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

		if len(task.Name) < 1 {
			utils.Respond(
				w,
				utils.PrepareData(
					http.StatusBadRequest,
					"Please provide task name",
					nil,
				),
			)
			return
		}

		if task.MaxCount == 0 {
			task.MaxCount = 1
		}

		task.Count = 0

		if err = collection.AddTask(task); err != nil {
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
				"Task successfully added to collection",
				nil,
			),
		)
		return
	}
}
