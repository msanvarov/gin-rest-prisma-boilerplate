package forms

// RegistrationForm structure for defining the type of payload to valid against when registering.
type RegistrationForm struct {
	Name     string `json:"name" binding:"required,max=60"`
	Username string `json:"username" binding:"required,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginForm structure for defining the type of payload to valid against when logging in.
type LoginForm struct {
	Username string `json:"username" binding:"required,max=20"`
	Password string `json:"password" binding:"required"`
}
