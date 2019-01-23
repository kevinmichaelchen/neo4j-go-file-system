package grpc

import (
	"github.com/kevinmichaelchen/neo4j-go-file-system/pb"
	"github.com/kevinmichaelchen/neo4j-go-file-system/user"
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"
)

func CreateUser(userService user.Service, ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	u, svcErr := userService.CreateUser(user.User{
		EmailAddress: in.User.EmailAddress,
		FullName:     in.User.FullName,
	})
	if svcErr != nil {
		return nil, status.Error(svcErr.GrpcCode, svcErr.ErrorMessage)
	}
	return &pb.UserResponse{User: toGrpcUser(u)}, nil
}

func GetUser(userService user.Service, ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	u, svcErr := userService.GetUser(user.User{
		ResourceID: in.User.Id,
	})
	if svcErr != nil {
		return nil, status.Error(svcErr.GrpcCode, svcErr.ErrorMessage)
	}
	return &pb.UserResponse{User: toGrpcUser(u)}, nil
}

func UpdateUser(userService user.Service, ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	u, svcErr := userService.UpdateUser(toUser(in.User))
	if svcErr != nil {
		return nil, status.Error(svcErr.GrpcCode, svcErr.ErrorMessage)
	}
	return &pb.UserResponse{User: toGrpcUser(u)}, nil
}

func DeleteUser(userService user.Service, ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	u, svcErr := userService.DeleteUser(user.User{
		ResourceID: in.User.Id,
	})
	if svcErr != nil {
		return nil, status.Error(svcErr.GrpcCode, svcErr.ErrorMessage)
	}
	return &pb.UserResponse{User: toGrpcUser(u)}, nil
}

func toUser(u *pb.User) user.User {
	return user.User{
		ResourceID:   u.Id,
		EmailAddress: u.EmailAddress,
		FullName:     u.FullName,
	}
}

func toGrpcUser(u *user.User) *pb.User {
	return &pb.User{
		Id:           u.ResourceID,
		EmailAddress: u.EmailAddress,
		FullName:     u.FullName,
	}
}
