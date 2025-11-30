#! /bin/bash
source .env
mkdir -p Day$1
cp -n utils/template.go Day$1/main.go
cp -n utils/utils.go Day$1
touch Day$1/test.txt
curl -b "session=$SESSION" https://adventofcode.com/2024/day/$1/input > Day$1/input.txt
