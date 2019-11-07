<p align="center">  
  <a href="http://golang.org" target="blank"><img src="https://cacophony.org.nz/sites/default/files/gopher.png" width="200" alt="Nest Logo" /></a>  
  <a href="https://gin-gonic.com/" target="blank"><img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" width="200" alt="Nest Logo" /></a>  
</p>  
  
Go is an open-source programming language that makes it easy to build simple, reliable, and efficient software. Gin is a web framework for Go to help accelerate web development.

[![GoDoc](https://godoc.org/github.com/gin-gonic/gin?status.svg)](https://godoc.org/github.com/gin-gonic/gin)
[![Build Status](https://travis-ci.org/msanvarov/gin-rest-prisma-boilerplate.svg?branch=master)](https://travis-ci.org/msanvarov/gin-rest-prisma-boilerplate)

### 📚 Description

This boilerplate is made to leverage the Gin framework and quickly prototype backend applications. It comes with database, logging, security, and authentication features out of the box.

---

### 🍬 Features

- Based on [Gin](https://github.com/gin-gonic/gin).

- [Prisma ORM](https://www.prisma.io/) for Mongo.

- [Gin Sessions](https://github.com/gin-contrib/sessions) for Redis.

- [Gin Authz](https://github.com/gin-contrib/authz) for RBAC management. Utilizes [Casbin](https://github.com/casbin/casbin) in the backend.

- [Viper](https://github.com/spf13/viper) for configurations.

---

### 🌱 Project Structure

A quick synopsis of the folder structure in this project.

```txt
.
├── Dockerfile
├── Gopkg.lock        // dep files
├── Gopkg.toml        // dep files
├── Makefile
├── README.md
├── config
│   └── config.go     // viper module to read yaml file
├── config.yaml       // web app configuration
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
│   └── user.go       // payloads definitions
├── go.mod
├── go.sum
├── main.go
├── model.conf        // casbin configs
├── policy.csv        // casbin configs
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
    ├── error.go      // gin global error handler
    ├── passwords.go  // password hashing util
    └── rbac.go       // gin authz configuration
```

---

### 🛠️ Prerequisites

#### 🐳 Docker

- Please make sure to have `Docker Desktop` setup on any preferred operating system to quickly compose the required dependencies. Then follow the docker procedure outlined below.

- To get familiar with Prisma. A detailed Prisma setup tutorial can be found [here](https://www.prisma.io/docs/get-started/01-setting-up-prisma-new-database-GO-g002/).

- Redis configuration can be found in the [configuration yaml file](https://github.com/msanvarov/gin-rest-prisma-boilerplate/blob/master/config.yaml#L10-L14).

**Note: Docker Desktop comes free on both Mac and Windows, but when on Windows, it only supports Windows 10 Pro. A workaround is to get [Docker Toolbox](https://docs.docker.com/toolbox/toolbox_install_windows/) which will bypass the Windows 10 Pro prerequisite by executing Docker in a VM.**

#### 🧰 Node

- The [Prisma CLI](https://www.prisma.io/docs/prisma-cli-and-configuration/using-the-prisma-cli-alx4/) is essential for streamlining workflows, and has to be downloaded using NPM. Which means that [Node](https://nodejs.org/en/download/) is a requirement.

- Once Node is installed, the following command will proceed in downloading the Prisma CLI:

```bash
// prisma cli
$ npm install -g prisma
```

---

### 🚀 Deployment

- If needed, replace the existing config variables in the [config.yaml](https://github.com/msanvarov/gin-rest-prisma-boilerplate/blob/master/config.yaml) file with preferred configuration.

  - Changing the `server.env : "test"` to `server.env : "dev"` yields better logging that can come of use when developing.

#### 🐳 Developing Inside Docker

To achieve the full Docker experience, VSCode permits the development of source code to happen exclusively in a Docker container. For more information on how to achieve this, [please read the following documentation](https://code.visualstudio.com/docs/remote/containers).

- To bring up the web application and all of its dependencies in Docker execute one of the following commands:

```bash
# runs in detached mode
$ docker-compose up -d

# without detaching
$ docker-compose up
```

- The following command will set up the project for you (creating the Docker containers, and starting the web application).  
  The web application and Prisma will then be exposed to http://localhost:9000 and http://localhost:4466 respectively.

**Please beware, each time a change to the code occurs, the container must be rebuilt.**

#### 🏡 Developing Locally Outside of Docker

Developing the web application locally can be opted for over developing in a container. In this circumstance, the application dependencies such as Prisma, Mongo, and Redis will still require Docker to run. Mainly the limitation is that Prisma can't be set up locally like Mongo and Redis.

- Execute the following command to run the application dependencies in Docker while leaving the source code to be developed locally:

```bash
# runs the application locally with only dependencies executing in docker
$ make ensure-deps

# start the binary
$ ./gin-rest-prisma-boilerplate
```

---

### 🔒 Environment Configuration

By default, the application leverages the [viper](https://github.com/spf13/viper) module, which can read every configuration variable from the [`config.yaml`](https://github.com/msanvarov/gin-rest-prisma-boilerplate/blob/master/config.yaml) file.

**server.env** - the application environment it will be executing as, either in development, production, or testing. Options: `dev`, `test`, or `prod`.

**server.port** - the default port to expose the application to.

**session.name** - the name of the session for redis.

**redis.idle_connections** - the number of idle connections redis should support (default is 10).

**redis.network_type** - redis network type, default is "tcp" but "udp" is also supported.

**redis.address** - the URL to the main redis endpoint.

**redis.secret_key** - secret key for redis store.

**redis.password** - redis password for authentication.

---

### 🏗️ Choosing between Dep and Go Modules

**There is an option to choose between Dep or Go modules as the preferred package manager for Golang. By default, Go modules are utilized for their dependency management.**

#### 📦 Dep over Go Modules:

- Dep is a package manager for Go. It aids in managing packages for any Go application. To get dep, please type in the following command:

  `$ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh`

- Removing Go Module relateed files involves simply deleting the `go.mod` and `go.sum` files, and run `dep ensure` to verify vendor files are in order.

**Note: If on Windows, please use Git Bash or WSL where curl is included by default.**

---

### ❓Why both Redis and Mongo?

The design behind making the session management, Redis based, instead of Mongo based, came down to understanding that constant reads and writes to a database for cookie management were redundant. The focus was to leave the persistent data in Mongo and less important session-based data in Redis.

---

### ✅ Testing

Depending on where the development is occuring; docker or not, tests can be executed through the Docker shell or locally.

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
