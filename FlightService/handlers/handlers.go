package handlers

import (
	"SkyTicket/FlightService/entity"
	"SkyTicket/FlightService/flight"
	repository "SkyTicket/FlightService/repo"
	"SkyTicket/pkg/request"
	"SkyTicket/proto/pb"
	"context"
	"database/sql"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FlightHandler struct {
	pb.UnimplementedFlightManagerServer
	flightRepo repository.FlightRepository
	//bookingClient pb.BookingManagerClient
}

func NewFlightHandler(flightRepo repository.FlightRepository) (*FlightHandler, error) {
	return &FlightHandler{
		flightRepo: flightRepo,
		//bookingClient: bookingClient,
	}, nil
}

func (h *FlightHandler) SearchFlight(ctx context.Context, req *pb.SearchFlightRequest) (*pb.SearchFlightResponse, error) {
	paging := request.Paging{
		Page:  req.Page,
		Limit: req.Limit,
	}
	list, pg, err := h.flightRepo.SearchFlight(ctx, req.From, req.To, req.DepartureDate.AsTime(), req.ArrivalDate.AsTime(), paging)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// append ent type to protobuf response
	var flightPbRes []*pb.Flight
	for _, fl := range list {
		flightPb := &pb.Flight{}
		flightPb.Id = fl.Id
		flightPb.Name = fl.Name
		flightPb.From = fl.From
		flightPb.To = fl.To
		flightPb.DepartureDate = timestamppb.New(fl.DepartureDate)
		flightPb.ArrivalDate = timestamppb.New(fl.ArrivalDate)
		flightPb.AvailableFirstSlot = int64(fl.AvailableFirstSlot)
		flightPb.AvailableEconomySlot = int64(fl.AvailableEconomySlot)
		flightPb.Status = string(fl.Status)
		flightPb.CreatedAt = timestamppb.New(fl.CreatedAt)
		flightPb.UpdatedAt = timestamppb.New(fl.UpdatedAt)

		flightPbRes = append(flightPbRes, flightPb)
	}

	pRes := &pb.SearchFlightResponse{
		FlightList: flightPbRes,
		Total:      pg.Total,
		Page:       pg.Page,
	}
	return pRes, nil
}

func (h *FlightHandler) CreateFlight(ctx context.Context, req *pb.Flight) (*pb.Flight, error) {
	fl := &entity.Flight{
		Name:                 req.Name,
		From:                 req.From,
		To:                   req.To,
		DepartureDate:        req.DepartureDate.AsTime(),
		ArrivalDate:          req.ArrivalDate.AsTime(),
		AvailableFirstSlot:   int(req.AvailableFirstSlot),
		AvailableEconomySlot: int(req.AvailableEconomySlot),
		Status:               req.Status,
	}

	f, err := h.flightRepo.CreateFlight(ctx, fl)
	if err != nil {
		return nil, err
	}

	pRes := &pb.Flight{}
	if err := copier.Copy(&pRes, f); err != nil {
		return nil, err
	}
	pRes.Id = f.Id

	return pRes, nil
}

func (h *FlightHandler) UpdateFlight(ctx context.Context, req *pb.Flight) (*pb.Flight, error) {
	fl, err := h.flightRepo.GetFlightById(ctx, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "flight not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	if fl.Status == flight.StatusCancel {
		return nil, status.Error(codes.InvalidArgument, "flight was canceled, can't update status as cancel")
	} else if fl.Status == flight.StatusArrived {
		return nil, status.Error(codes.InvalidArgument, "flight was arrived, can't update status as arrived")
	}

	// Update only the fields that are provided in the request
	if req.Name != "" {
		fl.Name = req.Name
	}
	if req.DepartureDate != nil {
		fl.DepartureDate = req.DepartureDate.AsTime()
	}
	if req.ArrivalDate != nil {
		fl.ArrivalDate = req.ArrivalDate.AsTime()
	}
	if req.AvailableFirstSlot != 0 {
		fl.AvailableFirstSlot = int(req.AvailableFirstSlot)
	}
	if req.AvailableEconomySlot != 0 {
		fl.AvailableEconomySlot = int(req.AvailableEconomySlot)
	}
	if req.Status != "" {
		fl.Status = req.Status
	}

	updatedFlight, err := h.flightRepo.UpdateFlight(ctx, fl.Id, fl)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	//switch updatedFlight.Status {
	//case flight.StatusCancel, flight.StatusArrived:
	//	_, err := h.bookingClient.UpdateBookingStatus(ctx, &pb.UpdateBookingStatusRequest{
	//		FlightId: updatedFlight.Id,
	//		Status:   updatedFlight.Status,
	//	})
	//	if err != nil {
	//		return nil, status.Error(codes.Internal, err.Error())
	//	}
	//}

	// Prepare response
	response := &pb.Flight{
		Id:                   updatedFlight.Id,
		Name:                 updatedFlight.Name,
		DepartureDate:        timestamppb.New(updatedFlight.DepartureDate),
		ArrivalDate:          timestamppb.New(updatedFlight.ArrivalDate),
		AvailableFirstSlot:   int64(updatedFlight.AvailableFirstSlot),
		AvailableEconomySlot: int64(updatedFlight.AvailableEconomySlot),
		Status:               updatedFlight.Status,
	}

	return response, nil
}
func (h *FlightHandler) GetFlight(ctx context.Context, req *pb.GetFlightRequest) (*pb.GetFlightResponse, error) {
	fl, err := h.flightRepo.GetFlight(ctx, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "flight not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	pRes := &pb.GetFlightResponse{
		FlightList: []*pb.Flight{
			{
				Id:                   fl.Id,
				Name:                 fl.Name,
				From:                 fl.From,
				To:                   fl.To,
				DepartureDate:        timestamppb.New(fl.DepartureDate),
				ArrivalDate:          timestamppb.New(fl.ArrivalDate),
				AvailableFirstSlot:   int64(fl.AvailableFirstSlot),
				AvailableEconomySlot: int64(fl.AvailableEconomySlot),
				Status:               fl.Status,
				CreatedAt:            timestamppb.New(fl.CreatedAt),
				UpdatedAt:            timestamppb.New(fl.UpdatedAt),
			},
		},
	}
	return pRes, nil
}
func (h *FlightHandler) GetFlightById(ctx context.Context, req *pb.GetFlightRequest) (*pb.GetFlightRequest, error) {
	fl, err := h.flightRepo.GetFlightById(ctx, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "flight not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	pRes := &pb.GetFlightRequest{
		Id:   fl.Id,
		Name: fl.Name,
	}
	return pRes, nil
}

func (h *FlightHandler) ListFlight(ctx context.Context, req *pb.ListFlightRequest) (*pb.ListFlightResponse, error) {
	paging := request.Paging{
		Page:  req.Page,
		Limit: req.Limit,
	}

	list, pg, err := h.flightRepo.ListFlight(ctx, paging)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// append ent type to protobuf response
	var flightPbRes []*pb.Flight
	for _, fl := range list {
		flightPb := &pb.Flight{}
		flightPb.Id = fl.Id
		flightPb.Name = fl.Name
		flightPb.From = fl.From
		flightPb.To = fl.To
		flightPb.DepartureDate = timestamppb.New(fl.DepartureDate)
		flightPb.ArrivalDate = timestamppb.New(fl.ArrivalDate)
		flightPb.AvailableFirstSlot = int64(fl.AvailableFirstSlot)
		flightPb.AvailableEconomySlot = int64(fl.AvailableEconomySlot)
		flightPb.Status = string(fl.Status)
		flightPb.CreatedAt = timestamppb.New(fl.CreatedAt)
		flightPb.UpdatedAt = timestamppb.New(fl.UpdatedAt)

		flightPbRes = append(flightPbRes, flightPb)
	}

	pRes := &pb.ListFlightResponse{
		FlightList: flightPbRes,
		Total:      pg.Total,
		Page:       pg.Page,
	}
	return pRes, nil
}
func (h *FlightHandler) UpdateFlightSlot(ctx context.Context, req *pb.UpdateFlightSlotRequest) (*pb.Flight, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Internal, "failed to get metadata from context")
	}

	updateType, ok := md["update"]
	if !ok || len(updateType) == 0 {
		return nil, status.Error(codes.InvalidArgument, "update type not provided in metadata")
	}

	fl, err := h.flightRepo.GetFlightById(ctx, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "flight not found")
		}
		return nil, status.Error(codes.Internal, "failed to get flight: "+err.Error())
	}

	switch updateType[0] {
	case "create":
		if req.TicketId == 2 {
			fl.AvailableEconomySlot--
		} else {
			fl.AvailableFirstSlot--
		}
	case "delete":
		if req.TicketId == 2 {
			fl.AvailableEconomySlot++
		} else {
			fl.AvailableFirstSlot++
		}
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid update type provided")
	}

	updatedFlight, err := h.flightRepo.UpdateFlight(ctx, req.Id, fl)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to update flight: "+err.Error())
	}

	return &pb.Flight{Id: updatedFlight.Id}, nil
}

func (h *FlightHandler) DeleteFlight(ctx context.Context, req *pb.DeleteFlightRequest) (*empty.Empty, error) {
	err := h.flightRepo.DeleteFlight(ctx, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "flight not found")
		}
		return nil, err
	}
	return nil, nil
}
