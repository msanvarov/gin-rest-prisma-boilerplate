package auth

import (
	"context"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/gin-rest-prisma-boilerplate/db"
	"github.com/gin-rest-prisma-boilerplate/forms"
	"github.com/gin-rest-prisma-boilerplate/prisma-client"
	"github.com/gin-rest-prisma-boilerplate/utils"
	"log"
	"net/http"
	"strings"
)

type AuthenticationController struct{}

var (
	client   = db.DB()
	contextB = context.Background()
)

func (AuthenticationController) Register(c *gin.Context) {
	var registrationPayload forms.RegistrationForm
	if validationErr := c.BindJSON(&registrationPayload); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"operation": "Tried to bind request body to expected JSON fields.", "error": validationErr.Error(), "status": http.StatusBadRequest})
		return
	}
	if hashedPass, hashErr := utils.EncryptPassword(registrationPayload.Password); hashErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"operation": "Tried to hash payload password.", "error": hashErr.Error(), "status": http.StatusInternalServerError})
		return
	} else {
		registrationPayload.Password = hashedPass
	}

	user, err := client.CreateUser(prisma.UserCreateInput{
		Email: registrationPayload.Email,
		Name: registrationPayload.Name,
		Username: registrationPayload.Username,
		Password: registrationPayload.Password,
	}, ).Exec(contextB)

	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authH := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(authH) != 2 || authH[0] != "Basic" {
			GlobalErrorHandler(c, http.StatusUnauthorized, "Unauthorized")
			return
		}
		payload, _ := base64.StdEncoding.DecodeString(authH[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !authenticateUser(pair[0], pair[1]) {
			GlobalErrorHandler(c,http.StatusUnauthorized, "Unauthorized")
			return
		}
		c.Next()
	}
}

func authenticateUser(username, password string) bool {
	if user, err := client.User(prisma.UserWhereUniqueInput{Username: &username}).Exec(contextB); err != nil {
		log.Fatal(err)
		return false
	} else {
		return utils.CheckPassword(password, user.Password)
	}
}
