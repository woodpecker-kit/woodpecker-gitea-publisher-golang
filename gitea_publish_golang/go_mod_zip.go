package gitea_publish_golang

import "golang.org/x/mod/modfile"

type GoModZip struct {
	modRootPath string

	modFile     *modfile.File
	versionName string

	zipPackageFilePath string

	GoModZipFunc GoModZipFunc `json:"-"`
}

type GoModZipFunc interface {
	GetModFile() *modfile.File

	GetVersionName() string

	GetGoModPackageName() string

	GetGoModeGoVersion() string

	CreateGoModeZipPackageFile(targetPath string, version string) error

	GetZipPackageFilePath() (string, error)
}

func NewGoModZip(modRootPath string) (*GoModZip, error) {
	modZip := GoModZip{
		modRootPath: modRootPath,
	}
	errFetch := modZip.fetchGoModFile()
	if errFetch != nil {
		return nil, errFetch
	}

	return &modZip, nil
}
