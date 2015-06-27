package api

import (
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"

	mw "github.com/reinbach/zenpager/middleware"
)

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
	api.Get("/groups/", GroupList)
	api.Get("/groups/:id", GroupItem)
	api.Post("/groups/", GroupAdd)
	api.Put("/groups/:id", GroupUpdate)
	api.Patch("/groups/:id", GroupUpdate)
	api.Delete("/groups/:id", GroupDelete)

	return api
}
