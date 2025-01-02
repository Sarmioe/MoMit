.PHONY: all client server clean
all: client server

client:
    cd .\MoMitClient\
	go build .\main.go .\utils.go
    echo "Client built in MoMitClient floder"
server:
	cd .\MoMitServer\
    go build .\main.go .\utils.go
	echo "Server built in MoMitServer floder"
debuild:
	rm -rf MoMitClient/MoMitClient MoMitServer/MoMitServer
	echo "Deleted build files sussessfully"