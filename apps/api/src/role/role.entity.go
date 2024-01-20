package permission

import (
	"time"

	"github.com/uptrace/bun"
)

type Role struct {
	bun.BaseModel `bun:"table:roles"`

	ID          int64  `bun:",nullzero,pk,autoincrement"`
	Name        string `bun:",unique"`
	DisplayName string
	Description bool
	CreatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
