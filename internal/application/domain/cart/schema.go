package cart

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

func Schema() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Unique(),
		field.Time("ModifiedAt"),
	}
}

func Relations() []ent.Edge {
	return nil
}
