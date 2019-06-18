#!/usr/bin/env bash

set -e

[[ $# == 1 ]] && [[ $1 == 'clear' ]] && rm -rf ./build && exit 1

[[ $# != 1 ]] || [[ $1 != "build" ]] && echo "build: bash $0 build" && echo "clear: bash $0 clear" && exit 1

rm -rf ./build && mkdir -p build/{bin,conf}
go build -o DCache-Cli main.go
mv DCache-Cli build/bin/
cp conf/conf.toml build/conf/
