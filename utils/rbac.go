package utils

import (
	"net/http"

	"github.com/casbin/casbin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// IBasicAuthorizer interface.
type IBasicAuthorizer interface {
	GetRoleName(c *gin.Context) interface{}
	CheckPermission(c *gin.Context) bool
	RequirePermission(c *gin.Context)
}

// BasicAuthorizer stores the casbin enforcer
type BasicAuthorizer struct {
	Enforcer *casbin.Enforcer
}

// GetRoleName gets the user name from the request.
// Currently, only HTTP basic authentication is supported
func (BasicAuthorizer) GetRoleName(c *gin.Context) interface{} {
	session := sessions.Default(c)
	role := session.Get("role")
	if role != nil && role != "" {
		return role
	}
	return "ANON"
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *BasicAuthorizer) CheckPermission(c *gin.Context) bool {
	role := a.GetRoleName(c)
	method := c.Request.Method
	path := c.Request.URL.Path
	return a.Enforcer.Enforce(role, path, method)
}

// RequirePermission returns the 403 Forbidden to the client
func (BasicAuthorizer) RequirePermission(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}
