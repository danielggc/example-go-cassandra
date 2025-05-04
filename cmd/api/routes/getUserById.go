package routes

import (
	"broker/cmd/api/dto"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	. "github.com/samber/mo"
)

func (app *Config) GetUserById(w http.ResponseWriter, r *http.Request) {

	maybeUserId := Try(func() (Option[string], error) {
		uuid := chi.URLParam(r, "id")
		if uuid == "" {
			return None[string](), fmt.Errorf("error empty the  param userId ")

		}
		return Some(uuid), nil

	})
	if maybeUserId.IsError() {
		http.Error(w, maybeUserId.Error().Error(), 400)
		return
	}

	app.UserRepository.GetUserById(maybeUserId.MustGet().MustGet()).Match(
		func(user Option[dto.User]) (Option[dto.User], error) {
			if user.IsAbsent() {
				http.Error(w, "error null user", http.StatusInternalServerError)
				return None[dto.User](), fmt.Errorf("user null ")
			}
			err := json.NewEncoder(w).Encode(user.MustGet())
			if err != nil {
				http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			}
			return user, nil
		},
		func(err error) (Option[dto.User], error) {
			http.Error(w, "error getting the user  :"+err.Error(), http.StatusInternalServerError)
			return None[dto.User](), err
		},
	)
}
