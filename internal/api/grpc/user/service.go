package user

import (
	desc "github.com/ruknez/auth/pkg/user_v1"
)

// UserV1GrpcService реализация ручке grpc.
type UserV1GrpcService struct {
	desc.UnimplementedUserV1Server
}

// NewUserV1GrpcService возвращает сервис реализующий хэндлеры.
func NewUserV1GrpcService() *UserV1GrpcService {
	return &UserV1GrpcService{}
}
