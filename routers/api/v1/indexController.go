package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jayxtt999/ip-ddos-sz/pkg/setting"
	"github.com/jayxtt999/ip-ddos-sz/service"
)

type IndexController struct {
	BaseController
}

func (con IndexController) JdBlock(c *gin.Context) {
	ip := c.PostForm("ip")
	check := con.CheckIp(ip)
	if check == false {
		con.Error(c, "IP 格式错误", 0)
		return
	}
	//构造连接实例，如下进行了4种连接，分别为 centos[ssh] centos[telnet] h3c[ssh] h3c[telnet]
	//它们最终都需要执行 `BlockIp` 方法

	setting.Cfg.Section("centos_ssh").MapTo(setting.OsSetting)
	centosSsh := service.NewRouteCentOSSsh(ip, *setting.OsSetting)

	setting.Cfg.Section("centos_telnet").MapTo(setting.OsSetting)
	centosTelnet := service.NewRouteCentOSTelnet(ip, *setting.OsSetting)

	setting.Cfg.Section("h3c_ssh").MapTo(setting.OsSetting)
	h3cSsh := service.NewRouteH3cSsh(ip, *setting.OsSetting)

	setting.Cfg.Section("h3c_telnet").MapTo(setting.OsSetting)
	h3cTelnet := service.NewRouteH3cTelnet(ip, *setting.OsSetting)

	subject := service.NewRouteSubject()
	subject.Add(centosSsh)
	subject.Add(centosTelnet)
	subject.Add(h3cSsh)
	subject.Add(h3cTelnet)
	res := subject.BlockIp(ip)
	var msg = "未知"
	switch res {
	case 0:
		msg = "部分成功"
		break

	case -1:
		msg = "失败"
		break

	case 1:
		msg = "成功"
		break
	default:
		break
	}
	con.Success(c, msg)

}

//
//  JdFree
//  @Description:
//  @receiver con
//  @param c
//
func (con IndexController) JdFree(c *gin.Context) {
	//todo
}

//
//  BgBlock
//  @Description:
//  @receiver con
//  @param c
//
func (con IndexController) BgBlock(c *gin.Context) {
	//todo

}

//
//  BgFree
//  @Description:
//  @receiver con
//  @param c
//
func (con IndexController) BgFree(c *gin.Context) {
	//todo

}

//
//  BgBatchFree
//  @Description:
//  @receiver con
//  @param c
//
func (con IndexController) BgBatchFree(c *gin.Context) {
	//todo

}

//
//  AdService
//  @Description:
//  @receiver con
//  @param c
//
func (con IndexController) AdService(c *gin.Context) {
	//todo

}
