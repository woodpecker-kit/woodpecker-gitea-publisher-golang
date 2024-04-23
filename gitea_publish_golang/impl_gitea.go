package gitea_publish_golang

import (
	"errors"
	"fmt"
	"github.com/sinlov-go/go-common-lib/pkg/filepath_plus"
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

	packageGoUpload, errPackageGoUpload := pc.PackageGoUpload()
	if errPackageGoUpload != nil {
		return errPackageGoUpload
	}
	saveUploadFilePath := filepath.Join(p.Settings.ResultUploadRootPath, p.Settings.ResultUploadFileName)
	errSaveResult := filepath_plus.WriteFileAsJsonBeauty(saveUploadFilePath, packageGoUpload, false)
	if errSaveResult != nil {
		return errSaveResult
	}

	return nil
}
