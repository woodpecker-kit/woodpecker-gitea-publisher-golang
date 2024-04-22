package constant

import (
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_steps_transfer"
)

const (
	// NameCliPluginStepsTransferFilePath
	//  @Description: Steps transfer file path
	//  @Usage: Steps transfer file path
	//  @Default: wd_steps_transfer.DefaultKitStepsFileName as `.woodpecker_kit.steps.transfer`
	//  @EnvKey: WOODPECKER_KIT_STEPS_TRANSFER_FILE_PATH
	NameCliPluginStepsTransferFilePath = "settings.woodpecker-kit-steps-transfer-file-path"

	// EnvKeyPluginStepsTransferFilePath
	//  @Description: Steps transfer file path
	//  @Usage: Steps transfer file path
	//  @Default: wd_steps_transfer.DefaultKitStepsFileName as `.woodpecker_kit.steps.transfer`
	EnvKeyPluginStepsTransferFilePath = "PLUGIN_WOODPECKER_KIT_STEPS_TRANSFER_FILE_PATH"

	// NameCliPluginStepsTransferDisableOut
	//  @Description: Steps transfer write disable out
	//  @Usage: Steps transfer write
	NameCliPluginStepsTransferDisableOut = "settings.woodpecker-kit-steps-transfer-disable-out"

	// EnvKeyPluginStepsTransferDisableOut
	//  @Description: Steps transfer write disable out
	EnvKeyPluginStepsTransferDisableOut = "PLUGIN_WOODPECKER_KIT_STEPS_TRANSFER_DISABLE_OUT"
)

// CommonFlag
// Other modules also have flags
func CommonFlag() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    wd_flag.NameCliPluginDebug,
			Usage:   "Provides the debug flag. This value is true when the is open debug mode",
			EnvVars: []string{wd_flag.EnvKeyPluginDebug},
		},
	}
}

func HideCommonGlobalFlag() []cli.Flag {
	return []cli.Flag{
		&cli.UintFlag{
			Name:    wd_flag.NameCliPluginTimeoutSecond,
			Usage:   "command timeout setting second",
			Hidden:  true,
			Value:   10,
			EnvVars: []string{wd_flag.EnvKeyPluginTimeoutSecond},
		},

		&cli.StringFlag{
			Name:    NameCliPluginStepsTransferFilePath,
			Usage:   "Steps transfer file path",
			Hidden:  true,
			Value:   wd_steps_transfer.DefaultKitStepsFileName,
			EnvVars: []string{EnvKeyPluginStepsTransferFilePath},
		},

		&cli.BoolFlag{
			Name:    NameCliPluginStepsTransferDisableOut,
			Usage:   "Steps transfer write disable out",
			Hidden:  true,
			EnvVars: []string{EnvKeyPluginStepsTransferDisableOut},
		},
	}
}
