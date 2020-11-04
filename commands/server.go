package commands

import (
    "github.com/mix-go/mix-grpc-skeleton/globals"
    pb "github.com/mix-go/mix-grpc-skeleton/protos"
    "github.com/mix-go/mix-grpc-skeleton/services"
    "google.golang.org/grpc"
    "net"
    "os"
    "os/signal"
    "strings"
    "syscall"
)

const Addr = ":8181"

type GrpcServerCommand struct {
}

func (t *GrpcServerCommand) Main() {

    logger := globals.Logger()

    listener, err := net.Listen("tcp", Addr)
    if err != nil {
        panic(err)
    }

    ch := make(chan os.Signal)
    signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-ch
        logger.Info("Server shutdown")
        if err := listener.Close(); err != nil {
            panic(err)
        }
    }()

    s := grpc.NewServer()
    pb.RegisterUserServer(s, &services.UserService{})
    logger.Infof("Server run %s", Addr)
    if err := s.Serve(listener); err != nil && !strings.Contains(err.Error(), "use of closed network connection") {
        panic(err)
    }
}
