.PHONY: release run install all dep clean dockerbuild dockerrun dockertest

OUT = ${GOPATH}/bin/gin-rest-prisma-boilerplate

all:
	go build -v -o ${OUT} .

install:
	which dep 2> /dev/null || echo "dep is not installed. Please run\n\tmake dep" && exit 1
	dep ensure

clean:
	rm ${OUT}

dep:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

run: all
	${OUT}

release:
	go build -ldflags="-w -s" -o ${OUT} .

dockerbuild:
	docker build -t grp-boilerplate .

dockerrun:
	docker run --rm grp-boilerplate

dockertest:
	docker build -t grp-boilerplate-test -f test.Dockerfile .
