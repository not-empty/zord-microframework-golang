package user

import userDomain "go-skeleton/internal/application/domain/user"

type Response struct {
	Data userDomain.User
}
