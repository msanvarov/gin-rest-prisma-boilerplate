.PHONY: build clean test

OUT = ${GOPATH}/bin/gin-rest-prisma-boilerplate

build:
	go build -v -o ${OUT} .

clean:
	rm ${OUT}

test:
	go test -v ./tests/*
