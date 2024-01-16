package user

import domain "go-skeleton/internal/application/domain/user"

type Response struct {
    Data   domain.User `json:"data"`
}
