#!/bin/bash

set -e

# running tests
function run-test() {
  pushd utils
  go test "$@"
  popd
  _log_success "All Tests finished successfully ;)"
}
