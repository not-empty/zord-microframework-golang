package user

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

func Schema() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Unique(),
		field.String("name").NotEmpty(),
		field.String("email").Unique().NotEmpty(),
		field.String("password").NotEmpty(),
	}
}

func Relations(cartType interface{}) []ent.Edge {
	return []ent.Edge{
		edge.To("cart", cartType),
	}
}
