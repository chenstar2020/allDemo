package middleware

type HandlerFunc func(ctx *Context)

type Context struct {
	handlers []HandlerFunc
	index int
}

func newContext()*Context{
	return &Context{
		index: -1,
	}
}


func (c *Context)Next(){
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++{
		f := c.handlers[c.index]
		f(c)
	}
}

func (c *Context)Use(f HandlerFunc){
	c.handlers = append(c.handlers, f)
}