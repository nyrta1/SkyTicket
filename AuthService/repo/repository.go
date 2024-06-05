package repo

import (
	"SkyTicket/AuthService/entity"
	"database/sql"

	"golang.org/x/net/context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, u *entity.User) (*entity.User, error)
	GetUser(ctx context.Context, id int64) (*entity.User, error)
	ListUser(ctx context.Context) ([]*entity.User, error)
	//ChangeUserRole(ctx context.Context, email string) (*entity.User, error)
	DeleteUser(ctx context.Context, id int64) error
	UpdateUser(ctx context.Context, id int64, u *entity.User) (*entity.User, error)
	UpdateUserPassword(ctx context.Context, id int64, hashPass string) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	//CreateUserRole(ctx context.Context, u *entity.User) (*entity.User, error)
}
type Models struct {
	User entity.UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		User: entity.UserModel{Db: db},
	}
}
