#!/usr/bin/env bash

export DEP_VERSION=0.4.1
curl -L -s https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 -o dep
chmod +x dep
./dep ensure
