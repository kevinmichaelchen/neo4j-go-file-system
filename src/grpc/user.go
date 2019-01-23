package grpc

import (
	"github.com/kevinmichaelchen/neo4j-go-file-system/pb"
	"github.com/kevinmichaelchen/neo4j-go-file-system/user"
	"golang.org/x/net/context"
)

func CreateUser(userService user.Service, ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	u, svcError := userService.CreateUser(user.User{
		EmailAddress: in.User.EmailAddress,
		FullName:     in.User.FullName,
	})
	if svcError.Error != nil {
		return nil, svcError.Error
	}
	return &pb.CreateUserResponse{User: &pb.User{
		UserID:       u.ResourceID,
		EmailAddress: u.EmailAddress,
		FullName:     u.FullName,
	}}, nil
}

func GetUser(userService user.Service, ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return nil, nil
}

func UpdateUser(userService user.Service, ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return nil, nil
}

func DeleteUser(userService user.Service, ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return nil, nil
}
