#!/bin/bash

set -e

# releasing a new version
function release() {
  run-test
  build
  goreleaser release --rm-dist
  _log_success "Release finished successfully ;)"
}
