package service

import "hotgo/internal/service/ipmaker"

// IPMaker add ons.
type IPMaker interface {
	Upload() ipmaker.IUpload
	User() ipmaker.IAppUser
	Ip() ipmaker.IIp
	Story() ipmaker.IStory
}

var localIPMaker IPMaker

func RegisterIPMaker(i IPMaker) {
	localIPMaker = i
}

func IpMaker() IPMaker {
	if localIPMaker == nil {
		panic("implement not found for interface IPMaker, forgot register?")
	}
	return localIPMaker
}
