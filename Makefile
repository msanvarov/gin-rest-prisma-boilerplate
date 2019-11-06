.PHONY: build clean test docker-run

OUT = ${GOPATH}/bin/gin-rest-prisma-boilerplate

build:
	go build -v -o ${OUT} .

clean:
	rm ${OUT}

test:
	go test -v ./tests/*

docker-run:
	docker-compose up && docker exec -it "cd prisma; prisma deploy"