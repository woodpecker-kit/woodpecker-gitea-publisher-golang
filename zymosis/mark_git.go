package zymosis

import (
	"embed"
	"path"
)

const (
	resMarkFolder   = "go_zymosis_record"
	resMarkFileName = "git_rev_parse"
	defaultGitHash  = "0000000"
)

var (
	//go:embed go_zymosis_record
	embedResMark     embed.FS
	markGitHeadShort = ""
)

func MainProgramRes() string {
	if markGitHeadShort == "" {
		resMarkFileContent, err := embedResMark.ReadFile(path.Join(resMarkFolder, resMarkFileName))
		if err == nil {
			markGitHeadShort = string(resMarkFileContent)
		} else {
			markGitHeadShort = defaultGitHash
		}
	}

	return markGitHeadShort
}
