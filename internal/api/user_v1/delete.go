package user_v1

import (
	"context"
	"fmt"
	"log/slog"

	desc "auth/pkg/user_v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete удаление пользователя из системы по его идентификатору.
func (s *UserV1GrpcService) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	_ = req

	slog.Info("Delete method", "req", fmt.Sprintf("%+v", req))
	return nil, nil
}
