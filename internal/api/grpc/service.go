package grpc

import (
	desc "auth/pkg/note_v1"
)

// NoteV1GrpcService реализация ручке grpc.
type NoteV1GrpcService struct {
	desc.UnimplementedNoteV1Server
}

// NewNoteV1GrpcService возвращает сервис реализующий хэндлеры.
func NewNoteV1GrpcService() *NoteV1GrpcService {
	return &NoteV1GrpcService{}
}
