package template

import (
	"net/http"

	"github.com/zenazn/goji/web"

	"git.ironlabs.com/greg/zenpager/session"
)

type Context struct {
	Values map[string]interface{}
}

func NewContext(c *web.C, w http.ResponseWriter, r *http.Request) *Context {
	ctx := &Context{}
	// check for messages
	ctx.Add("Messages", session.GetMessages(c, w, r))
	// check for form errors
	return ctx
}

func (c *Context) Add(key string, value interface{}) {
	if c.Values == nil {
		c.Values = make(map[string]interface{}, 1)
	}
	c.Values[key] = value
}
