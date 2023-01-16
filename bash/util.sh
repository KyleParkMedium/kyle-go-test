#!/bin/bash

function C() {

    echo ${FUNCNAME[0]}
    func=$1

    shift $(expr $OPTIND)
    while getopts p:a: OPT; do
        case $OPT in
        p)
            partition=$OPTARG
            string+="\"partition\":\"$partition\","
            ;;
        a)
            amount=$OPTARG
            string+="\"amount\":$amount"
            ;;
        esac
    done
    # string-=','
    # args="'{"Args":["$1", "{\"partition\":\"imsyToken\""]}'"
    param="{\"Args\":[\"$func\", \"{$string}\"]}"

    echo $param

}

function A() {

    count=1
    while [ -n "$1" ]; do
        echo $1
        # echo "parameter #$count : $1"
        # ((count = count + 1))

        shift
    done

    # inputToMap

}

function B() {

    while [ -n "$1" ]; do
        while getopts u:a:f: flag; do
            case "${flag}" in
            u) username=${OPTARG} ;;
            a) age=${OPTARG} ;;
            f) fullname=${OPTARG} ;;
            esac
        done
    done

    echo "Username: $username"
    echo "Age: $age"
    echo "Full Name: $fullname"
}

# function inputToMap() {

# }

function main() {

    # CreateWallet
    C CreateWallet -p 'imsyToken' -a 25
    # A a b c
    # B -f 'John Smith' -a 25 -u john
}

main
