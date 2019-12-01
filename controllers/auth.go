// Package controllers includes a collection of controller
//  structures for executing function logic.
package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/msanvarov/gin-rest-prisma-boilerplate/db"
	"github.com/msanvarov/gin-rest-prisma-boilerplate/forms"
	"github.com/msanvarov/gin-rest-prisma-boilerplate/prisma-client"
	"github.com/msanvarov/gin-rest-prisma-boilerplate/utils"
)

// IAuthenticationController interface.
type IAuthenticationController interface {
	GetSessionData(c *gin.Context)
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

// AuthenticationController for authentication logic.
type AuthenticationController struct{}

var (
	prismaClient = db.GetDB()
	contextB     = context.Background()
)

// GetSessionData method is responsible for retrieving session data once authenticated.
// @summary Fetch session data
// @tags auth
// @produce  json
// @success 200 {string} string	"Fetch Session Data Request Completed"
// @failure 400 {object} utils.HTTPError
// @router /api/v1/session [get]
func (AuthenticationController) GetSessionData(c *gin.Context) {
	session := sessions.Default(c)
	uuid := session.Get("uuid")
	if uuid != nil {
		c.JSON(http.StatusOK, gin.H{"uuid": uuid, "username": session.Get("username"),
			"email": session.Get("email")})
	} else {
		utils.CreateError(c, http.StatusBadRequest,
			"Failed to fetch session data. Make sure to be logged in.")
	}
}

// Register method provides registration logic for when onboarding onto the api.
// @summary Register
// @tags auth
// @produce  json
// @param name body forms.RegistrationForm true "Registration Payload"
// @success 200 {string} string	"Registration Request Completed"
// @failure 400 {object} utils.HTTPError
// @router /api/v1/register [post]
func (AuthenticationController) Register(c *gin.Context) {
	var registrationPayload forms.RegistrationForm
	if validationErr := c.BindJSON(&registrationPayload); validationErr != nil {
		utils.CreateError(c, http.StatusBadRequest, validationErr.Error())
		return
	}
	hashedPass, _ := utils.EncryptPassword(registrationPayload.Password)
	registrationPayload.Password = hashedPass
	user, prismaError := prismaClient.CreateUser(prisma.UserCreateInput{
		Email:    registrationPayload.Email,
		Name:     registrationPayload.Name,
		Username: registrationPayload.Username,
		Password: registrationPayload.Password,
		Role:     prisma.RoleDefault,
	}).Exec(contextB)

	if prismaError != nil {
		log.Print(prismaError)
		utils.CreateError(c, http.StatusNotAcceptable, "Failed to create profile in database. Please try again.")
		return
	}
	// setting session keys
	sessionSaveError := setSessionKeys(c, user)
	if sessionSaveError != nil {
		utils.CreateError(c, http.StatusInternalServerError, sessionSaveError.Error())
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"name":     user.Name,
		"username": user.Username,
		"role":     user.Role,
	})
}

// Login method provides login logic for when signing into the api.
// @summary Login
// @tags auth
// @produce  json
// @param name body forms.LoginForm true "Login Payload"
// @success 200 {object} prisma.User
// @failure 400 {object} utils.HTTPError
// @router /api/v1/login [post]
func (AuthenticationController) Login(c *gin.Context) {
	var loginPayload forms.LoginForm
	if validationErr := c.BindJSON(&loginPayload); validationErr != nil {
		utils.CreateError(c, http.StatusBadRequest, validationErr.Error())
	}

	if user, userLookupError := prismaClient.User(
		prisma.UserWhereUniqueInput{Username: &loginPayload.Username}).Exec(contextB); userLookupError != nil {
		log.Println(userLookupError)
		utils.CreateError(c, http.StatusBadRequest,
			fmt.Sprintf(
				"The profile with the username: %s doesn't exist. Please register before trying to login.",
				loginPayload.Username))
	} else {
		if passwordMatch := utils.CheckPassword(loginPayload.Password, user.Password); passwordMatch != true {
			utils.CreateError(c, http.StatusNotAcceptable, "Password entered is incorrect. Please try again.")
		} else {
			// setting session keys
			sessionSaveError := setSessionKeys(c, user)
			if sessionSaveError != nil {
				utils.CreateError(c, http.StatusInternalServerError, sessionSaveError.Error())
				c.Abort()
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "Logged in.",
				"user":    user,
			})
		}
	}
}

// Logout method provides logic for logging out once authenticated.
// @summary Logout
// @tags auth
// @produce  text/plain
// @success 200 {string} string	"Logout Request Completed"
// @failure 400 {object} utils.HTTPError
// @router /api/v1/logout [post]
func (AuthenticationController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	fmt.Print(session)
	session.Clear()
	if sessionSaveError := session.Save(); sessionSaveError != nil {
		log.Panic(sessionSaveError)
		utils.CreateError(c, http.StatusInternalServerError, "Failed to logout. Please try again.")
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Logged out..."})
}

// SetSessionKeys helper to populate redis session with user keys.
func setSessionKeys(c *gin.Context, user *prisma.User) error {
	session := sessions.Default(c)
	session.Set("uuid", user.ID)
	session.Set("email", user.Email)
	session.Set("username", user.Username)
	session.Set("role", string(user.Role))
	if sessionSaveError := session.Save(); sessionSaveError != nil {
		return sessionSaveError
	}
	return nil
}
