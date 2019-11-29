# 🍱 Gin Rest Prisma Boilerplate

<p align="center">  
  <a href="http://golang.org" target="blank"><img src="https://cacophony.org.nz/sites/default/files/gopher.png" width="200" alt="Nest Logo" /></a>  
  <a href="https://gin-gonic.com/" target="blank"><img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" width="200" alt="Nest Logo" /></a>  
</p>  
  
Golang, is a statically typed, compiled programming language designed at Google that aids in building simple, reliable, and efficient software. Gin is a full-featured web framework for Go that achieves outstanding performance.

[![GoDoc](https://godoc.org/github.com/gin-gonic/gin?status.svg)](https://godoc.org/github.com/gin-gonic/gin)
[![codebeat badge](https://codebeat.co/badges/c9d048b7-5d3b-416a-ab64-b9f510f947ed)](https://codebeat.co/projects/github-com-msanvarov-gin-rest-prisma-boilerplate-master)
[![Build Status](https://travis-ci.org/msanvarov/gin-rest-prisma-boilerplate.svg?branch=master)](https://travis-ci.org/msanvarov/gin-rest-prisma-boilerplate)
[![GitHub license](https://img.shields.io/github/license/msanvarov/gin-rest-prisma-boilerplate)](https://github.com/msanvarov/gin-rest-prisma-boilerplate/blob/master/LICENSE)

## 📚 Description

This boilerplate leverages the Gin framework to quickly prototype backend applications. It comes with database, logging, security, and authentication features out of the box.

---


## 🍬 Features

- Based on [Gin](https://github.com/gin-gonic/gin).

- [Prisma ORM](https://www.prisma.io/) for Mongo. But can support MYSQL/PostgreSQL and Amazon Aurora.

- [Gin Sessions](https://github.com/gin-contrib/sessions) for Redis.

- [Gin Authz](https://github.com/gin-contrib/authz) for role based access management. Internally, utilizing the powerful authentication library [Casbin](https://github.com/casbin/casbin).

- [Viper](https://github.com/spf13/viper) for working with yaml configurations.

---

## 🛠️ Prerequisites

### 🐳 Docker

Please make sure to have Docker Desktop operational on the preferred operating system of choice to quickly get started. To get started, please see the following [link](https://www.docker.com/products/docker-desktop).

> **Note: Despite the fact that Docker Desktop comes free for both Mac and Windows, it only supports the Pro edition of Windows 10. A common workaround is to get [Docker Toolbox](https://docs.docker.com/toolbox/toolbox_install_windows/) which will bypass the Windows 10 Pro restriction by executing Docker in a VM.**

### 🧰 Node

The [Prisma CLI](https://www.prisma.io/docs/prisma-cli-and-configuration/using-the-prisma-cli-alx4/) is essential for streamlining workflows in managing and deploying Prisma services. The CLI can to be downloaded using `npm`; which requires [NodeJS](https://nodejs.org/en/download/).

Then, the following command will download the Prisma command line interface:

```bash
// prisma cli
$ npm install -g prisma
```

---

## 🔨 Getting Started

If need be, replace the existing configuration variables in the [config.yaml](https://github.com/msanvarov/gin-rest-prisma-boilerplate/blob/master/config.yaml) file with the preferred configuration values.

> Changing the `server.env : "test"` to `server.env : "dev"` yields better logging that can come of use when developing.

### 🐳 Developing Inside Docker

VSCode permits the development of source code to happen exclusively in a Docker container. For more information on how this works, [please read the following documentation](https://code.visualstudio.com/docs/remote/containers).

This boilerplate comes with a `.devcontainer` configuration enabling such a feature.

To get started:

1. Clone this repository.
2. Open VSCode and download the [Remote-Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers).
3. Press <kbd>F1</kbd> and select the **Remote-Containers: Open Folder in Container...** command.
4. Select the cloned copy of this folder, wait for the container to start.
5. Run `make prisma-deploy` or `prisma deploy`.
6. Run `make dev`
7. Start developing!

### ⛲ Developing Locally Outside of Docker

Developing locally is still possible but requires some tweaks. The application dependencies such as Prisma, Mongo, and Redis will still require Docker to run. Mainly because Prisma can't be set up locally like Mongo and Redis. But the code itself will not be containerized.

- Execute the following command to run the application dependencies in Docker without building the web application container:

```bash
# runs the application locally with only dependencies executing in docker
$ make develop-locally

# entrypoint for web application
$ make dev
```

### Why both Redis and Mongo❓

The design behind making the session management, Redis based, instead of Mongo based, came down to understanding that constant reads and writes to a database for cookie management were redundant and ineffective. The focus was to leave the persistent data in Mongo and less important session-based data in Redis.

---

## 🔒 Environment Configuration

As mentioned before, this application leverages the [Viper](https://github.com/spf13/viper) module, which can read in configuration variables from the [config.yaml](https://github.com/msanvarov/gin-rest-prisma-boilerplate/blob/master/config.yaml) file.

This is a breakdown of the variables:

**server.env** - the application environment it will be executing in, either in development, production, or testing. Options: `dev`, `test`, or `prod`.

**server.port** - the default port to expose the application to.

**session.name** - the name of the session for redis.

**redis.idle_connections** - the number of idle connections redis should support (default is 10).

**redis.network_type** - redis network type, default is "tcp" but "udp" is also supported.

**redis.address** - the URL to redis endpoint.

**redis.secret_key** - secret key to the redis store.

**redis.password** - redis password for authentication.

---

## 📦 Choosing between Dep and Go Modules

**One can choose to use Dep over Go Modules as their preferred package manager for Golang.**

- Dep is a package manager for Go. It aids in managing packages for any Go application.

Getting stated with Dep:

On Mac:

```bash
$ brew install dep
$ brew upgrade dep
```

Other Platforms:

```bash
# downloads dep package manager
$ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```

- Navigate to the project directory, initialize Dep by running `dep init` then remove the `go.mod` and `go.sum` files.

```bash
# initializing dep
$ dep init

# removing the Go Modules files
$ rm go.*
```

> **Note: On Windows, please use Git Bash or WSL where curl is included by default.**

---

## 🧪 Testing

Depending on where the development is occurring; in docker or not, tests can be executed through the Docker shell or locally.

```bash
$ go test -v ./tests/*
```

---

## 📝 License

GRPB is [MIT licensed](https://github.com/msanvarov/gin-rest-prisma-boilerplate/blob/master/LICENSE).

[Author: Sal Anvarov](https://msanvarov.github.io/personal-portfolio/).
