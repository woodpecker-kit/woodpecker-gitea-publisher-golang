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
	errLocalPackageGoFetch := pc.LocalPackageGoFetch(p.Settings.findOutGoModPath)
	if errLocalPackageGoFetch != nil {
		return errLocalPackageGoFetch
	}

	_, errRemotePackageGoFetch := pc.RemotePackageGoFetch(p.Settings.PublishPackageVersion)
	if errRemotePackageGoFetch != nil {
		if !errors.Is(errRemotePackageGoFetch, ErrPackageNotExist) {
			return fmt.Errorf(" RemotePackageGoFetch error: %s", errRemotePackageGoFetch)
		}
	}

	errCreateGoModZip := pc.CreateGoModZip(
		p.Settings.PublishPackageVersion,
		p.Settings.ZipTargetRootPath,
		p.Settings.findOutGoModPath,
		p.Settings.PublishRemovePaths,
	)
	if errCreateGoModZip != nil {
		p.cleanZipTargetRootPath()
		return errCreateGoModZip
	}
	wd_log.Infof("create go mod zip to: %s", p.Settings.ZipTargetRootPath)

	packageGoUpload, errPackageGoUpload := pc.PackageGoUpload()
	if errPackageGoUpload != nil {
		p.cleanZipTargetRootPath()
		return errPackageGoUpload
	}
	if p.Settings.DryRun {
		p.cleanZipTargetRootPath()
		return nil
	}

	wd_log.DebugJsonf(packageGoUpload, "packageGoUpload res")

	saveUploadFilePath := filepath.Join(p.Settings.resultRootFullPath, p.Settings.ResultUploadFileName)
	errSaveResult := filepath_plus.WriteFileAsJsonBeauty(saveUploadFilePath, packageGoUpload, true)
	if errSaveResult != nil {
		p.cleanZipTargetRootPath()
		return errSaveResult
	}
	wd_log.Infof("save upload result to: %s", saveUploadFilePath)
	p.cleanZipTargetRootPath()
	return nil
}

func (p *GiteaPublishGolang) cleanZipTargetRootPath() {
	if p.Settings.ZipTargetRootPath == "" {
		return
	}
	if filepath_plus.PathExistsFast(p.Settings.ZipTargetRootPath) {
		errRemove := filepath_plus.RmDir(p.Settings.ZipTargetRootPath, true)
		if errRemove != nil {
			wd_log.Warnf("cleanZipTargetRootPath remove %s err: %v", p.Settings.ZipTargetRootPath, errRemove)
		}
		wd_log.Debugf("cleanZipTargetRootPath remove success at path: %s", p.Settings.ZipTargetRootPath)
	} else {
		wd_log.Debugf("cleanZipTargetRootPath not exists at path: %s", p.Settings.ZipTargetRootPath)
	}
}
