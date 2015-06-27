package api

import (
	"net/http"
	"strconv"

	"github.com/zenazn/goji/web"

	"github.com/reinbach/zenpager/database"
	"github.com/reinbach/zenpager/models"
	"github.com/reinbach/zenpager/utils"
)

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
	user := models.User{
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
