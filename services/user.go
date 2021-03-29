package services

import (
    "context"
    "github.com/mix-go/grpc-skeleton/globals"
    "github.com/mix-go/grpc-skeleton/models"
    pb "github.com/mix-go/grpc-skeleton/protos"
    "time"
)

type UserService struct {
}

func (t *UserService) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
    db := globals.Gorm()
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
