package plugin

import (
	"fmt"
	"github.com/sinlov-go/go-common-lib/pkg/string_tools"
	"github.com/sinlov-go/go-common-lib/pkg/struct_kit"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_steps_transfer"
	"os"
	"strconv"
	"strings"
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
	if p.Settings.StepsTransferDemo {
		var readConfigData Settings
		errLoad := wd_steps_transfer.In(p.Settings.RootPath, p.Settings.StepsTransferPath, *p.woodpeckerInfo, StepsTransferMarkDemoConfig, &readConfigData)
		if errLoad != nil {
			return nil
		}
		wd_log.VerboseJsonf(readConfigData, "load steps transfer config mark [ %s ]", StepsTransferMarkDemoConfig)
	}
	// change or remove or this code end
	return nil
}

func (p *Plugin) checkArgs() error {

	errCheck := argCheckInArr("build status", p.woodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineStatus, pluginBuildStateSupport)
	if errCheck != nil {
		return errCheck
	}

	return nil
}

func argCheckInArr(mark string, target string, checkArr []string) error {
	if !(string_tools.StringInArr(target, checkArr)) {
		return fmt.Errorf("not support %s now [ %s ], must in %v", mark, target, checkArr)
	}
	return nil
}

func checkEnvNotEmpty(keys []string) error {
	for _, env := range keys {
		if os.Getenv(env) == "" {
			return fmt.Errorf("check env [ %s ] must set, now is empty", env)
		}
	}
	return nil
}

// doBiz
//
//	replace this code with your plugin implementation
func (p *Plugin) doBiz() error {

	if p.Settings.DryRun {
		wd_log.Verbosef("dry run, skip some biz code, more info can open debug by flag [ %s ]", wd_flag.EnvKeyPluginDebug)
		return nil
	}

	// change or remove or this code start
	printBasicEnv(p)
	if len(p.Settings.NotEmptyEnvKeys) > 0 {
		errCheck := checkEnvNotEmpty(p.Settings.NotEmptyEnvKeys)
		if errCheck != nil {
			return errCheck
		}
	}
	// change or remove or this code end
	return nil
}

func (p *Plugin) saveStepsTransfer() error {
	// change or remove this code

	if p.Settings.StepsOutDisable {
		wd_log.Debugf("steps out disable by flag [ %v ], skip save steps transfer", p.Settings.StepsOutDisable)
		return nil
	}

	// change or remove or this code start
	if p.Settings.StepsTransferDemo {
		transferAppendObj, errSave := wd_steps_transfer.Out(p.Settings.RootPath, p.Settings.StepsTransferPath, *p.woodpeckerInfo, StepsTransferMarkDemoConfig, p.Settings)
		if errSave != nil {
			return errSave
		}
		wd_log.VerboseJsonf(transferAppendObj, "save steps transfer config mark [ %s ]", StepsTransferMarkDemoConfig)
	}
	// change or remove or this code end
	return nil
}

// change or remove or method start

func printBasicEnv(p *Plugin) {
	var sb strings.Builder
	_, _ = fmt.Fprint(&sb, "-> just print basic env:\n")
	paddingMax := strconv.Itoa(p.Settings.PaddingLeftMax)

	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCiSystemUrl, p.woodpeckerInfo.CiSystemInfo.CiSystemUrl)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCiSystemHost, p.woodpeckerInfo.CiSystemInfo.CiSystemHost)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCiMachine, p.woodpeckerInfo.CiSystemInfo.CiMachine)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCiSystemPlatform, p.woodpeckerInfo.CiSystemInfo.CiSystemPlatform)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCiSystemVersion, p.woodpeckerInfo.CiSystemInfo.CiSystemVersion)

	appendStrBuilderNewLine(&sb)

	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCiWorkflowName, p.woodpeckerInfo.CurrentInfo.CurrentWorkflowInfo.CiWorkflowName)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyWoodpeckerBackend, p.woodpeckerInfo.CiSystemInfo.WoodpeckerBackend)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCiMachine, p.woodpeckerInfo.CiSystemInfo.CiMachine)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCiSystemPlatform, p.woodpeckerInfo.CiSystemInfo.CiSystemPlatform)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyRepositoryCiName, p.woodpeckerInfo.RepositoryInfo.CIRepoName)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyRepositoryCiOwner, p.woodpeckerInfo.RepositoryInfo.CIRepoOwner)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCommitCiCommitBranch, p.woodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitBranch)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCommitCiCommitRef, p.woodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitRef)

	appendStrBuilderNewLine(&sb)

	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentPipelineNumber, p.woodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineNumber)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentPipelineEvent, p.woodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineEvent)

	switch p.woodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineEvent {
	case wd_info.EventPipelineTag:
		appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCommitCiCommitTag, p.woodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitTag)
	case wd_info.EventPipelinePullRequest:
		appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCommitCiCommitPullRequest, p.woodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitPullRequest)
		appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCommitCiCommitPullRequestLabels, p.woodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitPullRequestLabels)
		appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCommitCiCommitSourceBranch, p.woodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitSourceBranch)
		appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCommitCiCommitTargetBranch, p.woodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitTargetBranch)
	case wd_info.EventPipelinePullRequestClose:
		appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCommitCiCommitPullRequest, p.woodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitPullRequest)
		appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCommitCiCommitPullRequestLabels, p.woodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitPullRequestLabels)
		appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCommitCiCommitSourceBranch, p.woodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitSourceBranch)
		appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCommitCiCommitTargetBranch, p.woodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitTargetBranch)
	case wd_info.EventPipelineRelease:
		appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentCommitCiCommitPreRelease, strconv.FormatBool(p.woodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitPreRelease))
	}

	appendStrBuilderNewLine(&sb)

	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentPipelineUrl, p.woodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineUrl)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyCurrentPipelineForgeUrl, p.woodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineForgeUrl)

	appendStrBuilderNewLine(&sb)

	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyPreviousCiCommitBranch, p.woodpeckerInfo.PreviousInfo.PreviousCommitInfo.CiPreviousCommitBranch)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyPreviousCiCommitRef, p.woodpeckerInfo.PreviousInfo.PreviousCommitInfo.CiPreviousCommitRef)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyPreviousCiPipelineEvent, p.woodpeckerInfo.PreviousInfo.PreviousPipelineInfo.CiPreviousPipelineEvent)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyPreviousCiPipelineStatus, p.woodpeckerInfo.PreviousInfo.PreviousPipelineInfo.CiPreviousPipelineStatus)
	appendEnvStrBuilder(&sb, paddingMax, wd_flag.EnvKeyPreviousCiPipelineUrl, p.woodpeckerInfo.PreviousInfo.PreviousPipelineInfo.CiPreviousPipelineUrl)

	if len(p.Settings.EnvPrintKeys) > 0 {
		appendStrBuilderNewLine(&sb)
		_, _ = fmt.Fprint(&sb, "-> start print keys env:\n")
		for _, key := range p.Settings.EnvPrintKeys {
			appendEnvStrBuilder(&sb, paddingMax, key, os.Getenv(key))
		}
		_, _ = fmt.Fprint(&sb, "-> end print keys env\n")
	}

	wd_log.Verbosef("%s", sb.String())
}

func appendStrBuilderNewLine(sb *strings.Builder) {
	_, _ = fmt.Fprintf(sb, "\n")
}

func appendEnvStrBuilder(sb *strings.Builder, paddingMax string, key string, value string) {
	_, _ = fmt.Fprintf(sb, "%-"+paddingMax+"s %s\n", key, value)
}

// change or remove or method end
