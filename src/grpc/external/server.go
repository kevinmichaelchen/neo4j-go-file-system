package external

import (
	"fmt"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/kevinmichaelchen/neo4j-go-file-system/file"
	"github.com/kevinmichaelchen/neo4j-go-file-system/folder"
	"github.com/kevinmichaelchen/neo4j-go-file-system/move"
	"github.com/kevinmichaelchen/neo4j-go-file-system/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	Port          int
	MoveService   move.Service
	FileService   file.Service
	FolderService folder.Service
}

func (s *Server) EmitEvent(ctx context.Context, in *pb.FileRequest) (*pb.FileResponse, error) {
	switch e := in.FileEvent.(type) {
	case *pb.FileRequest_CreateEvent:
		return CreateFile(s.FileService, ctx, e.CreateEvent)
	case *pb.FileRequest_UpdateEvent:
		return UpdateFile(s.FileService, ctx, e.UpdateEvent)
	case *pb.FileRequest_DeleteEvent:
		return DeleteFile(s.FileService, ctx, e.DeleteEvent)
	}
	return nil, status.Error(codes.InvalidArgument, "Unsupported event type")
}

func (s *Server) CreateFile(ctx context.Context, in *pb.CreateFileRequest) (*pb.FileResponse, error) {
	return CreateFile(s.FileService, ctx, in)
}

func (s *Server) GetFile(ctx context.Context, in *pb.GetFileRequest) (*pb.FileResponse, error) {
	return GetFile(s.FileService, ctx, in)
}

func (s *Server) UpdateFile(ctx context.Context, in *pb.UpdateFileRequest) (*pb.FileResponse, error) {
	return UpdateFile(s.FileService, ctx, in)
}

func (s *Server) DeleteFile(ctx context.Context, in *pb.DeleteFileRequest) (*pb.FileResponse, error) {
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
	pb.RegisterFileServiceServer(server, s)
	pb.RegisterFolderServiceServer(server, s)

	// Register reflection service on gRPC server.
	reflection.Register(server)
	log.Println("Registered gRPC services...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
