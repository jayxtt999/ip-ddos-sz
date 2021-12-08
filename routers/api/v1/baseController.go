package v1

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

type BaseController struct {

}

//
//  GetIp
//  @Description: 获取ip参数
//  @receiver con
//  @param c
//  @return string
//
func (con BaseController) GetIp(c *gin.Context) string {
	ip :=  c.PostForm("ip")
	return ip
}

//
//  CheckIp
//  @Description: 验证ip是否合法
//  @receiver con
//  @param ip
//  @return bool
//
func (con BaseController) CheckIp(ip string) bool {
	address := net.ParseIP(ip)
	if address == nil {
		return false
	}else {
		return true
	}
}

//
//  Success
//  @Description: 通用成功返回
//  @receiver con
//  @param c
//  @param message
//
func (con BaseController) Success(c *gin.Context,message string) {
	c.JSON(http.StatusOK, gin.H{"message": message, "code": 1})
}

//
//  Error
//  @Description: 通用错误返回
//  @receiver con
//  @param c
//  @param message
//  @param code
//
func (con BaseController) Error(c *gin.Context,message string,code int) {
	c.JSON(http.StatusOK, gin.H{"message": message, "code": code})
}
