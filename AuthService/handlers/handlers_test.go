package handlers

import (
	"SkyTicket/AuthService/repo"
	"SkyTicket/proto/pb"
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_GetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userRepo := repo.NewModels(db).User
	userHandler, _ := NewUserHandler(&userRepo)

	ctx := context.Background()
	userID := int64(1)

	mock.ExpectQuery("SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = \\$1").
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).
			AddRow(userID, "John Doe", "johndoe@example.com", "securepassword", time.Now(), time.Now()))

	req := &pb.GetUserRequest{Id: userID}
	resp, err := userHandler.GetUser(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, userID, resp.Id)
	assert.Equal(t, "John Doe", resp.Name)
	assert.Equal(t, "johndoe@example.com", resp.Email)
}

func TestUserHandler_GetUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userRepo := repo.NewModels(db).User
	userHandler, _ := NewUserHandler(&userRepo)

	ctx := context.Background()
	userID := int64(1)

	mock.ExpectQuery("SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = \\$1").
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).
			AddRow(userID, "John Doe", "johndoe@example.com", "securepassword", time.Now(), time.Now()))

	req := &pb.GetUserRequest{Id: userID}
	resp, err := userHandler.GetUser(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, userID, resp.Id)
	assert.Equal(t, "John Doe", resp.Name)
	assert.Equal(t, "johndoe@example.com", resp.Email)
}
