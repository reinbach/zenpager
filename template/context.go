package template

import (
	"net/http"

	"git.ironlabs.com/greg/zenpager/session"
)

type Context struct {
	Values map[string]interface{}
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	ctx := &Context{}
	msg := session.GetMessages(w, r)
	ctx.Add("Msg", msg)
	return ctx
}

func (c *Context) Add(key string, value interface{}) {
	if c.Values == nil {
		c.Values = make(map[string]interface{}, 1)
	}
	c.Values[key] = value
}
