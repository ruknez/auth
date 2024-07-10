package user

import (
	"context"
	"fmt"
	"log/slog"

	desc "github.com/ruknez/auth/pkg/user_v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

// Update ручка обновления информации о пользователе по его идентификатору.
func (s *GrpcServiceV1) Update(_ context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	slog.Info("Update method", "req", fmt.Sprintf("%+v", req))

	return nil, nil
}
