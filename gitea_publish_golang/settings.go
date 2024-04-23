package gitea_publish_golang

import "github.com/woodpecker-kit/woodpecker-tools/wd_info"

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

		PublishPackageGoPath string
		// findOutGoModPath is the path to find go.mod file will change by check args
		findOutGoModPath     string
		PublishRemovePaths   []string
		ResultUploadRootPath string
		ResultUploadFileName string

		// PublishPackageVersion is the version to publish this by check args success will init by tag or latest
		PublishPackageVersion string

		// ZipTargetRootPath    is the root path of zip target with check args success will init by
		// tempDir/woodpecker-gitea-publisher-golang/{repo-hostname}/{owner}/{repo}/{build_number}
		ZipTargetRootPath string
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
