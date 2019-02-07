package users

import (
	"encoding/json"
	"net/http"

	"github.com/go-xorm/xorm"

	UsersModel "github.com/learning/project/api/models/users"
)

// Index - returns users
func Index(db *xorm.Engine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var findBy UsersModel.User
		users, err := UsersModel.Index(db, &findBy)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		packet, err := json.Marshal(users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(packet)
	}
}
