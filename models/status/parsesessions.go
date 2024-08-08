package models

import models "github.com/SymbioSix/ProgressieAPI/models/auth"

type ParseSessionsForAuth struct {
	AccessToken  string                  `json:"access_token"`
	RefreshToken string                  `json:"refresh_token"`
	ExpiredAt    int64                   `json:"expired_at"`
	ExpiresIn    int                     `json:"expires_in"`
	TokenType    string                  `json:"token_type"`
	Data         models.UserRoleResponse `json:"data"`
}
