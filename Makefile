CURRENT_DIR=$(shell pwd)

gen-proto-module:
	./script/genproto.sh ${CURRENT_DIR}

rm-proto-module:
	sudo rm -rf genproto

swag-init:
	swag init -g api/api.go -o api/docs

run:
	go run cmd/main.go
