package task

import (
	"net/http"

	"github.com/andreylm/process-manager/utils"

	"github.com/andreylm/process-manager/models/task"
)

// DeleteAll - deletes all tasks
func DeleteAll(collection *task.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collection.Tasks = nil
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
