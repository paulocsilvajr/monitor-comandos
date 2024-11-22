#!/bin/bash

sudo apt install -qqq curl jq

resultado=$(curl -s localhost:8080/"$1")

grep "404 page not found" <(echo "$resultado") ||
    echo "$resultado" | jq --raw-output ."$1"
