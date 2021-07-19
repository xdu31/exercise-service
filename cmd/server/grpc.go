package main

import (
	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"

	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	grpc_preprocess "github.com/infobloxopen/protoc-gen-preprocess/middleware"

	pb "github.com/xdu31/exercise-service/pkg/server/pb"
	"github.com/xdu31/exercise-service/pkg/server/svc"
)

// NewGRPCServer creates a new GRPC Server
func NewGRPCServer(logger *logrus.Logger) (*grpc.Server, error) {
	interceptors := []grpc.UnaryServerInterceptor{
		grpc_preprocess.UnaryServerInterceptor(),                    // preprocessing middleware
		grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logger)), // logging middleware
		gateway.UnaryServerInterceptor(),                            // collection operators middleware
		AuthZUnaryServerInterceptor(),
	}

	options := grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(interceptors...))
	// create new gRPC grpcServer with middleware chain
	grpcServer := grpc.NewServer(options)

	// register service implementation with the grpcServer
	pb.RegisterExerciseServiceServer(grpcServer, svc.NewExerciseServer())

	return grpcServer, nil
}
