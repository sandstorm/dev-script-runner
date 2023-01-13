#!/bin/bash
############################## DEV_SCRIPT_MARKER ##############################
# This script is used to document and run recurring tasks in development.     #
#                                                                             #
# You can run your tasks using the script `./dev some-task`.                  #
# You can install the Sandstorm Dev Script Runner and run your tasks from any #
# nested folder using `dev run some-task`.                                    #
###############################################################################

set -e

######### TASKS #########

function foo() {
  echo "test"
}

function switch-dev-binary() {
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

function restore-dev-binary() {
  if test -f "/usr/local/bin/dev_back"; then
      echo "/usr/local/bin/dev_back exists."
      echo "restoring /usr/local/bin/dev "
      rm -f /usr/local/bin/dev || true
      mv /usr/local/bin/dev_back /usr/local/bin/dev || true
    else
      echo "no /usr/local/bin/dev_back found. Nothing to restore!"
  fi
}

function build() {
  go build main
}

function setup() {
  # As the setup typically is more complex we recommend using a separate
  # file `dev_setup.sh`
  ./dev_setup.sh
}

####### Utilities #######

_log_success() {
  printf "\033[0;32m${1}\033[0m\n"
}
_log_warning() {
  printf "\033[1;33m%s\033[0m\n" "${1}"
}
_log_error() {
  printf "\033[0;31m%s\033[0m\n" "${1}"
}

# THIS NEEDS TO BE LAST!!!
# this will run your tasks
"$@"
