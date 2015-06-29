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

	// contact groups
	api.Get("/groups/", ContactGroupList)
	api.Get("/groups/:id", ContactGroupItem)
	api.Post("/groups/", ContactGroupAdd)
	api.Put("/groups/:id", ContactGroupUpdate)
	api.Patch("/groups/:id", ContactGroupUpdate)
	api.Delete("/groups/:id", ContactGroupDelete)

	// contact groups contacts
	api.Get("/groups/:id/contacts/", ContactGroupContacts)
	api.Post("/groups/:id/contacts/", ContactGroupContactAdd)
	api.Delete("/groups/:id/contacts/", ContactGroupRemoveContact)

	return api
}
