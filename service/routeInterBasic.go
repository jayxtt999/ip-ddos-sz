package service

import "github.com/jayxtt999/ip-ddos-sz/pkg/setting"

type RouteBasic struct {
	config *setting.Os
	RouteInterFace
}

func (rb RouteBasic)setConfig(config *setting.Os)  {
	rb.config = config
}

