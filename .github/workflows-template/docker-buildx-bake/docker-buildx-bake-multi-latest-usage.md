## need `New repository secret`

- file `docker-buildx-bake-multi-latest.yml`
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

  docker-buildx-bake-multi-latest:
    name: docker-buildx-bake-multi-latest
    needs:
      - version
    uses: ./.github/workflows/docker-buildx-bake-multi-latest.yml
    # if: ${{ ( github.event_name == 'push' && github.ref == 'refs/heads/main' ) || github.base_ref == 'main' }}
    with:
      ghcr_package_owner_name: ${{ github.repository_owner }} # required for ghcr.io
      # push_remote_flag: ${{ github.ref == 'refs/heads/main' }}
      push_remote_flag: ${{ github.event.pull_request.merged == true }}
    secrets:
      DOCKERHUB_TOKEN: "${{ secrets.DOCKERHUB_TOKEN }}"
```

- `push_remote_flag` default is `false`