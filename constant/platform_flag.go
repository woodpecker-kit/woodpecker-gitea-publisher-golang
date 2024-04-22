package constant

import "github.com/urfave/cli/v2"

type PlatformConfig struct {
	Windows bool
	Linux   bool
	MacOS   bool
}

// PlatformFlags
//
//	@Description: PlatformFlags for all subcommand
func PlatformFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:  "windows",
			Usage: "windows platform",
		},

		&cli.BoolFlag{
			Name:  "linux",
			Usage: "linux platform",
		},

		&cli.BoolFlag{
			Name:  "macos",
			Usage: "macos platform",
		},
	}
}

// BindPlatformConfig
//
//	@Description: BindPlatformConfig for all subcommand
func BindPlatformConfig(c *cli.Context) *PlatformConfig {
	return &PlatformConfig{
		Windows: c.Bool("windows"),
		Linux:   c.Bool("linux"),
		MacOS:   c.Bool("macos"),
	}
}
