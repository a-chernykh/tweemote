package server

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/andreychernih/tweemote/ctx"
	"bitbucket.org/andreychernih/tweemote/services"
	"github.com/asaskevich/govalidator"
)

type UserForm struct {
	Email                string `json:"email" valid:"required,email"`
	Password             string `json:"password" valid:"required,passwordConfirmation~Password confirmation does not match"`
	PasswordConfirmation string `json:"password_confirmation" valid:"required"`
}

type CurrentUserResponse struct {
	Email string `json:"email"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var u UserForm

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, errorJson(err), 400)
		return
	}

	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		http.Error(w, errorJson(err), 422)
		return
	}

	_, err = services.RegisterUser(u.Email, u.Password)
	if err != nil {
		http.Error(w, errorJson(err), 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetCurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	user := ctx.UserFromContext(r.Context())

	var cur CurrentUserResponse
	cur.Email = user.Email

	renderJson(cur, w)
}
