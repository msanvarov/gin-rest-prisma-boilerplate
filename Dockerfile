FROM        golang:1.13-buster as build
WORKDIR     /go/src/app
COPY        . .
ENV         GOOS linux
ENV         GOARCH amd64
RUN         curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh \
                && echo 'Running dep ensure' \
                && dep ensure \
                && echo 'Running go build' \
                && go build -v -ldflags="-w -s" -o /go/bin/app .

FROM        gcr.io/distroless/base as runtime
WORKDIR     /out
COPY        --from=build /go/bin/app .
COPY        --from=build /go/src/app/config.yaml .
ENTRYPOINT  [ "./app" ]
