package process

import (
	"net/http"

	"github.com/andreylm/process-manager/models/process"
)

// Create - create new process
func Create(collection *process.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
