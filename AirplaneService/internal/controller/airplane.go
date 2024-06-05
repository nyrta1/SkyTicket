package controller

import (
	"AirplaneService/internal/entity"
	"AirplaneService/internal/grpc"
	"AirplaneService/internal/repository"
	"context"
	"database/sql"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AirplaneHandler struct {
	grpc.UnimplementedAirplaneServiceServer
	airplaneRepo repository.AirplaneRepository
}

func NewBookingHandler(bookingRepo repository.AirplaneRepository) (*AirplaneHandler, error) {
	return &AirplaneHandler{
		airplaneRepo: bookingRepo,
	}, nil
}

func (h *AirplaneHandler) GetAirplane(ctx context.Context, airplaneInput *entity.GetAirplaneRequest) (*entity.Airplane, error) {
	airplane, err := h.airplaneRepo.GetById(ctx, airplaneInput.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "airplane not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return airplane, nil
}

func (h *AirplaneHandler) ListAirplanes(ctx context.Context, eb *emptypb.Empty) (*entity.ListAirplanesResponse, error) {
	airplanes, err := h.airplaneRepo.GetAll(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "airplanes not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	airplanesResponse := &entity.ListAirplanesResponse{
		Airplanes: airplanes,
	}

	return airplanesResponse, nil
}

func (h *AirplaneHandler) CreateAirplane(ctx context.Context, request *entity.CreateAirplaneRequest) (*entity.Airplane, error) {
	if err := h.airplaneRepo.Add(ctx, request.Airplane); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return request.Airplane, nil
}

func (h *AirplaneHandler) UpdateAirplane(ctx context.Context, request *entity.UpdateAirplaneRequest) (*entity.Airplane, error) {
	if err := h.airplaneRepo.UpdateById(ctx, request.Airplane.Id, request.Airplane); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return request.Airplane, nil
}

func (h *AirplaneHandler) DeleteAirplane(ctx context.Context, request *entity.DeleteAirplaneRequest) (*emptypb.Empty, error) {
	if err := h.airplaneRepo.DeleteById(ctx, request.Id); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return nil, nil
}
