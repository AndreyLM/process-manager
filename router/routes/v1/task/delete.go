package task

import (
	"net/http"

	"github.com/andreylm/process-manager/utils"

	"github.com/gorilla/mux"

	"github.com/andreylm/process-manager/models/task"
)

// Delete - deletes task
func Delete(collection *task.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		taskName := vars["task"]

		if !collection.TaskExist(taskName) {
			utils.Respond(
				w,
				utils.PrepareData(
					http.StatusNotFound,
					"Task does not exist",
					nil,
				),
			)
			return
		}

		collection.RemoveTask(taskName)
		utils.Respond(
			w,
			utils.PrepareData(
				http.StatusOK,
				"Task was successfully deleted",
				nil,
			),
		)
		return
	}
}
