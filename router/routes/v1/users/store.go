package users

import (
	"encoding/json"
	"net/http"

	"github.com/go-xorm/xorm"

	UsersModel "github.com/learning/project/api/models/users"
)

// Store - store user
func Store(db *xorm.Engine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userToStore UsersModel.User
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&userToStore)

		if err = UsersModel.Store(db, &userToStore); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		packet, err := json.Marshal(userToStore)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(packet)
	}
}
