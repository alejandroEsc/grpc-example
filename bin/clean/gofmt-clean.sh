#!/bin/bash

# from http://github.com/kubernetes/kubernetes/hack/verify-gofmt.sh

# This script is really designed to use after gofmt was found to fail and you want to
# go ahead with suggestions. Its worth while having this hear in case we want to do
# something slightly different in the future.

set -o errexit
set -o nounset
set -o pipefail

ROOT=$(dirname "${BASH_SOURCE}")/../..

cd "${ROOT}"

gofmt=$(which gofmt)
if [[ ! -x "${gofmt}" ]]; then
  echo "could not find gofmt, please verify your GOPATH"
  exit 1
fi

source "${ROOT}/bin/common.sh"

# gofmt exits with non-zero exit code if it finds a problem unrelated to
# formatting (e.g., a file does not parse correctly). Without "|| true" this
# would have led to no useful error message from gofmt, because the script would
# have failed before getting to the "echo" in the block below.
diff=$( echo `valid_go_files` | xargs ${gofmt} -s -w 2>&1) || true
if [[ -n "${diff}" ]]; then
  echo "${diff}"
  exit 1
fi