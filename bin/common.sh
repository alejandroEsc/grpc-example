#!/usr/bin/env bash

# Some useful colors.
declare -r color_start="\033["
declare -r color_red="${color_start}0;31m"
declare -r color_yellow="${color_start}0;33m"
declare -r color_green="${color_start}0;32m"
declare -r color_norm="${color_start}0m"

function warn {
  echo -e "\033[1;33mWARNING: $1\033[0m"
}

function error {
  echo -e "\033[0;31mERROR: $1\033[0m"
}

function inf {
  echo -e "\033[0;32m$1\033[0m"
}

packages() {
  echo "./client/ ./configs/ ./server/ ./server-gw/"
}

valid_go_files() {
  git ls-files "**/*.go" "*.go" | grep -v -e "vendor"
}
