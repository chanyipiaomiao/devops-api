#!/bin/bash

# 程序名称
APP_NAME=devops-api

# 程序的版本
APP_VERSION=1.0.0

# 编译时的Go版本
GO_VERSION=`go version | awk '{print $3}'`

# Git 提交时的ID
COMMIT_HASH=`git rev-parse HEAD 2>/dev/null`

# 编译日期时间
BUILD_DATE=`date "+%F %H:%M:%S"`

goreleaser