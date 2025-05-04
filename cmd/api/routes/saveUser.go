package routes

import (
	"broker/cmd/api/dto"
	"encoding/json"
	"net/http"

	. "github.com/samber/mo"
)

func (app *Config) SaveUser(w http.ResponseWriter, r *http.Request) {
	maybeUser := Try(func() (Option[dto.User], error) {
		var user dto.User
		err := json.NewDecoder(r.Body).Decode(&user)
		return Some(user), err
	}).Map(func(value Option[dto.User]) (Option[dto.User], error) {
		err := value.MustGet().Validate()
		if err != nil {
			return None[dto.User](), err
		}
		return value, nil
	})
	if maybeUser.IsError() {
		http.Error(w, maybeUser.Error().Error(), 404)
		return
	}

	app.UserRepository.SaveUser(maybeUser.MustGet().MustGet()).Match(
		func(value Option[dto.User]) (Option[dto.User], error) {
			w.Write([]byte("sucess"))
			return value, nil
		},
		func(err error) (Option[dto.User], error) {
			http.Error(w, "error save the user ", http.StatusInternalServerError)
			return None[dto.User](), err
		},
	)

}
