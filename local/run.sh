#!/bin/bash

BOARD="arduino:avr:nano:cpu=atmega328P"

for i in "$@"
    do
        case $i in
        --board=*)
            BOARD="${i#*=}"
        ;;
        *)
        BUILDPATH="${i#*=}"
        # unknown option
        ;;
    esac
done

arduino --verify --pref sketchbook.path=$(pwd) --board $BOARD $BUILDPATH
arduino --upload --pref sketchbook.path=$(pwd) --board $BOARD $BUILDPATH
