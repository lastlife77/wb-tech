package service

type Auth interface {
	SignIn(login string, password string) (bool, error)
}

type AuthService struct {
}

func (AuthService) SignIn(a Auth, login string, password string) (bool, error) {
	return a.SignIn(login, password)
}
