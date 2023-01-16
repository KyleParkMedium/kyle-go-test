#!/bin/bash

# a=1
function local() {

    local a=2
}

function main() {
    local
    echo $a
}

main
