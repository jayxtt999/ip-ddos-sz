package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jayxtt999/ip-ddos-sz/pkg/setting"
	"github.com/jayxtt999/ip-ddos-sz/pkg/tools"
	"net/http"
	"strings"
)

//
//  Middleware
//  @Description: 验证访问ip是否合法，token 等
//  @return gin.HandlerFunc
//
func  Middleware() gin.HandlerFunc {

	return func(c *gin.Context){
		//验证请求来源是否合法
		checkIp := setting.ServerSetting.CheckAllowIp
		ipList := setting.ServerSetting.AllowIpList
		if checkIp == true && ipList!=""{
			ip := GetRequestIP(c)
			ipListArr := strings.Split(ip,",")
			if !tools.InArray(ip,ipListArr){
				c.Abort()
				c.JSON(http.StatusForbidden, gin.H{"message": "forbidden access", "code": 0})
				return
			}
		}
		//todo 验证token
		/*token := c.Request.Header["Token"][0]
		if token != "aaaa"{
			c.Abort()
			c.JSON(http.StatusForbidden, gin.H{"message": "token error", "code": 0})
			return
		}*/
	}

}

func GetRequestIP(c *gin.Context)string{
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}