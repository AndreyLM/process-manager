package task

import (
	"net/http"

	"github.com/andreylm/process-manager/models/task"
	"github.com/andreylm/process-manager/utils"
)

// List - gets list of tasks
func List(collection *task.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collection.RemoveExpired()
		utils.Respond(
			w,
			utils.PrepareData(
				http.StatusOK,
				"Tasks list",
				collection.Tasks,
			),
		)

		return
	}
}
