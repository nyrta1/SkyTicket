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

type CountryHandler struct {
	pb.UnimplementedCountryServiceServer
	countryRepo repository.CountryRepository
}

func NewCountryHandler(countryRepo repository.CountryRepository) (*CountryHandler, error) {
	return &CountryHandler{
		countryRepo: countryRepo,
	}, nil
}

func (h *CountryHandler) CreateCountry(ctx context.Context, req *pb.CreateCountryRequest) (*pb.Country, error) {
	country := &pb.Country{Name: req.Name, Prefix: req.Prefix}

	if err := h.countryRepo.Add(ctx, country); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return country, nil
}

func (h *CountryHandler) GetCountry(ctx context.Context, req *pb.GetCountryRequest) (*pb.Country, error) {
	country, err := h.countryRepo.GetById(ctx, req.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "country not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return country, nil
}

func (h *CountryHandler) ListCountries(ctx context.Context, req *emptypb.Empty) (*pb.ListCountriesResponse, error) {
	countries, err := h.countryRepo.GetAll(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "countries not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &pb.ListCountriesResponse{
		Countries: countries,
	}

	return response, nil
}

func (h *CountryHandler) UpdateCountry(ctx context.Context, req *pb.UpdateCountryRequest) (*pb.Country, error) {
	country := &pb.Country{
		Name:   req.Name,
		Prefix: req.Prefix,
	}

	if err := h.countryRepo.UpdateById(ctx, req.Id, country); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return country, nil
}

func (h *CountryHandler) DeleteCountry(ctx context.Context, req *pb.DeleteCountryRequest) (*emptypb.Empty, error) {
	if err := h.countryRepo.DeleteById(ctx, req.Id); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return nil, nil
}
