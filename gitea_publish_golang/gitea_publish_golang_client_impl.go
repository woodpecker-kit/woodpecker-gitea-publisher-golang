package gitea_publish_golang

import (
	"fmt"
	"github.com/sinlov-go/go-common-lib/pkg/filepath_plus"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"golang.org/x/mod/module"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	ErrPackageNotExist = fmt.Errorf("RemotePackageGoFetch not exist, code 404")
)

func (p *publishGolangClient) LocalPackageGoFetch(goModRootPath string) error {
	goModZip, errNewGoModZip := NewGoModZip(goModRootPath)
	if errNewGoModZip != nil {
		return fmt.Errorf("can not load go.mod by path: %s NewGoModZip err: %v", goModRootPath, errNewGoModZip)
	}
	p.goModZip = goModZip
	return nil
}

func (p *publishGolangClient) RemotePackageGoFetch(version string) (*GiteaPackageInfo, error) {
	if p.goModZip == nil {
		return nil, fmt.Errorf("RemotePackageGoFetch go.mod not loaded by LocalPackageGoFetch")
	}
	pkgName := p.goModZip.GetGoModPackageName()
	errEsCapePkgName := escapeValidatePathSegments(&pkgName)
	if errEsCapePkgName != nil {
		return nil, fmt.Errorf("RemotePackageGoFetch escapeValidatePathSegments error: %s", errEsCapePkgName)
	}
	pkgVersion := version
	errEsCapePkgNamePackageVersion := escapeValidatePathSegments(&pkgVersion)
	if errEsCapePkgNamePackageVersion != nil {
		return nil, fmt.Errorf("RemotePackageGoFetch escapeValidatePathSegments error: %s", errEsCapePkgNamePackageVersion)
	}

	apiPath := fmt.Sprintf("/api/v1/packages/%s/%s/%s/%s", p.owner, "go", pkgName, pkgVersion)

	wd_log.Debugf("try RemotePackageGoFetch apiPath: %s", apiPath)

	var giteaPackage GiteaPackageInfo
	resp, err := p.ApiGiteaGet(apiPath, nil, &giteaPackage)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return nil, ErrPackageNotExist
		}
		return nil, fmt.Errorf("check package type [ %s ] [ %s ] err: %v", "go", apiPath, err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("check package type [ %s ] [ %s ] errcode: %v", "go", apiPath, resp.StatusCode)
	}
	return &giteaPackage, nil
}

func (p *publishGolangClient) DeletePackageGoFetch(version string) error {
	if p.goModZip == nil {
		return fmt.Errorf("DeletePackageGoFetch go.mod not loaded by LocalPackageGoFetch")
	}
	pkgName := p.goModZip.GetGoModPackageName()
	errEsCapePkgName := escapeValidatePathSegments(&pkgName)
	if errEsCapePkgName != nil {
		return fmt.Errorf("DeletePackageGoFetch escapeValidatePathSegments error: %s", errEsCapePkgName)
	}
	pkgVersion := version
	errEsCapePkgNamePackageVersion := escapeValidatePathSegments(&pkgVersion)
	if errEsCapePkgNamePackageVersion != nil {
		return fmt.Errorf("DeletePackageGoFetch escapeValidatePathSegments error: %s", errEsCapePkgNamePackageVersion)
	}
	apiPath := fmt.Sprintf("/api/v1/packages/%s/%s/%s/%s", p.owner, "go", pkgName, pkgVersion)

	wd_log.Debugf("try DeletePackageGoFetch apiPath: %s", apiPath)
	resp, errApi := p.ApiGiteaDelete(apiPath, nil, nil)
	if errApi != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return ErrPackageNotExist
		}
		return fmt.Errorf("DeletePackageGoFetch apiPath: %s, err: %v", apiPath, errApi)
	}
	if resp.StatusCode != http.StatusNoContent {
		return nil
	}
	return nil
}

