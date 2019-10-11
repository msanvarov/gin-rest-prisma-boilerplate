<p align="center">  
  <a href="http://golang.org" target="blank"><img src="https://cacophony.org.nz/sites/default/files/gopher.png" width="200" alt="Nest Logo" /></a>  
  <a href="http://golang.org" target="blank"><img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" width="200" alt="Nest Logo" /></a>  
</p>  
  
Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. Gin is a web framework written in Go (Golang).  

[![Travis](https://travis-ci.org/msanvarov/gin-rest-prisma-boilerplate.svg?branch=master)](https://travis-ci.org/msanvarov/gin-rest-prisma-boilerplate)
[![GoDoc](https://godoc.org/github.com/gin-gonic/gin?status.svg)](https://godoc.org/github.com/gin-gonic/gin)  

  ### 📚 Description  
  
This boilerplate is made to quickly prototype backend applications. It comes with database, logging, security, and authentication features out of the box.  
  
---  
  
### 🛠️ Prerequisites  

#### Docker 🐳  
  
- Please make sure to have docker desktop setup on any preferred operating system to quickly compose the required dependencies. Then follow the docker procedure outlined below.
  
- A detailed Prisma setup tutorial can be found [here](https://www.prisma.io/docs/get-started/01-setting-up-prisma-existing-database-GO-g003/).  

- Redis configuration can be found in the [configuration yaml file](https://github.com/msanvarov/gin-rest-prisma-boilerplate/blob/master/config.yaml#L10)

**Note: Docker Desktop comes free on both Mac and Windows, but it only works with Windows 10 Pro. A workaround is to get  [Docker Toolbox](https://docs.docker.com/toolbox/toolbox_install_windows/)  which will bypass the Windows 10 Pro prerequisite by executing in a VM.**
  
### Dep 📦   

- Dep is a package manager for Go. It aids in managing packages for any golang application. To get dep, please type in the following command: `$ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh` 
  
**Note: If on Windows, please use Git Bash or WSL where curl is included by default.**   

---  
  
### 🚀 Deployment  
  
- If needed, replace the existing config variables in the [config.yaml](https://github.com/msanvarov/gin-rest-prisma-boilerplate/blob/master/config.yaml) file.  
	- Please change the `server.env : "test"` to `server.env : "dev"` for better logging.  
  
- Install project dependencies using `dep ensure`  
  
- Execute the following commands in-app directory:  

``` bash
# creates and loads the docker container with required configuration  
$ docker-compose up -d
# starts the prisma server  
$ cd prisma && prisma deploy
```

**Note: Please make sure prisma, redis, and mongo are deployed and running on localhost before starting the web server.**
  
- The following command will set up the project for you (building the Docker images, starting the web stack).   
The web application and Prisma will be exposed to http://localhost:9000 and http://localhost:4466 respectively. 

- Since all the dependency are up and running by now, all is left is to start the gin web server by typing in the following command:
`go run main.go`
  
### 🔒 Environment Configuration  
  
By default, the application comes with a config module that can read every configuration variable from the `config` yaml file.  
  
**server.env** - the application environment it will be executing as, either in development, production, or testing. Options: `dev`, `test`, or `prod`.   
  
**server.port** - the default port to expose the application to.
  
**session.name** - the name of session for redis session manager.  
  
**redis.idle_connections** - the number of idle connections redis should support (default is 10).

**redis.network_type** - redis network type, default is "tcp" but "udp" is also supported.
  
**redis.address** - the URL to the main redis endpoint.

**redis.secret_key** - secret key for redis store.

**redis.password** - redis password for authentication.
  
---  

### ❓Why both Redis and Mongo?

The philosophy behind making the session management, Redis based, came down to understanding that constant reads and writes to a database for cookie management were redundant. The focus was to leave the persistent data in Mongo and less important session-based data in Redis.

  
### ✅ Testing  

```bash
# integration tests
$ go test -v ./tests/*```  
```

---  

### 👥  Support

Gin can grow thanks to the sponsors and support by backers. If you'd like to join them, please [read more here](https://github.com/gin-gonic/gin).

---

### License

Gin is [MIT licensed](https://github.com/gin-gonic/gin/blob/master/LICENSE).

[Author](https://msanvarov.github.io/personal-portfolio/)
