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

# Some Docs
function foo_bar() {
  echo "#### TEST ####"
}

function foo-foo(){
  echo "#### TEST ####"
}

function foo3 () {
  echo "#### TEST ####"
}

function foo1() {
  echo "#### TEST ####"
  _log_success "SUCCESS"
  _log_warning "WARNING"
  _log_success "Arguments"
  _log_success '  $0: '"$0"
  _log_success '  $1: '"$1"
  _log_success '  $2: '"$2"
}

# Some Docs
function test() {
  echo "#### TEST ####"
  _log_success "SUCCESS"
  _log_warning "WARNING"
  _log_success "Arguments"
  echo "$@"
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

