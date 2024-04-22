package plugin

import "github.com/woodpecker-kit/woodpecker-tools/wd_info"

const (
	// change or remove settings config const start

	// StepsTransferMarkDemoConfig
	// steps transfer key
	StepsTransferMarkDemoConfig = "demo_config"

	// change or remove settings config const end
)

type (
	// Settings plugin private config
	Settings struct {
		Debug             bool
		TimeoutSecond     uint
		StepsTransferPath string
		StepsOutDisable   bool
		RootPath          string

		// change or remove this config demo start
		NotEmptyEnvKeys   []string
		EnvPrintKeys      []string
		PaddingLeftMax    int
		StepsTransferDemo bool
		// change or remove this config demo end

		DryRun bool
	}
)

var (

	// change or remove settings config check args start

	// pluginBuildStateSupport
	pluginBuildStateSupport = []string{
		wd_info.BuildStatusCreated,
		wd_info.BuildStatusRunning,
		wd_info.BuildStatusSuccess,
		wd_info.BuildStatusFailure,
		wd_info.BuildStatusError,
		wd_info.BuildStatusKilled,
	}

	// change or remove settings config check args end
)
