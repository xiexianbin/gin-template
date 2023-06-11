# gin-incubator

[![build-test](https://github.com/xiexianbin/gin-incubator/actions/workflows/workflow.yaml/badge.svg)](https://github.com/xiexianbin/gin-incubator/actions/workflows/workflow.yaml)
[![GoDoc](https://godoc.org/github.com/xiexianbin/gin-incubator?status.svg)](https://pkg.go.dev/github.com/xiexianbin/gin-incubator)
[![Go Report Card](https://goreportcard.com/badge/github.com/xiexianbin/gin-incubator)](https://goreportcard.com/report/github.com/xiexianbin/gin-incubator)

golang gin web framework demo

## usage

- release

```
git tag v0.1.0
git push origin --tags
```

- download
```
curl -Lfs -o main https://github.com/xiexianbin/gin-incubator/releases/latest/download/main-{linux|darwin|windows}
chmod +x main
./main
```

- [gin example](https://gin-gonic.com/docs/examples/)
  - https://github.com/gin-gonic/examples
- [swag example](https://github.com/swaggo/swag/tree/master/example)


	app := gin.Default()

	// serve static favicon file from a location relative to main.go directory
	//app.StaticFile("/favicon.ico", "./.assets/favicon.ico")
	app.StaticFile("/favicon.ico", "./favicon.ico")

	app.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello favicon.")
	})


	router.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		// Multipart form
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, "get form err: %s", err.Error())
			return
		}
		files := form.File["files"]

		for _, file := range files {
			filename := filepath.Base(file.Filename)
			if err := c.SaveUploadedFile(file, filename); err != nil {
				c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
				return
			}
		}

		c.String(http.StatusOK, "Uploaded successfully %d files with fields name=%s and email=%s.", len(files), name, email)
	})


POST   /post/test?id=1234&page=1  HTTP/1.1
id := c.Query("id") //查询请求URL后面的参数
page := c.DefaultQuery("page", "0") //查询请求URL后面的参数，如果没有填写默认值
name := c.PostForm("name") //从表单中查询参数

/////////////////////////////////
//POST和PUT主体参数优先于URL查询字符串值。
name := c.Request.FormValue("name") 

//返回POST并放置body参数，URL查询参数被忽略
name := c.Request.PostFormValue("name")

//从表单中查询参数，如果没有填写默认值  
message := c.DefaultPostForm("message", "aa") 
复制代码
假如gin定义的路由路径为:

router.POST("/post/:uuid", func(c *gin.Context){
    ...
}
则获取uuid的值方法为

uuid := c.Param("uuid") //取得URL中参数
其他:

复制代码
s, _ := c.Get("current_manager") //从用户上下文读取值      

var u User

//从http.Request中读取值到User结构体中，手动确定绑定类型binding.Form
err1 := c.BindWith(&u, binding.Form) 

//从http.Request中读取值到User结构体中,根据请求方法类型和请求内容格式类型自动确定绑定类型
err2 := c.Bind(&u)

//从session中读取值
//用户上下文和session生命周期不同，每一次请求会生成一个对应的上下文，一次http请求结束，该次请求的上下文结束，一般来说session(会话)会留存一段时间
//session(会话)中一般保存用户登录状态等信息，context(上下文)主要用于在一次http请求中，在中间件(流)中进行信息传递
user := sessions.Default(c).get("user") 


https://www.tizi365.com/archives/257.html

	// curl http://127.0.0.1:8080/ping
	//{"message":"pong"}
	

func initMiddleware(c *gin.Context) {
	fmt.Println("initMiddleware")
}
