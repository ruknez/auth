package grpc

import (
	"context"

	desc "auth/pkg/note_v1"
)

// Create ручка создания нового пользователя в системе.
func (s *serviceGrpc) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	_ = req
	return nil, nil
}
