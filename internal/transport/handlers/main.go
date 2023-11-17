package handlers

import (
	"context"
	"fmt"

	"github.com/abdukhashimov/go_api_mono_repo/generated/user_service"
	"github.com/abdukhashimov/go_api_mono_repo/internal/config"
	"github.com/abdukhashimov/go_api_mono_repo/pkg/wrapper"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	mainGrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New(ctx context.Context, cfg *config.Config) *runtime.ServeMux {
	gwMux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(wrapper.CustomMatcher),
	)
	connPingService, err := mainGrpc.Dial(
		fmt.Sprintf("%s:%d", cfg.Grpc.Host, cfg.Grpc.Port),
		mainGrpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil
	}

	if err := user_service.RegisterUserServiceHandler(ctx, gwMux, connPingService); err != nil {
		return nil
	}

	return gwMux
}

func makeHost(host string, port int32) string {
	return host + ":" + fmt.Sprintf("%d", port)
}
