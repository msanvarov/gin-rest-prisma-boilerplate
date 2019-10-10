
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
  
#### Docker, Mongo, and Redis 🐳  
  
-   Please make sure to have docker desktop setup on any preferred operating system to quickly compose the required dependencies. Then follow the docker procedure outlined below.

- Please make sure to have MongoDB locally, or utilize Mongo on the cloud by configuring a cluster in atlas. Then grab the connection string and modify the following [line](https://github.com/msanvarov/gin-rest-prisma-boilerplate/blob/master/docker-compose.yml#L17).  
  
- Since this boilerplate comes with Prisma, Prisma must be setup and configured. A detailed Prisma setup tutorial can be found [here](https://www.prisma.io/docs/get-started/01-setting-up-prisma-existing-database-GO-g003/).  

- Prisma is required for session management. It can be setup using `choco`, `brew` or manually using `wget` or `curl`. Please see the following documentation on [setup instructions](https://redis.io/topics/quickstart).

**Note**: Docker Desktop comes free on both Mac and Windows, but it only works with Windows 10 Pro. A workaround is to get  [Docker Toolbox](https://docs.docker.com/toolbox/toolbox_install_windows/)  which will bypass the Windows 10 Pro prerequisite by executing in a VM.
  
### Dep 📦   

- Dep is a package manager for Go. It aids in managing packages for any golang application. To get dep, please type in the following command: `$ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh` 
  
**Note: If on Windows, please use Git Bash or WSL where curl is included by default**   

---  
  
### 🚀 Deployment  
  
- If needed, replace the existing config variables in the [env-config.yaml](https://github.com/msanvarov/gin-rest-prisma-boilerplate/blob/master/env-config.yaml) file.  
	- Please change the `server.env : "test"` to `server.env : "dev"` for better logging.  
  
- Install project dependencies using `dep ensure`  
  
- Execute the following commands in-app directory:  

``` bash
# creates and loads the docker container with required configuration  
$ docker-compose up -d
# starts the prisma server  
$ cd prisma && prisma deploy
```

**Note: Please make sure redis is deployed and running on localhost before starting the go server**
  
- The following command will set up the project for you (building the Docker images, starting docker-compose stack).   
The web application and Prisma will be exposed to http://localhost:9000 and http://localhost:4466 respectively. 
  
### 🔒 Environment Configuration  
  
By default, the application comes with a config module that can read every configuration variable from the `env-config` yaml file.  
  
**server.env** - the application environment it will be executing as, either in development, production, or testing. Options: `dev`, `test`, or `prod`.   
  
**server.port** - the default port to expose the application to.
  
**session.name** - the name of session for redis session manager.  
  
**redis.idle_connections** - the number of idle connections redis should support (default is 10).

**redis.network_type** - redis network type, default is "tcp" but "udp" is also supported.
  
**redis.address** - the URL to the main redis endpoint.

**redis.secret_key** - secret key for redis store.

**redis.password** - redis password for authentication.
  
---  
  
### ✅ Testing  

```bash
# integration tests
$ go test -v ./tests/*```  
```

---  

### 👥  Support

Gin can grow thanks to the sponsors and support by backers. If you'd like to join them, please [read more here](https://github.com/gin-gonic/gin).

----------

### License

Gin is [MIT licensed](https://github.com/gin-gonic/gin/blob/master/LICENSE).

[Author](https://msanvarov.github.io/personal-portfolio/)