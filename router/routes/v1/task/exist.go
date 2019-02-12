package task

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

		if collection.TaskExist(taskName) {
			utils.Respond(
				w,
				utils.PrepareData(
					http.StatusOK,
					"Success",
					true,
				),
			)
		} else {
			utils.Respond(
				w,
				utils.PrepareData(
					http.StatusOK,
					"Success",
					false,
				),
			)
		}

		return
	}
}
