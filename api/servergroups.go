package api

import (
	"net/http"
	"strconv"

	"github.com/zenazn/goji/web"

	"github.com/reinbach/zenpager/database"
	"github.com/reinbach/zenpager/models"
	"github.com/reinbach/zenpager/utils"
)

func ServerGroupList(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	d := models.ServerGroupGetAll(db)

	res := utils.Response{Result: "success", Data: d}
	utils.EncodePayload(w, http.StatusOK, res)
}

func ServerGroupItem(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Server Group not found.")
		return
	}

	group := models.ServerGroup{ID: id}
	group.Get(db)

	if group.Name == "" {
		utils.NotFoundResponse(w, "Server Group not found")
		return
	}

	res := utils.Response{Result: "success", Data: group}
	utils.EncodePayload(w, http.StatusOK, res)
}

func ServerGroupAdd(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	g := models.ServerGroup{}
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

	gl := models.ServerGroupGetAll(db)
	f := false
	for _, v := range gl {
		if v.Name == g.Name {
			f = true
		}
	}
	if f {
		utils.BadRequestResponse(
			w,
			"Server Group with this name already exists.",
		)
		return
	}
	if g.Create(db) {
		m := utils.Message{
			Type:    "success",
			Content: "Server Group data added.",
		}
		res := utils.Response{Result: "success", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusAccepted, res)
	} else {
		m := utils.Message{
			Type:    "danger",
			Content: "Failed to update Server Group.",
		}
		res := utils.Response{Result: "error", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusInternalServerError, res)
	}
}

func ServerGroupUpdate(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	// get id of server group to be updated
	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Server Group not found.")
		return
	}

	// set server with current data
	g := models.ServerGroup{ID: id}
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
			Content: "Server Group data updated.",
		}
		res := utils.Response{Result: "success", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusAccepted, res)
	} else {
		m := utils.Message{
			Type:    "danger",
			Content: "Failed to update Server Group.",
		}
		res := utils.Response{Result: "error", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusInternalServerError, res)
	}
}

func ServerGroupDelete(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Server Group not found.")
		return
	}

	g := models.ServerGroup{ID: id}
	g.Delete(db)

	res := utils.Response{Result: "success"}
	utils.EncodePayload(w, http.StatusOK, res)
}

func ServerGroupServers(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Server Group not found.")
		return
	}

	g := models.ServerGroup{ID: id}
	g.Get(db)

	if g.Name == "" {
		utils.NotFoundResponse(w, "Server Group not found")
		return
	}

	g.GetServers(db)

	res := utils.Response{Result: "success", Data: g}
	utils.EncodePayload(w, http.StatusOK, res)
}

func ServerGroupAddServer(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Server Group not found.")
		return
	}

	g := models.ServerGroup{ID: id}
	g.Get(db)

	if g.Name == "" {
		utils.NotFoundResponse(w, "Server Group not found")
		return
	}

	server := models.Server{}
	if err := utils.DecodePayload(r, &server); err != nil {
		utils.BadRequestResponse(w, "Data appears to be invalid.")
		return
	}
	server.Get(db)
	if server.Name == "" {
		utils.NotFoundResponse(w, "Server not found")
		return
	}

	s := g.AddServer(db, &server)
	if s == true {
		g.Servers = append(g.Servers, server)
		res := utils.Response{Result: "success", Data: g}
		utils.EncodePayload(w, http.StatusOK, res)
	} else {
		utils.BadRequestResponse(w, "Failed to add server to group.")
	}
}

func ServerGroupRemoveServer(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Server Group not found.")
		return
	}

	sid, err := strconv.ParseInt(c.URLParams["sid"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Server not found.")
		return
	}

	g := models.ServerGroup{ID: id}
	server := models.Server{ID: sid}

	g.Get(db)
	if g.Name == "" {
		utils.NotFoundResponse(w, "Server Group not found")
		return
	}

	s := g.RemoveServer(db, &server)
	if s == true {
		cs := []models.Server{}
		for _, ct := range g.Servers {
			if ct.ID != server.ID {
				cs = append(cs, ct)
			}
		}
		g.Servers = cs
		res := utils.Response{Result: "success", Data: g}
		utils.EncodePayload(w, http.StatusOK, res)
	} else {
		utils.BadRequestResponse(w, "Failed to remove server from group.")
	}
}
