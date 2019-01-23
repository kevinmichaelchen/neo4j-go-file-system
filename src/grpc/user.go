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
	return &pb.CreateUserResponse{User: toGrpcUser(u)}, nil
}

func GetUser(userService user.Service, ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	u, svcError := userService.GetUser(user.User{
		ResourceID: in.UserID,
	})
	if svcError.Error != nil {
		return nil, svcError.Error
	}
	return &pb.GetUserResponse{User: toGrpcUser(u)}, nil
}

func UpdateUser(userService user.Service, ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	u, svcError := userService.UpdateUser(toUser(in.User))
	if svcError.Error != nil {
		return nil, svcError.Error
	}
	return &pb.UpdateUserResponse{User: toGrpcUser(u)}, nil
}

func DeleteUser(userService user.Service, ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	u, svcError := userService.DeleteUser(user.User{
		ResourceID: in.UserID,
	})
	if svcError.Error != nil {
		return nil, svcError.Error
	}
	return &pb.DeleteUserResponse{User: toGrpcUser(u)}, nil
}

func toUser(u *pb.User) user.User {
	return user.User{
		ResourceID:   u.UserID,
		EmailAddress: u.EmailAddress,
		FullName:     u.FullName,
	}
}

func toGrpcUser(u *user.User) *pb.User {
	return &pb.User{
		UserID:       u.ResourceID,
		EmailAddress: u.EmailAddress,
		FullName:     u.FullName,
	}
}
