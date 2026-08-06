# golanghelloworld

[![Go Reference](https://pkg.go.dev/badge/github.com/alrayyes/golanghelloworld.svg)](https://pkg.go.dev/github.com/alrayyes/golanghelloworld)
[![Go Report Card](https://goreportcard.com/badge/github.com/alrayyes/golanghelloworld)](https://goreportcard.com/report/github.com/alrayyes/golanghelloworld)
[![codecov](https://codecov.io/gh/alrayyes/golanghelloworld/graph/badge.svg?token=LMBZHSBSSD)](https://codecov.io/gh/alrayyes/golanghelloworld)

Just a silly hello world project.

## Prerequisites

- [Go](https://go.dev/dl/) 1.24+
- [Bun](https://bun.sh) (JS tooling: prettier, markdownlint, commitlint, lefthook)
- [Docker](https://www.docker.com/) (used by lint-staged/lefthook for Dockerfile linting via hadolint and shellcheck)

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
golangci-lint run       # Go
bun prettier:lint .     # formatting
bun markdown:lint .     # Markdown
```

Dockerfiles are linted with [hadolint](https://github.com/hadolint/hadolint) via `docker compose` (see [`docker-compose.yml`](./docker-compose.yml)), and shell scripts are linted with [shellcheck](https://www.shellcheck.net/) in CI.

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
