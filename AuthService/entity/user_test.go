package entity

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUserModel_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	userModel := UserModel{Db: db}

	user := &User{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "securepassword",
	}

	mock.ExpectQuery("INSERT INTO users").
		WithArgs(user.Name, user.Email, user.Password, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	createdUser, err := userModel.CreateUser(ctx, user)
	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, int64(1), createdUser.ID)
}

func TestUserModel_GetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	userModel := UserModel{Db: db}

	user := &User{
		ID:       1,
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "securepassword",
	}

	mock.ExpectQuery("SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = \\$1").
		WithArgs(user.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).
			AddRow(user.ID, user.Name, user.Email, user.Password, time.Now(), time.Now()))

	fetchedUser, err := userModel.GetUser(ctx, user.ID)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedUser)
	assert.Equal(t, user.ID, fetchedUser.ID)
	assert.Equal(t, user.Email, fetchedUser.Email)
}

func TestUserModel_DeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	userModel := UserModel{Db: db}

	user := &User{
		ID: 1,
	}

	mock.ExpectExec("DELETE FROM users WHERE id = \\?").
		WithArgs(user.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = userModel.DeleteUser(ctx, user.ID)
	assert.NoError(t, err)
}

func TestUserModel_UpdateUserPassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	userModel := UserModel{Db: db}

	user := &User{
		ID:       1,
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "securepassword",
	}

	newPassword := "newsecurepassword"

	mock.ExpectExec("UPDATE users SET password = \\?, updated_at = \\? WHERE id = \\?").
		WithArgs(newPassword, sqlmock.AnyArg(), user.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery("SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = \\$1").
		WithArgs(user.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).
			AddRow(user.ID, user.Name, user.Email, newPassword, time.Now(), time.Now()))

	updatedUser, err := userModel.UpdateUserPassword(ctx, user.ID, newPassword)
	assert.NoError(t, err)
	assert.Equal(t, newPassword, updatedUser.Password)
}