func (p *publishGolangClient) CreateGoModZip(version string, zipRootPath string, goModRootPath string, removePath []string) error {
	if p.goModZip == nil {
		return fmt.Errorf("CreateGoModZip go.mod not loaded by LocalPackageGoFetch")
	}
	if len(removePath) > 0 {
		wd_log.Debugf("try CreateGoModZip removePath: %v", removePath)
		for _, removePathItem := range removePath {
			removeFullPath := strings.Replace(removePathItem, "/", string(filepath.Separator), -1)
			removeFullPath = filepath.Join(goModRootPath, removeFullPath)
			errRemovePath := filepath_plus.RmDir(removeFullPath, true)
			if errRemovePath != nil {
				return fmt.Errorf("CreateGoModZip removePath: %s, error: %s", removeFullPath, errRemovePath)
			}
			wd_log.Infof("CreateGoModZip removePath success: %s", removeFullPath)
		}
	}

	if p.dryRun && version == "latest" {
		wd_log.Infof("CreateGoModZip dryRun mode")
		wd_log.Infof("CreateGoModZip version: %s\n", version)
		wd_log.Infof("CreateGoModZip zipRootPath: %s\n", zipRootPath)
		wd_log.Infof("CreateGoModZip goModRootPath: %s\n", goModRootPath)
		p.zipUploadPath = filepath.Join(zipRootPath, fmt.Sprintf("%s.zip", version))
		return nil
	}

	errCreateZip := p.goModZip.CreateGoModeZipPackageFile(zipRootPath, version)
	if errCreateZip != nil {
		return fmt.Errorf("CreateGoModZip CreateGoModeZipPackageFile err: %v", errCreateZip)
	}
	zipFilePath, errZipFilePath := p.goModZip.GetZipPackageFilePath()
	if errZipFilePath != nil {
		return fmt.Errorf("CreateGoModZip GetZipPackageFilePath err: %v", errZipFilePath)
	}
	p.zipUploadPath = zipFilePath
	return nil
}

func (p *publishGolangClient) PackageGoUpload() (*PublishPackageGoInfo, error) {
	if p.zipUploadPath == "" {
		return nil, fmt.Errorf("PackageGoUpload upload file not loaded by CreateGoModZip")
	}
	pkgName := p.goModZip.GetGoModPackageName()
	pkgVersion := p.tag
	goVersion := p.goModZip.GetGoModeGoVersion()
	res := &PublishPackageGoInfo{
		ModVersion: module.Version{
			Path:    pkgName,
			Version: pkgVersion,
		},
	}
	wd_log.Debugf("try PackageGoUpload outZipPath: %s", p.zipUploadPath)
	uploadPath := fmt.Sprintf("/api/packages/%s/go/upload", p.owner)
	if p.dryRun {
		wd_log.Infof("PackageGoUpload dryRun")
		wd_log.Infof("PackageGoUpload upload file Path: %s\n", p.zipUploadPath)
		wd_log.Infof("PackageGoUpload upload url Path: %s\n", uploadPath)
		wd_log.Infof("PackageGoUpload pkgName: %s\n", pkgName)
		wd_log.Infof("PackageGoUpload pkgVersion: %s\n", pkgVersion)
		wd_log.Infof("PackageGoUpload goVersion: %s\n", goVersion)
		return res, nil
	}

	fileBodyIO, errOpen := os.Open(p.zipUploadPath)
	if errOpen != nil {
		return res, fmt.Errorf("open zip file %s , error: %s", p.zipUploadPath, errOpen)
	}
	defer func(fileBodyIO *os.File) {
		errFileBodyIO := fileBodyIO.Close()
		if errOpen != nil {
			log.Fatalf("try ResourcesPostFile file IO Close err: %v", errFileBodyIO)
		}
	}(fileBodyIO)

	statusCode, errPutGoPackage := p.ApiGiteaStatusCode(http.MethodPut, uploadPath, nil, fileBodyIO)
	if errPutGoPackage != nil {
		return res, fmt.Errorf("PackageGoUpload upload file %s , error: %s", p.zipUploadPath, errPutGoPackage)
	}
	if statusCode != http.StatusCreated {
		return res, fmt.Errorf("do put go package [ %s ] errcode: %v, zip_path: %s", uploadPath, statusCode, p.zipUploadPath)
	}
	res.Version = pkgVersion
	res.PackageName = pkgName
	res.GoModGoVersion = goVersion
	res.UploadTimeUnix = time.Now().Unix()
	res.HostName = p.targetHostName

	pkgNameUrl := pkgName
	pkgVersionUrl := pkgVersion
	_ = escapeValidatePathSegments(&pkgNameUrl, &pkgVersionUrl)

	res.PackagePageUrl = fmt.Sprintf("%s/%s/-/packages/%s/%s/%s", p.GetBaseUrl(), p.owner, "go", pkgNameUrl, pkgVersionUrl)
	return res, nil
}

// escapeValidatePathSegments is a help function to validate and encode url path segments
//
//nolint:golint,unused
func escapeValidatePathSegments(seg ...*string) error {
	for i := range seg {
		if seg[i] == nil || len(*seg[i]) == 0 {
			return fmt.Errorf("path segment [%d] is empty", i)
		}
		*seg[i] = url.PathEscape(*seg[i])
	}
	return nil
}
