package middleware

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %v\n", err)
				debug.PrintStack()

				if c.IsAborted() {
					c.Status(200)
				}
				switch errStr := err.(type) {
				case string:
					p := strings.Split(errStr, "#")
					if len(p) == 3 && p[0] == "CustomError" {
						statusCode, e := strconv.Atoi(p[1])
						if e != nil {
							break
						}
						c.Status(statusCode)
						fmt.Println(
							time.Now().Format("\n 2006-01-02 15:04:05.9999"),
							"[ERROR]",
							c.Request.Method,
							c.Request.URL,
							statusCode,
							c.Request.RequestURI,
							c.ClientIP(),
							p[2],
						)
						c.JSON(http.StatusOK, gin.H{
							"code": statusCode,
							"msg":  p[2],
						})
					}
				default:
					panic(err)
				}
			}
		}()

		c.Next()
	}
}
