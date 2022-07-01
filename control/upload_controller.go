package control

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func UploadsController(c *gin.Context) {
	if path := c.Param("path"); path != "" { //获取路径
		target := filepath.Join(GetUploadsDir(), path)
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attachment; filename="+path)
		c.Header("Content-Type", "application/octet-stream") //binary application/octet-stream 传输二进制流，可以是任意类型
		c.File(target) //给前端发送一个任意类型文件
	} else {
		c.Status(http.StatusNotFound)
	}
}
