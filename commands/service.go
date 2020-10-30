package commands

import (
    "context"
    "github.com/mix-go/mix-grpc-skeleton/globals"
    "github.com/mix-go/mix-grpc-skeleton/models"
    pb "github.com/mix-go/mix-grpc-skeleton/protos"
    "time"
)

type UserService struct {
}

func (t *UserService) Add(context.Context, *pb.AddRequest) (*pb.AddResponse, error) {
    db := globals.DB()
    user := models.User{
        ID:       0,
        Name:     "",
        CreateAt: time.Time{},
    }
    if err := db.Create(&user).Error; err != nil {
        return nil, err
    }
    resp := pb.AddResponse{
        ErrorCode:    0,
        ErrorMessage: "",
        OrderId:      user.ID,
    }
    return &resp, nil
}
