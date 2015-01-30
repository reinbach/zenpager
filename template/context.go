package template

type Context struct {
	Values map[string]string
}

func (c *Context) Add(key, value string) {
	if c.Values == nil {
		c.Values = make(map[string]string)
	}
	c.Values[key] = value
}
