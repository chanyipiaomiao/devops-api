#!/bin/bash

# 编译时的Go版本
GO_VERSION=`go version | awk '{print $3}'`

GO_VERSION goreleaser --rm-dist