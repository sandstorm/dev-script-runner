#!/bin/bash
############################## DEV_SCRIPT_MARKER ##############################

set -e

######### TASKS #########

# one line comment block - 1
function one-line-comment-block() {
    echo "Some Code"
}

# multiline comment block - 1
# multiline comment block - 2
# multiline comment block - 3
function multiline-comment-block() {
    echo "Some Code"
}

# multiline comment block with empty lines - 1
#
# multiline comment block with empty lines - 2
#
# multiline comment block with empty lines - 3
function empty-lines-in-comment-block() {
    echo "Some Code"
}

#leading space missing
#   keep leading-spaces used for indents
#
#
#
# comment block end
function leading-spaces-or-none-in-comment-block() {
    echo "Some Code"
}
