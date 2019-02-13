package process

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/andreylm/process-manager/models/task"
	"github.com/andreylm/process-manager/utils"
)

// Info - gets task's info
func Info(collection *task.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		taskName := vars["name"]
		processUUID := vars["uuid"]

		task := collection.GetTask(taskName)

		if task == nil {
			utils.Respond(
				w,
				utils.PrepareData(
					http.StatusNotFound,
					"Task does not exist. Create one.",
					nil,
				),
			)
			return
		}

		task.RemoveExpiredProcesses()
		for _, val := range task.Processes {
			if val.UUID.String() == processUUID {
				utils.Respond(
					w,
					utils.PrepareData(
						http.StatusOK,
						"Process info",
						val,
					),
				)

				return
			}
		}

		utils.Respond(
			w,
			utils.PrepareData(
				http.StatusOK,
				"Task info",
				task,
			),
		)

		return
	}
}
