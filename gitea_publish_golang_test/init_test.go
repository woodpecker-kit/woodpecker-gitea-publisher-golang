package gitea_publish_golang_test

import (
	"fmt"
	"github.com/sinlov-go/unittest-kit/env_kit"
	"github.com/sinlov-go/unittest-kit/unittest_file_kit"
	"github.com/woodpecker-kit/woodpecker-gitea-publisher-golang/gitea_publish_golang"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_steps_transfer"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

const (
	keyEnvDebug  = "CI_DEBUG"
	keyEnvCiNum  = "CI_NUMBER"
	keyEnvCiKey  = "CI_KEY"
	keyEnvCiKeys = "CI_KEYS"

	envKeyProjectRootPath = "CI_PROJECT_ROOT_PATH"

	mockVersion = "v1.0.0"
	mockName    = "woodpecker-gitea-publisher-golang"
)

var (
	// testBaseFolderPath
	//  test base dir will auto get by package init()
	testBaseFolderPath = ""
	testGoldenKit      *unittest_file_kit.TestGoldenKit

	// mustSetInCiEnvList
	//  for check set in CI env not empty
	mustSetInCiEnvList = []string{
		gitea_publish_golang.EnvGiteaPublishGolangApiKey,
	}

	// mustSetArgsAsEnvList
	mustSetArgsAsEnvList = []string{
		gitea_publish_golang.EnvGiteaPublishGolangApiKey,
	}

	valEnvTimeoutSecond             uint
	valEnvPluginDebug               = false
	valEnvGiteaPublishGolangApiKey  = ""
	valEnvGiteaPubGolangBaseUrl     = ""
	valEnvGiteaPubGolangInsecure    = false
	valEnvGiteaPubGolangDryRun      = true
	valEnvGiteaPubGolangPathGo      = ""
	valGiteaReleaseExistDo          = ""
	valEnvGiteaPubGolangRemovePaths = []string{
		"dist",
	}
	valEnvGiteaPubGolangLatest = true

	// CI Test Env

	varProjectRootPath = "" // change by env:CI_PROJECT_ROOT_PATH
	valCiRepoName      = ""
	valCiRepoOwner     = ""
	valCiSystemHost    = ""
	valCiSystemUrl     = ""
	valCiForgeType     = "gitea"
	valCiForgeUrl      = ""
)

func init() {
	testBaseFolderPath, _ = getCurrentFolderPath()
	wd_log.SetLogLineDeep(2)
	// if open wd_template please open this
	//wd_template.RegisterSettings(wd_template.DefaultHelpers)

	testGoldenKit = unittest_file_kit.NewTestGoldenKit(testBaseFolderPath)

	varProjectRootPath = env_kit.FetchOsEnvStr(envKeyProjectRootPath, "")
	valCiRepoName = env_kit.FetchOsEnvStr(wd_flag.EnvKeyRepositoryCiName, "")
	valCiRepoOwner = env_kit.FetchOsEnvStr(wd_flag.EnvKeyRepositoryCiOwner, "")
	valCiSystemHost = env_kit.FetchOsEnvStr(wd_flag.EnvKeyCiSystemHost, "")
	valCiSystemUrl = env_kit.FetchOsEnvStr(wd_flag.EnvKeyCiSystemUrl, "")
	valCiForgeType = env_kit.FetchOsEnvStr(wd_flag.EnvKeyCiForgeType, "gitea")
	valCiForgeUrl = env_kit.FetchOsEnvStr(wd_flag.EnvKeyCiForgeUrl, "")

	valEnvTimeoutSecond = uint(env_kit.FetchOsEnvInt(wd_flag.EnvKeyPluginTimeoutSecond, 10))
	valEnvPluginDebug = env_kit.FetchOsEnvBool(wd_flag.EnvKeyPluginDebug, false)

	valEnvGiteaPublishGolangApiKey = env_kit.FetchOsEnvStr(gitea_publish_golang.EnvGiteaPublishGolangApiKey, "")
	valEnvGiteaPubGolangBaseUrl = env_kit.FetchOsEnvStr(gitea_publish_golang.EnvGiteaPubGolangBaseUrl, "")
	valEnvGiteaPubGolangInsecure = env_kit.FetchOsEnvBool(gitea_publish_golang.EnvGiteaPubGolangInsecure, false)
	valEnvGiteaPubGolangDryRun = env_kit.FetchOsEnvBool(gitea_publish_golang.EnvGiteaPubGolangDryRun, true)
	valEnvGiteaPubGolangPathGo = env_kit.FetchOsEnvStr(gitea_publish_golang.EnvGiteaPubGolangPathGo, "")
	valGiteaReleaseExistDo = env_kit.FetchOsEnvStr(gitea_publish_golang.EnvGiteaReleaseExistsDo, gitea_publish_golang.GiteaReleaseExistsDoFail)
	removePathsEnv := env_kit.FetchOsEnvStringSlice(gitea_publish_golang.EnvGiteaPubGolangRemovePaths)
	if len(removePathsEnv) > 0 {
		valEnvGiteaPubGolangRemovePaths = removePathsEnv
	}
}

// test case basic tools start
// getCurrentFolderPath
//
//	can get run path this golang dir
func getCurrentFolderPath() (string, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("can not get current file info")
	}
	return filepath.Dir(file), nil
}

