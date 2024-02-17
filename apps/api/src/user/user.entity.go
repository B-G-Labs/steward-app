package user

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID        int64     `bun:",nullzero,pk,autoincrement" json:"id"`
	Name      string    `bun:",unique" json:"name" validate:"required,alphanum,min=4,max=20"`
	Password  string    `json:"password" validate:"required,min=8,max=20"`
	Active    bool      `json:"active" validate:"boolean" default:"true"`
	CreatedAt time.Time `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" bun:",nullzero,notnull,default:current_timestamp"`
}
