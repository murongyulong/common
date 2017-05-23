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
	Port        int
	Mount       string
}

func (this *App) String() string {
	return fmt.Sprintf(
		"<Name:%s, Memory:%d, InstanceCnt:%d, Image:%s, Status:%d, Port:%d, Mount:%s>",
		this.Name,
		this.Memory,
		this.InstanceCnt,
		this.Image,
		this.Status,
		this.Port,
		this.Mount
	)
}