package client

import (
	"fmt"
	"github.com/jayxtt999/ip-ddos-sz/pkg/logging"
	"github.com/jayxtt999/ip-ddos-sz/pkg/telnet"
	"runtime"
	"time"
)
//
//  TelnetInf
//  @Description: telnet 接口
//
type TelnetInf struct {
	host      string
	port      int
	user      string
	pass      string
	telnetClient *telnet.Conn
}

//
//  Config
//  @Description: telnet接口配置
//  @receiver inf
//  @param host
//  @param port
//  @param user
//  @param pass
//
func (inf *TelnetInf) Config(host string, port int, user string, pass string) {
	inf.host = host
	inf.port = port
	inf.user = user
	inf.pass = pass
}


//
//  checkErr
//  @Description: 验证错误
//  @param err
//  @return bool
//

func (inf *TelnetInf)checkErr(err error) bool {
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
//
//  expect
//  @Description:预期返回
//  @receiver inf
//  @param t
//  @param d
//  @return bool
//
func (inf *TelnetInf)expect(t *telnet.Conn, d ...string) bool {
	if inf.checkErr(t.SetReadDeadline(time.Now().Add(timeout))) {
		return false
	}
	if inf.checkErr(t.SkipUntil(d...)) {
		return false
	}
	return true
}

//
//  sendLn
//  @Description: 发送命令
//  @receiver inf
//  @param t
//  @param s
//  @return bool
//
func (inf *TelnetInf)sendLn(t *telnet.Conn, s string) bool {
	if inf.checkErr(t.SetWriteDeadline(time.Now().Add(timeout))) {
		return false
	}
	buf := make([]byte, len(s)+1)
	copy(buf, s)
	buf[len(s)] = '\n'
	_, err := t.Write(buf)
	if inf.checkErr(err) {
		return false
	}
	return true
}

