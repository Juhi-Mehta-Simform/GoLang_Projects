package controller

import (
	"github.com/gin-gonic/gin"
	"go/Gin/service"
)

var (
	loginService = service.NewLoginService()
	jwtService   = service.NewJWTService()
	credential   = Credentials{
		Username: "Juhi",
		Password: "juhi0604",
	}
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTservice
}

type Credentials struct {
	Username string
	Password string
}

func NewLoginController(loginService service.LoginService, jWtService service.JWTservice) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	err := ctx.ShouldBind(&credential)
	if err != nil {
		panic(err)
	}
	isAuthenticated := controller.loginService.Login(credential.Username, credential.Password)
	if isAuthenticated {
		return controller.jwtService.GenerateToken(credential.Username, true)
	}
	return ""
}
