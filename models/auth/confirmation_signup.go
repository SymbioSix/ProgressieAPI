package models

type ConfirmationSignup struct {
	Type       string `query:"type"`
	TokenHash  string `query:"token_hash"`
	RedirectTo string `query:"redirect_to"`
}
