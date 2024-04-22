package plugin

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
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
			Usage:   fmt.Sprintf("publish go package remove paths, this path under %s, vars like dist,target/os", CliNameGiteaPubGolangRemovePaths),
			EnvVars: []string{EnvGiteaPubGolangRemovePaths},
		},
	}
}

const (
	CliNameGiteaPubGolangTimeoutSecond = "settings.gitea-publish-golang-timeout-second"
	EvnGiteaPubGolangTimeoutSecond     = "PLUGIN_GITEA_PUBLISH_GOLANG_TIMEOUT_SECOND"
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
	}
}

func BindCliFlags(c *cli.Context,
	debug bool,
	cliName, cliVersion string,
	wdInfo *wd_info.WoodpeckerInfo,
	rootPath,
	stepsTransferPath string, stepsOutDisable bool,
) (*Plugin, error) {

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

		PublishPackageGo:   c.String(CliNameGiteaPubGolangPathGo),
		PublishRemovePaths: c.StringSlice(CliNameGiteaPubGolangRemovePaths),
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

	p := Plugin{
		Name:           cliName,
		Version:        cliVersion,
		woodpeckerInfo: wdInfo,
		wdShortInfo:    &infoShort,
		Settings:       config,
	}

	return &p, nil
}
