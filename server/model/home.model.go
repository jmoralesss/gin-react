package model

type LoginAttempt struct {
	Email    string
	Password string
}

type RegistrationAttempt struct {
	Email    string
	Password string
}

type RegistrationResponse struct {
	Email   string
	Success bool
}
