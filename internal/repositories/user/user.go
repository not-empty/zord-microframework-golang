package user

import (
	"context"
	"fmt"
	"go-skeleton/internal/application/domain/user"
	"go-skeleton/pkg/ent"
	"go-skeleton/pkg/idCreator"
)

type Repository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) *Repository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) Create(ctx context.Context, newUser user.User) (error, *user.User) {
	save, err := r.client.User.
		Create().
		SetID(idCreator.NewIdCreator().Create()).
		SetEmail(newUser.Email).
		SetName(newUser.Name).
		SetPassword(newUser.Password).
		AddCart().
		Save(ctx)

	fmt.Println(save, err)

	if err != nil {
		return err, nil
	}

	return nil, &user.User{
		Id:       save.ID,
		Name:     save.Name,
		Email:    save.Email,
		Password: save.Password,
	}
}
