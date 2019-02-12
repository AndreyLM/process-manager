package process

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/andreylm/process-manager/models/task"
	"github.com/andreylm/process-manager/utils"
)

// Exist - checks if task exist
func Exist(collection *task.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		taskName := vars["task"]
		processUUID := mux.Vars(r)["uuid"]

		if !collection.TaskExist(taskName) {
			utils.Respond(
				w,
				utils.PrepareData(
					http.StatusNotFound,
					"Task does not exist",
					false,
				),
			)
		}
		t := collection.GetTask(taskName)
		t.RemoveExpiredProcesses()

		for _, val := range t.Processes {
			if val.UUID.String() == processUUID {
				utils.Respond(
					w,
					utils.PrepareData(
						http.StatusOK,
						"Process exist",
						true,
					),
				)
				return
			}
		}

		utils.Respond(
			w,
			utils.PrepareData(
				http.StatusOK,
				"Process does not exist",
				false,
			),
		)
		return
	}
}
