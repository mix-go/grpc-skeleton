package services

import (
    "context"
    "github.com/mix-go/mix-grpc-skeleton/globals"
    "github.com/mix-go/mix-grpc-skeleton/models"
    pb "github.com/mix-go/mix-grpc-skeleton/protos"
    "time"
)

type UserService struct {
}

func (t *UserService) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
    db := globals.DB()
    user := models.User{
        Name:     in.Name,
        CreateAt: time.Now(),
    }
    if err := db.Create(&user).Error; err != nil {
        return nil, err
    }
    resp := pb.AddResponse{
        ErrorCode:    0,
        ErrorMessage: "",
        UserId:       user.ID,
    }
    return &resp, nil
}
