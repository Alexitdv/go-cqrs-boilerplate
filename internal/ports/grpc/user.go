package grpc

import (
	"context"

	"github.com/google/uuid"

	"boilerplate/internal/domain/user"
	pb "boilerplate/internal/ports/grpc/server/boilderplate/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (g *GrpcServer) GetUser(ctx context.Context, in *pb.Identifier) (*pb.User, error) {
	id, err := uuid.Parse(in.Uuid)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "uuid is not correct")
	}
	u, err := g.app.Queries.GetUser.Handle(ctx, id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return QueryUserToPBUser(u)
}

func (g *GrpcServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.SimpleResponse, error) {
	ok, err := g.app.Commands.Login.Handle(ctx, in.Phone, in.Password)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	if !ok {
		return nil, status.Errorf(codes.PermissionDenied, "Incorrect login or password")
	}
	return &pb.SimpleResponse{Status: "ok"}, nil
}

func (g *GrpcServer) AddUser(ctx context.Context, pbUser *pb.User) (*pb.Identifier, error) {
	opts := user.NewOptions(
		user.WithName(pbUser.Name.GetValue()),
		user.WithLastName(pbUser.LastName.GetValue()),
		user.WithEmail(pbUser.Email.GetValue()),
		user.WithPassword(pbUser.Password.GetValue()),
		user.WithPhone(pbUser.Phone.GetValue()),
	)
	if pbUser.Uuid != nil {
		opts.Append(user.WithID(pbUser.Uuid.GetValue()))
	}
	u, err := user.NewUser(opts...)
	if err != nil {
		return nil, err
	}

	id, err := g.app.Commands.AddUser.Handle(ctx, u)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.Identifier{Uuid: id}, nil
}

func (g *GrpcServer) UpdateUser(ctx context.Context, pbUser *pb.User) (*emptypb.Empty, error) {
	if pbUser.Uuid == nil {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "UUID is empty. ")
	}
	opts := user.NewOptions(
		user.WithID(pbUser.Uuid.GetValue()),
	)
	if pbUser.Name != nil {
		opts.Append(user.WithName(pbUser.Name.GetValue()))
	}
	if pbUser.LastName != nil {
		opts.Append(user.WithLastName(pbUser.LastName.GetValue()))
	}
	if pbUser.Email != nil {
		opts.Append(user.WithEmail(pbUser.Email.GetValue()))
	}
	if pbUser.Phone != nil {
		opts.Append(user.WithPhone(pbUser.Phone.GetValue()))
	}

	u, err := user.NewUser(opts...)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = g.app.Commands.UpdateUser.Handle(ctx, u)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (g *GrpcServer) DeleteUser(ctx context.Context, id *pb.Identifier) (*emptypb.Empty, error) {
	if err := g.app.Commands.DeleteUser.Handle(ctx, id.Uuid); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (g *GrpcServer) UpdatePassword(ctx context.Context, in *pb.UpdatePasswordRequest) (*emptypb.Empty, error) {
	if err := g.app.Commands.UpdatePassword.Handle(ctx, in.Phone, in.OldPassword, in.NewPassword); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