// test case basic tools end

func envCheck(t *testing.T) bool {

	if valEnvPluginDebug {
		wd_log.OpenDebug()
	}

	// most CI system will set env CI to true
	envCI := env_kit.FetchOsEnvStr("CI", "")
	if envCI == "" {
		t.Logf("not in CI system, skip envCheck")
		return false
	}
	t.Logf("check env for CI system")
	return env_kit.MustHasEnvSetByArray(t, mustSetInCiEnvList)
}

func envMustArgsCheck(t *testing.T) bool {
	for _, item := range mustSetArgsAsEnvList {
		if os.Getenv(item) == "" {
			t.Logf("plasee set env: %s, than run test\nfull need set env %v", item, mustSetArgsAsEnvList)
			return true
		}
	}
	return false
}

func generateTransferStepsOut(plugin gitea_publish_golang.GiteaPublishGolang, mark string, data interface{}) error {
	_, err := wd_steps_transfer.Out(plugin.Settings.RootPath, plugin.Settings.StepsTransferPath, plugin.GetWoodPeckerInfo(), mark, data)
	return err
}

func mockPluginSettings() gitea_publish_golang.Settings {
	// all mock settings can set here
	settings := gitea_publish_golang.Settings{
		// use env:PLUGIN_DEBUG
		Debug:             valEnvPluginDebug,
		TimeoutSecond:     valEnvTimeoutSecond,
		RootPath:          testGoldenKit.GetTestDataFolderFullPath(),
		StepsTransferPath: wd_steps_transfer.DefaultKitStepsFileName,
	}

	settings.DryRun = valEnvGiteaPubGolangDryRun
	settings.GiteaApiKey = valEnvGiteaPublishGolangApiKey
	settings.GiteaBaseUrl = valEnvGiteaPubGolangBaseUrl
	settings.GiteaInsecure = valEnvGiteaPubGolangInsecure
	settings.PublishPackageGoPath = valEnvGiteaPubGolangPathGo
	settings.PublishRemovePaths = valEnvGiteaPubGolangRemovePaths
	settings.GiteaReleaseExistDo = valGiteaReleaseExistDo

	return settings

}

func mockPluginWithSettings(t *testing.T, woodpeckerInfo wd_info.WoodpeckerInfo, settings gitea_publish_golang.Settings) gitea_publish_golang.GiteaPublishGolang {
	p := gitea_publish_golang.GiteaPublishGolang{
		Name:    mockName,
		Version: mockVersion,
	}

	// mock woodpecker info
	//t.Log("mockPluginWithStatus")

	p.SetWoodpeckerInfo(woodpeckerInfo)

	if p.ShortInfo().Build.WorkSpace != "" {
		settings.RootPath = p.ShortInfo().Build.WorkSpace
	}

	p.Settings = settings
	return p
}
