package gitea_publish_golang

import (
	"fmt"
	"golang.org/x/mod/modfile"
	"golang.org/x/mod/zip"
	"os"
	"path/filepath"
)

func (z *GoModZip) CreateGoModeZipPackageFile(targetPath string) error {
	if z.modFile == nil {
		return fmt.Errorf("check error at CreateGoModeZipPackageFile modFile is nil")
	}
	if z.versionName == "" {
		return fmt.Errorf("check error at CreateGoModeZipPackageFile versionName is empty")
	}

	outPath := filepath.Join(targetPath, fmt.Sprintf("%s.zip", z.versionName))
	outFile, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("check error at CreateGoModeZipPackageFile create zip file: %s", err)
	}

	err = zip.CreateFromDir(outFile, z.modFile.Module.Mod, z.modRootPath)
	if err != nil {
		return fmt.Errorf("check error at CreateGoModeZipPackageFile CreateFromDir err: %s", err)
	}

	z.zipPackageFilePath = outPath

	return nil
}

func (z *GoModZip) GetZipPackageFilePath() (string, error) {
	if z.zipPackageFilePath == "" {
		return "", fmt.Errorf("please run CreateGoModeZipPackageFile first")
	}
	return z.zipPackageFilePath, nil
}

func (z *GoModZip) GetModFile() *modfile.File {
	return z.modFile
}

func (z *GoModZip) GetVersionName() string {
	return z.versionName
}

func (z *GoModZip) GetGoModPackageName() string {
	return z.modFile.Module.Mod.Path
}

func (z *GoModZip) fetchGoModFile() error {
	if !modfile.IsDirectoryPath(z.modRootPath) {
		return fmt.Errorf("check error at fetchGoModFile not is go mod root path: %s", z.modRootPath)
	}
	goModPath := filepath.Join(z.modRootPath, "go.mod")

	goModData, errReadFile := os.ReadFile(goModPath)
	if errReadFile != nil {
		return fmt.Errorf("check error at fetchGoModFile read go.mod: %v", errReadFile)
	}
	goModFile, errParse := modfile.Parse(goModPath, goModData, nil)
	if errParse != nil {
		return fmt.Errorf("check error at CreateGoModZipFromDir parse go.mod: %s", errParse)
	}
	if goModFile == nil {
		return fmt.Errorf("check error at fetchGoModFile parse go.mod is nil")
	}

	z.modFile = goModFile
	z.versionName = goModFile.Module.Mod.Version

	return nil
}
