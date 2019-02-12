package process

import (
	"net/http"

	"github.com/andreylm/process-manager/utils"

	"github.com/gorilla/mux"

	"github.com/andreylm/process-manager/models/task"
)

// DeleteAll - create new process
func DeleteAll(collection *task.Collection) http.HandlerFunc {
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

		t := collection.GetTask(taskName)
		t.Processes = nil

		utils.Respond(
			w,
			utils.PrepareData(
				http.StatusOK,
				"All processes were deleted",
				nil,
			),
		)
		return
	}
}
