package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"

	"github.com/reinbach/zenpager/database"
	"github.com/reinbach/zenpager/utils"
)

type Message struct {
	Type    string
	Content string
}

type Response struct {
	Result   string
	Messages []Message
	ID       int64
}

func Routes() *web.Mux {
	api := web.New()
	api.Use(middleware.SubRouter)
	api.Use(utils.ApplicationJSON)

	api.Post("/login", Login)
	api.Get("/logout", Logout)

	// user
	// api.Get("/user/", UserList)
	// api.Get("/user/:id", UserItem)
	// api.Post("/user/", UserAdd)
	// api.Put("/user/:id", UserUpdate)
	api.Patch("/user/:id", UserPartialUpdate)
	// api.Delete("/user/:id", UserDelete)
	// api.Get("/user", http.RedirectHandler("/user/", 301))

	return api
}

func Login(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)
	var res = Response{}
	var user User
	var m Message
	if err := utils.DecodePayload(r, &user); err != nil {
		m = Message{
			Type:    "danger",
			Content: "Data appears to be invalid.",
		}
		res = Response{
			Result:   "error",
			Messages: []Message{m},
		}
		utils.EncodePayload(w, http.StatusBadRequest, res)
		return
	}
	errors := user.Validate(true)
	if len(errors) > 0 {
		res = Response{Result: "error", Messages: errors}
		utils.EncodePayload(w, http.StatusBadRequest, res)
		return
	}
	log.Println("user before login: ", user)
	if user.Login(db) {
		log.Println("user after login: ", user)
		// TODO create and store token
		w.Header().Set("X-Access-Token", "<token>")
		m = Message{Type: "success", Content: "Successfully logged in"}
		res = Response{
			Result:   "success",
			Messages: []Message{m},
			ID:       user.ID,
		}
		utils.EncodePayload(w, http.StatusOK, res)
		return
	}
	m = Message{
		Type:    "danger",
		Content: "Username and/or password is invalid.",
	}
	res = Response{
		Result:   "error",
		Messages: []Message{m},
	}
	utils.EncodePayload(w, http.StatusBadRequest, res)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	//TODO remove token
	m := Message{Type: "success", Content: "Signed Out"}
	res := Response{Result: "success", Messages: []Message{m}}
	b, err := json.Marshal(res)
	if err != nil {
		log.Println("Failed to encode response: ", err)
	}
	w.Write(b)
}

func UserPartialUpdate(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)
	var res = Response{}
	var m Message

	// get id of user to be updated
	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		m = Message{Type: "danger", Content: "User not found."}
		res = Response{Result: "error", Messages: []Message{m}}
		utils.EncodePayload(w, http.StatusNotFound, res)
		return
	}

	// set user with current data
	user := User{
		ID: id,
	}
	user.Get(db)

	// update data with new data and ensure it is valid
	if err := utils.DecodePayload(r, &user); err != nil {
		m = Message{Type: "danger", Content: "Data appears to be invalid."}
		res = Response{Result: "error", Messages: []Message{m}}
		utils.EncodePayload(w, http.StatusBadRequest, res)
		return
	}
	errors := user.Validate(false)
	if len(errors) > 0 {
		res = Response{Result: "error", Messages: errors}
		utils.EncodePayload(w, http.StatusBadRequest, res)
		return
	}

	if user.Update(db) {
		m = Message{Type: "success", Content: "User data updated."}
		res = Response{Result: "success", Messages: []Message{m}}
		utils.EncodePayload(w, http.StatusAccepted, res)
	} else {
		m = Message{Type: "danger", Content: "Failed to update user."}
		res = Response{Result: "error", Messages: []Message{m}}
		utils.EncodePayload(w, http.StatusInternalServerError, res)
	}
}
