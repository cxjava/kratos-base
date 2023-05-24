package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

func TimeFields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Immutable().Comment("创建时间"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
		field.Int64("created_by").Default(0).Comment("创建者"),
		field.Int64("updated_by").Default(0).Comment("更新者"),
		field.Bool("is_deleted").Optional().Nillable().Comment("是否删除"),
		field.Time("deleted_at").Optional().Nillable().Comment("删除时间"),
		field.Int64("deleted_by").Optional().Nillable().Comment("删除者"),
	}
}

func IdFields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
	}
}

func VersionField() []ent.Field {
	return []ent.Field{
		field.Int64("version").
			UpdateDefault(func() int64 {
				return time.Now().UnixNano()
			}).
			DefaultFunc(func() int64 {
				return time.Now().UnixNano()
			}).Comment("Unix time of when the latest update occurred"),
	}
}
