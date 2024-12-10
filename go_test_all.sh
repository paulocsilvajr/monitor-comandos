#!/bin/bash

echo -n "Deseja limpar o cache:[sim/Não] "
read -r confirmacao

case $confirmacao in
    S | s | Sim | sim)
        go clean -cache
        ;;
    *)
        echo "N"
        ;;
esac

go test -cover ./...
