#!/bin/bash

# Init path
cd "$(dirname "$0")" || exit 1
cd ..

go run ./...