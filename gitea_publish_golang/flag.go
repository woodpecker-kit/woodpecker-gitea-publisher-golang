package gitea_publish_golang

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"path/filepath"
)

const (
	CliNameGiteaPublishGolangApiKey = "settings.gitea-publish-golang-api-key"
	EnvGiteaPublishGolangApiKey     = "PLUGIN_GITEA_PUBLISH_GOLANG_API_KEY"

	CliNameGiteaPubGolangBaseUrl = "settings.gitea-publish-golang-base-url"
	EnvGiteaPubGolangBaseUrl     = "PLUGIN_GITEA_PUBLISH_GOLANG_BASE_URL"

	CliNameGiteaPubGolangInsecure = "settings.gitea-publish-golang-insecure"
	EnvGiteaPubGolangInsecure     = "PLUGIN_GITEA_PUBLISH_GOLANG_INSECURE"

	CliNameGiteaPubGolangDryRun = "settings.gitea-publish-golang-dry-run"
	EnvGiteaPubGolangDryRun     = "PLUGIN_GITEA_PUBLISH_GOLANG_DRY_RUN"

	CliNameGiteaPubGolangPathGo = "settings.gitea-publish-golang-path-go"
	EnvGiteaPubGolangPathGo     = "PLUGIN_GITEA_PUBLISH_GOLANG_PATH_GO"

	CliNameGiteaPubGolangRemovePaths = "settings.gitea-publish-golang-remove-paths"
	EnvGiteaPubGolangRemovePaths     = "PLUGIN_GITEA_PUBLISH_GOLANG_REMOVE_PATHS"

	CliNameGiteaPubGolangUpdateResultRootPath = "settings.gitea-publish-golang-update-result-root-path"
	EnvGiteaPubGolangUpdateResultRootPath     = "PLUGIN_GITEA_PUBLISH_GOLANG_UPDATE_RESULT_ROOT_PATH"

	CliNameGiteaPubGolangUpdateResultFileName = "settings.gitea-publish-golang-update-result-file-name"
	EnvGiteaPubGolangUpdateResultFileName     = "PLUGIN_GITEA_PUBLISH_GOLANG_UPDATE_RESULT_FILE_NAME"
)

// GlobalFlag
// Other modules also have flags
func GlobalFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    CliNameGiteaPublishGolangApiKey,
			Usage:   "gitea api key, Required",
			EnvVars: []string{EnvGiteaPublishGolangApiKey},
		},
		&cli.StringFlag{
			Name:    CliNameGiteaPubGolangBaseUrl,
			Usage:   fmt.Sprintf("gitea base url, when `%s` is `gitea`, and this flag is empty, will try get from `%s`", wd_flag.EnvKeyCiForgeType, wd_flag.EnvKeyCiForgeUrl),
			EnvVars: []string{EnvGiteaPubGolangBaseUrl},
		},
		&cli.BoolFlag{
			Name:    CliNameGiteaPubGolangInsecure,
			Usage:   "visit base-url via insecure https protocol",
			EnvVars: []string{EnvGiteaPubGolangInsecure},
		},
		&cli.BoolFlag{
			Name:    CliNameGiteaPubGolangDryRun,
			Usage:   "dry run mode, will not publish",
			EnvVars: []string{EnvGiteaPubGolangDryRun},
		},
		&cli.StringFlag{
			Name:    CliNameGiteaPubGolangPathGo,
			Usage:   "publish go package is dir to find go.mod, if not set will use git root path, gitea 1.20.1+ support",
			EnvVars: []string{EnvGiteaPubGolangPathGo},
		},
		&cli.StringSliceFlag{
			Name:    CliNameGiteaPubGolangRemovePaths,
			Usage:   "publish go package remove paths, this path under [ gitea-publish-golang-path-go ], vars like dist,target/os",
			Value:   cli.NewStringSlice("dist"),
			EnvVars: []string{EnvGiteaPubGolangRemovePaths},
		},
		&cli.StringFlag{
			Name:    CliNameGiteaPubGolangUpdateResultRootPath,
			Usage:   "update result root path, `this path must under workspace`, default is: dist",
			Value:   "dist",
			EnvVars: []string{EnvGiteaPubGolangUpdateResultRootPath},
		},
		&cli.StringFlag{
			Name:    CliNameGiteaPubGolangUpdateResultFileName,
			Usage:   "update result file name, default is: go-mod-upload.json",
			Value:   "go-mod-upload.json",
			EnvVars: []string{EnvGiteaPubGolangUpdateResultFileName},
		},
	}
}

