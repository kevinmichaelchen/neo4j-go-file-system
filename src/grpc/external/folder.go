package external

import (
	"github.com/kevinmichaelchen/neo4j-go-file-system/folder"
	"github.com/kevinmichaelchen/neo4j-go-file-system/pb"
	"github.com/kevinmichaelchen/neo4j-go-file-system/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"
)

func CreateFolder(folderService folder.Service, ctx context.Context, in *pb.CreateFolderRequest) (*pb.FolderResponse, error) {
	userID := getUserID(ctx)
	f, svcErr := folderService.CreateFolder(service.CreateUserContext(userID), folder.FolderInput{
		Name:     in.Name,
		ParentID: optString(in.ParentID),
		// TODO revision ID
	})
	if svcErr != nil {
		return nil, status.Error(svcErr.GrpcCode, svcErr.ErrorMessage)
	}
	return &pb.FolderResponse{Folder: toGrpcFolder(f)}, nil
}

func GetFolder(folderService folder.Service, ctx context.Context, in *pb.GetFolderRequest) (*pb.FolderResponse, error) {
	userID := getUserID(ctx)
	f, svcErr := folderService.GetFolder(service.CreateUserContext(userID), folder.FolderInput{ResourceID: in.FolderID})
	if svcErr != nil {
		return nil, status.Error(svcErr.GrpcCode, svcErr.ErrorMessage)
	}
	return &pb.FolderResponse{Folder: toGrpcFolder(f)}, nil
}

func UpdateFolder(folderService folder.Service, ctx context.Context, in *pb.UpdateFolderRequest) (*pb.FolderResponse, error) {
	userID := getUserID(ctx)
	f, svcErr := folderService.UpdateFolder(service.CreateUserContext(userID), folder.FolderInput{
		ResourceID: in.FolderID,
		Name:       in.Name,
		ParentID:   optString(in.ParentID),
		// TODO revision ID
	})
	if svcErr != nil {
		return nil, status.Error(svcErr.GrpcCode, svcErr.ErrorMessage)
	}
	return &pb.FolderResponse{Folder: toGrpcFolder(f)}, nil
}

func DeleteFolder(folderService folder.Service, ctx context.Context, in *pb.DeleteFolderRequest) (*pb.FolderResponse, error) {
	userID := getUserID(ctx)
	f, svcErr := folderService.DeleteFolder(service.CreateUserContext(userID), folder.FolderInput{ResourceID: in.FolderID})
	if svcErr != nil {
		return nil, status.Error(svcErr.GrpcCode, svcErr.ErrorMessage)
	}
	return &pb.FolderResponse{Folder: toGrpcFolder(f)}, nil
}

func toGrpcFolder(f *folder.Folder) *pb.Folder {
	return &pb.Folder{
		FolderID: f.ResourceID.String(),
		ParentID: f.ParentID.String(),
		Name:     f.Name,
		// TODO revisionID
	}
}
