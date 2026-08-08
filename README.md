# golanghelloworld

[![Go Reference](https://pkg.go.dev/badge/github.com/alrayyes/golanghelloworld.svg)](https://pkg.go.dev/github.com/alrayyes/golanghelloworld)
[![Go Report Card](https://goreportcard.com/badge/github.com/alrayyes/golanghelloworld)](https://goreportcard.com/report/github.com/alrayyes/golanghelloworld)
[![codecov](https://codecov.io/gh/alrayyes/golanghelloworld/graph/badge.svg?token=LMBZHSBSSD)](https://codecov.io/gh/alrayyes/golanghelloworld)

Just a silly hello world project.

## Prerequisites

- [Go](https://go.dev/dl/) 1.24+
- [Bun](https://bun.sh) (JS tooling: biome, markdownlint, commitlint, lefthook)
- [Docker](https://www.docker.com/) (used by lefthook to run hadolint against Dockerfiles)

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
```

[Biome](https://biomejs.dev) is configured in [`biome.json`](./biome.json). It replaces the
old prettier setup, including the `prettier-plugin-sort-json` key sorting — that lives on as
Biome's `useSortedKeys` assist, switched off for `package.json` so its conventional key order
survives. Biome has no YAML support, so `.yml` files are no longer auto-formatted.

Dockerfiles are linted with [hadolint](https://github.com/hadolint/hadolint) via `docker compose` (see [`docker-compose.yml`](./docker-compose.yml)).

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
