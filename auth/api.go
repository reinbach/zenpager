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

type Response struct {
	Result   string
	Messages []string
}

func Routes() *web.Mux {
	api := web.New()
	api.Use(middleware.SubRouter)
	api.Use(utils.ApplicationJSON)

	api.Post("/login", Login)
	api.Get("/logout", Logout)

	return api
}

func Login(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)
	var res = Response{}
	var user User
	if err := utils.DecodePayload(r, &user); err != nil {
		res = Response{Result: "error", Messages: []string{"Data appears to be invalid."}}
		utils.EncodePayload(w, http.StatusBadRequest, res)
		return
	}
	errors := user.Validate()
	if len(errors) > 0 {
		res = Response{Result: "error", Messages: errors}
		utils.EncodePayload(w, http.StatusBadRequest, res)
		return
	}
	if user.Login(db) {
		//TODO create token and add to response
		w.Header().Set("X-Access-Token", "<token>")
		res = Response{
			Result:   "success",
			Messages: []string{"Successfully logged in"},
		}
		utils.EncodePayload(w, http.StatusOK, res)
		return
	}
	res = Response{Result: "error", Messages: []string{"Username and/or password is invalid."}}
	utils.EncodePayload(w, http.StatusBadRequest, res)
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
