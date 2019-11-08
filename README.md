<p align="center">  
  <a href="http://golang.org" target="blank"><img src="https://cacophony.org.nz/sites/default/files/gopher.png" width="200" alt="Nest Logo" /></a>  
  <a href="https://gin-gonic.com/" target="blank"><img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" width="200" alt="Nest Logo" /></a>  
</p>  
  
Golang, is a statically typed, compiled programming language designed at Google that aids in building simple, reliable, and efficient software. Gin is a full-featured web framework for Go that achieves outstanding performance.

[![GoDoc](https://godoc.org/github.com/gin-gonic/gin?status.svg)](https://godoc.org/github.com/gin-gonic/gin)
[![codebeat badge](https://codebeat.co/badges/c9d048b7-5d3b-416a-ab64-b9f510f947ed)](https://codebeat.co/projects/github-com-msanvarov-gin-rest-prisma-boilerplate-master)
[![Build Status](https://travis-ci.org/msanvarov/gin-rest-prisma-boilerplate.svg?branch=master)](https://travis-ci.org/msanvarov/gin-rest-prisma-boilerplate)

### 📚 Description

This boilerplate leverages the Gin framework to quickly prototype backend applications. It comes with database, logging, security, and authentication features out of the box.

---

### 🍬 Features

- Based on [Gin](https://github.com/gin-gonic/gin).

- [Prisma ORM](https://www.prisma.io/) for Mongo. But can support MYSQL/PostgreSQL and Amazon Aurora.

- [Gin Sessions](https://github.com/gin-contrib/sessions) for Redis.

- [Gin Authz](https://github.com/gin-contrib/authz) for role based access management. Internally, utilizing the powerful authentication library [Casbin](https://github.com/casbin/casbin).

- [Viper](https://github.com/spf13/viper) for working with yaml configurations.

---

### 🌱 Project Structure

A quick synopsis of the folder structure.

```txt
.
├── Dockerfile
├── Makefile
├── README.md
├── config
│   └── config.go     // viper module to read yaml file
├── config.yaml       // yaml file for web app configuration
├── controllers
│   └── auth.go       // authentication controller
├── db
│   └── db.go         // prisma client instance
├── docker
│   ├── replace.awk
│   ├── run.sh
│   ├── to-docker.txt
│   ├── to-local.txt
│   └── wait-for-it.sh
├── docker-compose.override.yml
├── docker-compose.yml
├── forms
│   └── user.go       // payload definitions
├── go.mod
├── go.sum
├── main.go
├── model.conf        // casbin model configuration
├── policy.csv        // casbin policy configuration
├── prisma
│   ├── datamodel.prisma
│   └── prisma.yml    // prisma configs
├── prisma-client
│   └── prisma.go     // generated prisma client
├── renovate.json
├── router
│   └── router.go     // application router
├── tests
│   └── auth_test.go
└── utils
    ├── error.go      // global http error handler
    ├── passwords.go  // password hashing utility
    └── rbac.go       // gin authz middleware configuration
```

---

### 🛠️ Prerequisites

#### 🐳 Docker

- Please make sure to have `Docker Desktop` operational on the preferred operating system of choice to quickly get started. Then follow the docker procedure outlined below.

- To get familiar with Prisma, a detailed guide on setting Prisma up can be found [here](https://www.prisma.io/docs/get-started/01-setting-up-prisma-new-database-GO-g002/).

- To configure Redis, edit the following [yaml file](https://github.com/msanvarov/gin-rest-prisma-boilerplate/blob/master/config.yaml#L10-L14).

**Note: Despite the fact that Docker Desktop comes free for both Mac and Windows, it only supports the Pro edition of Windows 10. A common workaround is to get [Docker Toolbox](https://docs.docker.com/toolbox/toolbox_install_windows/) which will bypass the Windows 10 Pro restriction by executing Docker in a VM.**

#### 🧰 Node

- The [Prisma CLI](https://www.prisma.io/docs/prisma-cli-and-configuration/using-the-prisma-cli-alx4/) is essential for streamlining workflows for managing and deploying Prisma services. The CLI can to be downloaded using `npm`; which makes [Node](https://nodejs.org/en/download/) a requirement.

- On acquisition of NodeJS, the following command will download the Prisma CLI:

```bash
// prisma cli
$ npm install -g prisma
```

---

### 🚀 Deployment

- If need be, replace the existing configuration variables in the [config.yaml](https://github.com/msanvarov/gin-rest-prisma-boilerplate/blob/master/config.yaml) file with the preferred configuration values.

  - Changing the `server.env : "test"` to `server.env : "dev"` yields better logging that can come of use when developing.

#### 🐳 Developing Inside Docker

To achieve the full Docker experience, VSCode permits the development of source code to happen exclusively in a Docker container. For more information on how this works, [please read the following documentation](https://code.visualstudio.com/docs/remote/containers).

In attempts to embrace best programming practices, this boilerplate comes with an option to enable development in docker.

- To bring up the web application and all of its dependencies in Docker execute one of the following commands:

```bash
# runs in detached mode
$ docker-compose up -d

# without detaching
$ docker-compose up
```

- The web application and Prisma will then be exposed to http://localhost:9000 and http://localhost:4466 respectively.

**Please beware, each time a change to the code base occurs, the main container must be rebuilt.**

#### ⛲ Developing Locally Outside of Docker

Developing the web application locally can be chosen over developing in a container. In this case, the application dependencies such as Prisma, Mongo, and Redis will still require Docker to run. Mainly because Prisma can't be set up locally like Mongo and Redis.

- Execute the following command to run the application dependencies in Docker without building the web application container:

```bash
# runs the application locally with only dependencies executing in docker
$ make ensure-deps

# entrypoint for web application
$ go run main.go
```

---

### 🔒 Environment Configuration

As mentioned before, this application leverages the [Viper](https://github.com/spf13/viper) module, which can read in configuration variables from the [`config.yaml`](https://github.com/msanvarov/gin-rest-prisma-boilerplate/blob/master/config.yaml) file.

**server.env** - the application environment it will be executing in, either in development, production, or testing. Options: `dev`, `test`, or `prod`.

**server.port** - the default port to expose the application to.

**session.name** - the name of the session for redis.

**redis.idle_connections** - the number of idle connections redis should support (default is 10).

**redis.network_type** - redis network type, default is "tcp" but "udp" is also supported.

**redis.address** - the URL to redis endpoint.

**redis.secret_key** - secret key to the redis store.

**redis.password** - redis password for authentication.

---

### 📦 Choosing between Dep and Go Modules

**One can choose to use Dep over Go Modules as their preferred package manager for Golang.**

- Dep is a package manager for Go. It aids in managing packages for any Go application. To get stated with dep, please enter in the following command:

```bash
# downloads dep package manager
$ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```

- Then initialize dep by running `dep init` and remove the `go.mod` and `go.sum` files.

```bash
# initializing dep
$ dep init

# removing the Go Modules files
$  rm go.mod go.sum
```

**Note: On Windows, please use Git Bash or WSL where curl is included by default.**

---

### ❓Why both Redis and Mongo?

The design behind making the session management, Redis based, instead of Mongo based, came down to understanding that constant reads and writes to a database for cookie management were redundant and ineffective. The focus was to leave the persistent data in Mongo and less important session-based data in Redis.

---

### ✅ Testing

Depending on where the development is occurring; in docker or not, tests can be executed through the Docker shell or locally.

- ☁️ Test Execution When Developing in Docker:

```bash
# docker execution
$ docker exec -it gin-rest-prisma-boilerplate_app_1 go test -v ./tests/*
```

- 💻 Test Execution When Developing Locally:

```bash
# non-docker execution
$ go test -v ./tests/*
```

---

### License

Gin is [MIT licensed](https://github.com/gin-gonic/gin/blob/master/LICENSE).

[Author](https://msanvarov.github.io/personal-portfolio/)
