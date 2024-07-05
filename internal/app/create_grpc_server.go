package app

import (
	"fmt"
	"log/slog"
	"net"

	desc "auth/pkg/user_v1"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// StartGrpcServer запускает grpc сервис.
func StartGrpcServer(transport net.Listener, userV1 desc.UserV1Server) error {
	s := grpc.NewServer()
	reflection.Register(s)

	desc.RegisterUserV1Server(s, userV1)

	slog.Info(fmt.Sprintf("GRPC server start at %s", transport.Addr()))

	return errors.Wrap(s.Serve(transport), "failed to start grpc server")
}
