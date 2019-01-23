package grpc

import (
	"github.com/kevinmichaelchen/neo4j-go-file-system/organization"
	"github.com/kevinmichaelchen/neo4j-go-file-system/pb"
	"golang.org/x/net/context"
)

func CreateOrganization(organizationService organization.Service, ctx context.Context, in *pb.CreateOrganizationRequest) (*pb.CreateOrganizationResponse, error) {
	u, svcError := organizationService.CreateOrganization(organization.Organization{
		Name: in.Name,
	})
	if svcError.Error != nil {
		return nil, svcError.Error
	}
	return &pb.CreateOrganizationResponse{Organization: toGrpcOrganization(u)}, nil
}

func GetOrganization(organizationService organization.Service, ctx context.Context, in *pb.GetOrganizationRequest) (*pb.GetOrganizationResponse, error) {
	u, svcError := organizationService.GetOrganization(organization.Organization{
		ResourceID: in.OrganizationID,
	})
	if svcError.Error != nil {
		return nil, svcError.Error
	}
	return &pb.GetOrganizationResponse{Organization: toGrpcOrganization(u)}, nil
}

func UpdateOrganization(organizationService organization.Service, ctx context.Context, in *pb.UpdateOrganizationRequest) (*pb.UpdateOrganizationResponse, error) {
	u, svcError := organizationService.UpdateOrganization(toOrganization(in.Organization))
	if svcError.Error != nil {
		return nil, svcError.Error
	}
	return &pb.UpdateOrganizationResponse{Organization: toGrpcOrganization(u)}, nil
}

func DeleteOrganization(organizationService organization.Service, ctx context.Context, in *pb.DeleteOrganizationRequest) (*pb.DeleteOrganizationResponse, error) {
	u, svcError := organizationService.DeleteOrganization(organization.Organization{
		ResourceID: in.OrganizationID,
	})
	if svcError.Error != nil {
		return nil, svcError.Error
	}
	return &pb.DeleteOrganizationResponse{Organization: toGrpcOrganization(u)}, nil
}

func toOrganization(u *pb.Organization) organization.Organization {
	return organization.Organization{
		ResourceID: u.OrganizationID,
		Name:       u.Name,
	}
}

func toGrpcOrganization(u *organization.Organization) *pb.Organization {
	return &pb.Organization{
		OrganizationID: u.ResourceID,
		Name:           u.Name,
	}
}
