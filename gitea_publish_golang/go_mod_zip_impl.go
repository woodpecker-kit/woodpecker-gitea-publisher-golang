package gitea_publish_golang

import (
	"fmt"
	"github.com/sinlov-go/go-common-lib/pkg/filepath_plus"
	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
	"os"
	"path/filepath"
)

func (z *GoModZip) CreateGoModeZipPackageFile(targetPath string, version string) error {
	if z.modFile == nil {
		return fmt.Errorf("check error at CreateGoModeZipPackageFile modFile is nil")
	}
	if version == "" {
		return fmt.Errorf("check error at CreateGoModeZipPackageFile version is empty")
	}

	if !filepath_plus.PathExistsFast(targetPath) {
		errMakeZipRootPath := filepath_plus.Mkdir(targetPath)
		if errMakeZipRootPath != nil {
			return fmt.Errorf("check error at CreateGoModeZipPackageFile make zip root path: %s", errMakeZipRootPath)
		}
	}

	writable, errDirWritable := isDirectoryWritable(targetPath)
	if errDirWritable != nil {
		return fmt.Errorf("check error at CreateGoModeZipPackageFile check targetPath %s writable: %v", targetPath, errDirWritable)
	}
	if !writable {
		return fmt.Errorf("check error at CreateGoModeZipPackageFile targetPath not writable: %s", targetPath)

	}

	outPath := filepath.Join(targetPath, fmt.Sprintf("%s.zip", version))
	outFile, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("check error at CreateGoModeZipPackageFile create zip file: %s", err)
	}

	modVersion := module.Version{
		Path:    z.modFile.Module.Mod.Path,
		Version: version,
	}

	err = zip.CreateFromDir(outFile, modVersion, z.modRootPath)
	if err != nil {
		return fmt.Errorf("CreateGoModeZipPackageFile zip.CreateFromDir err: %s", err)
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

func (z *GoModZip) GetGoModeGoVersion() string {
	return z.modFile.Go.Version
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

// isDirectoryWritable checks if the given directory path is writable on the current platform.
func isDirectoryWritable(dirPath string) (bool, error) {
	info, err := os.Stat(dirPath)
	if err != nil {
		return false, fmt.Errorf("failed to get directory info: %w", err)
	}

	if !info.IsDir() {
		return false, fmt.Errorf("%s is not a directory", dirPath)
	}

	// Attempt to open the directory for writing. This operation should fail if the directory is not writable.
	// Use the O_RDONLY flag with O_CREATE to ensure that we're only checking permissions and not actually creating a file.
	file, err := os.OpenFile(filepath.Join(dirPath, ".write_check"), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		if os.IsPermission(err) {
			return false, nil // Directory is not writable due to permission issues
		}
		return false, fmt.Errorf("failed to check directory writability: %w", err)
	}
	defer file.Close() // Close the temporary file after the check

	// If we reach here, the directory is writable, so remove the temporary file
	err = os.Remove(file.Name())
	if err != nil {
		return false, fmt.Errorf("failed to remove temporary file: %w", err)
	}

	return true, nil
}
