package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type Log struct {
	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var LogSetting = &Log{}

type Server struct {
	RunMode         string
	HttpPort        int
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	RuntimeRootPath string
	CheckAllowIp    bool
	AllowIpList     string
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

var OsSetting = &Os{}

type Os struct {
	Host    string
	Port    int
	Timeout int
	User    string
	Pass    string
	Prompt  string
}

var Cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}
	mapTo("server", ServerSetting)
	mapTo("log", LogSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second

}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := Cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
