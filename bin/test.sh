#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

my_dir=$(dirname "${BASH_SOURCE}")


source ${my_dir}/../bin/common.sh


# First pass at testing.
# we will used grpcc to test the endpoint if not found do a full stop.
# this assumes a running instance of your service

inf "grpcc --insecure --proto "${my_dir}/../api/hello.proto" --address ${SERVER_ADDRESS}  --exec ${my_dir}/../bin/test_apis.js"
grpcc --insecure --proto "${my_dir}/../api/hello.proto" --address ${SERVER_ADDRESS}  --exec ${my_dir}/../bin/test_apis.js
