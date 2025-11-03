package auth

type GitAuth struct {
}

func (*GitAuth) SignIn(login string) (bool, error) {
	// some logic

	return true, nil
}

type GitAdapter struct {
	Git *GitAuth
}

func (a *GitAdapter) SignIn(login string, password string) (bool, error) {

	return a.Git.SignIn(login)
}
