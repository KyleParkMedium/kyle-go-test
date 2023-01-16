#!/bin/bash
while getopts ":ab:cd" opt; do
    case $opt in
    a) echo "Found the option -a" ;;
    b) echo "Found the option -b, parameter $OPTARG" ;;
    c) echo "Found the option -c" ;;
    d) echo "Found the option -d" ;;
    *) echo "$opt is not the option" ;;
    esac
done

echo $OPTIND
