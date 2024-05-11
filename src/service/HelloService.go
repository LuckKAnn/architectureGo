package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
)

type AddParam struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func AddBodyParam(c *gin.Context) {
	var addParam AddParam
	if err := c.ShouldBindJSON(&addParam); err != nil {
		c.JSON(400, gin.H{
			"ERROR": "Caculate Failure",
		})
		return
	}
	c.JSON(200, gin.H{"data": addParam.Y + addParam.X})
}

func SayHello(c *gin.Context) {
	ip := getHostIp()
	fmt.Println()
	c.String(200, "Hello, LuckKun V1 , IP:"+ip)
}
func getHostIp() string {
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("get current host ip err: ", err)
		return ""
	}
	var ip string
	for _, address := range addrList {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
				break
			}
		}
	}
	return ip
}
