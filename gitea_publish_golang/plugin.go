package gitea_publish_golang

import (
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
)

type (
	// GiteaPublishGolang gitea_publish_golang all config
	GiteaPublishGolang struct {
		Name           string
		Version        string
		woodpeckerInfo *wd_info.WoodpeckerInfo
		wdShortInfo    *wd_short_info.WoodpeckerInfoShort
		onlyArgsCheck  bool

		Settings Settings

		FuncPlugin FuncGiteaPublishGolang `json:"-"`
	}
)

type FuncGiteaPublishGolang interface {
	ShortInfo() wd_short_info.WoodpeckerInfoShort

	SetWoodpeckerInfo(info wd_info.WoodpeckerInfo)
	GetWoodPeckerInfo() wd_info.WoodpeckerInfo

	OnlyArgsCheck()

	Exec() error

	loadStepsTransfer() error
	checkArgs() error
	saveStepsTransfer() error
}
