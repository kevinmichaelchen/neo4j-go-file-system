package grpc

import (
	"github.com/kevinmichaelchen/neo4j-go-file-system/organization"
	"github.com/kevinmichaelchen/neo4j-go-file-system/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"
)

func CreateOrganization(organizationService organization.Service, ctx context.Context, in *pb.OrganizationCrudRequest) (*pb.OrganizationResponse, error) {
	o, svcErr := organizationService.CreateOrganization(toOrganization(in.Organization))

	if svcErr != nil {
		return nil, status.Error(svcErr.GrpcCode, svcErr.ErrorMessage)
	}

	return &pb.OrganizationResponse{Organization: toGrpcOrganization(o)}, nil
}

func GetOrganization(organizationService organization.Service, ctx context.Context, in *pb.OrganizationCrudRequest) (*pb.OrganizationResponse, error) {
	o, svcErr := organizationService.GetOrganization(organization.Organization{
		ResourceID: in.Organization.Id,
	})
	if svcErr != nil {
		return nil, status.Error(svcErr.GrpcCode, svcErr.ErrorMessage)
	}
	return &pb.OrganizationResponse{Organization: toGrpcOrganization(o)}, nil
}

func UpdateOrganization(organizationService organization.Service, ctx context.Context, in *pb.OrganizationCrudRequest) (*pb.OrganizationResponse, error) {
	o, svcErr := organizationService.UpdateOrganization(toOrganization(in.Organization))
	if svcErr != nil {
		return nil, status.Error(svcErr.GrpcCode, svcErr.ErrorMessage)
	}
	return &pb.OrganizationResponse{Organization: toGrpcOrganization(o)}, nil
}

func DeleteOrganization(organizationService organization.Service, ctx context.Context, in *pb.OrganizationCrudRequest) (*pb.OrganizationResponse, error) {
	o, svcErr := organizationService.DeleteOrganization(organization.Organization{
		ResourceID: in.Organization.Id,
	})
	if svcErr != nil {
		return nil, status.Error(svcErr.GrpcCode, svcErr.ErrorMessage)
	}
	return &pb.OrganizationResponse{Organization: toGrpcOrganization(o)}, nil
}

func AddUserToOrganization(organizationService organization.Service, ctx context.Context, in *pb.AddUserToOrganizationRequest) (*pb.OrganizationResponse, error) {
	return nil, nil
}

func RemoveUserFromOrganization(organizationService organization.Service, ctx context.Context, in *pb.RemoveUserFromOrganizationRequest) (*pb.OrganizationResponse, error) {
	return nil, nil
}

func toOrganization(in *pb.Organization) organization.Organization {
	return organization.Organization{
		ResourceID: in.Id,
		Name:       in.Name,
	}
}

func toGrpcOrganization(in *organization.Organization) *pb.Organization {
	return &pb.Organization{
		Id:   in.ResourceID,
		Name: in.Name,
	}
}
