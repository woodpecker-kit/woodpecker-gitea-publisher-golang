package gitea_publish_golang

import (
	"fmt"
	"github.com/sinlov-go/gitea-client-wrapper/gitea_token_client"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"net/url"
)

type publishGolangClient struct {
	gitea_token_client.GiteaTokenClient

	dryRun bool
	owner  string
	repo   string
	tag    string
	// tagTarget
	//is the branch or commit sha to tag
	tagTarget string

	zipTargetRootPath string

	// goModZip by go.mod fetch
	goModZip *GoModZip
	// zipUploadPath wait to upload zip path
	zipUploadPath  string
	targetHostName string
}

type GiteaPublishGolangClient interface {
	LocalPackageGoFetch(goModRootPath string) error

	RemotePackageGoFetch(version string) (*GiteaPackageInfo, error)

	CreateGoModZip(version string, zipRootPath string, goModRootPath string, removePath []string) error

	PackageGoUpload() (*PublishPackageGoInfo, error)
}

var (
	ErrMissingTag = fmt.Errorf("NewGiteaPublishGolangClientByWoodpeckerShort missing tag, please check now in tag build")
)

// NewGiteaPublishGolangClientByWoodpeckerShort
// create gitea publish golang client by wd_short_info.WoodpeckerInfoShort
func NewGiteaPublishGolangClientByWoodpeckerShort(info wd_short_info.WoodpeckerInfoShort, config Settings) (GiteaPublishGolangClient, error) {
	if info.Build.Event != wd_info.EventPipelineTag && !config.DryRun {
		return nil, ErrMissingTag
	}

	parse, errUrlParse := url.Parse(config.GiteaBaseUrl)
	if errUrlParse != nil {
		return nil, errUrlParse
	}
	hostName := parse.Hostname()

	pc := &publishGolangClient{
		dryRun:            config.DryRun,
		owner:             info.Repo.OwnerName,
		repo:              info.Repo.ShortName,
		tag:               info.Build.Tag,
		tagTarget:         info.Build.CommitBranch,
		zipTargetRootPath: config.ZipTargetRootPath,
		targetHostName:    hostName,
	}

	errNewClient := pc.NewClientWithHttpTimeout(config.GiteaBaseUrl, config.GiteaApiKey, config.GiteaTimeoutSecond, config.GiteaInsecure)
	if errNewClient != nil {
		return nil, errNewClient
	}
	wd_log.Debug("gitea client created success")
	return pc, nil
}
