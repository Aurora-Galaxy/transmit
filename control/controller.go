package control

import (
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// GetUploadsDir 获取执行文件所在目录并在路径添加uploads
func GetUploadsDir() (uploads string) {
	exe, err := os.Executable() //获取当前文件执行路径
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(exe) //获取当前文件的执行目录
	uploads = filepath.Join(dir, "uploads") //获取uploads目录
	return
}


// QrcodesController 生成二维码
func QrcodesController(c *gin.Context) {
	if content := c.Query("content"); content != "" { //获取路径
		png, err := qrcode.Encode(content, qrcode.Medium, 256) 	//生成二维码
		if err != nil {
			log.Fatal(err)
		}
		c.Data(http.StatusOK, "image/png", png) 	//返回二维码给前端
	} else {
		c.Status(http.StatusBadRequest)
	}
}

