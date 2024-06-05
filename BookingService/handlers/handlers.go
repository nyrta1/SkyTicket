package handlers

import (
	"SkyTicket/pb"
	repository "SkyTicket/repo"
	"context"
	"database/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookingHandler struct {
	pb.UnimplementedBookingManagerServer
	bookingRepo repository.BookingRepository
}

func NewBookingHandler(bookingRepo repository.BookingRepository) (*BookingHandler, error) {
	return &BookingHandler{
		bookingRepo: bookingRepo,
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
		Id:         booking.ID,
		CustomerId: booking.CustomerID,
		FlightId:   booking.FlightID,
		Code:       booking.Code,
		Status:     booking.Status,
		TicketId:   booking.TicketID,
		CreatedAt:  timestamppb.New(booking.CreatedAt),
		UpdatedAt:  timestamppb.New(booking.UpdatedAt),
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
			Id:         bk.ID,
			CustomerId: bk.CustomerID,
			FlightId:   bk.FlightID,
			Code:       bk.Code,
			Status:     bk.Status,
			TicketId:   bk.TicketID,
			CreatedAt:  timestamppb.New(bk.CreatedAt),
			UpdatedAt:  timestamppb.New(bk.UpdatedAt),
		})
	}

	res := &pb.ListBookingResponse{
		BookingList: bookings,
	}
	return res, nil
}

//func (h *BookingHandler) CreateBooking(ctx context.Context, b *pb.CreateBookingRequest) (*pb.Booking, error) {
//	// get metadata from client
//	md, _ := metadata.FromIncomingContext(ctx)
//
//	// Guest Create Booking
//	if md["user"][0] == "0" {
//		// check if ticket type valid
//		_, err := h.bookingRepo.GetTicketClass(ctx, b.TicketType)
//		if err != nil {
//			if err == sql.ErrNoRows {
//				return nil, status.Error(codes.NotFound, "ticket not valid")
//			}
//			return nil, status.Error(codes.Internal, err.Error())
//		}
//		flight, err := h.flightClient.GetFlight(ctx, &pb.GetFlightRequest{Name: b.FlightName})
//		if err != nil {
//			if err == sql.ErrNoRows {
//				return nil, status.Error(codes.NotFound, "flight not found")
//			}
//			return nil, status.Error(codes.Internal, err.Error())
//		}
//
//		// Check if slot is full or not
//		if b.TicketType == 1 && flight.AvailableFirstSlot == 0 {
//			return nil, status.Error(codes.Unavailable, "first class seat is full")
//		}
//		if b.TicketType == 2 && flight.AvailableEconomySlot == 0 {
//			return nil, status.Error(codes.Unavailable, "economy class seat is full")
//		}
//
//		u, _ := h.userClient.GetUserByEmail(ctx, &pb.GetUserByEmailRequest{Email: b.Email})
//		if u != nil {
//			return nil, status.Error(codes.AlreadyExists, "email already belongs to a user, please login to book flight")
//		}
//		customer, err := h.customerClient.CreateCustomer(ctx, &pb.Customer{
//			Name:           b.Name,
//			Email:          b.Email,
//			Address:        b.Address,
//			PhoneNumber:    b.PhoneNumber,
//			IdentifyNumber: b.IdentifyNumber,
//			DateOfBirth:    b.DateOfBirth,
//			MemberCode:     b.MemberCode,
//		})
//		if err != nil {
//			return nil, err
//		}
//
//		// repo
//		booking := &ent.Booking{
//			CustomerID: customer.Id,
//			FlightID:   flight.Id,
//			Status:     "Scheduled",
//			TicketID:   b.TicketType,
//		}
//		createdBooking, err := h.bookingRepo.CreateBooking(ctx, booking)
//		if err != nil {
//			return nil, err
//		}
//
//		md := metadata.Pairs("update", "create")
//		ctx = metadata.NewOutgoingContext(ctx, md)
//		// call update with metadata to check if update from create-booking (delete available slot)
//		_, err = h.flightClient.UpdateFlightSlot(ctx, &pb.UpdateFlightSlotRequest{
//			Id:         createdBooking.FlightID,
//			TicketType: createdBooking.TicketID,
//		})
//		if err != nil {
//			return nil, err
//		}
//
//		// protobuf response
//		pRes := &pb.Booking{
//			Id:         createdBooking.ID,
//			CustomerId: createdBooking.CustomerID,
//			FlightId:   createdBooking.FlightID,
//			Code:       createdBooking.Code,
//			Status:     createdBooking.Status.String(),
//			TicketId:   createdBooking.TicketID,
//			CreatedAt:  timestamppb.New(createdBooking.CreatedAt),
//			UpdatedAt:  timestamppb.New(createdBooking.UpdatedAt),
//		}
//		return pRes, nil
//	}
//
//	// User Create Booking
//	customerId := md["user"][0]
//	// get customer information associated with user
//	conVId, _ := strconv.Atoi(customerId)
//	c, err := h.customerClient.GetCustomer(ctx, &pb.GetCustomerRequest{Id: int64(conVId)})
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return nil, status.Error(codes.NotFound, "customer not found")
//		}
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	// get flight
//	flight, err := h.flightClient.GetFlight(ctx, &pb.GetFlightRequest{Name: b.FlightName})
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return nil, status.Error(codes.NotFound, "flight not found")
//		}
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	// check ticket valid
//	_, err = h.bookingRepo.GetTicketClass(ctx, b.TicketType)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return nil, status.Error(codes.NotFound, "ticket not valid")
//		}
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	// Check if slot is full
//	if b.TicketType == 1 && flight.AvailableFirstSlot == 0 {
//		return nil, status.Error(codes.Unavailable, "first class Seat is full")
//	}
//	if b.TicketType == 2 && flight.AvailableEconomySlot == 0 {
//		return nil, status.Error(codes.Unavailable, "economy class seat is full")
//	}
//
//	// repo
//	booking := &ent.Booking{
//		CustomerID: c.Id,
//		FlightID:   flight.Id,
//		Status:     "Scheduled",
//		TicketID:   b.TicketType,
//	}
//	createdBooking, err := h.bookingRepo.CreateBooking(ctx, booking)
//	if err != nil {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	md = metadata.Pairs("update", "create")
//	ctx = metadata.NewOutgoingContext(ctx, md)
//
//	// call update with metadata to check if update from create-booking (delete available slot)
//	_, err = h.flightClient.UpdateFlightSlot(ctx, &pb.UpdateFlightSlotRequest{
//		Id:         createdBooking.FlightID,
//		TicketType: createdBooking.TicketID,
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	// protobuf response
//	pRes := &pb.Booking{
//		Id:         createdBooking.ID,
//		CustomerId: createdBooking.CustomerID,
//		FlightId:   createdBooking.FlightID,
//		Code:       createdBooking.Code,
//		Status:     createdBooking.Status.String(),
//		TicketId:   createdBooking.TicketID,
//		CreatedAt:  timestamppb.New(createdBooking.CreatedAt),
//		UpdatedAt:  timestamppb.New(createdBooking.UpdatedAt),
//	}
//	return pRes, nil
//}

