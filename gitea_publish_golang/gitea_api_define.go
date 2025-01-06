package gitea_publish_golang

import "code.gitea.io/sdk/gitea"

type GiteaPackageInfo struct {
	Id      uint64 `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Version string `json:"version"`

	HtmlUrl    string           `json:"html_url"`
	Creator    gitea.User       `json:"creator"`
	Owner      gitea.User       `json:"owner"`
	Repository gitea.Repository `json:"repository"`
}
