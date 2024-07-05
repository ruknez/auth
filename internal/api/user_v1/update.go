package user_v1

import (
	"context"
	"fmt"
	"log/slog"

	desc "auth/pkg/user_v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

// Update ручка обновления информации о пользователе по его идентификатору.
func (s *UserV1GrpcService) Update(_ context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	_ = req

	slog.Info("Update method", "req", fmt.Sprintf("%+v", req))
	return nil, nil
}
