package handlers

import (
	"SkyTicket/BookingService/entity"
	repository "SkyTicket/BookingService/repo"
	"SkyTicket/proto/pb"
	"context"
	"database/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

type BookingHandler struct {
	pb.UnimplementedBookingManagerServer
	bookingRepo    repository.BookingRepository
	userClient     pb.UserManagerClient
	airplaneClient pb.AirplaneServiceClient
}

func NewBookingHandler(bookingRepo repository.BookingRepository, userClient pb.UserManagerClient, airplaneClient pb.AirplaneServiceClient) (*BookingHandler, error) {
	return &BookingHandler{
		bookingRepo:    bookingRepo,
		userClient:     userClient,
		airplaneClient: airplaneClient,
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

func (h *BookingHandler) CreateBooking(ctx context.Context, b *pb.CreateBookingRequest) (*pb.Booking, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	customerId := md["user"][0]
	conVId, _ := strconv.Atoi(customerId)
	u, err := h.userClient.GetUser(ctx, &pb.GetUserRequest{Id: int64(conVId)})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "customer not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	airplane := h.airplaneClient.GetAirplane(ctx, &pb.GetAirplaneRequest{Id: id})
	h.airplaneClient.
		// check ticket valid
		_, err = h.bookingRepo.GetTicketClass(ctx, b.TicketType)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "ticket not valid")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	if b.TicketType == 1 && flight.AvailableFirstSlot == 0 {
		return nil, status.Error(codes.Unavailable, "first class Seat is full")
	}
	if b.TicketType == 2 && flight.AvailableEconomySlot == 0 {
		return nil, status.Error(codes.Unavailable, "economy class seat is full")
	}

	booking := &entity.Booking{
		UserID:   u.Id,
		FlightID: flight.Id,
		Status:   "Scheduled",
		TicketID: b.TicketType,
	}
	createdBooking, err := h.bookingRepo.CreateBooking(ctx, booking)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	md = metadata.Pairs("update", "create")
	ctx = metadata.NewOutgoingContext(ctx, md)

	// call update with metadata to check if update from create-booking (delete available slot)
	_, err = h.flightClient.UpdateFlightSlot(ctx, &pb.UpdateFlightSlotRequest{
		Id:         createdBooking.FlightID,
		TicketType: createdBooking.TicketID,
	})
	if err != nil {
		return nil, err
	}

	// protobuf response
	pRes := &pb.Booking{
		Id:        createdBooking.ID,
		UserId:    createdBooking.UserID,
		FlightId:  createdBooking.FlightID,
		Code:      createdBooking.Code,
		Status:    createdBooking.Status,
		TicketId:  createdBooking.TicketID,
		CreatedAt: timestamppb.New(createdBooking.CreatedAt),
		UpdatedAt: timestamppb.New(createdBooking.UpdatedAt),
	}
	return pRes, nil
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
