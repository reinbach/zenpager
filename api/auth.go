package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zenazn/goji/web"

	"github.com/reinbach/zenpager/database"
	"github.com/reinbach/zenpager/models"
	"github.com/reinbach/zenpager/utils"
)

func Login(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)
	var res = utils.Response{}
	var user models.User
	var m utils.Message
	if err := utils.DecodePayload(r, &user); err != nil {
		utils.BadRequestResponse(w, "Data appears to be invalid.")
		return
	}
	errors := user.Validate(true)
	if len(errors) > 0 {
		res = utils.Response{Result: "error", Messages: errors}
		utils.EncodePayload(w, http.StatusBadRequest, res)
		return
	}
	if user.Login(db) {
		t, err := user.AddToken(db)
		if err != nil {
			utils.BadRequestResponse(w, "Failed to login.")
			return
		}
		w.Header().Set("X-Access-Token", t.Token)
		m = utils.Message{Type: "success", Content: "Successfully logged in"}
		res = utils.Response{
			Result:   "success",
			Messages: []utils.Message{m},
			Data:     user.ID,
		}
		utils.EncodePayload(w, http.StatusOK, res)
		return
	}
	utils.BadRequestResponse(w, "Username and/or password is invalid.")
}

func Logout(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)
	var t = r.Header.Get("X-Access-Token")

	if t != "undefined" {
		models.RemoveToken(t, db)
	}

	m := utils.Message{Type: "success", Content: "Signed Out"}
	res := utils.Response{Result: "success", Messages: []utils.Message{m}}
	b, err := json.Marshal(res)
	if err != nil {
		log.Println("Failed to encode response: ", err)
	}
	w.Write(b)
}
