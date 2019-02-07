package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/go-xorm/xorm"

	UsersModel "github.com/learning/project/api/models/users"
)

// Update - update user
func Update(db *xorm.Engine) http.HandlerFunc {
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

		var userToUpdate UsersModel.User
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&userToUpdate)

		if err = UsersModel.Update(db, int64(id), &userToUpdate); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		packet, err := json.Marshal(userToUpdate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(packet)
	}
}
