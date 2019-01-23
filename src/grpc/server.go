package grpc

import (
	"fmt"
	"log"
	"net"
	"strings"

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

type Server struct {
	Port                int
	UserService         user.Service
	OrganizationService organization.Service
	MoveService         move.Service
	FileService         file.Service
	FolderService       folder.Service
}

func (s *Server) CreateOrganization(ctx context.Context, in *pb.OrganizationCrudRequest) (*pb.OrganizationResponse, error) {
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

func (s *Server) CreateUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	return CreateUser(s.UserService, ctx, in)
}

func (s *Server) GetUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	return GetUser(s.UserService, ctx, in)
}

func (s *Server) UpdateUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	return UpdateUser(s.UserService, ctx, in)
}

func (s *Server) DeleteUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
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

func (s *Server) CreateFolder(ctx context.Context, in *pb.CreateFolderRequest) (*pb.FolderResponse, error) {
	return CreateFolder(s.FolderService, ctx, in)
}

func (s *Server) GetFolder(ctx context.Context, in *pb.GetFolderRequest) (*pb.FolderResponse, error) {
	return GetFolder(s.FolderService, ctx, in)
}

func (s *Server) UpdateFolder(ctx context.Context, in *pb.UpdateFolderRequest) (*pb.FolderResponse, error) {
	return UpdateFolder(s.FolderService, ctx, in)
}

func (s *Server) DeleteFolder(ctx context.Context, in *pb.DeleteFolderRequest) (*pb.FolderResponse, error) {
	return DeleteFolder(s.FolderService, ctx, in)
}

// optString is a utility function we use for passing "optional" strings to our service layer
// Absent string fields come in as the empty string by default.
// I'd prefer to work with string pointers rather than empty strings.
// -Kevin
func optString(s string) *string {
	if strings.TrimSpace(s) == "" {
		return nil
	}
	return &s
}

func getUserID(ctx context.Context) int {
	// TODO properly fetch user ID from headers or metadata
	return 1
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
	pb.RegisterOrganizationServiceServer(server, s)
	pb.RegisterFileServiceServer(server, s)
	pb.RegisterFolderServiceServer(server, s)

	// Register reflection service on gRPC server.
	reflection.Register(server)
	log.Println("Registered gRPC services...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
