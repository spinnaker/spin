#!/usr/bin/env bash

go get -u github.com/golang/dep/...
dep ensure -v
