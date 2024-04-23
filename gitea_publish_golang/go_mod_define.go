package gitea_publish_golang

import (
	"golang.org/x/mod/module"
)

type PublishPackageGoInfo struct {
	ModVersion     module.Version `json:"-"`
	Version        string
	PackageName    string
	GoModGoVersion string

	HostName       string
	PackagePageUrl string
	UploadTimeUnix int64
}
