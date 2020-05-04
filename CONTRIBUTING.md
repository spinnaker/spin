# Contributing

Interested in contributing to Spinnaker? Please review the [contribution documentation](https://www.spinnaker.io/community/contributing/).

## Setup

### Go

[Install Go 1.13.x](https://golang.org/doc/install).

### Go modules

Clone the repository to a directory outside of your GOPATH:

```bash
$ git clone https://github.com/spinnaker/spin
```

Afterward, use `go build` to build the program. This will automatically fetch dependencies.

```bash
$ go build
```

Upon first build, you may see output while the `go` tool fetches dependencies.

To verify dependencies match checksums under go.sum, run `go mod verify`.

To clean up any old, unused go.mod or go.sum lines, run `go mod tidy`.


## Running Spin

Run using

```bash
./spin <cmds> <flags>
```


## Running tests

Test using

```bash
go test -v ./...
```

from the root `spin/` directory.

## Updating the Gate API

Spin CLI uses [Swagger](https://swagger.io/) to generate the API client library for [Gate](https://github.com/spinnaker/gate).

To make it easier, without having to have a full local Spinnaker build environment, we use Docker as to cache dependencies and standardize versions.

Build the `spinnaker/swagger` image locally:
```bash
docker build -t spinnaker/swagger -f Dockerfile.swagger .
```

Run the `spinnaker/swagger` image locally to generate a new `gateapi/` from the Gate source code:
```bash
docker run --rm -v "$(pwd):/workspace" --workdir "/workspace" spinnaker/swagger
```

Once generated, commit the changes to a new branch and create a PR.
