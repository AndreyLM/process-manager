package task

import (
	"encoding/json"
	"net/http"

	"github.com/andreylm/process-manager/models/task"
	"github.com/andreylm/process-manager/utils"
)

// List - gets list of tasks
func List(collection *task.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collection.RemoveExpired()

		res, _ := json.Marshal(collection)
		newCollection := task.Collection{}
		json.Unmarshal(res, &newCollection)
		for int := range newCollection.Tasks {
			newCollection.Tasks[int].Processes = nil
		}

		utils.Respond(
			w,
			utils.PrepareData(
				http.StatusOK,
				"Tasks list",
				newCollection.Tasks,
			),
		)

		return
	}
}
