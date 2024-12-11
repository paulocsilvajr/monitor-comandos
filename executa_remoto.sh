#!/bin/bash

jq --version > /dev/null && curl --version > /dev/null ||
    sudo apt install -qqq curl jq

# se a rota inicia com /, remove ela
if [[ "$1" =~ ^/ ]]; then
    rota=${1#/}
else
    rota=$1
fi

resultado=$(curl -s localhost:8080/"$rota")

grep "404 page not found" <(echo "$resultado") ||
    echo "$resultado" | jq --raw-output
