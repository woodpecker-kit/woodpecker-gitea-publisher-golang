---
name: woodpecker-gitea-publisher-golang
description: woodpecker gitea_publish_golang template
author: woodpecker-kit
tags: [ environment, woodpecker-gitea-publisher-golang ]
containerImage: woodpecker-kit/woodpecker-gitea-publisher-golang
containerImageUrl: https://hub.docker.com/r/woodpecker-kit/woodpecker-gitea-publisher-golang
url: https://github.com/woodpecker-kit/woodpecker-gitea-publisher-golang
icon: https://raw.githubusercontent.com/woodpecker-kit/woodpecker-gitea-publisher-golang/main/doc/logo.svg
---

woodpecker-gitea-publisher-golang

plugin as https://woodpecker-ci.org/ for https://docs.gitea.com/usage/packages/go/ to publisher golang package

## Features

- [x] push to [gitea Go Package Registry](https://docs.gitea.com/usage/packages/go/)
- [x] use TempDir to generate go mod zip file
- [x] support dry-run mode
- [x] auto get gitea base URL by `CI_FORGE_URL`
- [x] support [Monorepo](https://en.wikipedia.org/wiki/Monorepo)
    - use `settings.gitea-publish-golang-path-go` to publish dir go.mod one by one
- [x] out publish info json file default `dist/go-mod-upload.json`
    - use `settings.gitea-publish-golang-update-result-root-path` to change out root path, default `dist`
    - use `settings.gitea-publish-golang-update-result-file-name` to change out file name, default `go-mod-upload.json`

## Settings

| Name                                           | Required | Default value        | Description                                                                                                 |
|------------------------------------------------|----------|----------------------|-------------------------------------------------------------------------------------------------------------|
| `debug`                                        | **no**   | *false*              | open debug log or open by env `PLUGIN_DEBUG`                                                                |
| `gitea-publish-golang-api-key`                 | **yes**  |                      | gitea api key, Required                                                                                     |
| `gitea-publish-golang-base-url`                | **no**   |                      | gitea base url, default by `CI_FORGE_URL`                                                                   |
| `gitea-publish-golang-insecure`                | **no**   | *false*              | gitea insecure enable                                                                                       |
| `gitea-publish-golang-dry-run`                 | **no**   | *false*              | dry run mode                                                                                                |
| `gitea-publish-golang-path-go`                 | **no**   |                      | publish go package is dir to find go.mod, will append project root path, default is this project root path  |
| `gitea-publish-golang-remove-paths`            | **no**   | `["dist"]`           | publish go package remove paths, this path under `gitea-publish-golang-path-go`, default will remove `dist` |
| `gitea-publish-golang-update-result-root-path` | **no**   | *dist*               | out result root path append CI Workspace, default `dist`                                                    |
| `gitea-publish-golang-update-result-file-name` | **no**   | *go-mod-upload.json* | out file name, default `go-mod-upload.json`                                                                 |

**Hide Settings:**

| Name                                        | Required | Default value                    | Description                                                                      |
|---------------------------------------------|----------|----------------------------------|----------------------------------------------------------------------------------|
| `timeout_second`                            | **no**   | *10*                             | command timeout setting by second                                                |
| `gitea-publish-golang-timeout-second`       | **no**   | *60*                             | gitea release api timeout second, default 60, less 30                            |
| `woodpecker-kit-steps-transfer-file-path`   | **no**   | `.woodpecker_kit.steps.transfer` | Steps transfer file path, default by `wd_steps_transfer.DefaultKitStepsFileName` |
| `woodpecker-kit-steps-transfer-disable-out` | **no**   | *false*                          | Steps transfer write disable out                                                 |

## Example

- workflow with backend `docker`

[![docker hub version semver](https://img.shields.io/docker/v/sinlov/woodpecker-gitea-publisher-golang?sort=semver)](https://hub.docker.com/r/sinlov/woodpecker-gitea-publisher-golang/tags?page=1&ordering=last_updated)
[![docker hub image size](https://img.shields.io/docker/image-size/sinlov/woodpecker-gitea-publisher-golang)](https://hub.docker.com/r/sinlov/woodpecker-gitea-publisher-golang)
[![docker hub image pulls](https://img.shields.io/docker/pulls/sinlov/woodpecker-gitea-publisher-golang)](https://hub.docker.com/r/sinlov/woodpecker-gitea-publisher-golang/tags?page=1&ordering=last_updated)

```yml
labels:
  backend: docker
steps:
  woodpecker-gitea-publisher-golang:
    image: sinlov/woodpecker-gitea-publisher-golang:latest
    pull: false
    settings:
      # debug: true
      # gitea-publish-golang-dry-run: true # dry run mode
      gitea-publish-golang-api-key: # gitea api key, Required
        from_secret: gitea_api_key_release
      gitea-publish-golang-path-go: "" # publish go package is dir to find go.mod, will append project root path, default is this project root path
      gitea-publish-golang-remove-paths: # publish go package remove paths, this path under `gitea-publish-golang-path-go`, default will remove `dist`
        - "dist"
      # gitea-publish-golang-update-result-root-path: "dist" # out result root path append CI Workspace, default `dist`
      # gitea-publish-golang-update-result-file-name: "go-mod-upload.json" # out file name, default `go-mod-upload.json`
```

- workflow with backend `local`, must install at local and effective at evn `PATH`
    - can download by [github release](https://github.com/woodpecker-kit/woodpecker-gitea-publisher-golang/releases)
- install at ${GOPATH}/bin, latest

```bash
go install -a github.com/woodpecker-kit/woodpecker-gitea-publisher-golang/cmd/woodpecker-gitea-publisher-golang@latest
```

[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/woodpecker-kit/woodpecker-gitea-publisher-golang)](https://github.com/woodpecker-kit/woodpecker-gitea-publisher-golang/tags)
[![GitHub release)](https://img.shields.io/github/v/release/woodpecker-kit/woodpecker-gitea-publisher-golang)](https://github.com/woodpecker-kit/woodpecker-gitea-publisher-golang/releases)

- install at ${GOPATH}/bin, v1.0.0

```bash
go install -v github.com/woodpecker-kit/woodpecker-gitea-publisher-golang/cmd/woodpecker-gitea-publisher-golang@v1.0.0
```

```yml
labels:
  backend: local
steps:
  woodpecker-gitea-publisher-golang:
    image: woodpecker-gitea-publisher-golang
    settings:
      # debug: true
      # gitea-publish-golang-dry-run: true # dry run mode
      gitea-publish-golang-api-key: # gitea api key, Required
        from_secret: gitea_api_key_release
      gitea-publish-golang-path-go: "" # publish go package is dir to find go.mod, will append project root path, default is this project root path
      gitea-publish-golang-remove-paths: # publish go package remove paths, this path under `gitea-publish-golang-path-go`, default will remove `dist`
        - "dist"
      # gitea-publish-golang-update-result-root-path: "dist" # out result root path append CI Workspace, default `dist`
      # gitea-publish-golang-update-result-file-name: "go-mod-upload.json" # out file name, default `go-mod-upload.json`
```

- full config

```yaml
labels:
  backend: docker
steps:
  woodpecker-gitea-publisher-golang:
    image: sinlov/woodpecker-gitea-publisher-golang:latest
    pull: false
    settings:
      debug: true
      gitea-publish-golang-dry-run: true # dry run mode
      gitea-publish-golang-timeout-second: 120 # gitea release api timeout second, default 60, less 30 
      gitea-publish-golang-api-key: # gitea api key, Required
        from_secret: gitea_api_key_release
      gitea-publish-golang-base-url: "https://gitea.example.com" # default by CI_FORGE_URL auto to find
      gitea-publish-golang-insecure: true #  gitea insecure enable
      gitea-publish-golang-path-go: "sub-go" # publish go package is dir to find go.mod, will append project root path, default is this project root path
      gitea-publish-golang-remove-paths: # publish go package remove paths, this path under `gitea-publish-golang-path-go`, default will remove `dist`
        - "dist"
      gitea-publish-golang-update-result-root-path: "dist" # out result root path append CI Workspace, default `dist`
      gitea-publish-golang-update-result-file-name: "go-mod-upload.json" # out file name, default `go-mod-upload.json`
```

## Notes

- go mod zip file will generate by TempDir and try to remove after publish
- template dir path graph `tempDir/woodpecker-gitea-publisher-golang/{repo-hostname}/{owner}/{repo}/{build_number}`

## Known limitations

1. go mod zip file generate by `zip.CreateFromDir`, this method will check as [https://semver.org/](https://semver.org/), so this kit use `CI_COMMIT_TAG` to generate version
2. open `gitea-publish-golang-dry-run` mode can pass all check, but not publish to gitea
3. but without event tag will use version mark as `latest`, this will not pass `zip.CreateFromDir`check, so will not publish to gitea
