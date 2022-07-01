package server

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net/http"
	"project/transmit/control"
	"project/transmit/webSocket"
	"strings"
)

//go:embed frontend/dist/*
var FS embed.FS

func Run() {
	const port string = "27149"
	hub := webSocket.NewHub()
	go hub.Run()
	go func() {
		gin.SetMode(gin.DebugMode) //调试模式
		gin.DisableConsoleColor()
		router := gin.Default()
		staticFile , _ := fs.Sub(FS , "frontend/dist")  //将打包后的文件变成一个结构化的目录
		router.POST("/api/v1/texts", control.TextController)
		router.POST("/api/v1/files", control.FilesController)
		router.GET("/uploads/:path", control.UploadsController)
		router.GET("/api/v1/addresses",control.AddressController)
		router.GET("/api/v1/qrcodes",control.QrcodesController)
		router.GET("/ws", func(c *gin.Context) {
			webSocket.HttpController(c, hub)
		})
		router.StaticFS("/static",http.FS(staticFile))
		//如果请求的不是设定的路径，则执行下面代码，提供更好的用户体验
		router.NoRoute(func(c *gin.Context) {
			path := c.Request.URL.Path  //获取用户访问的路径
			if strings.HasPrefix(path, "/static/") {  //判断路径前是否有static（自己指定）
				reader, err := staticFile.Open("index.html")  //如果找不到指定的文件，就按照index.html统一渲染
				if err != nil {
					log.Fatal(err)
				}
				defer reader.Close()
				stat, err := reader.Stat()  //计算（统计）文件的长度
				if err != nil {
					log.Fatal(err)
				}
				c.DataFromReader(http.StatusOK, stat.Size(), "text/html;charset=utf-8", reader, nil)
			} else {
				c.Status(http.StatusNotFound)
			}
		})
		router.Run(":" + port) //gin在这会执行完毕，同时主进程也会结束，所以在这起协程
	}()
}
