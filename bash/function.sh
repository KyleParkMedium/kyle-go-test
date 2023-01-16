#!/bin/bash

# while getopts ":F:u:h:o:s:f:t:r:d:p:a:b:z:" OPT; do
#     case $OPT in
#     b)
#         echo "check"
#         ;;
#     esac
# done
# shift "$((OPTIND - 1))"

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
    while [ "${1:-}" != "" ]; do
        case "$1" in
        -F | --Function)
            shift
            function=$1
            ;;
        -u | --user)
            shift
            who=$1
            ;;
        -h | --holder)
            shift
            tokenHolder0=$1
            tokenHolder=$(cat ./log/${tokenHolder0}_Address.log)
            string+="\\\"tokenHolder\\\":\\\"$tokenHolder\\\","
            ;;
        -o | --owner)
            shift
            owner0=$1
            owner=$(cat ./log/${owner0}_Address.log)
            string+="\\\"owner\\\":\\\"$owner\\\","
            ;;
        -s | --spender)
            shift
            spender0=$1
            spender=$(cat ./log/${spender0}_Address.log)
            string+="\\\"spender\\\":\\\"$spender\\\","
            ;;
        -f | --from)
            shift
            from0=$1
            from=$(cat ./log/${from0}_Address.log)
            string+="\\\"from\\\":\\\"$from\\\","
            ;;
        -t | --to)
            shift
            to0=$1
            to=$(cat ./log/${to0}_Address.log)
            string+="\\\"to\\\":\\\"$to\\\","
            ;;
        -r | --recipient)
            shift
            recipient0=$1
            recipient=$(cat ./log/${recipient0}_Address.log)
            string+="\\\"recipient\\\":\\\"$recipient\\\","
            ;;
        -d | --drop)
            shift
            recipient0=$1
            recipient=$(cat ./log/${recipient0}_Address.log)
            array+="\\\"${recipient0}\\\":\\\"$recipient\\\","
            ;;
        -p | --partition)
            shift
            partition=$1
            string+="\\\"partition\\\":\\\"$partition\\\","
            ;;
        -a | --amount)
            shift
            amount=$1
            string+="\\\"amount\\\":$amount,"
            ;;
        -b | --bookmark)
            shift
            bookmark=$1
            string+="\\\"bookmark\\\":\\\"$bookmark\\\","
            ;;
        -p | --pageSize)
            shift
            pageSize=$1
            string+="\\\"pageSize\\\":$pageSize,"
            ;;
        *)
            usage
            ;;
        esac
        shift
    done
}

set_options "$@"

sleep 3

$function
