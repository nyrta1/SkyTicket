package handlers

import (
	"SkyTicket/AuthService/entity"
	"SkyTicket/AuthService/internal/auth"
	"SkyTicket/AuthService/middleware"
	"SkyTicket/AuthService/repo"
	"SkyTicket/proto/pb"
	"context"
	"database/sql"
	_ "database/sql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

type UserHandler struct {
	pb.UnimplementedUserManagerServer
	userRepo repo.UserRepository
}

func JWTUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	email, err := middleware.JWTMiddleware(ctx)
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, "email", email)

	return handler(ctx, req)
}

func NewUserHandler(userRepo repo.UserRepository) (*UserHandler, error) {
	return &UserHandler{
		userRepo: userRepo,
	}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	email, err := middleware.JWTMiddleware(ctx)
	if err != nil {
		return nil, err
	}

	u, err := h.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	pUser := &pb.User{
		Id:        u.ID,
		Email:     u.Email,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}
	return pUser, nil
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error hashing password: "+err.Error())
	}

	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	createdUser, err := h.userRepo.CreateUser(ctx, user)
	//_, err = h.userRepo.CreateUserRole(ctx, user)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error creating user: "+err.Error())
	}

	response := &pb.User{
		Id:    createdUser.ID,
		Name:  createdUser.Name,
		Email: createdUser.Email,
	}
	return response, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
	hashedPassword, err := auth.HashPassword(req.User.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error hashing password: "+err.Error())
	}
	user := &entity.User{
		Name:     req.User.Name,
		Email:    req.User.Email,
		Password: hashedPassword,
	}

	updatedUser, err := h.userRepo.UpdateUser(ctx, req.User.Id, user)
	pRes := &pb.User{
		Id:       updatedUser.ID,
		Email:    updatedUser.Email,
		Password: updatedUser.Password,
	}
	return pRes, nil
}

func (h *UserHandler) UpdateUserPassword(ctx context.Context, req *pb.UpdateUserPasswordRequest) (*pb.User, error) {
	md, _ := metadata.FromIncomingContext(ctx)

	if md["user"][0] != "0" {
		idConV, _ := strconv.Atoi(md["user"][0])
		u, err := h.userRepo.GetUser(ctx, int64(idConV))
		if err != nil {
			return nil, err
		}
		if !auth.CheckPasswordHash(req.OldPassword, u.Password) {
			return nil, status.Error(codes.InvalidArgument, "old password is not correct")
		}

		passHashed, err := auth.HashPassword(req.Password)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		userUpdated, err := h.userRepo.UpdateUserPassword(ctx, u.ID, passHashed)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return &pb.User{Id: userUpdated.ID}, nil
	}
	if md["user"][0] == "0" {
		u, err := h.userRepo.UpdateUserPassword(ctx, req.Id, req.Password)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return &pb.User{Id: u.ID}, nil
	}
	return nil, nil
}

func (h *UserHandler) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.User, error) {
	u, err := h.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "user not found")
		}
	}

	pRes := &pb.User{
		Id:        u.ID,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}
	return pRes, nil
}

func (h *UserHandler) Login(ctx context.Context, loginRequest *pb.LoginRequest) (*pb.LoginResponse, error) {
	u, err := h.userRepo.GetUserByEmail(ctx, loginRequest.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !auth.CheckPasswordHash(loginRequest.Password, u.Password) {
		return nil, status.Error(codes.Unauthenticated, "wrong username or password")
	}
	token, err := auth.GenerateToken(loginRequest.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	//_, err = h.userRepo.ChangeUserRole(ctx, loginRequest.Email)
	//if err != nil {
	//	log.Printf("Error changing user role: %v", err)
	//}

	pRes := &pb.LoginResponse{
		JwtToken: token,
	}
	return pRes, nil
}

func (h *UserHandler) ParseToken(ctx context.Context, tokenRequest *pb.ParseTokenRequest) (*pb.ParseTokenResponse, error) {
	email, err := auth.ParseToken(tokenRequest.Token)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	u, err := h.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	pUser := &pb.User{
		Id:        u.ID,
		Email:     u.Email,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}
	pRes := &pb.ParseTokenResponse{
		User: pUser,
	}
	return pRes, nil
}
