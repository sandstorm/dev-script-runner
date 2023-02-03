#!/bin/bash
############################## DEV_SCRIPT_MARKER ##############################

source ./tasks_from_imports__import1.sh
source ./tasks_from_imports__import2.sh
source ./nested_imports/tasks_from_imports__import3.sh

set -e

######### TASKS #########

function task() {
    echo "Some Code"
}
