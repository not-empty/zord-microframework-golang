package schema

import (
	"entgo.io/ent"
	"go-skeleton/internal/application/domain/user"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return user.Schema()
}

// Edges of the User.
func (s User) Edges() []ent.Edge {
	return user.Relations(Cart.Type)
}
