package external

import (
	"github.com/kevinmichaelchen/neo4j-go-file-system/file"
	"github.com/kevinmichaelchen/neo4j-go-file-system/pb"
	"github.com/kevinmichaelchen/neo4j-go-file-system/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"
)

func CreateFile(fileService file.Service, ctx context.Context, in *pb.CreateFileRequest) (*pb.FileResponse, error) {
	userID := getUserID(ctx)
	f, svcErr := fileService.CreateFile(service.CreateUserContext(userID), file.FileInput{
		ParentID: in.ParentID,
		Name:     in.Name,
	})
	if svcErr != nil {
		return nil, status.Error(svcErr.GrpcCode, svcErr.ErrorMessage)
	}
	return &pb.FileResponse{File: toGrpcFile(f)}, nil
}

func GetFile(fileService file.Service, ctx context.Context, in *pb.GetFileRequest) (*pb.FileResponse, error) {
	userID := getUserID(ctx)
	f, svcErr := fileService.GetFile(service.CreateUserContext(userID), file.FileInput{ResourceID: in.FileID})
	if svcErr != nil {
		return nil, status.Error(svcErr.GrpcCode, svcErr.ErrorMessage)
	}
	return &pb.FileResponse{File: toGrpcFile(f)}, nil
}

func UpdateFile(fileService file.Service, ctx context.Context, in *pb.UpdateFileRequest) (*pb.FileResponse, error) {
	userID := getUserID(ctx)
	f, svcErr := fileService.UpdateFile(service.CreateUserContext(userID), file.FileInput{
		ResourceID: in.FileID,
		ParentID:   in.ParentID,
		Name:       in.Name,
	})
	if svcErr != nil {
		return nil, status.Error(svcErr.GrpcCode, svcErr.ErrorMessage)
	}
	return &pb.FileResponse{File: toGrpcFile(f)}, nil
}

func DeleteFile(fileService file.Service, ctx context.Context, in *pb.DeleteFileRequest) (*pb.FileResponse, error) {
	userID := getUserID(ctx)
	f, svcErr := fileService.DeleteFile(service.CreateUserContext(userID), file.FileInput{
		ResourceID: in.FileID,
	})
	if svcErr != nil {
		return nil, status.Error(svcErr.GrpcCode, svcErr.ErrorMessage)
	}
	return &pb.FileResponse{File: toGrpcFile(f)}, nil
}

func toGrpcFile(f *file.File) *pb.File {
	return &pb.File{
		Id:       f.ResourceID.String(),
		ParentID: f.ParentID.String(),
		Name:     f.Name,
		// TODO revisionID
	}
}
