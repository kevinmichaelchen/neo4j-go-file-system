package grpc

import (
	"github.com/google/uuid"
	"github.com/kevinmichaelchen/neo4j-go-file-system/file"
	"github.com/kevinmichaelchen/neo4j-go-file-system/pb"
	"github.com/kevinmichaelchen/neo4j-go-file-system/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateFile(fileService file.Service, ctx context.Context, in *pb.CreateFileRequest) (*pb.CreateFileResponse, error) {
	return nil, nil
}

func GetFile(fileService file.Service, ctx context.Context, in *pb.GetFileRequest) (*pb.GetFileResponse, error) {
	fileID, err := uuid.Parse(in.FileID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid file ID")
	}
	// TODO fetch userID from the context
	userID := 11
	f, svcErr := fileService.GetFile(service.CreateUserContext(userID), fileID)
	if svcErr != nil {
		return nil, svcErr.Error
	}
	return &pb.GetFileResponse{File: &pb.File{
		FileID:   f.ResourceID.String(),
		ParentID: f.ParentID.String(),
		Name:     f.Name,
		// TODO revisionID
	}}, nil
}

func UpdateFile(fileService file.Service, ctx context.Context, in *pb.UpdateFileRequest) (*pb.UpdateFileResponse, error) {
	return nil, nil
}

func DeleteFile(fileService file.Service, ctx context.Context, in *pb.DeleteFileRequest) (*pb.DeleteFileResponse, error) {
	return nil, nil
}