package user

import (
	"context"
)

type User struct {
	Id       string
	Name     string
	Email    string
	Password string
}

type Repo interface {
	Create(context.Context, User) (error, *User)
}
