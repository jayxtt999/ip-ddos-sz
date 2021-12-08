package client

import (
	"fmt"
	"github.com/jayxtt999/ip-ddos-sz/pkg/logging"
	"github.com/scottkiss/gosshtool"
	"runtime"
)
//
//  SshInf
//  @Description: ssh 接口
//
type SshInf struct {
	host      string
	port      int
	user      string
	pass      string
	sshClient *gosshtool.SSHClient
}
//
//  Config
//  @Description: ssh接口配置
//  @receiver inf
//  @param host
//  @param port
//  @param user
//  @param pass
//
func (inf *SshInf) Config(host string, port int, user string, pass string) {
	inf.host = host
	inf.port = port
	inf.user = user
	inf.pass = pass
}

func (inf *SshInf)checkErr(err error) bool {
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			emsg := fmt.Sprintf("file:%s, line:%d, error:%s", file, line, err.Error())
			logging.TagError(inf.host,emsg)
		} else {
			logging.TagError(inf.host,err)
		}
		return true
	}
	return false
}