//func (h *BookingHandler) CancelBooking(ctx context.Context, req *pb.CancelBookingRequest) (*pb.Booking, error) {
//	booking, err := h.bookingRepo.GetBookingByCode(ctx, req.BookingCode)
//	if err != nil {
//		return nil, status.Error(codes.NotFound, "booking not found")
//	}
//
//	if booking.Status == "Cancel" {
//		return nil, status.Error(codes.Aborted, "booking was canceled")
//	}
//
//	// repo update
//	booking.Status = "Cancel"
//	updatedBooking, err := h.bookingRepo.UpdateBooking(ctx, booking.ID, booking)
//	if err != nil {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	md := metadata.Pairs("update", "delete")
//	ctx = metadata.NewOutgoingContext(ctx, md)
//	// call update with metadata to check if update from delete-booking (add more available slot)
//	_, err = h.flightClient.UpdateFlightSlot(ctx, &pb.UpdateFlightSlotRequest{
//		Id:         updatedBooking.FlightID,
//		TicketType: updatedBooking.TicketID,
//	})
//	if err != nil {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	// protobuf response
//	pRes := &pb.Booking{
//		Id:         updatedBooking.ID,
//		CustomerId: updatedBooking.CustomerID,
//		FlightId:   updatedBooking.FlightID,
//		Code:       updatedBooking.Code,
//		Status:     updatedBooking.Status.String(),
//		CreatedAt:  timestamppb.New(updatedBooking.CreatedAt),
//		UpdatedAt:  timestamppb.New(updatedBooking.UpdatedAt),
//	}
//	return pRes, nil
//}
