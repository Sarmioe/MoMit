set GOOS=windows
set GOARCH=amd64
go build -o MoMitServer-win-amd64.exe

set GOOS=windows
set GOARCH=arm64
go build -o MoMitServer-win-arm64.exe

set GOOS=windows
set GOARCH=386
go build -o MoMitServer-win-i386.exe

set GOOS=linux
set GOARCH=amd64
go build -o MoMitServer-linux-amd64

set GOOS=linux
set GOARCH=arm64
go build -o MoMitServer-linux-arm64

set GOOS=linux
set GOARCH=386
go build -o MoMitServer-linux-i386

set GOOS=linux
set GOARCH=riscv64
go build -o MoMitServer-linux-riscv64

set GOOS=linux
set GOARCH=mips
go build -o MoMitServer-linux-mips

set GOOS=linux
set GOARCH=mips64
go build -o MoMitServer-linux-mips64

set GOOS=linux
set GOARCH=mipsle
go build -o MoMitServer-linux-mipsle

set GOOS=linux
set GOARCH=mips64le
go build -o MoMitServer-linux-mips64le

set GOOS=linux
set GOARCH=loong64
go build -o MoMitServer-linux-loong64

set GOOS=darwin
set GOARCH=amd64
go build -o MoMitServer-darwin-amd64

set GOOS=darwin
set GOARCH=arm64
go build -o MoMitServer-darwin-arm64
