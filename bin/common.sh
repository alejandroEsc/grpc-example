#!/usr/bin/env bash

function warn {
  echo -e "\033[1;33mWARNING: $1\033[0m"
}

function error {
  echo -e "\033[0;31mERROR: $1\033[0m"
}

function inf {
  echo -e "\033[0;32m$1\033[0m"
}

function checkPreReqs {
    # space seperated list
    PRE_REQS=$1

    for pr in $PRE_REQS
    do
      if ! which $pr >/dev/null 2>&1
      then
        echo >&2 "prerequisite application called '${pr}' is not found on this system"
        return=1
      fi
    done

    return 0
}
