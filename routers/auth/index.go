package routers

import (
	"github.com/SymbioSix/ProgressieAPI/services/auth"
	"github.com/gofiber/fiber/v3"
)

type AuthRouter struct {
	authRouter auth.AuthController
}

func NewRouteAuthController(authRouter auth.AuthController) AuthRouter {
	return AuthRouter{authRouter}
}

func (ar *AuthRouter) AuthRoutes(rg fiber.Router) {
	router := rg.Group("auth")

	router.Post("/signin-email-password", ar.authRouter.SignInWithEmailPassword)
	router.Post("/signup-email-password", ar.authRouter.SignUpWithEmailPassword)
	router.Post("/signout", ar.authRouter.SignOut)
	router.Get("/verify-signup", ar.authRouter.VerifySignUp)
	router.Post("/send-forgot-password-email", ar.authRouter.SendForgotPasswordEmail)
	router.Get("/verify-password-recovery", ar.authRouter.VerifyPasswordRecovery)
	router.Put("/update-user-password", ar.authRouter.UpdateUserPassword)
	router.Get("/failed", ar.authRouter.FailedAuthService)
}
