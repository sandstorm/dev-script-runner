#!/bin/bash
############################## DEV_SCRIPT_MARKER ##############################
# This script is used to document and run recurring tasks in development.     #
#                                                                             #
# You can run your tasks using the script `./dev some-task`.                  #
# You can install the Sandstorm Dev Script Runner and run your tasks from any #
# nested folder using `dev run some-task`.                                    #
###############################################################################

source ./dev_utilities.sh
source ./dev_tasks_testing.sh
source ./dev_tasks_release.sh

set -e

######### TASKS #########

# install dev dependencies
function setup() {
  brew install go
  brew install goreleaser/tap/goreleaser
}

# build
function build() {
  go build main
}

# exposes ./main binary globally
# we rename the original file and copy a fresh build to `/usr/local/bin/`
function switch-binary() {
    echo "-----------> creating build"
    build
    echo "-----------> replacing binary"
    if test -f "/usr/local/bin/dev_back"; then
      rm -f /usr/local/bin/dev || true
    else
      mv /usr/local/bin/dev /usr/local/bin/dev_back || true
    fi
    mv ./main /usr/local/bin/dev || true
}

# restores original dev binary
#
# we check if the dev binary was backed up and replace the current dev with
# this backup ;)
function restore-binary() {
  if test -f "/usr/local/bin/dev_back"; then
      echo "/usr/local/bin/dev_back exists."
      echo "restoring /usr/local/bin/dev "
      rm -f /usr/local/bin/dev || true
      mv /usr/local/bin/dev_back /usr/local/bin/dev || true
    else
      echo "no /usr/local/bin/dev_back found. Nothing to restore!"
  fi
}

_log_green "------------------------- RUNNING TASK: $1 -------------------------"

# THIS NEEDS TO BE LAST!!!
# this will run your tasks
"$@"
