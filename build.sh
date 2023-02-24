#!/bin/bash

rm -rf out
mkdir out
go build -o out
cp -r ./public ./out
cp config.yml ./out
