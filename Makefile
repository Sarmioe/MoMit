.PHONY: all client server clean

all: client server

client:
	cd ./MoMitClient && \
	go build -o MoMitClient
	@echo "Client built in MoMitClient folder"

server:
	cd ./MoMitServer && \
	go build -o MoMitServer
	@echo "Server built in MoMitServer folder"

clean:
	rm -rf MoMitClient/MoMitClient MoMitServer/MoMitServer
	@echo "Deleted build files successfully"
