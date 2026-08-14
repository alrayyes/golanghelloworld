# golanghelloworld

[![Build](https://github.com/alrayyes/golanghelloworld/actions/workflows/build.yml/badge.svg)](https://github.com/alrayyes/golanghelloworld/actions/workflows/build.yml)
[![Tests](https://github.com/alrayyes/golanghelloworld/actions/workflows/test.yml/badge.svg)](https://github.com/alrayyes/golanghelloworld/actions/workflows/test.yml)
[![Codecov](https://codecov.io/gh/alrayyes/golanghelloworld/graph/badge.svg?token=LMBZHSBSSD)](https://codecov.io/gh/alrayyes/golanghelloworld)
[![Go Reference](https://pkg.go.dev/badge/github.com/alrayyes/golanghelloworld.svg)](https://pkg.go.dev/github.com/alrayyes/golanghelloworld)

Just a silly hello world project.

## Prerequisites

[Go](https://go.dev/dl/) 1.24+, and nothing else to run it.

Working on it needs more — the tooling list, and everything else about changing
this, is in [CONTRIBUTING.md](CONTRIBUTING.md).

## Usage

```shell
go run ./cmd/golanghelloworld
```

The greeting itself lives in [`internal/greeting`](./internal/greeting) so it can be imported and tested without going through the binary, and `cmd/golanghelloworld` is only the wiring that prints it. It sits under `internal/` because nothing outside this module has any business importing a hello world: Go enforces that, and it keeps the package free to change without being somebody's API.

## Contributing

Everything about working on this — the tooling, the git hooks, what each linter
is for, and how a release is cut — is in [CONTRIBUTING.md](CONTRIBUTING.md).
Short version: `bun install`, branch, and commit under
[Conventional Commits](https://www.conventionalcommits.org/), because those
subjects pick the next version number.

## License

[GPL-3.0-only](./LICENSE)
