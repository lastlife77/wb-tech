package auth

type LoginAuth struct {
}

func (*LoginAuth) SignIn(login string, password string) (bool, error) {
	// some logic

	return true, nil
}
