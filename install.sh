#!/bin/bash

SCRIPT_PATH=${SCRIPT_PATH:-$(dirname $(readlink -e $0))}
MAIN_PATH=$SCRIPT_PATH

ENV_FILE_NAME=".env"
ENV_FILE=${ENV_FILE:-"$MAIN_PATH/$ENV_FILE_NAME"}

APP_NAME=npulse-agent
APP_FILE=${MAIN_PATH}/${APP_NAME}

BIN_PATH=/usr/local/bin

function check_root
{
    local isExit=${1:-true}
    local ret=1

    if [ "$EUID" -ne 0 ]; then
        echo "Please run as root"
        ret=0
        if [ $isExit ==  true ]; then
            exit 1
        fi
        return $ret
    fi

    return $ret
}

function env_file_load
{
    local isExit=${1:-true}
    local ret=1

    if [ -e $ENV_FILE ]; then
        source $ENV_FILE
    else
        ret=0
        echo "Please set up your .env file before starting your environment."
        if [ $isExit ==  true ]; then
            exit 1
        fi
    fi

    return $ret
}

function main
{
    check_root

    cp $APP_FILE $BIN_PATH
    chmod 755 ${BIN_PATH}/${APP_NAME}
}

main