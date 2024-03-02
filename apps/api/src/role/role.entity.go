package role

import (
	"time"

	"github.com/uptrace/bun"
)

type Role struct {
	bun.BaseModel `bun:"table:roles"`

	ID          int64  `bun:",nullzero,pk,autoincrement" json:"id"`
	Name        string `bun:",unique" validate:"required,max=40" json:"name"`
	DisplayName string `validate:"required,max=40" json:"display_name"`
	Description string	 `validate:"max=255" json:"description"`
	CreatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`
}
