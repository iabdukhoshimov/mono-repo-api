package service

import (
	"context"
	"encoding/json"

	"github.com/abdukhashimov/go_api_mono_repo/generated/user_service"
	"github.com/abdukhashimov/go_api_mono_repo/internal/core/repository"
	"github.com/abdukhashimov/go_api_mono_repo/internal/core/repository/psql/sqlc"
	"github.com/abdukhashimov/go_api_mono_repo/internal/transport/queue/distributor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	user_service.UnimplementedUserServiceServer
	store           repository.Store
	taskDistributor distributor.TaskDistributor
}

func NewUserService(store repository.Store, taskDistributor distributor.TaskDistributor) *UserService {
	return &UserService{
		store:           store,
		taskDistributor: taskDistributor,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *user_service.User) (*user_service.User, error) {

	err := s.store.CreateUser(ctx, sqlc.CreateUserParams{
		Name:  req.Name,
		Email: req.Email,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err.Error())
	}

	// TODO: db tx
	taskPayload := &distributor.PayloadSendVerifyEmail{
		Username: req.Name,
	}
	err = s.taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to distribute task: %s", err.Error())
	}

	return &user_service.User{
		Name:  req.Name,
		Email: req.Email,
	}, nil
}

func (s *UserService) GetUserById(ctx context.Context, req *user_service.UserPk) (*user_service.User, error) {

	user, err := s.store.GetUserById(ctx, req.Id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %s", err.Error())
	}
	response := user_service.User{}

	if userInBytes, err := json.Marshal(user); err == nil {
		json.Unmarshal(userInBytes, &response)
	} else {
		return nil, status.Errorf(codes.Internal, "failed to marshal user: %s    ", err.Error())
	}

	return &response, nil
}
