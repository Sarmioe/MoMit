set GOOS=windows
set GOARCH=amd64
go build -o MoMitClient-win-amd64.exe

set GOOS=windows
set GOARCH=arm64
go build -o MoMitClient-win-arm64.exe

set GOOS=windows
set GOARCH=386
go build -o MoMitClient-win-i386.exe

set GOOS=linux
set GOARCH=amd64
go build -o MoMitClient-linux-amd64

set GOOS=linux
set GOARCH=arm64
go build -o MoMitClient-linux-arm64

set GOOS=linux
set GOARCH=386
go build -o MoMitClient-linux-i386

set GOOS=linux
set GOARCH=riscv64
go build -o MoMitClient-linux-riscv64

set GOOS=linux
set GOARCH=mips
go build -o MoMitClient-linux-mips

set GOOS=linux
set GOARCH=mips64
go build -o MoMitClient-linux-mips64

set GOOS=linux
set GOARCH=mipsle
go build -o MoMitClient-linux-mipsle

set GOOS=linux
set GOARCH=mips64le
go build -o MoMitClient-linux-mips64le

set GOOS=linux
set GOARCH=loong64
go build -o MoMitClient-linux-loong64

set GOOS=darwin
set GOARCH=amd64
go build -o MoMitClient-darwin-amd64

set GOOS=darwin
set GOARCH=arm64
go build -o MoMitClient-darwin-arm64
