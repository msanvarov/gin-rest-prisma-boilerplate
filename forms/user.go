package forms

// RegistrationForm structure for defining the type of payload to valid against when registering.
type RegistrationForm struct {
	Name     string `json:"name" binding:"required,max=60" example:"account name"`
	Username string `json:"username" binding:"required,max=20" example:"account_mock"`
	Email    string `json:"email" binding:"required,email" example:"account@mock.com"`
	Password string `json:"password" binding:"required" example:"password"`
}

// LoginForm structure for defining the type of payload to valid against when logging in.
type LoginForm struct {
	Username string `json:"username" binding:"required,max=20" example:"account_mock"`
	Password string `json:"password" binding:"required" example:"password"`
}
