.PHONY: build clean test run ensure-deps prisma-endpoint-to-docker prisma-endpoint-to-local compose-deps prisma-deploy

OUT = ${GOPATH}/bin/gin-rest-prisma-boilerplate

build:
	go build -v -o ${OUT} .

clean:
	rm ${OUT}

test:
	go test -v ./tests/*

run:
	go run main.go

compose-deps:
	docker-compose -f docker-compose.override.yml up -d

prisma-endpoint-to-local:
	awk -f ./docker/replace.awk ./docker/to-local.txt ./prisma/prisma.yml | tee ./prisma/prisma.yml

prisma-endpoint-to-docker:
	awk -f ./docker/replace.awk ./docker/to-docker.txt ./prisma/prisma.yml | tee ./prisma/prisma.yml

prisma-deploy:
	prisma deploy 

ensure-deps:
	make compose-deps && make prisma-endpoint-to-local && make prisma-deploy