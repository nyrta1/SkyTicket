package controller

import (
	"SkyTicket/AirplaneService/internal/repository"
	"SkyTicket/proto/pb"
	"context"
	"database/sql"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ManufacturerHandler struct {
	pb.UnimplementedManufacturerServiceServer
	manufacturerRepo repository.ManufacturerRepository
}

func NewManufacturerHandler(manufacturerRepo repository.ManufacturerRepository) (*ManufacturerHandler, error) {
	return &ManufacturerHandler{
		manufacturerRepo: manufacturerRepo,
	}, nil
}

func (h *ManufacturerHandler) CreateManufacturer(ctx context.Context, req *pb.CreateManufacturerRequest) (*pb.ManufacturerResponse, error) {
	if err := h.manufacturerRepo.Add(ctx, req.Manufacturer); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &pb.ManufacturerResponse{
		Manufacturer: req.Manufacturer,
	}

	return res, nil
}

func (h *ManufacturerHandler) GetManufacturer(ctx context.Context, req *pb.GetManufacturerRequest) (*pb.ManufacturerResponse, error) {
	manufacturer, err := h.manufacturerRepo.GetById(ctx, req.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "manufacturer not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &pb.ManufacturerResponse{
		Manufacturer: manufacturer,
	}

	return res, nil
}

func (h *ManufacturerHandler) ListManufacturers(ctx context.Context, req *emptypb.Empty) (*pb.ListManufacturersResponse, error) {
	manufacturers, err := h.manufacturerRepo.GetAll(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "manufacturers not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &pb.ListManufacturersResponse{Manufacturers: manufacturers}

	return res, nil
}

func (h *ManufacturerHandler) UpdateManufacturer(ctx context.Context, req *pb.UpdateManufacturerRequest) (*pb.ManufacturerResponse, error) {
	if err := h.manufacturerRepo.UpdateById(ctx, req.Manufacturer.Id, req.Manufacturer); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &pb.ManufacturerResponse{
		Manufacturer: req.Manufacturer,
	}

	return res, nil
}

func (h *ManufacturerHandler) DeleteManufacturer(ctx context.Context, req *pb.DeleteManufacturerRequest) (*emptypb.Empty, error) {
	if err := h.manufacturerRepo.DeleteById(ctx, req.Id); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return nil, nil
}
