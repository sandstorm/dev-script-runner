#!/bin/bash
############################## DEV_SCRIPT_MARKER ##############################

set -e

######### TASKS #########

function task-not-starting-with-underscore() {
    echo "Some Code"
}

function task-not-starting-with-underscore_() {
    echo "Some Code"
}

function task-not-starting-with_underscore() {
    echo "Some Code"
}

function _task-with-underscore() {
    echo "Some Code"
}

function __task-with-underscore() {
    echo "Some Code"
}

function ___task-with-underscore() {
    echo "Some Code"
}
