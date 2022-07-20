package grpc

import (
	query "boilerplate/internal/app/query/users"
	pb "boilerplate/internal/ports/grpc/server/boilderplate/v1"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

func QueryUserToPBUser(user *query.User) (*pb.User, error) {
	return &pb.User{
		Uuid:     &wrapperspb.StringValue{Value: user.UUID},
		Phone:    &wrapperspb.StringValue{Value: user.Phone},
		Name:     &wrapperspb.StringValue{Value: user.Name},
		LastName: &wrapperspb.StringValue{Value: user.LastName},
		Email:    &wrapperspb.StringValue{Value: user.Email},
	}, nil
}
