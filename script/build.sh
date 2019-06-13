#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

PROJECT_NAME="util"
NIGHTLY=""
VERSION=""
OUTPUT_DIR=${PWD}/binaryfiles
VERSION_REGEX="## ([0-9].[0-9].[0-9]) - \[(.*)\]"

# prepare 'files' dir for artifacts
if [[ ! -d ${OUTPUT_DIR} ]]; then
    echo "create dir ${OUTPUT_DIR}"
    mkdir ${OUTPUT_DIR}
fi


#while read LINE; do
#    if [[ ${LINE} =~ ${VERSION_REGEX} ]];then
#       if [[ ${VERSION} != "" ]]; then
#            break
#        fi
#        VERSION=${BASH_REMATCH[1]}
#        if [[ ${BASH_REMATCH[2]}  == "unreleased" ]]; then
#            NIGHTLY="nightly"
#        fi
#        break
#    fi
#done < CHANGELOG.md

if [[ ${NIGHTLY} != "" ]]; then
    VERSION=${VERSION}_${NIGHTLY}
fi

ERROR_LOG=error.log
rm -f ${ERROR_LOG}

COMMIT=`git rev-parse --short HEAD`
BRANCH=`git rev-parse --abbrev-ref HEAD`
BUILT=`date '+%Y-%m-%d %H:%M:%S'`
GO_VERSION=`go version`
LDFLAGS="-X util/version.Version=${VERSION} -X 'util//version.BuildTime=${BUILT}' -X util/version.GitBranch=${BRANCH} -X util/version.GitCommit=${COMMIT} -X 'util/version.GoVersion=${GO_VERSION#"go version "}'"
# go build -ldflags "-extldflags \"-static\" ${LDFLAGS}" -o ${OUTPUT_DIR}/hybird-${MODULE} ./cmd/${MODULE}/*.go

for MODULE in $@; do
    echo "start build ${MODULE}"
    go build -ldflags "-s -w ${LDFLAGS}" -o ${OUTPUT_DIR}/${PROJECT_NAME}-${MODULE} ./cmd/${MODULE}/*.go || echo "${MODULE}" > ${ERROR_LOG} &
done

# wait all build done
wait

if [[ -s ${ERROR_LOG} ]]; then
    exit 1
else
    exit 0
fi