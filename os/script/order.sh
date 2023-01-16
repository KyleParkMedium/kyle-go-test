#!/bin/bash

function A() {
    echo $*
    echo $#
    echo $0
}

function B() {
    echo $*
    echo $#
    echo $0
}

function main() {
    A aa bb cc
    B qq ww ee
}

main
