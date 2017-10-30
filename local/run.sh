#!/bin/bash

BOARD="arduino:avr:nano:cpu=atmega328"

for i in "$@"
    do
        case $i in
        --board=*)
            BOARD="${i#*=}"
        ;;
        *)
        PATH="${i#*=}"
        # unknown option
        ;;
    esac
done

arduino --verify --pref sketchbook.path=$(pwd) --board $BOARD $PATH
arduino --upload --pref sketchbook.path=$(pwd) --board $BOARD $PATH