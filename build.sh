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

# 编译时传入到代码中的一些变量
LDFLAGS="-X \"devops-api/common.CommitHash=${COMMIT_HASH}\" \
         -X \"devops-api/common.BuildDate=${BUILD_DATE}\" \
         -X \"devops-api/common.AppVersion=${APP_VERSION}\" \
         -X \"devops-api/common.AppName=${APP_NAME}\" \
         -X \"devops-api/common.GoVersion=${GO_VERSION}\""


DIST=dist
DIST_APP=${DIST}/${APP_NAME}

[[ ! -e ${DIST_APP} ]] && mkdir -p ${DIST_APP}

DEST=('linux:amd64' 'linux:386' 'darwin:amd64')
# DEST=('linux:amd64')

function build(){
    for i in ${DEST[@]};do
        goarch=${i#*:}
        goos=${i%%:*}
        export GOOS=$goos GOARCH=$goarch
        TARNAME=${APP_NAME}-${APP_VERSION}_${GOOS}_${GOARCH}.tar.gz
        CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags "${LDFLAGS}" -o ${DIST_APP}/${APP_NAME}
        cp -rf conf ${DIST_APP}
        cd ${DIST}
        tar zcf ${TARNAME} ${APP_NAME}
        cd ..
    done
    rm -rf ${DIST_APP}
}

build
