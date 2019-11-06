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

### 🛠️ Prerequisites

#### 🐳 Docker

- Please make sure to have `Docker Desktop` setup on any preferred operating system to quickly compose the required dependencies. Then follow the docker procedure outlined below.

- To get familiar with Prisma. A detailed Prisma setup tutorial can be found [here](https://www.prisma.io/docs/get-started/01-setting-up-prisma-new-database-GO-g002/).

- Redis configuration can be found in the [configuration yaml file](https://github.com/msanvarov/gin-rest-prisma-boilerplate/blob/master/config.yaml#L10-L14).

**Note: Docker Desktop comes free on both Mac and Windows, but when on Windows, it only supports Windows 10 Pro. A workaround is to get [Docker Toolbox](https://docs.docker.com/toolbox/toolbox_install_windows/) which will bypass the Windows 10 Pro prerequisite by executing Docker in a VM.**

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

#### 🏡 Developing Outside of Docker

Developing the web application locally can be opted for over developing in a container. In this circumstance, the application dependencies such as Prisma, Mongo, and Redis will still require Docker to run. Mainly the limitation is that Prisma can't be set up locally like Mongo and Redis.

- Execute one of the following commands to run the application dependencies in Docker while leaving the source code to be developed locally:

```bash
# runs in detached mode
$ docker-compose -f docker-compose.override.yml up -d
# without detaching
$ docker-compose -f docker-compose.override.yml up
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

The design behind making the session management, Redis based, instead of Mongo based, came down to understanding that constant reads and writes to a database for cookie management were redundant. The focus was to leave the persistent data in Mongo and less important session-based data in Redis. Not to mention the performance benefits that Redis provides over Mongo based queries.

---

### ✅ Testing

Depending on which setup was selected, tests can be executed through accessing the Docker shell or run locally.

#### ✅ + ☁️Test Execution When Developing in Docker

```bash
# docker execution
$ docker exec -it gin-rest-prisma-boilerplate_app_1 go test -v ./tests/*
```

#### ✅ + 💻Test Execution When Developing Locally

```bash
# non-docker execution
$ go test -v ./tests/*
```

---

### License

Gin is [MIT licensed](https://github.com/gin-gonic/gin/blob/master/LICENSE).

[Author](https://msanvarov.github.io/personal-portfolio/)
