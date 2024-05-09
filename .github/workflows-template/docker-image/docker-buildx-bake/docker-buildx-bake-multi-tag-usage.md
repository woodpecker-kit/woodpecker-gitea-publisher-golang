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

  ### deploy tag start

  docker-bake-multi-alpine-all-tag:
    name: docker-bake-multi-alpine-all-tag
    needs:
      - version
    uses: ./.github/workflows/docker-buildx-bake-multi-tag.yml
    if: startsWith(github.ref, 'refs/tags/')
    with:
      ghcr_package_owner_name: ${{ github.repository_owner }} # required for ghcr.io
      docker_bake_targets: 'image-alpine'
      docker-metadata-flavor-suffix: '-alpine' # default is '', can add as: -alpine -debian
      # push_remote_flag: true # default is true
    secrets:
      DOCKERHUB_TOKEN: "${{ secrets.DOCKERHUB_TOKEN }}"

  docker-bake-multi-basic-all-tag:
    name: docker-bake-multi-basic-all-tag
    needs:
      - version
      - docker-bake-multi-alpine-all-tag
    uses: ./.github/workflows/docker-buildx-bake-multi-tag.yml
    if: startsWith(github.ref, 'refs/tags/')
    with:
      ghcr_package_owner_name: ${{ github.repository_owner }} # required for ghcr.io
      docker_bake_targets: 'image-basic'
      docker-metadata-flavor-suffix: '' # default is '', can add as: -alpine -debian
      # push_remote_flag: true # default is true
    secrets:
      DOCKERHUB_TOKEN: "${{ secrets.DOCKERHUB_TOKEN }}"


  deploy-tag:
    needs:
      - version
      - docker-bake-multi-basic-all-tag
    name: deploy-tag
    uses: ./.github/workflows/deploy-tag.yml
    if: startsWith(github.ref, 'refs/tags/')
    secrets: inherit
    with:
      prerelease: true
      tag_name: ${{ needs.version.outputs.tag_name }}
      tag_changes: ${{ needs.version.outputs.cc_changes }}
      # download_artifact_name: foo-release

  ### deploy tag end
```

- `push_remote_flag` default is `false`