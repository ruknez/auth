package user

import (
	"context"
	"fmt"
	"log/slog"

	desc "github.com/ruknez/auth/pkg/user_v1"
)

// Create ручка создания нового пользователя в системе.
func (s *GrpcServiceV1) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	slog.Info("Create method", "req", fmt.Sprintf("%+v", req))

	return nil, nil
}
