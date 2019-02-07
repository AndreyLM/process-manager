package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/go-xorm/xorm"

	UsersModel "github.com/learning/project/api/models/users"
)

// Show - get user
func Show(db *xorm.Engine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		rawID := vars["id"]
		id, _ := strconv.Atoi(rawID)
		var findBy UsersModel.User
		findBy.ID = int64(id)

		users, err := UsersModel.Index(db, &findBy)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if len(users) == 0 {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		packet, err := json.Marshal(users[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(packet)
	}
}
