# golanghelloworld

[![Go Reference](https://pkg.go.dev/badge/github.com/alrayyes/golanghelloworld.svg)](https://pkg.go.dev/github.com/alrayyes/golanghelloworld)
[![Go Report Card](https://goreportcard.com/badge/github.com/alrayyes/golanghelloworld)](https://goreportcard.com/report/github.com/alrayyes/golanghelloworld)
[![codecov](https://codecov.io/gh/alrayyes/golanghelloworld/graph/badge.svg?token=LMBZHSBSSD)](https://codecov.io/gh/alrayyes/golanghelloworld)

Just a silly hello world project.

## Prerequisites

- [Go](https://go.dev/dl/) 1.24+
- [Bun](https://bun.sh) (JS tooling: biome, markdownlint, commitlint, lefthook)
- [Docker](https://www.docker.com/) (used by lefthook to run hadolint against Dockerfiles)
- [yamlfmt](https://github.com/google/yamlfmt), [yamllint](https://yamllint.readthedocs.io),
  [actionlint](https://github.com/rhysd/actionlint) and [typos](https://github.com/crate-ci/typos),
  run by the `pre-push` hook:

  ```shell
  go install github.com/google/yamlfmt/cmd/yamlfmt@v0.21.0
  go install github.com/rhysd/actionlint/cmd/actionlint@v1.7.12
  pipx install yamllint==1.38.0
  cargo install typos-cli --version 1.49.0
  ```

## Usage

```shell
go run .
```

## Development

Clone the repo and install dependencies:

```shell
bun install
```

This also installs the [lefthook](https://lefthook.dev) git hooks (`pre-commit`, `pre-push`, `commit-msg`) defined in [`lefthook.yml`](./lefthook.yml).

### Building

```shell
go build -v ./...
```

### Testing

```shell
go test -cover ./...
```

### Linting

```shell
golangci-lint run    # Go
bun biome:lint .     # JSON formatting and key ordering
bun markdown:lint .  # Markdown
yamlfmt -lint        # YAML formatting
yamllint --strict .  # YAML style
actionlint           # GitHub Actions workflows
typos                # spelling, everywhere
```

[Biome](https://biomejs.dev) is configured in [`biome.json`](./biome.json). It replaces the
old prettier setup, including the `prettier-plugin-sort-json` key sorting — that lives on as
Biome's `useSortedKeys` assist, switched off for `package.json` so its conventional key order
survives.

Biome has no YAML support, so YAML is handled by [yamlfmt](https://github.com/google/yamlfmt)
(formatting, see [`.yamlfmt`](./.yamlfmt)) and [yamllint](https://yamllint.readthedocs.io)
(style, see [`.yamllint.yml`](./.yamllint.yml)). The two disagree about inline comments by
default, so yamllint's `min-spaces-from-content` is lowered to 1 to match what yamlfmt writes —
otherwise each tool would undo the other on every run.

Dockerfiles are linted with [hadolint](https://github.com/hadolint/hadolint) via `docker compose` (see [`docker-compose.yml`](./docker-compose.yml)).

[gitleaks](https://github.com/gitleaks/gitleaks) scans the full history for committed secrets.
It runs in CI only — it needs an unshallow clone, which is too slow to put in a git hook.

### Commit messages

Commits must follow [Conventional Commits](https://www.conventionalcommits.org/), enforced locally by the `commit-msg` hook and in CI. Use the interactive prompt instead of writing messages by hand:

```shell
bun commit
```

## Releases

[release-please](https://github.com/googleapis/release-please) opens release PRs based on Conventional Commits history, and [GoReleaser](https://goreleaser.com/) (see [`.goreleaser.yaml`](./.goreleaser.yaml)) builds and publishes binaries when a version tag is pushed.

## Dependency updates

[Dependabot](./.github/dependabot.yml) keeps Go, Bun, Docker, and GitHub Actions dependencies up to date. PRs that pass CI are merged automatically (see [`auto-merge-dependabot.yml`](./.github/workflows/auto-merge-dependabot.yml)); failures notify via GitHub as usual.

## License

[GPL-3.0-only](./LICENSE)
