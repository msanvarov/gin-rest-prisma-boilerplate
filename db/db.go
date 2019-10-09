package db

import (
	"github.com/gin-rest-prisma-boilerplate/prisma-client"
	"sync"
)

var (
	client *prisma.Client
	once sync.Once)

func DB() *prisma.Client {
	once.Do(func() {
		client = prisma.New(nil)
	})
	return client
}
