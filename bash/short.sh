#!/bin/bash

function usage() {
    cat <<EOM
Usage: $0 [options] <url>
Options:
 -X, --request COMMAND   Specify request command to use
 -v, --verbose           Make the operation more talkative
EOM

    exit 1
}

function set_options() {
    while getopts "X:v" optname; do
        case $optname in
        X)
            request_method=$OPTARG
            ;;
        v)
            verbose_mode="true"
            ;;
        *)
            usage
            ;;
        esac
    done

    # 남아있는 인자 얻기, 남아있는 인자는 url이다
    shift $((OPTIND - 1))
    url=$@
}

set_options "$@"

echo "request_method='${request_method}'"
echo "verbose_mode='${verbose_mode}'"
echo "url='${url}'"
