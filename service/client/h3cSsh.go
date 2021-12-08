package client

import (
	"fmt"
	"github.com/jayxtt999/ip-ddos-sz/pkg/logging"
	"github.com/shenbowei/switch-ssh-go"
)

type H3cSsh struct {
	SshInf
}

//
//  Shell
//  @Description: 执行shell
//  @receiver t
//  @param cmds
//  @return bool
//
func(t *H3cSsh) Shell(cmds []string) bool {
	logging.TagInfo(t.SshInf.host,fmt.Sprintf("exec shell:%s",cmds))
	addr := fmt.Sprintf("%s:%d", t.SshInf.host, t.SshInf.port)
	result, err := ssh.RunCommandsWithBrand(t.SshInf.user, t.SshInf.pass, addr, ssh.H3C, cmds...)
	if t.checkErr(err) {
		return false
	}
	logging.TagInfo(t.SshInf.host,fmt.Sprintf("RunCommands result:%s", result))
	return true
}