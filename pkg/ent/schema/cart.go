package schema

import (
	"entgo.io/ent"
	"go-skeleton/internal/application/domain/cart"
)

// Cart holds the schema definition for the Cart entity.
type Cart struct {
	ent.Schema
}

// Fields of the Cart.
func (Cart) Fields() []ent.Field {
	return cart.Schema()
}

// Edges of the Cart.
func (Cart) Edges() []ent.Edge {
	return cart.Relations()
}
