package handlers

import (
	"SkyTicket/BookingService/entity"
	repository "SkyTicket/BookingService/repo"
	"SkyTicket/proto/pb"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	_ "google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	_ "strconv"
)

type BookingHandler struct {
	pb.UnimplementedBookingManagerServer
	bookingRepo  repository.BookingRepository
	userClient   pb.UserManagerClient
	flightClient pb.FlightManagerClient
}

func NewBookingHandler(bookingRepo repository.BookingRepository, flightClient pb.FlightManagerClient, userClient pb.UserManagerClient) (*BookingHandler, error) {
	return &BookingHandler{
		bookingRepo:  bookingRepo,
		userClient:   userClient,
		flightClient: flightClient,
	}, nil
}

func (h *BookingHandler) UpdateBookingStatus(ctx context.Context, req *pb.UpdateBookingStatusRequest) (*pb.UpdateBookingStatusResponse, error) {
	if req.Status == "Cancel" {
		err := h.bookingRepo.UpdateBookingStatus(ctx, req.FlightId, "Cancel")
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		pRes := &pb.UpdateBookingStatusResponse{Status: "Update Completed"}
		return pRes, nil
	} else if req.Status == "Arrived" {
		err := h.bookingRepo.UpdateBookingStatus(ctx, req.FlightId, "completed")
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		pRes := &pb.UpdateBookingStatusResponse{Status: "Update Completed"}
		return pRes, nil
	}
	return nil, nil
}

func (h *BookingHandler) GetBooking(ctx context.Context, b *pb.GetBookingRequest) (*pb.Booking, error) {
	booking, err := h.bookingRepo.GetBookingByCode(ctx, b.Code)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "booking not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &pb.Booking{
		Id:        booking.ID,
		UserId:    booking.UserID,
		FlightId:  booking.FlightID,
		Code:      booking.Code,
		Status:    booking.Status,
		TicketId:  booking.TicketID,
		CreatedAt: timestamppb.New(booking.CreatedAt),
		UpdatedAt: timestamppb.New(booking.UpdatedAt),
	}
	return res, nil
}

func (h *BookingHandler) ListBooking(ctx context.Context, req *pb.ListBookingRequest) (*pb.ListBookingResponse, error) {
	bks, err := h.bookingRepo.ListBooking(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "booking not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	var bookings []*pb.Booking
	for _, bk := range bks {
		bookings = append(bookings, &pb.Booking{
			Id:        bk.ID,
			UserId:    bk.UserID,
			FlightId:  bk.FlightID,
			Code:      bk.Code,
			Status:    bk.Status,
			TicketId:  bk.TicketID,
			CreatedAt: timestamppb.New(bk.CreatedAt),
			UpdatedAt: timestamppb.New(bk.UpdatedAt),
		})
	}

	res := &pb.ListBookingResponse{
		BookingList: bookings,
	}
	return res, nil
}
func (h *BookingHandler) CreateBooking(ctx context.Context, req *pb.CreateBookingRequest) (*pb.Booking, error) {
	// Log the received request
	log.Printf("Received booking request: %+v", req)

	// Retrieve user information
	_, err := h.userClient.GetUser(ctx, &pb.GetUserRequest{Id: req.UserId})
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	// Retrieve ticket information
	ticket, err := h.bookingRepo.GetTicketClass(ctx, req.TicketType)
	if err != nil {
		return nil, fmt.Errorf("failed to get ticket class: %v", err)
	}

	// Retrieve flight information
	flight, err := h.flightClient.GetFlight(ctx, &pb.GetFlightRequest{Id: req.FlightId})
	if err != nil {
		return nil, fmt.Errorf("failed to get flight: %v", err)
	}

	// Log flight information
	log.Printf("Flight information: %+v", flight)

	// Check available slots based on ticket type
	switch req.TicketType {
	case "first":
		if flight.AvailableFirstSlot == 0 {
			return nil, fmt.Errorf("first class seat is full")
		}
	case "economy":
		if flight.AvailableEconomySlot == 0 {
			return nil, fmt.Errorf("economy class seat is full")
		}
	default:
		return nil, fmt.Errorf("invalid ticket type")
	}

	// Create new booking
	newBooking := &entity.Booking{
		Code:     req.Code,
		UserID:   req.UserId,
		FlightID: req.FlightId,
		Status:   req.Status,
		TicketID: ticket.ID,
	}

	// Log new booking information
	log.Printf("New booking information: %+v", newBooking)

	// Create booking in the repository
	createdBooking, err := h.bookingRepo.CreateBooking(ctx, newBooking)
	if err != nil {
		return nil, fmt.Errorf("failed to create booking: %v", err)
	}

	// Log created booking information
	log.Printf("Created booking information: %+v", createdBooking)

	// Update flight slot
	_, err = h.flightClient.UpdateFlightSlot(ctx, &pb.UpdateFlightSlotRequest{
		Id:         createdBooking.FlightID,
		TicketType: createdBooking.TicketID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update flight slot: %v", err)
	}

	bookingRes := &pb.Booking{
		Id:        createdBooking.ID,
		UserId:    createdBooking.UserID,
		FlightId:  createdBooking.FlightID,
		Code:      createdBooking.Code,
		Status:    createdBooking.Status,
		TicketId:  createdBooking.TicketID,
		CreatedAt: &timestamp.Timestamp{Seconds: createdBooking.CreatedAt.Unix()},
		UpdatedAt: &timestamp.Timestamp{Seconds: createdBooking.UpdatedAt.Unix()},
	}
	return bookingRes, nil
}

func (h *BookingHandler) CancelBooking(ctx context.Context, req *pb.CancelBookingRequest) (*pb.Booking, error) {
	booking, err := h.bookingRepo.GetBookingByCode(ctx, req.BookingCode)
	if err != nil {
		return nil, status.Error(codes.NotFound, "booking not found")
	}

	if booking.Status == "Cancel" {
		return nil, status.Error(codes.Aborted, "booking was canceled")
	}
	booking.Status = "Cancel"

	updatedBooking, err := h.bookingRepo.UpdateBooking(ctx, booking.ID, booking)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	md := metadata.Pairs("update", "delete")
	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err = h.flightClient.UpdateFlightSlot(ctx, &pb.UpdateFlightSlotRequest{
		Id:         updatedBooking.FlightID,
		TicketType: updatedBooking.TicketID,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pRes := &pb.Booking{
		Id:        updatedBooking.ID,
		UserId:    updatedBooking.UserID,
		FlightId:  updatedBooking.FlightID,
		Code:      updatedBooking.Code,
		Status:    updatedBooking.Status,
		CreatedAt: timestamppb.New(updatedBooking.CreatedAt),
		UpdatedAt: timestamppb.New(updatedBooking.UpdatedAt),
	}
	return pRes, nil
}
