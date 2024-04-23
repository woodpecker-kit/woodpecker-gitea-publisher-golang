[![ci](https://github.com/woodpecker-kit/woodpecker-gitea-publisher-golang/workflows/ci/badge.svg)](https://github.com/woodpecker-kit/woodpecker-gitea-publisher-golang/actions/workflows/ci.yml)

[![go mod version](https://img.shields.io/github/go-mod/go-version/woodpecker-kit/woodpecker-gitea-publisher-golang?label=go.mod)](https://github.com/woodpecker-kit/woodpecker-gitea-publisher-golang)
[![GoDoc](https://godoc.org/github.com/woodpecker-kit/woodpecker-gitea-publisher-golang?status.png)](https://godoc.org/github.com/woodpecker-kit/woodpecker-gitea-publisher-golang)
[![goreportcard](https://goreportcard.com/badge/github.com/woodpecker-kit/woodpecker-gitea-publisher-golang)](https://goreportcard.com/report/github.com/woodpecker-kit/woodpecker-gitea-publisher-golang)

[![GitHub license](https://img.shields.io/github/license/woodpecker-kit/woodpecker-gitea-publisher-golang)](https://github.com/woodpecker-kit/woodpecker-gitea-publisher-golang)
[![codecov](https://codecov.io/gh/woodpecker-kit/woodpecker-gitea-publisher-golang/branch/main/graph/badge.svg)](https://codecov.io/gh/woodpecker-kit/woodpecker-gitea-publisher-golang)
[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/woodpecker-kit/woodpecker-gitea-publisher-golang)](https://github.com/woodpecker-kit/woodpecker-gitea-publisher-golang/tags)
[![GitHub release)](https://img.shields.io/github/v/release/woodpecker-kit/woodpecker-gitea-publisher-golang)](https://github.com/woodpecker-kit/woodpecker-gitea-publisher-golang/releases)

## for what

- this project used to woodpecker plugin

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/woodpecker-kit/woodpecker-gitea-publisher-golang)](https://github.com/woodpecker-kit/woodpecker-gitea-publisher-golang/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息

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
- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

## usage

- use this template, replace list below and add usage
    - `github.com/woodpecker-kit/woodpecker-gitea-publisher-golang` to your package name
    - `woodpecker-kit` to your owner name
    - `woodpecker-gitea-publisher-golang` to your project name

- use github action for this workflow push to docker hub, must add
    - variables `ENV_DOCKERHUB_OWNER` user of docker hub
    - variables `ENV_DOCKERHUB_REPO_NAME` repo name of docker hub
    - secrets `DOCKERHUB_TOKEN` token of docker hub user from [hub.docker](https://hub.docker.com/settings/security)

- check `docker-bake.hcl` config, change to your docker image

- if you use `wd_steps_transfer` just add `.woodpecker_kit.steps.transfer` at git ignore
- change code start with `// change or remove`

### workflow usage

- see [doc](doc/docs.md)

## Notice

- want dev this project, see [dev doc](doc/README.md)