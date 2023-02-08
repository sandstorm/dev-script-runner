#!/bin/bash

set -e

# running tests
function run-test() {
  pushd utils
  go test "$@"
  popd
  _log_green "All Tests finished successfully ;)"
}
