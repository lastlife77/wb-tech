package main

import (
	"github.com/lastlife77/wb-tech/l1/1.21/auth"
	"github.com/lastlife77/wb-tech/l1/1.21/service"
)

func main() {
	authService := &service.AuthService{}

	loginAuth := &auth.LoginAuth{}
	authService.SignIn(loginAuth, "login", "password")

	git := &auth.GitAdapter{
		Git: &auth.GitAuth{},
	}
	authService.SignIn(git, "login", "")
}
