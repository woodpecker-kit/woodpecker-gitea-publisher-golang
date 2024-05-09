## need `New repository secret`

- file `docker-buildx-bake-hubdocker-latest.yml`
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

  docker-bake-alpine-all-latest:
    name: docker-bake-alpine-all-latest
    needs:
      - version
    uses: ./.github/workflows/docker-buildx-bake-hubdocker-latest.yml
    if: ${{ github.ref_type != 'tag' }}
    # if: ${{ ( github.event_name == 'push' && github.ref == 'refs/heads/main' ) || github.base_ref == 'main' }}
    with:
      docker_bake_targets: 'image-alpine'
      push_remote_flag: false # for flavor latest only one image by `image-basic-all` so no need push remote
      docker-metadata-flavor-suffix: '-alpine'
    secrets:
      DOCKERHUB_TOKEN: "${{ secrets.DOCKERHUB_TOKEN }}"

  docker-bake-basic-all-latest:
    name: docker-bake-basic-all-latest
    needs:
      - version
    uses: ./.github/workflows/docker-buildx-bake-hubdocker-latest.yml
    if: ${{ github.ref_type != 'tag' }}
    # if: ${{ ( github.event_name == 'push' && github.ref == 'refs/heads/main' ) || github.base_ref == 'main' }}
    with:
      # push_remote_flag: ${{ github.ref == 'refs/heads/main' }}
      push_remote_flag: ${{ github.event.pull_request.merged == true }}
      docker_bake_targets: 'image-basic'
    secrets:
      DOCKERHUB_TOKEN: "${{ secrets.DOCKERHUB_TOKEN }}"
```

- `push_remote_flag` default is `false`