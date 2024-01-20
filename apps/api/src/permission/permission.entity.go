package permission

import (
	"time"

	role "api/src/role"
	user "api/src/user"

	"github.com/uptrace/bun"
)

type Permission struct {
	bun.BaseModel `bun:"table:permissions"`

	ID          int64  `bun:",nullzero,pk,autoincrement"`
	Name        string `bun:",unique"`
	DisplayName string
	Description bool
	Users       []user.User `bun:"m2m:user_permissions,join:Permission=User"`
	Roles       []role.Role `bun:"m2m:role_permissions,join:Permission=Role"`
	CreatedAt   time.Time   `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt   time.Time   `bun:",nullzero,notnull,default:current_timestamp"`
}

func Init(db *bun.DB) {
	// Register many to many model so bun can better recognize m2m relation.
	// This should be done before you use the model for the first time.
	db.RegisterModel((*UserPermission)(nil))
	db.RegisterModel((*RolePermission)(nil))
}

type UserPermission struct {
	PermissionId int64
	Permission   *Permission `bun:"rel:belongs-to,join:permission_id=permission.id"`
	UserId       int64
	User         *user.User `bun:"rel:belongs-to,join:user_id=user.id"`
}

type RolePermission struct {
	PermissionId int64
	Permission   *Permission `bun:"rel:belongs-to,join:permission_id=permission.id"`
	RoleId       int64
	Role         *role.Role `bun:"rel:belongs-to,join:role_id=role.id"`
}
