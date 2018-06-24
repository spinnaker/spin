#!/usr/bin/env bash

export DEP_VERSION=0.4.1
wget https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 && mv dep-linux-amd64 dep && chmod +x dep

./dep ensure
