package service

import (
	"github.com/jayxtt999/ip-ddos-sz/pkg/setting"
	"github.com/jayxtt999/ip-ddos-sz/service/client"
)

//
//  RouteCentOSTelnet
//  @Description: Centos telnet 方式的实现
//
type RouteCentOSTelnet struct {
	ip     string
	config setting.Os
	RouteBasic
}

func NewRouteCentOSTelnet(ip string, config setting.Os) *RouteCentOSTelnet {
	return &RouteCentOSTelnet{
		ip:     ip,
		config: config,
	}
}
//
//  BlockIp
//  @Description: 封堵，shell 是模拟的，一般常规linux做封堵是gobgp
//  @receiver centOs
//  @return bool
//
func (centOs RouteCentOSTelnet) BlockIp() bool {
	telnet := client.CentOsTelnet{}
	telnet.TelnetInf.Config(centOs.config.Host, centOs.config.Port, centOs.config.User, centOs.config.Pass)
	var r bool
	r = telnet.Login()
	if !r {
		return false
	}
	shell := "ls"//假设ls命令是封堵
	err := telnet.Exec(shell, '#')
	if err != nil {
		return false
	}
	telnet.Logout()

	return true
}
