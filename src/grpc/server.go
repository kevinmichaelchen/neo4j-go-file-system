package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/kevinmichaelchen/neo4j-go-file-system/file"
	"github.com/kevinmichaelchen/neo4j-go-file-system/folder"
	"github.com/kevinmichaelchen/neo4j-go-file-system/move"
	"github.com/kevinmichaelchen/neo4j-go-file-system/organization"

	"github.com/kevinmichaelchen/neo4j-go-file-system/pb"
	"github.com/kevinmichaelchen/neo4j-go-file-system/user"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// TODO use error codes https://github.com/grpc/grpc-go/blob/master/Documentation/rpc-errors.md
type Server struct {
	Port                int
	UserService         user.Service
	OrganizationService organization.Service
	MoveService         move.Service
	FileService         file.Service
	FolderService       folder.Service
}

func (s *Server) CreateOrganization(ctx context.Context, in *pb.CreateOrgRequest) (*pb.CreateOrgResponse, error) {
	return CreateOrganization(s.OrganizationService, ctx, in)
}

func (s *Server) GetOrganization(ctx context.Context, in *pb.OrganizationCrudRequest) (*pb.OrganizationResponse, error) {
	return GetOrganization(s.OrganizationService, ctx, in)
}

func (s *Server) UpdateOrganization(ctx context.Context, in *pb.OrganizationCrudRequest) (*pb.OrganizationResponse, error) {
	return UpdateOrganization(s.OrganizationService, ctx, in)
}

func (s *Server) DeleteOrganization(ctx context.Context, in *pb.OrganizationCrudRequest) (*pb.OrganizationResponse, error) {
	return DeleteOrganization(s.OrganizationService, ctx, in)
}

func (s *Server) AddUserToOrganization(ctx context.Context, in *pb.AddUserToOrganizationRequest) (*pb.OrganizationResponse, error) {
	return AddUserToOrganization(s.OrganizationService, ctx, in)
}

func (s *Server) RemoveUserFromOrganization(ctx context.Context, in *pb.RemoveUserFromOrganizationRequest) (*pb.OrganizationResponse, error) {
	return RemoveUserFromOrganization(s.OrganizationService, ctx, in)
}

func (s *Server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return CreateUser(s.UserService, ctx, in)
}

func (s *Server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return GetUser(s.UserService, ctx, in)
}

func (s *Server) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return UpdateUser(s.UserService, ctx, in)
}

func (s *Server) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return DeleteUser(s.UserService, ctx, in)
}

func (s *Server) CreateFile(ctx context.Context, in *pb.CreateFileRequest) (*pb.CreateFileResponse, error) {
	return CreateFile(s.FileService, ctx, in)
}

func (s *Server) GetFile(ctx context.Context, in *pb.GetFileRequest) (*pb.GetFileResponse, error) {
	return GetFile(s.FileService, ctx, in)
}

func (s *Server) UpdateFile(ctx context.Context, in *pb.UpdateFileRequest) (*pb.UpdateFileResponse, error) {
	return UpdateFile(s.FileService, ctx, in)
}

func (s *Server) DeleteFile(ctx context.Context, in *pb.DeleteFileRequest) (*pb.DeleteFileResponse, error) {
	return DeleteFile(s.FileService, ctx, in)
}

func (s *Server) Run() {
	address := fmt.Sprintf(":%d", s.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Starting gRPC server on %s...\n", address)
	server := grpc.NewServer()

	// Register our services
	pb.RegisterUserServiceServer(server, s)
	pb.RegisterFileServiceServer(server, s)
	pb.RegisterOrganizationServiceServer(server, s)

	// Register reflection service on gRPC server.
	reflection.Register(server)
	log.Println("Registered gRPC services...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
