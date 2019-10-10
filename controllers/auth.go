package auth

import (
	"context"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-rest-prisma-boilerplate/db"
	"github.com/gin-rest-prisma-boilerplate/forms"
	"github.com/gin-rest-prisma-boilerplate/prisma-client"
	"github.com/gin-rest-prisma-boilerplate/utils"
	"log"
	"net/http"
)

type AuthenticationController struct{}

var (
	client   = db.DB()
	contextB = context.Background()
)

func (AuthenticationController) GetSessionData(c *gin.Context) {
	session := sessions.Default(c)
	uuid := session.Get("uuid")
	if uuid != nil {
		c.JSON(http.StatusOK, gin.H{"uuid": uuid, "username": session.Get("username"),
			"email": session.Get("email")})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to fetch session data. Make sure to be logged on."})
		c.Abort()
		return
	}
}

func (AuthenticationController) Register(c *gin.Context) {
	var registrationPayload forms.RegistrationForm
	if validationErr := c.BindJSON(&registrationPayload); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"operation": "Tried to bind request body to expected JSON fields.",
			"error": validationErr.Error(), "status": http.StatusBadRequest})
		c.Abort()
		return
	}
	if hashedPass, hashErr := utils.EncryptPassword(registrationPayload.Password); hashErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"operation": "Tried to hash payload password.",
			"error": hashErr.Error(), "status": http.StatusInternalServerError})
		c.Abort()
		return
	} else {
		registrationPayload.Password = hashedPass
	}

	user, prismaError := client.CreateUser(prisma.UserCreateInput{
		Email:    registrationPayload.Email,
		Name:     registrationPayload.Name,
		Username: registrationPayload.Username,
		Password: registrationPayload.Password,
	}, ).Exec(contextB)

	if prismaError != nil {
		log.Println(prismaError)
		c.JSON(http.StatusNotAcceptable, gin.H{"message":
		"Failed to create the account. Please try another username."})
		c.Abort()
		return
	}
	// setting session keys
	session := sessions.Default(c)
	session.Set("uuid", user.ID)
	session.Set("email", user.Email)
	session.Set("username", user.Username)

	if sessionError := session.Save(); sessionError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message":
		"Failed to set session keys. Please try to register again."})
		c.Abort()
		log.Println(sessionError)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Registered.",
		"user":    user,
	})
}

func (AuthenticationController) Login(c *gin.Context) {
	var loginPayload forms.LoginForm
	if validationErr := c.BindJSON(&loginPayload); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"operation": "Tried to bind request body to expected JSON fields.",
			"error": validationErr.Error(), "status": http.StatusBadRequest})
		c.Abort()
		return
	}

	if user, err := client.User(
		prisma.UserWhereUniqueInput{Username: &loginPayload.Username}).Exec(contextB); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":
		fmt.Sprintf("The profile with the username: %s doesn't exist. Please register before trying to login.",
			loginPayload.Username)})
		log.Println(err)
	} else {
		if passwordMatch := utils.CheckPassword(loginPayload.Password, user.Password); passwordMatch != true {
			c.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid password details. Please try again."})
		} else {
			// setting session keys
			session := sessions.Default(c)
			session.Set("uuid", user.ID)
			session.Set("email", user.Email)
			session.Set("username", user.Name)
			if sessionError := session.Save(); sessionError != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to logout. Please try again."})
				c.Abort()
				log.Println(sessionError)
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "Logged in.",
				"user":    user,
			})
		}
	}
}

func (AuthenticationController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	if sessionError := session.Save(); sessionError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to logout. Please try again."})
		c.Abort()
		log.Println(sessionError)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Logged out..."})
	}
}
