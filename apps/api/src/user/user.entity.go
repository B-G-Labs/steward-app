package user

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID       int64  `bun:",nullzero,pk,autoincrement"`
	Name     string `bun:",unique"`
	Password string
}
