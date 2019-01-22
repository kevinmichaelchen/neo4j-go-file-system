package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/kevinmichaelchen/neo4j-go-file-system/file"
	"github.com/kevinmichaelchen/neo4j-go-file-system/folder"
	"github.com/kevinmichaelchen/neo4j-go-file-system/move"
	"github.com/kevinmichaelchen/neo4j-go-file-system/organization"

	pb "github.com/kevinmichaelchen/neo4j-go-file-system/pb"
	"github.com/kevinmichaelchen/neo4j-go-file-system/user"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	Port                int
	UserService         user.Service
	OrganizationService organization.Service
	MoveService         move.Service
	FileService         file.Service
	FolderService       folder.Service
}

func (s *Server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	u, svcError := s.UserService.CreateUser(user.User{
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

func (s *Server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return nil, nil
}

func (s *Server) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return nil, nil
}

func (s *Server) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return nil, nil
}

func (s *Server) Run() {
	address := fmt.Sprintf(":%d", s.Port)
	log.Printf("Serving gRPC on %s\n", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Starting gRPC server...")
	server := grpc.NewServer()

	// Register our services
	pb.RegisterUserServiceServer(server, s)

	// Register reflection service on gRPC server.
	reflection.Register(server)
	log.Println("Registered gRPC services...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
