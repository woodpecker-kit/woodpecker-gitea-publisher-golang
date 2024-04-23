package gitea_publish_golang

import (
	"errors"
	"fmt"
	"github.com/sinlov-go/go-common-lib/pkg/filepath_plus"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"path/filepath"
)

func (p *GiteaPublishGolang) publishByClient() error {
	pc, errNewReleaseClient := NewGiteaPublishGolangClientByWoodpeckerShort(p.ShortInfo(), p.Settings)
	if errNewReleaseClient != nil {
		return errNewReleaseClient

	}
	errLocalPackageGoFetch := pc.LocalPackageGoFetch(p.Settings.PublishPackageGoPath)
	if errLocalPackageGoFetch != nil {
		return errLocalPackageGoFetch
	}

	version := p.ShortInfo().Build.Tag
	if p.Settings.DryRun {
		version = "latest"
	}

	_, errRemotePackageGoFetch := pc.RemotePackageGoFetch(version)
	if errRemotePackageGoFetch != nil {
		if !errors.Is(errRemotePackageGoFetch, ErrPackageNotExist) {
			return fmt.Errorf(" RemotePackageGoFetch error: %s", errRemotePackageGoFetch)
		}
	}

	errCreateGoModZip := pc.CreateGoModZip(p.Settings.ZipTargetRootPath, p.Settings.PublishPackageGoPath, p.Settings.PublishRemovePaths)
	if errCreateGoModZip != nil {
		return errCreateGoModZip
	}
	wd_log.Infof("create go mod zip to: %s", p.Settings.ZipTargetRootPath)

	packageGoUpload, errPackageGoUpload := pc.PackageGoUpload()
	if errPackageGoUpload != nil {
		return errPackageGoUpload
	}
	if p.Settings.DryRun {
		return nil
	}

	wd_log.DebugJsonf(packageGoUpload, "packageGoUpload res")

	saveUploadFilePath := filepath.Join(p.Settings.ResultUploadRootPath, p.Settings.ResultUploadFileName)
	errSaveResult := filepath_plus.WriteFileAsJsonBeauty(saveUploadFilePath, packageGoUpload, false)
	if errSaveResult != nil {
		return errSaveResult
	}
	wd_log.Infof("save upload result to: %s", saveUploadFilePath)

	return nil
}
