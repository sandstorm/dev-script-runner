#!/bin/bash

set -e

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
