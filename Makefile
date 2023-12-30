format:
	gofmt -w ./

build:
	go build -o api ./api/.

builddocker:
	docker build -t qeditor .


#===============================================
.DEFAULT_GOAL := start
.PHONY: start
start:
	go run ./api/.