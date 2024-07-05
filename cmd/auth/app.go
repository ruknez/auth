package main

import (
	"fmt"

	"auth/internal/api/user_v1"
	"auth/internal/app"
	"auth/internal/app/grpc_transpotr"
)

func main() {
	noteV1 := user_v1.NewUserV1GrpcService()
	transport, err := grpc_transpotr.NewTransportService()
	if err != nil {
		panic(fmt.Sprint("Error creating transport service: ", err))
	}

	if err = app.StartGrpcServer(transport, noteV1); err != nil {
		panic(fmt.Sprint("Error starting grpc server: ", err))
	}
}
