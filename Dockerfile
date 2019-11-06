# Fetching Go
FROM golang:latest

# Label
LABEL maintainer="Sal Anvarov <msalanvarov@gmail.com>"

# Setting a current working directory
WORKDIR /go/src/app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy source files
COPY . .

# Environment setup
ENV GOOS linux
ENV GOARCH amd64

# Node PPA
RUN curl -sL https://deb.nodesource.com/setup_13.x | bash -

# Node 
RUN apt-get install nodejs

# Prisma CLI
RUN npm i -g prisma

# Building the Go app
RUN go build -v -o app .

# Expose default port 9000
EXPOSE 9000

# Starting the server
CMD  [ "./app" ]
