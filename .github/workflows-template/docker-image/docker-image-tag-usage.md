## need `New repository secret`

- file `docker-image-tag.yml`
- variables `ENV_DOCKERHUB_OWNER` for docker hub user
- variables `ENV_DOCKERHUB_REPO_NAME` for docker hub repo name
- secrets `DOCKERHUB_TOKEN` from [hub.docker](https://hub.docker.com/settings/security)
    - if close push remote can pass `DOCKERHUB_TOKEN` setting

## usage at github action

```yml
jobs:
  version:
    name: version
    uses: ./.github/workflows/version.yml

  docker-image-tag:
    name: docker-image-tag
    needs:
      - version
      - golang-ci
      - go-build-check-main
    uses: ./.github/workflows/docker-image-tag.yml
    if: startsWith(github.ref, 'refs/tags/')
    with:
      push_remote_flag: true # default is true
    secrets:
      DOCKERHUB_TOKEN: "${{ secrets.DOCKERHUB_TOKEN }}"
```

- `push_remote_flag` default is `false`