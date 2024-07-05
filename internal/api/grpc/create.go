package grpc

import (
	"context"
	"fmt"
	"log/slog"

	desc "auth/pkg/note_v1"
)

// Create ручка создания нового пользователя в системе.
func (s *NoteV1GrpcService) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	_ = req

	slog.Info("Create method", "req", fmt.Sprintf("%+v", req))
	return nil, nil
}
