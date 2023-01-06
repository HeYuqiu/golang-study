#!/bin/bash
## This script is only applicable to the debian series of operating systems
set -e -x -u

DISABLE_KUBE_PROXY="false"

function main() {
    echo "pre ${DISABLE_KUBE_PROXY}"
    source user-data-config.sh
    echo "after ${DISABLE_KUBE_PROXY}"
}

main "$@"
