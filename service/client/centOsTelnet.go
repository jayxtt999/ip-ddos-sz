package client

import (
	"fmt"
	"github.com/jayxtt999/ip-ddos-sz/pkg/logging"
	"github.com/jayxtt999/ip-ddos-sz/pkg/telnet"
	"time"
)

const timeout = 10 * time.Second

type CentOsTelnet struct {
	TelnetInf
}
//
//  Login
//  @Description: 登陆
//  @receiver t
//  @return bool
//
func (t *CentOsTelnet) Login() bool {

	addr := fmt.Sprintf("%s:%d", t.TelnetInf.host, t.TelnetInf.port)
	logging.TagInfo(t.TelnetInf.host,fmt.Sprintf("telnet login:%s:%d,user:%s", t.TelnetInf.host, t.TelnetInf.port, t.TelnetInf.user))

	c, err := telnet.Dial("tcp", addr)
	if t.checkErr(err) {
		return false
	}
	c.SetUnixWriteMode(true)
	var r bool
	r = t.expect(c, "login: ")
	if !r {
		return false
	}
	r = t.sendLn(c, t.TelnetInf.user)
	if !r {
		return false
	}
	t.expect(c, "ssword: ")
	if !r {
		return false
	}
	r = t.sendLn(c, t.TelnetInf.pass)
	if !r {
		return false
	}
	t.expect(c, "~]#")
	if !r {
		return false
	}
	t.TelnetInf.telnetClient = c
	return true

}
//
//  Exec
//  @Description: 执行shell
//  @receiver t
//  @param shell
//  @param delim
//  @return err
//
func (t *CentOsTelnet) Exec(shell string, delim byte) ( err error) {
	logging.TagInfo(t.TelnetInf.host,shell)
	c := t.TelnetInf.telnetClient
	var r bool
	r = t.sendLn(c, shell)
	if !r {
		return  err
	}
	data, err := c.ReadBytes(delim)
	logging.TagInfo(t.TelnetInf.host,data)
	if t.checkErr(err) {
		return  err
	}

	return nil
}
//
//  Logout
//  @Description: 退出
//  @receiver t
//  @return bool
//
func (t *CentOsTelnet) Logout() bool {
	c := t.TelnetInf.telnetClient
	r := t.sendLn(c, "exit")
	if !r {
		return false
	}
	return true
}
