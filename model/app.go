package model

import (
	"fmt"
)

const (
	AppStatus_Success             = 0
	AppStatus_CreateContainerFail = 1
	AppStatus_PullImageFail       = 2
	AppStatus_StartContainerFail  = 3
)

type App struct {
	Name        string
	Memory      int
	InstanceCnt int
	Image       string
	Status      int
	Port        string
	Mount       string
	CPUShares   int
	Id          string
}

func (this *App) String() string {
	return fmt.Sprintf(
		"<Name:%s, Memory:%d, InstanceCnt:%d, Image:%s, Status:%d, Port:%s, Mount:%s, CPUShares:%d,Id:%s>",
		this.Name,
		this.Memory,
		this.InstanceCnt,
		this.Image,
		this.Status,
		this.Port,
		this.Mount,
		this.CPUShares,
		this.Id,
	)
}
