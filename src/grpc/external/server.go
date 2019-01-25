package external

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	"strings"

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

func (s *Server) EmitEvent(stream pb.EventService_EmitEventServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		switch e := in.Event.(type) {
		case *pb.EventRequest_CreateFileEvent:
			log.Println("CREATE FILE EVENT")
			_, err := CreateFile(s.FileService, stream.Context(), e.CreateFileEvent)
			if err != nil {
				return err
			}
		case *pb.EventRequest_UpdateFileEvent:
			log.Println("UPDATE FILE EVENT")
		case *pb.EventRequest_DeleteFileEvent:
			log.Println("DELETE FILE EVENT")
		case *pb.EventRequest_CreateFolderEvent:
			log.Println("CREATE FOLDER EVENT")
		case *pb.EventRequest_UpdateFolderEvent:
			log.Println("UPDATE FOLDER EVENT")
		case *pb.EventRequest_DeleteFolderEvent:
			log.Println("DELETE FOLDER EVENT")
		default:
			return status.Error(codes.Unimplemented, "Unsupported event detected")
		}

		if err := stream.Send(&pb.EventResponse{Ok: true}); err != nil {
			return err
		}
	}
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
	pb.RegisterEventServiceServer(server, s)

	// Register reflection service on gRPC server.
	reflection.Register(server)
	log.Println("Registered gRPC services...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
