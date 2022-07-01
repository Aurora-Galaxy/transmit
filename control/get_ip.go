package control

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

// AddressController 获取电脑所在网络IP地址
func AddressController(c *gin.Context){
	address  , _ := net.InterfaceAddrs()  //获取所有地址
	var result []string
	for _ , addre := range address{
		ipnet , ok := addre.(*net.IPNet)  //断言出ip地址
		if ok && !ipnet.IP.IsLoopback(){ //判断是否是ipv4
			if ipnet.IP.To4() != nil{
				result = append(result , ipnet.IP.String())
			}
		}
	}
	c.JSON(http.StatusOK , gin.H{"addresses" : result})
}
