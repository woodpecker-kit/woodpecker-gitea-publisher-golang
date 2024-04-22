package gitea_publish_golang

import "github.com/woodpecker-kit/woodpecker-tools/wd_info"

const (
	// change or remove settings config const start

	// StepsTransferMarkDemoConfig
	// steps transfer key
	StepsTransferMarkDemoConfig = "demo_config"

	// change or remove settings config const end
)

type (
	// Settings gitea_publish_golang private config
	Settings struct {
		Debug             bool
		TimeoutSecond     uint
		StepsTransferPath string
		StepsOutDisable   bool
		RootPath          string

		DryRun             bool
		GiteaApiKey        string
		GiteaBaseUrl       string
		GiteaInsecure      bool
		GiteaTimeoutSecond uint

		PublishPackageGo   string
		PublishRemovePaths []string
	}
)

var (
	// pluginBuildStateSupport
	pluginBuildStateSupport = []string{
		wd_info.BuildStatusCreated,
		wd_info.BuildStatusRunning,
		wd_info.BuildStatusSuccess,
		wd_info.BuildStatusFailure,
		wd_info.BuildStatusError,
		wd_info.BuildStatusKilled,
	}
)
