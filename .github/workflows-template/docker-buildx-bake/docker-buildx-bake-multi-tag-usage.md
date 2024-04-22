## need `New repository secret`

- file `docker-buildx-bake-multi-tag.yml`
- variables `ENV_DOCKERHUB_OWNER` for docker hub user
- variables `ENV_DOCKERHUB_REPO_NAME` for docker hub repo name
- secrets `DOCKERHUB_TOKEN` from [hub.docker](https://hub.docker.com/settings/security)
    - if close push remote can pass `DOCKERHUB_TOKEN` setting

## usage at github action

```yml

permissions: # https://docs.github.com/actions/using-workflows/workflow-syntax-for-github-actions#permissions
  contents: write
  discussions: write
  packages: write

jobs:
  version:
    name: version
    uses: ./.github/workflows/version.yml

  docker-buildx-bake-multi-tag:
    name: docker-buildx-bake-multi-tag
    needs:
      - version
    uses: ./.github/workflows/docker-buildx-bake-multi-tag.yml
    if: startsWith(github.ref, 'refs/tags/')
    with:
      ghcr_package_owner_name: ${{ github.repository_owner }} # required for ghcr.io
      # push_remote_flag: true # default is true
    secrets:
      DOCKERHUB_TOKEN: "${{ secrets.DOCKERHUB_TOKEN }}"
```

- `push_remote_flag` default is `false`