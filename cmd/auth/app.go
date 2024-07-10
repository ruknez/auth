package main

import (
	"fmt"

	"github.com/ruknez/auth/internal/api/grpc/user"
	"github.com/ruknez/auth/internal/app"
	"github.com/ruknez/auth/internal/app/grpc_transport"
)

func main() {
	noteV1 := user.NewUserV1GrpcService()
	transport, err := grpc_transport.NewTransportService()
	if err != nil {
		panic(fmt.Errorf("creating transport service: %w", err))
	}

	if err = app.StartGrpcServer(transport, noteV1); err != nil {
		panic(fmt.Errorf("starting grpc server: %w", err))
	}
}
