package cli

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-gitea-publisher-golang/constant"
	"github.com/woodpecker-kit/woodpecker-gitea-publisher-golang/gitea_publish_golang"
	"github.com/woodpecker-kit/woodpecker-gitea-publisher-golang/internal/pkg_kit"
	"github.com/woodpecker-kit/woodpecker-gitea-publisher-golang/zymosis"
	"github.com/woodpecker-kit/woodpecker-tools/wd_urfave_cli_v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_urfave_cli_v2/cli_exit_urfave"
	"runtime"
	"time"
)

const (
	copyrightStartYear = "2024"
	defaultExitCode    = 1
)

func NewCliApp() *cli.App {
	cli_exit_urfave.ChangeDefaultExitCode(defaultExitCode)

	namePlugin := pkg_kit.GetPackageJsonName()
	versionPlugin := pkg_kit.GetPackageJsonVersionGoStyle(false)
	jsonAuthor := pkg_kit.GetPackageJsonAuthor()
	year := time.Now().Year()

	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Version = versionPlugin
	app.Name = namePlugin
	if pkg_kit.GetPackageJsonHomepage() != "" {
		app.Usage = fmt.Sprintf("see: %s", pkg_kit.GetPackageJsonHomepage())
	}
	app.Description = pkg_kit.GetPackageJsonDescription()

	var pkgBundlerInfo string
	pkgBundlerResourceCode := zymosis.MainProgramRes()
	if pkgBundlerResourceCode == "0000000" {
		pkgBundlerInfo = fmt.Sprintf("by: %s, run on %s %s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	} else {
		pkgBundlerInfo = fmt.Sprintf("by: %s, run on %s %s res: %s", runtime.Version(), runtime.GOOS, runtime.GOARCH, pkgBundlerResourceCode)
	}
	app.Copyright = fmt.Sprintf("© %s-%d %s %s",
		copyrightStartYear, year, jsonAuthor.Name, pkgBundlerInfo)

	author := &cli.Author{
		Name:  jsonAuthor.Name,
		Email: jsonAuthor.Email,
	}
	app.Authors = []*cli.Author{
		author,
	}

	flags := wd_urfave_cli_v2.UrfaveCliAppendCliFlags(
		wd_urfave_cli_v2.WoodpeckerUrfaveCliFlags(),
		constant.CommonFlag(),
		constant.HideCommonGlobalFlag(),
		gitea_publish_golang.GlobalFlag(),
		gitea_publish_golang.HideGlobalFlag(),
	)

	app.Flags = flags
	app.Before = GlobalBeforeAction
	app.Action = GlobalAction
	app.After = GlobalAfterAction

	return app
}
