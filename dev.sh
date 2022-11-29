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

function test() {
  echo "#### TEST ####"
  _log_success "SUCCESS"
  _log_warning "WARNING"
}

function setup() {
  # As the setup typically is more complex we recommend using a separate
  # file `dev_setup.sh`
  ./dev_setup.sh
}

function sometask() {
    # Most task will only require some steps. We recommend implementing them here
    _log_success "First Step of some task"
    _log_warning "TODO: implement more steps"
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
