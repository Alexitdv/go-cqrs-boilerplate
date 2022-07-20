package grpc

import (
	"boilerplate/internal/app"
	pb "boilerplate/internal/ports/grpc/server/boilderplate/v1"
)

type GrpcServer struct {
	app *app.Application
	pb.UnimplementedBoilerplateServer
}

func NewGrpcServer(application *app.Application) GrpcServer {
	return GrpcServer{app: application}
}
