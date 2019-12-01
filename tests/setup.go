package tests

import (
	"github.com/msanvarov/gin-rest-prisma-boilerplate/config"
	"github.com/msanvarov/gin-rest-prisma-boilerplate/db"
	"github.com/msanvarov/gin-rest-prisma-boilerplate/forms"
	"github.com/msanvarov/gin-rest-prisma-boilerplate/router"
	"github.com/spf13/viper"
)

var (
	cookie              string
	routing             = router.Router(TestConfiguration("../config"))
	client              = db.GetDB()
	registrationPayload = forms.RegistrationForm{
		Name:     "test",
		Username: "test",
		Email:    "test@test.com",
		Password: "test1234",
	}
	loginPayload = forms.LoginForm{Username: "test", Password: "test1234"}
)

// TestConfiguration combines the config functions to auto-configure the router
func TestConfiguration(fileName string) *viper.Viper {
	config.Configure(fileName)
	return config.GetConfiguration()
}
