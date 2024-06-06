package handlers

import (
	"SkyTicket/BookingService/entity"
	"SkyTicket/proto/pb"
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"net"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestBookingHandler_UpdateBookingStatus_Integration(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	defer db.Close()

	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	defer server.Stop()

	ctx := context.Background()

	bookingRepo := &entity.BookingModel{Db: db}

	bookingHandler := &BookingHandler{
		bookingRepo:  bookingRepo,
		userClient:   mockUserManagerClient(listener),
		flightClient: mockFlightManagerClient(listener),
	}

	req := &pb.UpdateBookingStatusRequest{
		Status:   "Cancel",
		FlightId: 1,
	}

	res, err := bookingHandler.UpdateBookingStatus(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestBookingHandler_GetBooking_Integration(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	defer db.Close()

	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	defer server.Stop()

	ctx := context.Background()

	bookingRepo := &entity.BookingModel{Db: db}

	bookingHandler := &BookingHandler{
		bookingRepo:  bookingRepo,
		userClient:   mockUserManagerClient(listener),
		flightClient: mockFlightManagerClient(listener),
	}

	expectedBooking := &entity.Booking{
		ID:        1,
		UserID:    1,
		FlightID:  1,
		Code:      "ABC123",
		Status:    "confirmed",
		TicketID:  1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	rows := sqlmock.NewRows([]string{"id", "user_id", "flight_id", "code", "status", "ticket_id", "created_at", "updated_at"}).
		AddRow(expectedBooking.ID, expectedBooking.UserID, expectedBooking.FlightID, expectedBooking.Code, expectedBooking.Status, expectedBooking.TicketID, expectedBooking.CreatedAt, expectedBooking.UpdatedAt)
	mock.ExpectQuery("SELECT id, user_id, flight_id, code, status, ticket_id, created_at, updated_at FROM booking WHERE code = ?").WithArgs("ABC123").WillReturnRows(rows)

	req := &pb.GetBookingRequest{
		Code: "ABC123",
	}

	res, err := bookingHandler.GetBooking(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBookingHandler_ListBooking_Integration(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	defer db.Close()

	bookingRepo := &entity.BookingModel{Db: db}

	bookingHandler := &BookingHandler{
		bookingRepo: bookingRepo,
	}

	expectedBookings := []*entity.Booking{
		{
			ID:        1,
			UserID:    1,
			FlightID:  1,
			Code:      "ABC123",
			Status:    "confirmed",
			TicketID:  1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			UserID:    2,
			FlightID:  2,
			Code:      "DEF456",
			Status:    "pending",
			TicketID:  2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	rows := sqlmock.NewRows([]string{"id", "user_id", "flight_id", "code", "status", "ticket_id", "created_at", "updated_at"})
	for _, booking := range expectedBookings {
		rows.AddRow(booking.ID, booking.UserID, booking.FlightID, booking.Code, booking.Status, booking.TicketID, booking.CreatedAt, booking.UpdatedAt)
	}
	mock.ExpectQuery("SELECT id, user_id, flight_id, code, status, ticket_id, created_at, updated_at FROM booking").WillReturnRows(rows)

	ctx := context.Background()
	res, err := bookingHandler.ListBooking(ctx, nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	assert.Len(t, res.BookingList, len(expectedBookings))
	for i, expectedBooking := range expectedBookings {
		assert.Equal(t, expectedBooking.ID, res.BookingList[i].Id)
		assert.Equal(t, expectedBooking.UserID, res.BookingList[i].UserId)
		assert.Equal(t, expectedBooking.FlightID, res.BookingList[i].FlightId)
		assert.Equal(t, expectedBooking.Code, res.BookingList[i].Code)
		assert.Equal(t, expectedBooking.Status, res.BookingList[i].Status)
		assert.Equal(t, expectedBooking.TicketID, res.BookingList[i].TicketId)
		assert.Equal(t, expectedBooking.CreatedAt.Unix(), res.BookingList[i].CreatedAt.Seconds)
		assert.Equal(t, expectedBooking.UpdatedAt.Unix(), res.BookingList[i].UpdatedAt.Seconds)
	}

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestBookingHandler_CreateBooking_Integration(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	defer db.Close()

	bookingRepo := &entity.BookingModel{Db: db}

	bookingHandler := &BookingHandler{
		bookingRepo: bookingRepo,
	}

	request := &pb.CreateBookingRequest{
		Code:      "ABC123",
		UserId:    1,
		FlightId:  1,
		Status:    "confirmed",
		TicketId:  1,
		CreatedAt: &timestamp.Timestamp{Seconds: time.Now().Unix()},
		UpdatedAt: &timestamp.Timestamp{Seconds: time.Now().Unix()},
	}

	mock.ExpectExec("INSERT INTO booking").WithArgs(request.Code, request.UserId, request.FlightId, request.Status, request.TicketId, request.CreatedAt, request.UpdatedAt).WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.Background()
	response, err := bookingHandler.CreateBooking(ctx, request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestBookingHandler_CancelBooking_Integration(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	defer db.Close()

	bookingRepo := &entity.BookingModel{Db: db}

	bookingHandler := &BookingHandler{
		bookingRepo: bookingRepo,
	}

	booking := &entity.Booking{
		ID:        1,
		UserID:    1,
		FlightID:  1,
		Code:      "ABC123",
		Status:    "confirmed",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	mock.ExpectQuery("SELECT").WithArgs("ABC123").WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "flight_id", "code", "status", "created_at", "updated_at"}).AddRow(booking.ID, booking.UserID, booking.FlightID, booking.Code, booking.Status, booking.CreatedAt, booking.UpdatedAt))

	mock.ExpectExec("UPDATE booking").WithArgs(booking.Status, booking.UpdatedAt, booking.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.Background()
	request := &pb.CancelBookingRequest{BookingCode: "ABC123"}
	response, err := bookingHandler.CancelBooking(ctx, request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func mockUserManagerClient(listener *bufconn.Listener) pb.UserManagerClient {
	conn, _ := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}), grpc.WithInsecure())
	return pb.NewUserManagerClient(conn)
}

func mockFlightManagerClient(listener *bufconn.Listener) pb.FlightManagerClient {
	conn, _ := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}), grpc.WithInsecure())
	return pb.NewFlightManagerClient(conn)
}
