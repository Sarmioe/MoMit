#!/bin/bash

GOOS=windows GOARCH=amd64 go build -o MoMitServer-win-amd64.exe
GOOS=windows GOARCH=arm64 go build -o MoMitServer-win-arm64.exe
GOOS=windows GOARCH=386 go build -o MoMitServer-win-i386.exe

GOOS=linux GOARCH=amd64 go build -o MoMitServer-linux-amd64
GOOS=linux GOARCH=arm64 go build -o MoMitServer-linux-arm64
GOOS=linux GOARCH=386 go build -o MoMitServer-linux-i386
GOOS=linux GOARCH=riscv64 go build -o MoMitServer-linux-riscv64
GOOS=linux GOARCH=mips go build -o MoMitServer-linux-mips
GOOS=linux GOARCH=mips64 go build -o MoMitServer-linux-mips64
GOOS=linux GOARCH=mipsle go build -o MoMitServer-linux-mipsle
GOOS=linux GOARCH=mips64le go build -o MoMitServer-linux-mips64le
GOOS=linux GOARCH=loong64 go build -o MoMitServer-linux-loong64

GOOS=darwin GOARCH=amd64 go build -o MoMitServer-darwin-amd64
GOOS=darwin GOARCH=arm64 go build -o MoMitServer-darwin-arm64
