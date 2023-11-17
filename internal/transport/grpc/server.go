package grpc

import (
	"github.com/abdukhashimov/go_api_mono_repo/generated/user_service"
	"github.com/abdukhashimov/go_api_mono_repo/internal/config"
	"github.com/abdukhashimov/go_api_mono_repo/internal/core/repository"
	"github.com/abdukhashimov/go_api_mono_repo/internal/core/service"
	"github.com/abdukhashimov/go_api_mono_repo/internal/transport/grpc/middleware"
	"github.com/abdukhashimov/go_api_mono_repo/internal/transport/queue/distributor"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func New(repo repository.Store, cfg *config.Config, taskDistributor distributor.TaskDistributor) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				middleware.GrpcLoggerMiddleware,
			),
		),
	)

	reflection.Register(grpcServer)
	user_service.RegisterUserServiceServer(grpcServer, service.NewUserService(repo, taskDistributor))
	return grpcServer
}
