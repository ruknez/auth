package grpc

import (
	desc "auth/pkg/note_v1"
)

type serviceGrpc struct {
	desc.UnimplementedNoteV1Server
}

// NewServiceGrpc возвращает сервис реализующий хэндлеры.
func NewServiceGrpc() *serviceGrpc {
	return &serviceGrpc{}
}
