package main

import (
	"fmt"
	"gee"
	"log"
	"net/http"
	"text/template"
	"time"
)

func onlyForV2(c *gee.Context) {
	t := time.Now()
	c.Faile(500, "internet server Error")
	log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
}

type student struct {
	Name string
	Age  int
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gee.Default()
	//注册hello 与之对应的handler
	// r.GET("/", func(w http.ResponseWriter, req *http.Request) {
	// 	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	// })
	// r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
	// 	for k, v := range req.Header {
	// 		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	// 	}
	// })

	// r.GET("/", func(c *gee.Context) {
	// 	c.HTML(http.StatusOK, "<h1>hello<h1>")
	// })
	// r.POST("/login", func(c *gee.Context) {
	// 	c.JSON(http.StatusOK, gee.H{
	// 		"username": c.PostForm("username"),
	// 		"passward": c.PostForm("passward"),
	// 	})
	// })

	// r.GET("/hi", func(c *gee.Context) {
	// 	c.String(http.StatusOK, "hi %s you are at %s\n", c.Query("name"), c.Path)
	// })
	// 	c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	// })
	// v1 := r.Group("/v1")
	// {
	// 	v1.GET("/", func(c *gee.Context) {
	// 		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	// 	})
	// 	v1.GET("/hello", func(c *gee.Context) {
	// 		// expect /hello?name=geektutu
	// 		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	// 	})
	// }
	// v2 := r.Group("/v2")
	// {
	// 	v2.GET("/hello/:name", func(c *gee.Context) {
	// 		// expect /hello/geektutu
	// 		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	// 	})
	// 	v2.POST("/login", func(c *gee.Context) {
	// 		c.JSON(http.StatusOK, gee.H{
	// 			"username": c.PostForm("username"),
	// 			"password": c.PostForm("password"),
	// 		})
	// 	})
	// }

	// //全局中间件
	// r.Use(gee.Logger())
	// r.GET("/", func(c *gee.Context) {
	// //	c.HTML(http.StatusOK, "<h1> hello Gee</h1>")
	// })
	// v2 := r.Group("/v2")
	// //给分组v2添加中间件
	// v2.Use(onlyForV2)
	// {
	// 	v2.GET("/hello/:name", func(c *gee.Context) {
	// 		c.String(http.StatusOK, "hello %s, you are at %s", c.Param("name"), c.Path)
	// 	})
	// }
	// r.Run(":8080")

	r.Use(gee.Logger())
	//funcmap
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	//加载模板进内存
	r.LoadHTMLGlob("templates/*")
	//将用户访问的资源地址映射到真实的资源地址，并返回文件
	r.Static("/assets", "./static")
	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
	r.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "gee",
			"now":   time.Now(),
		})
	})

	r.GET("/panic", func(c *gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})
	r.Run(":8080")
}