const (
	CliNameGiteaPubGolangTimeoutSecond = "settings.gitea-publish-golang-timeout-second"
	EvnGiteaPubGolangTimeoutSecond     = "PLUGIN_GITEA_PUBLISH_GOLANG_TIMEOUT_SECOND"

	CliNameGiteaPubGolangZipRootPath = "settings.gitea-publish-golang-zip-root-path"
	EnvGiteaPubGolangZipRootPath     = "PLUGIN_GITEA_PUBLISH_GOLANG_ZIP_ROOT_PATH"
)

func HideGlobalFlag() []cli.Flag {
	return []cli.Flag{
		&cli.UintFlag{
			Name:    CliNameGiteaPubGolangTimeoutSecond,
			Usage:   "gitea release api timeout second, default 60, less 30",
			Value:   60,
			Hidden:  true,
			EnvVars: []string{EvnGiteaPubGolangTimeoutSecond},
		},
		&cli.StringFlag{
			Name:    CliNameGiteaPubGolangZipRootPath,
			Usage:   "gitea publish golang zip root path, default is parent for CI_WORKSPACE",
			Hidden:  true,
			EnvVars: []string{EnvGiteaPubGolangZipRootPath},
		},
	}
}

func BindCliFlags(c *cli.Context,
	debug bool,
	cliName, cliVersion string,
	wdInfo *wd_info.WoodpeckerInfo,
	rootPath,
	stepsTransferPath string, stepsOutDisable bool,
) (*GiteaPublishGolang, error) {

	zipRootPath := c.String(CliNameGiteaPubGolangZipRootPath)
	if zipRootPath == "" {
		zipRootPath = filepath.Dir(rootPath)
	}

	saveUploadResultRoot := filepath.Join(rootPath, c.String(CliNameGiteaPubGolangUpdateResultRootPath))
	config := Settings{
		Debug:             debug,
		TimeoutSecond:     c.Uint(wd_flag.NameCliPluginTimeoutSecond),
		StepsTransferPath: stepsTransferPath,
		StepsOutDisable:   stepsOutDisable,
		RootPath:          rootPath,

		DryRun: c.Bool(CliNameGiteaPubGolangDryRun),

		GiteaApiKey:        c.String(CliNameGiteaPublishGolangApiKey),
		GiteaBaseUrl:       c.String(CliNameGiteaPubGolangBaseUrl),
		GiteaInsecure:      c.Bool(CliNameGiteaPubGolangInsecure),
		GiteaTimeoutSecond: c.Uint(CliNameGiteaPubGolangTimeoutSecond),

		PublishPackageGoPath: c.String(CliNameGiteaPubGolangPathGo),
		ZipTargetRootPath:    zipRootPath,
		PublishRemovePaths:   c.StringSlice(CliNameGiteaPubGolangRemovePaths),

		ResultUploadRootPath: saveUploadResultRoot,
		ResultUploadFileName: c.String(CliNameGiteaPubGolangUpdateResultFileName),
	}

	// set default TimeoutSecond
	if config.TimeoutSecond == 0 {
		config.TimeoutSecond = 10
	}
	if config.GiteaTimeoutSecond < 30 {
		config.GiteaTimeoutSecond = 30
	}

	// change or remove start

	wd_log.Debugf("args %s: %v", wd_flag.NameCliPluginTimeoutSecond, config.TimeoutSecond)

	infoShort := wd_short_info.ParseWoodpeckerInfo2Short(*wdInfo)

	p := GiteaPublishGolang{
		Name:           cliName,
		Version:        cliVersion,
		woodpeckerInfo: wdInfo,
		wdShortInfo:    &infoShort,
		Settings:       config,
	}

	return &p, nil
}
