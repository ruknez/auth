package user

import (
	desc "github.com/ruknez/auth/pkg/user_v1"
)

// GrpcServiceV1 реализация ручке grpc.
type GrpcServiceV1 struct {
	desc.UnimplementedUserV1Server
}

// NewUserV1GrpcService возвращает сервис реализующий хэндлеры.
func NewUserV1GrpcService() *GrpcServiceV1 {
	return &GrpcServiceV1{}
}
