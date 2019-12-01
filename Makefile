.PHONY: develop-locally build clean test run develop-locally prisma-endpoint-to-docker prisma-endpoint-to-local compose-deps prisma-deploy dev docs

OUT = ${GOPATH}/bin/gin-rest-prisma-boilerplate

develop-locally: compose-deps prisma-endpoint-to-local prisma-deploy 

build:
	go build -v -o ${OUT} .

clean:
	rm ${OUT}

test:
	go test -v ./tests/*

dev:
	fresh || go get -u github.com/pilu/fresh && fresh

docs:
	swag init || go get -u github.com/swaggo/swag/cmd/swag
	
run:
	go run main.go

compose-deps:
	docker-compose -f docker-compose.override.yml up -d

prisma-endpoint-to-local:
	awk -f ./docker/replace.awk ./docker/to-local.txt ./prisma/prisma.yml > ./docker/prisma.yml \
	&& mv ./docker/prisma.yml ./prisma/prisma.yml   

prisma-endpoint-to-docker:
	awk -f ./docker/replace.awk ./docker/to-docker.txt ./prisma/prisma.yml > ./docker/prisma.yml \
	&& mv ./docker/prisma.yml ./prisma/prisma.yml    

prisma-deploy:
	prisma deploy