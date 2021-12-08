package service

import (
	"github.com/jayxtt999/ip-ddos-sz/pkg/setting"
	"github.com/jayxtt999/ip-ddos-sz/service/client"
)
//
//  RouteH3cTelnet
//  @Description: H3c telnet方式的实现
//
type RouteH3cTelnet struct {
	ip     string
	config setting.Os
	RouteBasic
}

func NewRouteH3cTelnet(ip string, config setting.Os) *RouteH3cTelnet {
	return &RouteH3cTelnet{
		ip:     ip,
		config: config,
	}
}

func (h3c RouteH3cTelnet) BlockIp() bool {
	telnet := client.H3cTelnet{}
	telnet.Config(h3c.config.Host, h3c.config.Port, h3c.config.User, h3c.config.Pass)
	var r bool
	r = telnet.Login()
	if !r {
		return false
	}
	shell := "system-view"
	err := telnet.Exec(shell, ']')
	if err != nil {
		return false
	}
	shell = "quit"
	err = telnet.Exec("quit", '>')
	if err != nil {
		return false
	}
	r = telnet.Logout()
	if !r {
		return false
	}

	return true
}
