package service

import "fmt"

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

func NewLoginService() LoginService {
	return &loginService{
		authorizedUsername: "Juhi",
		authorizedPassword: "juhi0604",
	}
}

func (service *loginService) Login(username string, password string) bool {
	fmt.Println(username, password)
	return service.authorizedUsername == username &&
		service.authorizedPassword == password
}
