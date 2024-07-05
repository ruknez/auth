package app

import (
	"fmt"
	"log/slog"
	"net"

	desc "auth/pkg/note_v1"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// StartGrpcServer запускает grpc сервис.
func StartGrpcServer(transport net.Listener, noteV1 desc.NoteV1Server) error {
	s := grpc.NewServer()
	reflection.Register(s)

	desc.RegisterNoteV1Server(s, noteV1)

	slog.Info(fmt.Sprintf("GRPC server start at %s", transport.Addr()))

	return errors.Wrap(s.Serve(transport), "failed to start grpc server")
}
