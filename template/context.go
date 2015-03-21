package template

import (
	"net/http"

	"github.com/zenazn/goji/web"

	"github.com/reinbach/zenpager/session"
)

type Context struct {
	Values map[string]interface{}
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) Add(key string, value interface{}) {
	if c.Values == nil {
		c.Values = make(map[string]interface{}, 1)
	}
	c.Values[key] = value
}

func (c *Context) GetMessages(cweb web.C, w http.ResponseWriter, r *http.Request) {
	c.Add("Messages", session.GetMessages(&cweb, w, r))
}
