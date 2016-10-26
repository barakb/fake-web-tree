#!/usr/bin/env bash

rm -f fake-web-tree fake-web-tree.exe

echo "building linux"
go build -o fake-web-tree github.com/barakb/fake-web-tree/main

echo "building windows"
GOOS=windows GOARCH=386 go build -o fake-web-tree.exe github.com/barakb/fake-web-tree/main
