package controller

// func  router.GET("/panic", func(c *gin.Context) {
// 	// panic with a string -- the custom middleware could save this to a database or report it to the user
// 	panic("500 Internal Server Error")
// })

// Upload example
//
//	@Summary		Upload file
//	@Description	Upload file
//	@ID				file.upload
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			file	formData	file			true	"this is a test file"
//	@Success		200		{string}	string			"ok"
//	@Failure		400		{object}	web.APIError	"We need ID!!"
//	@Failure		404		{object}	web.APIError	"Can not find ID"
//	@Router			/file/upload [post]
// func Upload(w http.ResponseWriter, r *http.Request) {
// 	// write your code
// }

// func async(c *gin.Context) {
// 	// 需要搞一个副本
// 	copyContext := c.Copy()
// 	// 异步处理
// 	go func() {
// 		time.Sleep(3 * time.Second)
// 		log.Println("异步执行：" + copyContext.Request.URL.Path)
// 	}()
// 	c.JSON(200,gin.H{"msg":"异步请求成功"})
// })

// func (c *gin.Context) {
// 	time.Sleep(3 * time.Second)
// 	log.Println("同步执行：" + c.Request.URL.Path)
// 	c.JSON(200,gin.H{"msg":"同步请求成功"})
// })

// // FormFile returns the first file for the provided form key.
// // FormFile calls ParseMultipartForm and ParseForm if necessary.
// func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error) {
// 	if r.MultipartForm == multipartByReader {
// 			return nil, nil, errors.New("http: multipart handled by MultipartReader")
// 	}
// 	if r.MultipartForm == nil {
// 			err := r.ParseMultipartForm(defaultMaxMemory)
// 			if err != nil {
// 					return nil, nil, er
// 			}
// 	}
// 	if r.MultipartForm != nil && r.MultipartForm.File != nil {
// 			if fhs := r.MultipartForm.File[key]; len(fhs) > 0 {
// 					f, err := fhs[0].Open()
// 					return f, fhs[0], er
// 			}
// 	}
// 	return nil, nil, ErrMissingFile
// }

// func HandleNotFound(c *gin.Context) {
// 	global.NewResult(c).Error(404,"资源未找到")
// 	return
// }

// router := gin.Default()
// //处理异常
// router.NoRoute(HandleNotFound)
// router.NoMethod(HandleNotFound)
// //use middleware
// router.Use(middleware.MiddlewareTwo())
// router.Use(middleware.MiddlewareOne())
// router.Use(middleware.MiddlewareThree())
// router.Use(Recover)
