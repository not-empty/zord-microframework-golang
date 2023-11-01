package auth

import (
	"errors"
	"go-skeleton/internal/application/domain/auth"
)

type Request struct {
	Token auth.Token
	Err   error
}

func NewRequest(auth auth.Token) Request {
	return Request{
		Token: auth,
	}
}

func (r *Request) Validate() error {
	if err := r.tokenCreateRule(); err != nil {
		return err
	}
	return nil
}

func (r *Request) tokenCreateRule() error {
	if r.Token.Token == "" {
		return errors.New("Invalid Token")
	}
	if r.Token.Secret == "" {
		return errors.New("Invalid Secret")
	}
	if r.Token.Context == "" {
		return errors.New("Invalid Context")
	}
	return nil
}
