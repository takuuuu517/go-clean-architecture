package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").Comment("名前"),
		field.String("last_name").Comment("姓"),
		field.String("email").Unique().Comment("メールアドレス"),
		field.Time("created_at").Default(time.Now).Comment("作成日時"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("更新日時"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
