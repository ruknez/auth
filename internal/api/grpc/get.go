package grpc

import (
	"context"
	"fmt"
	"log/slog"

	desc "auth/pkg/note_v1"
)

// Get ручка получения информации о пользователе по его идентификатору.
func (s *NoteV1GrpcService) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	_ = req

	slog.Info("Get method", "req", fmt.Sprintf("%+v", req))
	return nil, nil
}
