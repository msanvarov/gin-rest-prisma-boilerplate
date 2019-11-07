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

# Make wait-for-it an executable
RUN chmod +x ./docker/wait-for-it.sh ./docker/run.sh

# Building the Go app
RUN go build -v -o app .

# Node PPA
RUN curl -sL https://deb.nodesource.com/setup_13.x | bash -

# Downloading Node
RUN apt install nodejs

# Prisma CLI
RUN npm i -g prisma

# Expose default port 9000
EXPOSE 9000
