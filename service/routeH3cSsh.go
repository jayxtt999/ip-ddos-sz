package service

import (
	"github.com/jayxtt999/ip-ddos-sz/pkg/setting"
	"github.com/jayxtt999/ip-ddos-sz/service/client"
)
//
//  RouteH3cSsh
//  @Description: H3c SSH方式的实现
//
type RouteH3cSsh struct {
	ip     string
	config setting.Os
	RouteBasic
}

func  NewRouteH3cSsh(ip string, config setting.Os) *RouteH3cSsh {
	return &RouteH3cSsh{
		ip:     ip,
		config: config,
	}
}

//
//  BlockIp
//  @Description: 封堵
//  @receiver h3c
//  @return bool
//
func (h3c RouteH3cSsh) BlockIp() bool {
	ssh := client.H3cSsh{}
	ssh.Config(h3c.config.Host, h3c.config.Port, h3c.config.User, h3c.config.Pass)
	var r bool
	cmds := make([]string, 0)
	cmds = append(cmds, "system-view")
	cmds = append(cmds, "quit")
	cmds = append(cmds, "quit")
	r = ssh.Shell(cmds)
	if !r {
		return false
	}
	return true
}
