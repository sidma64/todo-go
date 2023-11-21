#!make
include .env
export $(shell sed 's/=.*//' .env)

start:
	cd app; npm install
	cd app; npm run build
	go build -o build/out main.go
	./build/out
