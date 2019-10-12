FROM        golang:1.13-buster as build
WORKDIR     /go/src/app
COPY        . .
RUN         curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh \
                && echo 'Running dep ensure' \
                && dep ensure \
                && echo 'Running tests' \
                && go test -v ./tests/*
