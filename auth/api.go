package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	webctx "github.com/goji/context"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"

	"github.com/reinbach/zenpager/database"
	"github.com/reinbach/zenpager/utils"
)

func Routes() *web.Mux {
	api := web.New()
	api.Use(middleware.SubRouter)
	api.Use(utils.ApplicationJSON)

	api.Post("/login", Login)
	api.Get("/logout", Logout)

	return api
}

func UserRoutes() *web.Mux {
	api := web.New()
	api.Use(middleware.SubRouter)
	api.Use(utils.ApplicationJSON)

	// user
	api.Use(Middleware)
	// api.Get("/user/", UserList)
	// api.Get("/user/:id", UserItem)
	// api.Post("/user/", UserAdd)
	// api.Put("/user/:id", UserUpdate)
	api.Patch("/:id", UserPartialUpdate)
	// api.Delete("/user/:id", UserDelete)
	// api.Get("/user", http.RedirectHandler("/user/", 301))

	return api
}

func GetUser(c web.C) User {
	ctx := webctx.FromC(c)
	return ctx.Value("user").(User)
}

func Login(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)
	var res = utils.Response{}
	var user User
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
		RemoveToken(t, db)
	}

	m := utils.Message{Type: "success", Content: "Signed Out"}
	res := utils.Response{Result: "success", Messages: []utils.Message{m}}
	b, err := json.Marshal(res)
	if err != nil {
		log.Println("Failed to encode response: ", err)
	}
	w.Write(b)
}

func UserPartialUpdate(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)
	var res = utils.Response{}
	var m utils.Message

	// get id of user to be updated
	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		m = utils.Message{Type: "danger", Content: "User not found."}
		res = utils.Response{Result: "error", Messages: []utils.Message{m}}
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
		utils.BadRequestResponse(w, "Data appears to be invalid.")
		return
	}
	errors := user.Validate(false)
	if len(errors) > 0 {
		res = utils.Response{Result: "error", Messages: errors}
		utils.EncodePayload(w, http.StatusBadRequest, res)
		return
	}

	if user.Update(db) {
		m = utils.Message{Type: "success", Content: "User data updated."}
		res = utils.Response{Result: "success", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusAccepted, res)
	} else {
		m = utils.Message{Type: "danger", Content: "Failed to update user."}
		res = utils.Response{Result: "error", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusInternalServerError, res)
	}
}
