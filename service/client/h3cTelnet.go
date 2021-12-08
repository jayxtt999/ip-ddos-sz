package client

import (
	"fmt"
	"github.com/jayxtt999/ip-ddos-sz/pkg/logging"
	"github.com/jayxtt999/ip-ddos-sz/pkg/telnet"
)


type H3cTelnet struct {
	TelnetInf
}
//
//  Login
//  @Description: 登陆
//  @receiver t
//  @return bool
//
func (t *H3cTelnet) Login() bool {

	addr := fmt.Sprintf("%s:%d", t.host, t.port)
	logging.TagInfo(t.TelnetInf.host,fmt.Sprintf("telnet login:%s:%d,user:%s", t.host, t.port, t.user))
	c, err := telnet.Dial("tcp", addr)
	if t.checkErr(err) {
		return false
	}
	c.SetUnixWriteMode(true)
	var r bool
	r = t.expect(c, "name:")
	if !r {
		return false
	}
	r = t.sendLn(c, t.user)
	if !r {
		return false
	}
	t.expect(c, "word:")
	if !r {
		return false
	}
	r = t.sendLn(c, t.pass)
	if !r {
		return false
	}
	t.expect(c, ">")
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
func (t *H3cTelnet) Exec(shell string, delim byte) (err error) {
	logging.TagInfo(t.TelnetInf.host,shell)
	c := t.TelnetInf.telnetClient
	var r bool
	r = t.sendLn(c, shell)
	if !r {
		return err
	}
	data, err := c.ReadBytes(delim)
	logging.TagInfo(t.TelnetInf.host,string(data))
	if t.checkErr(err){
		return err
	}
	return nil
}
//
//  Logout
//  @Description: 退出
//  @receiver t
//  @return bool
//
func (t *H3cTelnet) Logout() bool {
	c := t.TelnetInf.telnetClient
	r := t.sendLn(c, "quit")
	if !r {
		return false
	}
	return true
}

