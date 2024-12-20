#!/bin/bash

source vars.sh

mkdir -v bin
go build -v -o $DIRETORIOBIN/$PROGRAMA
