package gee

import "strings"

type router struct {
	root     map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		root:     make(map[string]*node),
	}
}

//parsePattern 将路径以"/"分隔开
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)

	for _, part := range vs {
		// "/"的前面会产生一个 ""字符串
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern // eg:GET-/hello/name

	parts := parsePattern(pattern)
	_, ok := r.root[method]
	if !ok { //还没有以metohd为根的树
		r.root[method] = &node{}
	}
	r.root[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	//path 中是真实的路径，已经将占位符取代
	searchPath := parsePattern(path)

	params := make(map[string]string)
	root, ok := r.root[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchPath, 0)

	if n != nil {
		//n.pattern中仍保留占位符
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchPath[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchPath[index:], "/")
				break
			}
		}
	}
	return n, params

}

func (r *router) getRoutes(method string) []*node {
	root, ok := r.root[method]
	if !ok {
		return nil
	}
	nodes := make([]*node, 0)
	root.travel(&nodes)
	return nodes
}

func (r *router) handle(c *Context) {
	n, parms := r.getRoute(c.Method, c.Path)

	if n != nil {
		c.Params = parms
		key := c.Method + "-" + n.pattern
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(404, "NOT FOUND : %s \n", c.Path)
		})
	}
	//执行c的中间件和用户handler
	c.Next()
}
