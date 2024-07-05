package grpc

import (
	"context"

	desc "auth/pkg/note_v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

// Update ручка обновления информации о пользователе по его идентификатору.
func (s *serviceGrpc) Update(_ context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	_ = req

	return nil, nil
}
