package models

type SignUpRequest struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ConfirmationSignup struct {
	Type       string `query:"type"`
	TokenHash  string `query:"token_hash"`
	RedirectTo string `query:"redirect_to"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required"`
}

type UpdatePasswordAfterForgotPassword struct {
	NewPassword string `json:"new_password" binding:"required"`
}

type FailedAuth struct {
	Type string `query:"type"`
}
