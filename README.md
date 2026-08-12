# golanghelloworld

[![Build](https://github.com/alrayyes/golanghelloworld/actions/workflows/build.yml/badge.svg)](https://github.com/alrayyes/golanghelloworld/actions/workflows/build.yml)
[![Tests](https://github.com/alrayyes/golanghelloworld/actions/workflows/test.yml/badge.svg)](https://github.com/alrayyes/golanghelloworld/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/alrayyes/golanghelloworld/graph/badge.svg?token=LMBZHSBSSD)](https://codecov.io/gh/alrayyes/golanghelloworld)
[![Go Reference](https://pkg.go.dev/badge/github.com/alrayyes/golanghelloworld.svg)](https://pkg.go.dev/github.com/alrayyes/golanghelloworld)

Just a silly hello world project.

## Prerequisites

- [Go](https://go.dev/dl/) 1.24+
- [Bun](https://bun.sh) (JS tooling: biome, prettier, markdownlint, commitlint, lefthook)
- [Docker](https://www.docker.com/) (used by lefthook to run hadolint against Dockerfiles)
- [yamlfmt](https://github.com/google/yamlfmt), [yamllint](https://yamllint.readthedocs.io),
  [actionlint](https://github.com/rhysd/actionlint), [typos](https://github.com/crate-ci/typos)
  and [goreleaser](https://goreleaser.com), run by the `pre-push` hook:

  ```shell
  go install github.com/google/yamlfmt/cmd/yamlfmt@v0.21.0
  go install github.com/rhysd/actionlint/cmd/actionlint@v1.7.12
  go install github.com/goreleaser/goreleaser/v2@v2.17.1
  pipx install yamllint==1.38.0
  cargo install typos-cli --version 1.49.0
  ```

  These five are optional. `bun install` does not provide them, so the hook skips
  any that aren't on your `PATH` rather than failing the push — you can clone and
  contribute without installing all five. CI runs them unconditionally, so the
  check still can't be bypassed, it just moves later.

## Usage

```shell
go run ./cmd/golanghelloworld
```

The greeting itself lives in [`greeting`](./greeting), so it can be imported and tested without going through the binary. `cmd/golanghelloworld` is only the wiring that prints it.

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
golangci-lint run            # Go
bun biome:lint .             # JSON formatting and key ordering
bun prettier:lint "**/*.md"  # Markdown layout
bun markdown:lint .          # Markdown structure
yamlfmt -lint                # YAML formatting
yamllint --strict .          # YAML style
actionlint                   # GitHub Actions workflows
typos                        # spelling, everywhere
goreleaser check             # release config
```

[Biome](https://biomejs.dev) is configured in [`biome.json`](./biome.json) and owns every file
type it supports, including the JSON key sorting that used to come from
`prettier-plugin-sort-json` — that lives on as Biome's `useSortedKeys` assist, switched off for
`package.json` so its conventional key order survives.

[Prettier](https://prettier.io) fills the gap that leaves in prose: Markdown layout, which
markdownlint checks but never lays out. Prettier owns table alignment, so don't line a table up
by hand, it will just redo it. `proseWrap` is `preserve`, so your line breaks stay where you put
them. The `pre-commit` hook runs Prettier first and markdownlint second, and the markdownlint
rules that have an opinion about layout (MD004, MD007, MD012, MD049, MD050) are off in
[`.markdownlint.json`](./.markdownlint.json), so the two can't undo each other. What Prettier
does not touch is listed in [`.prettierignore`](./.prettierignore).

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

release-please runs in manifest mode: [`release-please-config.json`](./release-please-config.json) says how to release and [`.release-please-manifest.json`](./.release-please-manifest.json) holds the version we are on. That manifest is the file to correct by hand when a release goes wrong. Don't move the tag.

## Dependency updates

[Dependabot](./.github/dependabot.yml) keeps Go, Bun, Docker, and GitHub Actions dependencies up to date. PRs that pass CI are merged automatically (see [`auto-merge-dependabot.yml`](./.github/workflows/auto-merge-dependabot.yml)); failures notify via GitHub as usual.

## License

[GPL-3.0-only](./LICENSE)
