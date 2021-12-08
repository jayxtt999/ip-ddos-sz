package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jayxtt999/ip-ddos-sz/routers/api/v1"
	"github.com/jayxtt999/ip-ddos-sz/service"
	"net/http"
)


// InitRouter initialize routing information
func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//test...
	r.GET("/", func(c *gin.Context) {
		c.Abort()
		c.JSON(http.StatusOK, gin.H{"message": "test", "code": 0})
		return
	})
	//网关验证器
	r.Use(service.Middleware())

	//r.POST("/v1.0/firewall/execute_ip_traction", v1.IndexController{}.AdService)
	apiV1 := r.Group("/v1")
	{
		apiV1.POST("/jd_block", v1.IndexController{}.JdBlock)
		//apiV1.POST("/jd_free", v1.IndexController{}.JdFree)
		//apiV1.POST("/bg_block", v1.IndexController{}.BgBlock)
		//apiV1.POST("/bg_free", v1.IndexController{}.BgFree)
		//apiV1.POST("/bg_batch_free", v1.IndexController{}.BgBatchFree)
	}
	return r
}
