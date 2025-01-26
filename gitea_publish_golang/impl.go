package gitea_publish_golang

import (
	"fmt"
	"github.com/sinlov-go/go-common-lib/pkg/string_tools"
	"github.com/sinlov-go/go-common-lib/pkg/struct_kit"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"os"
	"path/filepath"
)

func (p *GiteaPublishGolang) ShortInfo() wd_short_info.WoodpeckerInfoShort {
	if p.wdShortInfo == nil {
		info2Short := wd_short_info.ParseWoodpeckerInfo2Short(*p.woodpeckerInfo)
		p.wdShortInfo = &info2Short
	}
	return *p.wdShortInfo
}

// SetWoodpeckerInfo
// also change ShortInfo() return
func (p *GiteaPublishGolang) SetWoodpeckerInfo(info wd_info.WoodpeckerInfo) {
	var newInfo wd_info.WoodpeckerInfo
	_ = struct_kit.DeepCopyByGob(&info, &newInfo)
	p.woodpeckerInfo = &newInfo
	info2Short := wd_short_info.ParseWoodpeckerInfo2Short(newInfo)
	p.wdShortInfo = &info2Short
}

func (p *GiteaPublishGolang) GetWoodPeckerInfo() wd_info.WoodpeckerInfo {
	return *p.woodpeckerInfo
}

func (p *GiteaPublishGolang) OnlyArgsCheck() {
	p.onlyArgsCheck = true
}

func (p *GiteaPublishGolang) Exec() error {
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

func (p *GiteaPublishGolang) loadStepsTransfer() error {
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

func (p *GiteaPublishGolang) checkArgs() error {

	//errCheck := argCheckInArr("build status", p.wdShortInfo.Build.Status, pluginBuildStateSupport)
	//if errCheck != nil {
	//	return errCheck
	//}

	errCheckGiteaReleaseSupport := argCheckInArr(CliNameGiteaReleaseExistsDo, p.Settings.GiteaReleaseExistDo, giteaReleaseExistDoSupport)
	if errCheckGiteaReleaseSupport != nil {
		return errCheckGiteaReleaseSupport
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

	version := p.ShortInfo().Build.Tag
	if version == "" {
		version = "latest"
	}

	p.Settings.PublishPackageVersion = version

	// append path
	p.Settings.findOutGoModPath = filepath.Join(p.Settings.RootPath, p.Settings.PublishPackageGoPath)
	p.Settings.resultRootFullPath = filepath.Join(p.Settings.RootPath, p.Settings.ResultUploadRootPath)

	tempDir := os.TempDir()
	shortInfo := p.ShortInfo()
	zipTempDir := filepath.Join(tempDir, "woodpecker-gitea-publisher-golang", shortInfo.Repo.Hostname, shortInfo.Repo.OwnerName, shortInfo.Repo.ShortName, shortInfo.Build.Number)
	wd_log.Debugf("zip target root path: %s", zipTempDir)
	p.Settings.ZipTargetRootPath = zipTempDir

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
func (p *GiteaPublishGolang) doBiz() error {

	err := p.publishByClient()
	if err != nil {
		return err
	}

	return nil
}

func (p *GiteaPublishGolang) saveStepsTransfer() error {
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
