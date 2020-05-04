#!/usr/bin/env bash

# Build swagger from the gate source, then build the Go SDK from swagger.

# docker build -t spinnaker/swagger -f Dockerfile.swagger .
# docker run --rm -v "$(pwd):/workspace" --workdir "/workspace" spinnaker/swagger

set -o errexit -o nounset -o pipefail -o posix

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd -P)"

GATE_PATH="${GATE_PATH:-/spinnaker/gate}"
GATE_VERSION="${GATE_VERSION:-master}"
CODEGEN_CLI_JAR="${CODEGEN_CLI_JAR:-/spinnaker/swagger-codegen-cli.jar}"

mkdir -o "${GATE_PATH}"
cd "${GATE_PATH}"
git clone https://github.com/spinnaker/gate .
cit checkout "${GATE_VERSION}"

# build gate from source and generate swagger/swagger.json
cd "${GATE_PATH}"
swagger/generate_swagger.sh

rm -r "${REPO_ROOT}/gateapi"

# generate the Go Gate SDK from swagger/swagger.json
java -jar "${CODEGEN_CLI_JAR}" generate \
  -i "${GATE_PATH}/swagger/swagger.json" \
  -l go \
  -o "${REPO_ROOT}/gateapi"

# format the Go code to avoid unnessesary diffs
cd "${REPO_ROOT}"
goimports -w -l .
