include .env
export

.PHONY: build
build: 
	cd ./function
	GOOS=linux go build -o main 
	cd ../
