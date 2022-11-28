3#!/bin/bash

# Setup Script to provide a great first run experience ;)

set -e

brew install go
# se our homebrew tap to always get the latest updates.
brew install goreleaser/tap/goreleaser

