package entity

import (
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"time"
)

type User struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `json:"name"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
type UserModel struct {
	Db *sql.DB
}

//
//func (r *UserModel) ChangeUserRole(ctx context.Context, email string) (*User, error) {
//	user, err := r.GetUserByEmail(ctx, email)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return nil, status.Error(codes.NotFound, "user not found")
//		}
//		return nil, fmt.Errorf("failed to get user by email: %w", err)
//	}
//	query := "UPDATE users_roles SET role = $1 WHERE user_id = $2 RETURNING id"
//	var id int64
//	err = r.Db.QueryRowContext(ctx, query, "customer", user.ID).Scan(&id)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return nil, status.Error(codes.NotFound, "no rows updated")
//		}
//		return nil, fmt.Errorf("failed to change user role: %w", err)
//	}
//	return &User{ID: id}, nil
//}

//func (r *UserModel) CreateUserRole(ctx context.Context, u *User) (*User, error) {
//	query := "INSERT INTO users_roles (user_id, role) VALUES ($1, $2) RETURNING id"
//	err := r.Db.QueryRowContext(ctx, query, u.ID, "user").Scan(&u.ID)
//	if err != nil {
//		return nil, fmt.Errorf("failed to insert user role: %w", err)
//	}
//	return u, nil
//}

func (r *UserModel) CreateUser(ctx context.Context, u *User) (*User, error) {
	query := "INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := r.Db.QueryRowContext(ctx, query, u.Name, u.Email, u.Password, time.Now(), time.Now()).Scan(&u.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}
	return u, nil
}
func (r *UserModel) GetUser(ctx context.Context, id int64) (*User, error) {
	query := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = $1"
	row := r.Db.QueryRowContext(ctx, query, id)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan user: %w", err)
	}
	return &user, nil
}

func (r *UserModel) ListUser(ctx context.Context) ([]*User, error) {
	query := "SELECT id, email, password, role, created_at, updated_at FROM users"
	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

//func (r *UserModel)(ctx context.Context, name string, error) {
//	query := "SELECT id, name FROM roles WHERE name = ?"
//	row := r.Db.QueryRowContext(ctx, query, name)
//
//	var r
//	if err := row.Scan(&role.ID, &role.Name); err != nil {
//		if err == sql.ErrNoRows {
//			return nil, nil
//		}
//		return nil, err
//	}
//
//	return &role, nil
//}

func (r *UserModel) DeleteUser(ctx context.Context, id int64) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.Db.ExecContext(ctx, query, id)
	return err
}

func (r *UserModel) UpdateUser(ctx context.Context, id int64, u *User) (*User, error) {
	query := "UPDATE users SET email = ?, password = ?, role = ?, updated_at = ? WHERE id = ?"
	_, err := r.Db.ExecContext(ctx, query, u.Email, u.Password, time.Now(), id)
	if err != nil {
		return nil, err
	}

	return r.GetUser(ctx, id)
}

func (r *UserModel) UpdateUserPassword(ctx context.Context, id int64, hashPass string) (*User, error) {
	query := "UPDATE users SET password = ?, updated_at = ? WHERE id = ?"
	_, err := r.Db.ExecContext(ctx, query, hashPass, time.Now(), id)
	if err != nil {
		return nil, err
	}

	return r.GetUser(ctx, id)
}

func (r *UserModel) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	query := "SELECT id, name ,email, password,created_at, updated_at FROM users WHERE email = $1"
	row := r.Db.QueryRowContext(ctx, query, email)

	var user User
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
