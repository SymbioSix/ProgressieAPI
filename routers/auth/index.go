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

func (ar *AuthRouter) AuthRoutes(rg *fiber.Group) {
	router := rg.Group("auth")

	router.Post("/loginemailpassword", ar.authRouter.SignInWithEmailPassword)
	router.Post("/signupemailpassword", ar.authRouter.SignUpWithEmailPassword)
	router.Post("/signout", ar.authRouter.SignOut)
}
