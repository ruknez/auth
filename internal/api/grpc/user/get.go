package user

import (
	"context"
	"fmt"
	"log/slog"

	desc "github.com/ruknez/auth/pkg/user_v1"
)

// Get ручка получения информации о пользователе по его идентификатору.
func (s *UserV1GrpcService) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	slog.Info("Get method", "req", fmt.Sprintf("%+v", req))

	return nil, nil
}
