package handlers

import (
	"SkyTicket/entity"
	"SkyTicket/internal"
	"SkyTicket/pb"
	repository "SkyTicket/repo"
	"context"
	"database/sql"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookingHandler struct {
	pb.UnimplementedBookingManagerServer
	bookingRepo repository.BookingRepository
	flightRepo  repository.FlightRepository
}

func NewBookingHandler(bookingRepo repository.BookingRepository, flightRepo repository.FlightRepository) (*BookingHandler, error) {
	return &BookingHandler{
		bookingRepo: bookingRepo,
		flightRepo:  flightRepo,
	}, nil
}

func (h *BookingHandler) SearchFlight(ctx context.Context, req *pb.SearchFlightRequest) (*pb.SearchFlightResponse, error) {

	list, err := h.flightRepo.SearchFlight(ctx, req.From, req.To, req.DepartureDate.AsTime(), req.ArrivalDate.AsTime())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var flightPbRes []*pb.Flight
	for _, fl := range list {
		flightPb := &pb.Flight{}
		flightPb.Id = fl.ID
		flightPb.Name = fl.Name
		flightPb.From = fl.From
		flightPb.To = fl.To
		flightPb.DepartureDate = timestamppb.New(fl.DepartureDate)
		flightPb.ArrivalDate = timestamppb.New(fl.ArrivalDate)
		flightPb.AvailableSlot = int64(fl.AvailableSlot)
		flightPb.Status = string(fl.Status)
		flightPb.CreatedAt = timestamppb.New(fl.CreatedAt)
		flightPb.UpdatedAt = timestamppb.New(fl.UpdatedAt)

		flightPbRes = append(flightPbRes, flightPb)
	}

	pRes := &pb.SearchFlightResponse{
		FlightList: flightPbRes,
	}
	return pRes, nil
}

func (h *BookingHandler) CreateFlight(ctx context.Context, req *pb.Flight) (*pb.Flight, error) {
	statusParse, _ := internal.ParseString(req.Status)

	fl := &entity.Flight{
		Name:          req.Name,
		From:          req.From,
		To:            req.To,
		DepartureDate: req.DepartureDate.AsTime(),
		ArrivalDate:   req.ArrivalDate.AsTime(),
		AvailableSlot: int(req.AvailableSlot),
		Status:        entity.Status(statusParse),
	}

	f, err := h.flightRepo.CreateFlight(ctx, fl)
	if err != nil {
		return nil, err
	}

	pRes := &pb.Flight{}
	if err := copier.Copy(&pRes, f); err != nil {
		return nil, err
	}
	pRes.Id = f.ID
	return pRes, nil
}

func (h *BookingHandler) UpdateFlight(ctx context.Context, req *pb.Flight) (*pb.Flight, error) {
	fl, err := h.flightRepo.GetFlightById(ctx, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "flight not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	statusParse, _ := internal.ParseString(req.Status)

	flReq := &entity.Flight{
		Name:          req.Name,
		DepartureDate: req.DepartureDate.AsTime(),
		ArrivalDate:   req.ArrivalDate.AsTime(),
		AvailableSlot: int(req.AvailableSlot),
		Status:        entity.Status(statusParse),
	}

	flightInput := internal.CheckFlightEmptyInput(flReq, fl)
	if fl.Status == entity.StatusCancel {
		return nil, status.Error(codes.InvalidArgument, "flight was canceled, can't update status as cancel")
	} else if fl.Status == entity.StatusArrived {
		return nil, status.Error(codes.InvalidArgument, "flight was arrived, can't update status as arrived")
	}

	f, err := h.flightRepo.UpdateFlight(ctx, fl.ID, flightInput)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	//if f.Status == entity.StatusCancel {
	//	_, err := h.bookingClient.UpdateBookingStatus(ctx, &pb.UpdateBookingStatusRequest{FlightId: f.ID, Status: flight.StatusCancel.String()})
	//	if err != nil {
	//		return nil, status.Error(codes.Internal, err.Error())
	//	}
	//} else if f.Status == flight.StatusArrived {
	//	_, err := h.bookingClient.UpdateBookingStatus(ctx, &pb.UpdateBookingStatusRequest{FlightId: f.ID, Status: flight.StatusArrived.String()})
	//	if err != nil {
	//		return nil, status.Error(codes.Internal, err.Error())
	//	}
	//}

	pRes := &pb.Flight{}
	if err := copier.Copy(&pRes, f); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	pRes.Id = f.ID
	pRes.DepartureDate = timestamppb.New(f.DepartureDate)
	pRes.ArrivalDate = timestamppb.New(f.ArrivalDate)

	return pRes, nil
}

func (h *BookingHandler) GetFlight(ctx context.Context, req *pb.GetFlightRequest) (*pb.Flight, error) {
	fl, err := h.flightRepo.GetFlight(ctx, req.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "flight not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	pRes := &pb.Flight{
		Id:            fl.ID,
		Name:          fl.Name,
		From:          fl.From,
		To:            fl.To,
		DepartureDate: timestamppb.New(fl.DepartureDate),
		ArrivalDate:   timestamppb.New(fl.ArrivalDate),
		AvailableSlot: int64(fl.AvailableSlot),
		Status:        string(fl.Status),
		CreatedAt:     timestamppb.New(fl.CreatedAt),
		UpdatedAt:     timestamppb.New(fl.UpdatedAt),
	}

	return pRes, nil
}

func (h *BookingHandler) GetFlightById(ctx context.Context, req *pb.GetFlightRequest) (*pb.Flight, error) {
	fl, err := h.flightRepo.GetFlightById(ctx, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "flight not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	pRes := &pb.Flight{
		Id:            fl.ID,
		Name:          fl.Name,
		From:          fl.From,
		To:            fl.To,
		DepartureDate: timestamppb.New(fl.DepartureDate),
		ArrivalDate:   timestamppb.New(fl.ArrivalDate),
		AvailableSlot: int64(fl.AvailableSlot),
		Status:        string(fl.Status),
		CreatedAt:     timestamppb.New(fl.CreatedAt),
		UpdatedAt:     timestamppb.New(fl.UpdatedAt),
	}
	return pRes, nil
}

func (h *BookingHandler) ListFlight(ctx context.Context, req *pb.ListFlightRequest) (*pb.ListFlightResponse, error) {

	list, err := h.flightRepo.ListFlight(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// append ent type to protobuf response
	var flightPbRes []*pb.Flight
	for _, fl := range list {
		flightPb := &pb.Flight{}
		flightPb.Id = fl.ID
		flightPb.Name = fl.Name
		flightPb.From = fl.From
		flightPb.To = fl.To
		flightPb.DepartureDate = timestamppb.New(fl.DepartureDate)
		flightPb.ArrivalDate = timestamppb.New(fl.ArrivalDate)
		flightPb.AvailableSlot = int64(fl.AvailableSlot)
		flightPb.Status = string(fl.Status)
		flightPb.CreatedAt = timestamppb.New(fl.CreatedAt)
		flightPb.UpdatedAt = timestamppb.New(fl.UpdatedAt)

		flightPbRes = append(flightPbRes, flightPb)
	}

	pRes := &pb.ListFlightResponse{
		FlightList: flightPbRes,
	}
	return pRes, nil
}

func (h *BookingHandler) UpdateFlightSlot(ctx context.Context, req *pb.UpdateFlightSlotRequest) (*pb.Flight, error) {
	md, _ := metadata.FromIncomingContext(ctx)

	if md["update"][0] == "create" {
		fl, _ := h.flightRepo.GetFlightById(ctx, req.Id)
		if req.TicketType == 2 {
			fl.AvailableSlot = fl.AvailableSlot - 1
		}

		updatedFlight, err := h.flightRepo.UpdateFlight(ctx, req.Id, fl)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		pRes := &pb.Flight{Id: updatedFlight.ID}
		return pRes, nil
	}

	if md["update"][0] == "delete" {
		fl, _ := h.flightRepo.GetFlightById(ctx, req.Id)
		if req.TicketType == 2 {
			fl.AvailableSlot = fl.AvailableSlot + 1
		}
		updatedFlight, err := h.flightRepo.UpdateFlight(ctx, req.Id, fl)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		pRes := &pb.Flight{Id: updatedFlight.ID}
		return pRes, nil
	}

	return nil, nil
}

func (h *BookingHandler) DeleteFlight(ctx context.Context, req *pb.DeleteFlightRequest) (*empty.Empty, error) {
	err := h.flightRepo.DeleteFlight(ctx, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "flight not found")
		}
		return nil, err
	}
	return nil, nil
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

func (h *BookingHandler) CreateBooking(ctx context.Context, b *pb.CreateBookingRequest) (*pb.Booking, error) {
	// get metadata from client
	md, _ := metadata.FromIncomingContext(ctx)
	// User Create Booking
	//customerId := md["user"][0]
	// get customer information associated with user
	//conVId, _ := strconv.Atoi(customerId)
	//c, err := h.customerClient.GetCustomer(ctx, &pb.GetCustomerRequest{Id: int64(conVId)})
	//if err != nil {
	//	if err == sql.ErrNoRows {
	//		return nil, status.Error(codes.NotFound, "customer not found")
	//	}
	//	return nil, status.Error(codes.Internal, err.Error())
	//}

	flight, err := h.GetFlight(ctx, &pb.GetFlightRequest{Name: b.FlightName})
	//if err != nil {
	//	if err == sql.ErrNoRows {
	//		return nil, status.Error(codes.NotFound, "flight not found")
	//	}
	//	return nil, status.Error(codes.Internal, err.Error())
	//}

	_, err = h.bookingRepo.GetTicketClass(ctx, b.TicketType)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "ticket not valid")
		}
		return nil, status.Error(codes.Internal, "error")
	}

	if flight.AvailableSlot == 0 {
		return nil, status.Error(codes.Unavailable, "Seat is full")
	}
	booking := &entity.Booking{
		//CustomerID: c.Id,
		FlightID: flight.Id,
		Status:   "Scheduled",
		TicketID: b.TicketType,
	}
	createdBooking, err := h.bookingRepo.CreateBooking(ctx, booking)
	if err != nil {
		return nil, status.Error(codes.Internal, "error")
	}

	md = metadata.Pairs("update", "create")
	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err = h.UpdateFlightSlot(ctx, &pb.UpdateFlightSlotRequest{
		Id:         createdBooking.FlightID,
		TicketType: createdBooking.TicketID,
	})
	if err != nil {
		return nil, err
	}

	pRes := &pb.Booking{
		Id:         createdBooking.ID,
		CustomerId: createdBooking.CustomerID,
		FlightId:   createdBooking.FlightID,
		Code:       createdBooking.Code,
		Status:     createdBooking.Status,
		TicketId:   createdBooking.TicketID,
		CreatedAt:  timestamppb.New(createdBooking.CreatedAt),
		UpdatedAt:  timestamppb.New(createdBooking.UpdatedAt),
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

	// repo update
	booking.Status = "Cancel"
	updatedBooking, err := h.bookingRepo.UpdateBooking(ctx, booking.ID, booking)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	md := metadata.Pairs("update", "delete")
	ctx = metadata.NewOutgoingContext(ctx, md)
	// call update with metadata to check if update from delete-booking (add more available slot)
	_, err = h.UpdateFlightSlot(ctx, &pb.UpdateFlightSlotRequest{
		Id:         updatedBooking.FlightID,
		TicketType: updatedBooking.TicketID,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pRes := &pb.Booking{
		Id:         updatedBooking.ID,
		CustomerId: updatedBooking.CustomerID,
		FlightId:   updatedBooking.FlightID,
		Code:       updatedBooking.Code,
		Status:     updatedBooking.Status,
		CreatedAt:  timestamppb.New(updatedBooking.CreatedAt),
		UpdatedAt:  timestamppb.New(updatedBooking.UpdatedAt),
	}
	return pRes, nil
}
