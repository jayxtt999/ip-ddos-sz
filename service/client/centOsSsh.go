package client

import (
	"fmt"
	"github.com/jayxtt999/ip-ddos-sz/pkg/logging"
	"github.com/scottkiss/gosshtool"
)

type CentOsSsh struct {
	SshInf
}
//
//  Login
//  @Description: 登陆
//  @receiver t
//  @return bool
//
func (t *CentOsSsh) Login() bool {

	sshConfig := &gosshtool.SSHClientConfig{
		User:     t.SshInf.user,
		Password: t.SshInf.pass,
		Host:     t.SshInf.host,
	}
	t.SshInf.sshClient = gosshtool.NewSSHClient(sshConfig)
	logging.TagInfo(t.SshInf.host,fmt.Sprintf("ssh login:%s:%d,user:%s", t.SshInf.host, t.SshInf.port, t.SshInf.user))
	stdout, stderr, _, err := t.SshInf.sshClient.Cmd("echo hi", nil, nil, 0)//`echo hi`简单验证是否登陆成功
	if t.SshInf.checkErr(err) {
		return false
	}
	logging.TagInfo(t.SshInf.host,fmt.Sprintf("stdout:%s,stderr:%s", stdout, stderr))
	//简单验证是否登陆成功
	if stdout != "hi\n" {
		logging.TagInfo(t.SshInf.host,"login err..")
		return false
	}
	return true

}
//
//  Exec
//  @Description: 执行shell
//  @receiver t
//  @param shell
//  @param delim
//  @return bool
//
func (t *CentOsSsh) Exec(shell string, delim byte) bool {

	logging.TagInfo(t.SshInf.host,fmt.Sprintf("exec shell:%s", shell))
	stdout, stderr, _, err := t.SshInf.sshClient.Cmd(shell, nil, nil, 0)
	logging.TagInfo(t.SshInf.host,fmt.Sprintf("stdout:%s,stderr:%s", stdout, stderr))
	if t.checkErr(err) {
		return false
	}
	return true
}

//
//  Logout
//  @Description: 退出
//  @receiver t
//  @return bool
//
func (t *CentOsSsh) Logout() bool {
	logging.TagInfo(t.SshInf.host,"exec shell:exit")
	stdout, stderr, _, err := t.SshInf.sshClient.Cmd("exit", nil, nil, 0)
	logging.TagInfo(t.SshInf.host,fmt.Sprintf("stdout:%s,stderr:%s", stdout, stderr))
	if t.checkErr(err) {
		return false
	}
	return true
}
