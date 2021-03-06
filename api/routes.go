package api

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"

	mw "github.com/reinbach/zenpager/middleware"
	"github.com/reinbach/zenpager/utils"
)

func Routes(p string) *web.Mux {
	api := web.New()
	api.Handle(fmt.Sprintf("%s/auth/*", p), AuthRoutes())
	api.Handle(fmt.Sprintf("%s/user/*", p), UserRoutes())
	api.Get("/user",
		http.RedirectHandler("/user/", 301))
	api.Handle(fmt.Sprintf("%s/contacts/*", p), ContactRoutes())
	api.Get("/contacts",
		http.RedirectHandler("/contacts/", 301))
	api.Handle(fmt.Sprintf("%s/servers/*", p), ServerRoutes())
	api.Get("/servers",
		http.RedirectHandler("/servers/", 301))
	api.Handle(fmt.Sprintf("%s/commands/*", p), CommandRoutes())
	api.Get("/commands",
		http.RedirectHandler("/commands/", 301))

	return api
}

func AuthRoutes() *web.Mux {
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
	api.Use(mw.Authenticate)
	// api.Get("/user/", UserList)
	// api.Get("/user/:id", UserItem)
	// api.Post("/user/", UserAdd)
	// api.Put("/user/:id", UserUpdate)
	api.Patch("/:id", UserPartialUpdate)
	// api.Delete("/user/:id", UserDelete)
	// api.Get("/user", http.RedirectHandler("/user/", 301))

	return api
}

func ContactRoutes() *web.Mux {
	api := web.New()
	api.Use(middleware.SubRouter)
	api.Use(mw.Authenticate)

	// contacts
	api.Get("/", ContactList)
	api.Get("/:id", ContactItem)
	api.Post("/", ContactAdd)
	api.Put("/:id", ContactUpdate)
	api.Patch("/:id", ContactUpdate)
	api.Delete("/:id", ContactDelete)
	// contacts groups
	api.Get("/:id/groups/", ContactGroups)

	// contact groups
	api.Get("/groups/", ContactGroupList)
	api.Get("/groups/:id", ContactGroupItem)
	api.Post("/groups/", ContactGroupAdd)
	api.Put("/groups/:id", ContactGroupUpdate)
	api.Patch("/groups/:id", ContactGroupUpdate)
	api.Delete("/groups/:id", ContactGroupDelete)
	// contact groups contacts
	api.Get("/groups/:id/contacts/", ContactGroupContacts)
	api.Post("/groups/:id/contacts/", ContactGroupAddContact)
	api.Delete("/groups/:id/contacts/:cid", ContactGroupRemoveContact)

	return api
}

func ServerRoutes() *web.Mux {
	api := web.New()
	api.Use(middleware.SubRouter)
	api.Use(mw.Authenticate)

	// servers
	api.Get("/", ServerList)
	api.Get("/:id", ServerItem)
	api.Post("/", ServerAdd)
	api.Put("/:id", ServerUpdate)
	api.Patch("/:id", ServerUpdate)
	api.Delete("/:id", ServerDelete)
	// servers groups
	api.Get("/:id/groups/", ServerGroups)

	// server groups
	api.Get("/groups/", ServerGroupList)
	api.Get("/groups/:id", ServerGroupItem)
	api.Post("/groups/", ServerGroupAdd)
	api.Put("/groups/:id", ServerGroupUpdate)
	api.Patch("/groups/:id", ServerGroupUpdate)
	api.Delete("/groups/:id", ServerGroupDelete)
	// server groups servers
	api.Get("/groups/:id/servers/", ServerGroupServers)
	api.Post("/groups/:id/servers/", ServerGroupAddServer)
	api.Delete("/groups/:id/servers/:sid", ServerGroupRemoveServer)

	return api
}

func CommandRoutes() *web.Mux {
	api := web.New()
	api.Use(middleware.SubRouter)
	api.Use(mw.Authenticate)

	// commands
	api.Get("/", CommandList)
	api.Get("/:id", CommandItem)
	api.Post("/", CommandAdd)
	api.Put("/:id", CommandUpdate)
	api.Patch("/:id", CommandUpdate)
	api.Delete("/:id", CommandDelete)
	// commands groups
	api.Get("/:id/groups/", CommandGroups)

	// command groups
	api.Get("/groups/", CommandGroupList)
	api.Get("/groups/:id", CommandGroupItem)
	api.Post("/groups/", CommandGroupAdd)
	api.Put("/groups/:id", CommandGroupUpdate)
	api.Patch("/groups/:id", CommandGroupUpdate)
	api.Delete("/groups/:id", CommandGroupDelete)
	// command groups commands
	api.Get("/groups/:id/commands/", CommandGroupCommands)
	api.Post("/groups/:id/commands/", CommandGroupAddCommand)
	api.Delete("/groups/:id/commands/:sid", CommandGroupRemoveCommand)

	return api
}
