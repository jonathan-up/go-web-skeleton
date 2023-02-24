#!/bin/bash

Root=$(pwd)
Dir="/out"
BuildDir="$Root""$Dir"

rm -rf "$BuildDir"
mkdir "$BuildDir"
go build -o "$BuildDir"
cp -r ./public "$BuildDir"
cp config.yml "$BuildDir"
