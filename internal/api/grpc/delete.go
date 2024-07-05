package grpc

import (
	"context"

	desc "auth/pkg/note_v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete удаление пользователя из системы по его идентификатору.
func (s *serviceGrpc) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	_ = req

	return nil, nil
}
