#!/bin/bash
############################## DEV_SCRIPT_MARKER ##############################
# This script is used to document and run recurring tasks in development.     #
#                                                                             #
# You can run your tasks using the script `./dev some-task`.                  #
# You can install the Sandstorm Dev Script Runner and run your tasks from any #
# nested folder using `dev some-task`.                                        #
# https://github.com/sandstorm/Sandstorm.DevScriptRunner                      #
###############################################################################

set -e

######### TASKS #########

# easy setup of the project
function setup() {
  # As the setup typically is more complex we recommend using a separate shell script
  ./dev_setup.sh
}

# sometask to help with something
#
# The first line of the comment will we used in the task overview.
# If you want to provide more details just add more lines ;)
function sometask() {
  # Most task will only require some steps. We recommend implementing them here
  _log_success "Some task"
  _log_warning "TODO: implement more steps"
}

# another task to help with something else
#
# The first line of the comment will we used in the task overview.
# If you want to provide more details just add more lines ;)
function taskwitharguments() {
  # You can access arguments using $@ array. The task name will not be part of the array
  _log_success "Task with arguments"
  _log_warning "TODO: implement more steps"
  _log_success "Arguments"
  _log_success '  $0: '"$0"
  _log_success '  $1: '"$1"
  _log_success '  $2: '"$2"
}

####### Utilities #######

_log_success() {
  printf "\033[0;32m%s\033[0m\n" "${1}"
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
