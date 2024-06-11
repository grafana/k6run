[![Go Reference](https://pkg.go.dev/badge/github.com/grafana/k6run.svg)](https://pkg.go.dev/github.com/grafana/k6run)
[![GitHub Release](https://img.shields.io/github/v/release/grafana/k6run)](https://github.com/grafana/k6run/releases/)
[![Go Report Card](https://goreportcard.com/badge/github.com/grafana/k6run)](https://goreportcard.com/report/github.com/grafana/k6run)
[![GitHub Actions](https://github.com/grafana/k6run/actions/workflows/test.yml/badge.svg)](https://github.com/grafana/k6run/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/grafana/k6run/graph/badge.svg?token=X9EUHZZADY)](https://codecov.io/gh/grafana/k6run)
![GitHub Downloads](https://img.shields.io/github/downloads/grafana/k6run/total)

<h1 name="title">k6run</h1>

**Launching k6 with extensions**

The purpose of k6run is to launch the k6 binary containing the extensions used by the k6 test script. For this purpose, k6run analyzes the k6 test scripts.

k6run is primarily used as a go library for [k6](https://github.com/grafana/k6) and [xk6](https://github.com/grafana/xk6). In addition, it also contains a command-line tool, which is suitable for launching k6 with extensions based on the dependencies of k6 test scripts.

The command line tool can be integrated into other command line tools as a subcommand. For this purpose, the library also contains the functionality of the command line tool as a factrory function that returns [cobra.Command](https://pkg.go.dev/github.com/spf13/cobra#Command).

## Install

Precompiled binaries can be downloaded and installed from the [Releases](https://github.com/grafana/k6run/releases) page.

If you have a go development environment, the installation can also be done with the following command:

```
go install github.com/grafana/k6run/cmd/k6run@latest
```

## Usage

<!-- #region cli -->
## k6run

Launching k6 with extensions

### Synopsis

Launching k6 containing the extensions used by the test script.


```
k6run [flags] [script-file]
```

### Flags

```
  -h, --help   help for k6run
```

<!-- #endregion cli -->

## Development

### Tasks

This section contains a description of the tasks performed during development. If you have the [xc (Markdown defined task runner)](https://github.com/joerdav/xc) command-line tool, individual tasks can be executed simply by using the `xc task-name` command.

<details><summary>Click to expand</summary>

#### readme

Update documentation in README.md.

```
go run ./tools/gendoc README.md
```

#### lint

Run the static analyzer.

```
golangci-lint run
```

#### test

Run the tests.

```
go test -count 1 -race -coverprofile=build/coverage.txt ./...
```

#### coverage

View the test coverage report.

```
go tool cover -html=build/coverage.txt
```

#### build

Build the executable binary.

This is the easiest way to create an executable binary (although the release process uses the goreleaser tool to create release versions).

```
go build -ldflags="-w -s" -o build/k6run ./cmd/k6run
```

#### snapshot

Creating an executable binary with a snapshot version.

The goreleaser command-line tool is used during the release process. During development, it is advisable to create binaries with the same tool from time to time.

```
goreleaser build --snapshot --clean --single-target -o build/k6run
```

#### clean

Delete the build directory.

```
rm -rf build
```

</details>
