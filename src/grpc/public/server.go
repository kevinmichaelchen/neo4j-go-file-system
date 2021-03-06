package public

import (
	"fmt"
	"log"
	"net"

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

	// Register reflection service on gRPC server.
	reflection.Register(server)
	log.Println("Registered gRPC services...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
