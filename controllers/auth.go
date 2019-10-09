package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gin-rest-prisma-boilerplate/forms"
	"github.com/gin-rest-prisma-boilerplate/prisma-client"
	"github.com/gin-rest-prisma-boilerplate/utils"
	"log"
	"net/http"
)

type AuthenticationController struct{}

var (
	client   = prisma.New(nil)
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
