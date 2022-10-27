package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

//Context 保存http头部信息
type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request

	Method     string
	Path       string
	StatusCode int
	Params     map[string]string
	//保存中间件
	handlers []HandlerFunc
	index    int
	//可调用engine中的模板处理程序
	engine *Engine
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Method: r.Method,
		Path:   r.URL.Path,
		index:  -1,
	}
}

func (c *Context) Faile(code int, err string) {
	//结束所有中间件和handler调用
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) Query(key string) string {
	//返回路径参数对应的值
	return c.Req.URL.Query().Get(key)
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

//设置头部信息
func (c *Context) SetHeader(key string, val string) {
	c.Writer.Header().Set(key, val)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	//TODO ?header
	c.SetHeader("content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	//成功的话就已经将obj解析成字符串写进c.Writer
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader("content-Type", "text/html")
	c.Status(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Faile(500, err.Error())
	}
}
