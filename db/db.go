package db

import (
	"sync"

	"github.com/msanvarov/gin-rest-prisma-boilerplate/prisma-client"
)

var (
	client *prisma.Client
	once   sync.Once
)

func DB() *prisma.Client {
	once.Do(func() {
		client = prisma.New(nil)
	})
	return client
}
