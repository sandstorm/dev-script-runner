#!/bin/bash
############################## DEV_SCRIPT_MARKER ##############################
# This script is used to document and run recurring tasks in development.     #
#                                                                             #
# You can run your tasks using the script `./dev some-task`.                  #
# You can install the Sandstorm Dev Script Runner and run your tasks from any #
# nested folder using `dev some-task`.                                        #
# https://github.com/sandstorm/Sandstorm.DevScriptRunner                      #
###############################################################################

source ./dev_utilities.sh

set -e

######### TASKS #########

# Easy setup of the project
function setup() {
  _log_green "Setting up your project"
  _log_yellow "TODO: implement more steps"
}

# Sometask to help with something
#
# The first line of the comment block will be used in the task overview.
# If you want to provide more details just add more lines ;)
function sometask {
  # Most task will only require some steps. We recommend implementing them here
  _log_green "Some task"
  _log_yellow "TODO: implement more steps"
}

# Another task to help with something else
#
# The first line of the comment block will be used in the task overview.
# If you want to provide more details just add more lines ;)
function taskwitharguments() {
  # You can access arguments using $@ array. The task name will not be part of the array
  _log_green "Task with arguments"
  _log_yellow "TODO: implement more steps"
  _log_green "Arguments"
  _log_green '  $0: '"$0"
  _log_green '  $1: '"$1"
  _log_green '  $2: '"$2"
}

_log_green "---------------------------- RUNNING TASK: $1 ----------------------------"

# THIS NEEDS TO BE LAST!!!
# this will run your tasks
"$@"
