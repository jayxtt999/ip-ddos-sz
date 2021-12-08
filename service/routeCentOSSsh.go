package service

import (
	"github.com/jayxtt999/ip-ddos-sz/pkg/setting"
	"github.com/jayxtt999/ip-ddos-sz/service/client"
)
//
//  RouteCentOSSsh
//  @Description: Centos SSH方式的实现
//
type RouteCentOSSsh struct {
	ip     string
	config setting.Os
	RouteBasic
}

func  NewRouteCentOSSsh(ip string, config setting.Os) *RouteCentOSSsh {
	return &RouteCentOSSsh{
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
func (centOs RouteCentOSSsh) BlockIp() bool {
	ssh := client.CentOsSsh{}
	ssh.SshInf.Config(centOs.config.Host, centOs.config.Port, centOs.config.User, centOs.config.Pass)
	var r bool
	r = ssh.Login()
	if !r {
		return false
	}
	shell := "ls"
	r = ssh.Exec(shell, '#')
	if !r {
		return false
	}
	r = ssh.Logout()
	if !r {
		return false
	}
	return true
}
