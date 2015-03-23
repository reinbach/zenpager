package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"

	"github.com/reinbach/zenpager/database"
	"github.com/reinbach/zenpager/form"
	"github.com/reinbach/zenpager/utils"
)

type Fields []form.Field

type Response struct {
	Result   string
	Messages []string
}

var (
	email = form.Field{
		Name:       "email",
		Required:   true,
		Validators: form.Validators{form.Email{}},
		Value:      "",
	}
	password = form.Field{
		Name:     "password",
		Required: true,
		Value:    "",
	}
)

func Routes() *web.Mux {
	api := web.New()
	api.Use(middleware.SubRouter)
	api.Use(utils.ApplicationJSON)

	api.Get("/login", Login)
	api.Get("/logout", Logout)

	return api
}

func Login(c web.C, w http.ResponseWriter, r *http.Request) {
	db := database.FromContext(c)
	fields := Fields{email, password}
	f := form.Validate(r, fields)
	res := Response{}
	if f.IsValid() == true {
		user := User{
			Email:    f.GetValue("email"),
			Password: f.GetValue("password"),
		}
		if user.Login(db) {
			//TODO create token and add to response
			w.Header().Set("X-Access-Token", "<token>")
			res = Response{
				Result:   "success",
				Messages: []string{"Successfully logged in"},
			}
		} else {
			res = Response{Result: "error", Messages: []string{"Invalid User"}}
		}
	} else {
		res = Response{Result: "error", Messages: f.Errors}
	}
	b, err := json.Marshal(res)
	if err != nil {
		log.Println("Failed to encode response: ", err)
	}
	w.Write(b)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	//TODO remove token
	res := Response{Result: "success", Messages: []string{"Signed Out"}}
	b, err := json.Marshal(res)
	if err != nil {
		log.Println("Failed to encode response: ", err)
	}
	w.Write(b)
}
