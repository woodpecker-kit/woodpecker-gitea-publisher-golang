## need `New repository secret`

- file `docker-image-latest.yml`
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

  docker-image-latest:
    name: docker-image-latest
    needs:
      - version
      - golang-ci
      - go-build-check-main
    uses: ./.github/workflows/docker-image-latest.yml
    if: startsWith(github.ref, 'refs/tags/')
    with:
      # push_remote_flag: ${{ github.ref == 'refs/heads/main' }}
      push_remote_flag: ${{ github.event.pull_request.merged == true }}
    secrets:
      DOCKERHUB_TOKEN: "${{ secrets.DOCKERHUB_TOKEN }}"
```

- `push_remote_flag` default is `false`