package commands

import (
    "context"
    "fmt"
    "github.com/mix-go/mix-grpc-skeleton/globals"
    pb "github.com/mix-go/mix-grpc-skeleton/protos"
    "google.golang.org/grpc"
    "net"
    "os"
    "os/signal"
    "strings"
    "syscall"
    "time"
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
    pb.RegisterUserServer(s, &UserService{})
    logger.Infof("Server run %s", Addr)
    if err := s.Serve(listener); err != nil && !strings.Contains(err.Error(), "use of closed network connection") {
        panic(err)
    }
}

type GrpcClientCommand struct {
}

func (t *GrpcClientCommand) Main() {
    ctx, _ := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
    conn, err := grpc.DialContext(ctx, Addr, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        panic(err)
    }
    defer func() {
        _ = conn.Close()
    }()

    cli := pb.NewUserClient(conn)
    req := pb.AddRequest{
        Name: "xiaoliu",
    }
    resp, err := cli.Add(ctx, &req)
    if err != nil {
        panic(err)
    }
    fmt.Println(fmt.Sprintf("Add order: %d", resp.OrderId))
}
