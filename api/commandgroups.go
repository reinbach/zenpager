package api

import (
	"net/http"
	"strconv"

	"github.com/zenazn/goji/web"

	"github.com/reinbach/zenpager/database"
	"github.com/reinbach/zenpager/models"
	"github.com/reinbach/zenpager/utils"
)

func CommandGroupList(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	d := models.CommandGroupGetAll(db)

	res := utils.Response{Result: "success", Data: d}
	utils.EncodePayload(w, http.StatusOK, res)
}

func CommandGroupItem(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Command Group not found.")
		return
	}

	group := models.CommandGroup{ID: id}
	group.Get(db)

	if group.Name == "" {
		utils.NotFoundResponse(w, "Command Group not found")
		return
	}

	res := utils.Response{Result: "success", Data: group}
	utils.EncodePayload(w, http.StatusOK, res)
}

func CommandGroupAdd(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	g := models.CommandGroup{}
	if err := utils.DecodePayload(r, &g); err != nil {
		utils.BadRequestResponse(w, "Data appears to be invalid.")
		return
	}

	errors := g.Validate()
	if len(errors) > 0 {
		res := utils.Response{Result: "error", Messages: errors}
		utils.EncodePayload(w, http.StatusBadRequest, res)
		return
	}

	gl := models.CommandGroupGetAll(db)
	f := false
	for _, v := range gl {
		if v.Name == g.Name {
			f = true
		}
	}
	if f {
		utils.BadRequestResponse(
			w,
			"Command Group with this name already exists.",
		)
		return
	}
	if g.Create(db) {
		m := utils.Message{
			Type:    "success",
			Content: "Command Group data added.",
		}
		res := utils.Response{Result: "success", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusAccepted, res)
	} else {
		m := utils.Message{
			Type:    "danger",
			Content: "Failed to update Command Group.",
		}
		res := utils.Response{Result: "error", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusInternalServerError, res)
	}
}

func CommandGroupUpdate(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	// get id of command group to be updated
	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Command Group not found.")
		return
	}

	// set command with current data
	g := models.CommandGroup{ID: id}
	g.Get(db)

	// update data with new data and ensure it is valid
	if err := utils.DecodePayload(r, &g); err != nil {
		utils.BadRequestResponse(w, "Data appears to be invalid.")
		return
	}
	errors := g.Validate()
	if len(errors) > 0 {
		res := utils.Response{Result: "error", Messages: errors}
		utils.EncodePayload(w, http.StatusBadRequest, res)
		return
	}

	if g.Update(db) {
		m := utils.Message{
			Type:    "success",
			Content: "Command Group data updated.",
		}
		res := utils.Response{Result: "success", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusAccepted, res)
	} else {
		m := utils.Message{
			Type:    "danger",
			Content: "Failed to update Command Group.",
		}
		res := utils.Response{Result: "error", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusInternalServerError, res)
	}
}

func CommandGroupDelete(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Command Group not found.")
		return
	}

	g := models.CommandGroup{ID: id}
	g.Delete(db)

	res := utils.Response{Result: "success"}
	utils.EncodePayload(w, http.StatusOK, res)
}

func CommandGroupCommands(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Command Group not found.")
		return
	}

	g := models.CommandGroup{ID: id}
	g.Get(db)

	if g.Name == "" {
		utils.NotFoundResponse(w, "Command Group not found")
		return
	}

	g.GetCommands(db)

	res := utils.Response{Result: "success", Data: g}
	utils.EncodePayload(w, http.StatusOK, res)
}

func CommandGroupAddCommand(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Command Group not found.")
		return
	}

	g := models.CommandGroup{ID: id}
	g.Get(db)

	if g.Name == "" {
		utils.NotFoundResponse(w, "Command Group not found")
		return
	}

	command := models.Command{}
	if err := utils.DecodePayload(r, &command); err != nil {
		utils.BadRequestResponse(w, "Data appears to be invalid.")
		return
	}
	command.Get(db)
	if command.Name == "" {
		utils.NotFoundResponse(w, "Command not found")
		return
	}

	s := g.AddCommand(db, &command)
	if s == true {
		g.Commands = append(g.Commands, command)
		res := utils.Response{Result: "success", Data: g}
		utils.EncodePayload(w, http.StatusOK, res)
	} else {
		utils.BadRequestResponse(w, "Failed to add command to group.")
	}
}

func CommandGroupRemoveCommand(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Command Group not found.")
		return
	}

	cid, err := strconv.ParseInt(c.URLParams["cid"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Command not found.")
		return
	}

	g := models.CommandGroup{ID: id}
	command := models.Command{ID: cid}

	g.Get(db)
	if g.Name == "" {
		utils.NotFoundResponse(w, "Command Group not found")
		return
	}

	s := g.RemoveCommand(db, &command)
	if s == true {
		cs := []models.Command{}
		for _, ct := range g.Commands {
			if ct.ID != command.ID {
				cs = append(cs, ct)
			}
		}
		g.Commands = cs
		res := utils.Response{Result: "success", Data: g}
		utils.EncodePayload(w, http.StatusOK, res)
	} else {
		utils.BadRequestResponse(w, "Failed to remove command from group.")
	}
}
