package grpc

import (
	"context"

	desc "auth/pkg/note_v1"
)

// Get ручка получения информации о пользователе по его идентификатору.
func (s *serviceGrpc) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	_ = req

	return nil, nil
}
