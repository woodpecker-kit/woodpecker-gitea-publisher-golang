package gitea_publish_golang

import (
	"fmt"
	"github.com/sinlov-go/go-common-lib/pkg/string_tools"
	"github.com/sinlov-go/go-common-lib/pkg/struct_kit"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
)

func (p *Plugin) ShortInfo() wd_short_info.WoodpeckerInfoShort {
	if p.wdShortInfo == nil {
		info2Short := wd_short_info.ParseWoodpeckerInfo2Short(*p.woodpeckerInfo)
		p.wdShortInfo = &info2Short
	}
	return *p.wdShortInfo
}

// SetWoodpeckerInfo
// also change ShortInfo() return
func (p *Plugin) SetWoodpeckerInfo(info wd_info.WoodpeckerInfo) {
	var newInfo wd_info.WoodpeckerInfo
	_ = struct_kit.DeepCopyByGob(&info, &newInfo)
	p.woodpeckerInfo = &newInfo
	info2Short := wd_short_info.ParseWoodpeckerInfo2Short(newInfo)
	p.wdShortInfo = &info2Short
}

func (p *Plugin) GetWoodPeckerInfo() wd_info.WoodpeckerInfo {
	return *p.woodpeckerInfo
}

func (p *Plugin) OnlyArgsCheck() {
	p.onlyArgsCheck = true
}

func (p *Plugin) Exec() error {
	errLoadStepsTransfer := p.loadStepsTransfer()
	if errLoadStepsTransfer != nil {
		return errLoadStepsTransfer
	}

	errCheckArgs := p.checkArgs()
	if errCheckArgs != nil {
		return fmt.Errorf("check args err: %v", errCheckArgs)
	}

	if p.onlyArgsCheck {
		wd_log.Info("only check args, skip do doBiz")
		return nil
	}

	err := p.doBiz()
	if err != nil {
		return err
	}
	errSaveStepsTransfer := p.saveStepsTransfer()
	if errSaveStepsTransfer != nil {
		return errSaveStepsTransfer
	}

	return nil
}

func (p *Plugin) loadStepsTransfer() error {
	// change or remove or this code start
	//if p.Settings.StepsTransferDemo {
	//	var readConfigData Settings
	//	errLoad := wd_steps_transfer.In(p.Settings.RootPath, p.Settings.StepsTransferPath, *p.woodpeckerInfo, StepsTransferMarkDemoConfig, &readConfigData)
	//	if errLoad != nil {
	//		return nil
	//	}
	//	wd_log.VerboseJsonf(readConfigData, "load steps transfer config mark [ %s ]", StepsTransferMarkDemoConfig)
	//}
	// change or remove or this code end
	return nil
}

func (p *Plugin) checkArgs() error {

	errCheck := argCheckInArr("build status", p.woodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineStatus, pluginBuildStateSupport)
	if errCheck != nil {
		return errCheck
	}

	if p.Settings.GiteaBaseUrl == "" {
		if p.woodpeckerInfo.CiForgeInfo.CiForgeType == "gitea" {
			wd_log.Debugf("when CiForgeType [ gitea ] woodpeckerInfo.CiForgeInfo.CiForgeUrl [ %s ] as GiteaBaseUrl", p.woodpeckerInfo.CiForgeInfo.CiForgeUrl)
			p.Settings.GiteaBaseUrl = p.woodpeckerInfo.CiForgeInfo.CiForgeUrl
		}

	}
	if p.Settings.GiteaBaseUrl == "" {
		return fmt.Errorf("check args [ %s ] set, now is empty, or can not get from CiForgeType [ gitea ] by env:CI_FORGE_URL", EnvGiteaPubGolangBaseUrl)
	}

	if p.Settings.GiteaApiKey == "" {
		return fmt.Errorf("check args [ %s ] must set, now is empty", CliNameGiteaPublishGolangApiKey)
	}

	return nil
}

func argCheckInArr(mark string, target string, checkArr []string) error {
	if !(string_tools.StringInArr(target, checkArr)) {
		return fmt.Errorf("not support %s now [ %s ], must in %v", mark, target, checkArr)
	}
	return nil
}

// doBiz
//
//	replace this code with your gitea_publish_golang implementation
func (p *Plugin) doBiz() error {

	if p.Settings.DryRun {
		wd_log.Verbosef("dry run, skip some biz code, more info can open debug by flag [ %s ]", wd_flag.EnvKeyPluginDebug)
		return nil
	}

	return nil
}

func (p *Plugin) saveStepsTransfer() error {
	// change or remove this code

	if p.Settings.StepsOutDisable {
		wd_log.Debugf("steps out disable by flag [ %v ], skip save steps transfer", p.Settings.StepsOutDisable)
		return nil
	}

	// change or remove or this code start
	//if p.Settings.StepsTransferDemo {
	//	transferAppendObj, errSave := wd_steps_transfer.Out(p.Settings.RootPath, p.Settings.StepsTransferPath, *p.woodpeckerInfo, StepsTransferMarkDemoConfig, p.Settings)
	//	if errSave != nil {
	//		return errSave
	//	}
	//	wd_log.VerboseJsonf(transferAppendObj, "save steps transfer config mark [ %s ]", StepsTransferMarkDemoConfig)
	//}
	// change or remove or this code end
	return nil
}
