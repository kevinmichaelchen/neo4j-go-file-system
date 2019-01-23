package grpc

import (
	"github.com/kevinmichaelchen/neo4j-go-file-system/organization"
	"github.com/kevinmichaelchen/neo4j-go-file-system/pb"
	"golang.org/x/net/context"
)

func CreateOrganization(organizationService organization.Service, ctx context.Context, in *pb.CreateOrgRequest) (*pb.CreateOrgResponse, error) {
	o, svcError := organizationService.CreateOrganization(toOrganization(in.Organization))
	if svcError.Error != nil {
		return nil, svcError.Error
	}
	// TODO this line is bombing w/ invalid memory address or nil pointer dereference
	return &pb.CreateOrgResponse{Organization: &pb.Organization{
		Id: o.ResourceID,
		Name: o.Name}}, nil
}

func GetOrganization(organizationService organization.Service, ctx context.Context, in *pb.OrganizationCrudRequest) (*pb.OrganizationResponse, error) {
	u, svcError := organizationService.GetOrganization(organization.Organization{
		ResourceID: in.Organization.Id,
	})
	if svcError.Error != nil {
		return nil, svcError.Error
	}
	return &pb.OrganizationResponse{Organization: toGrpcOrganization(u)}, nil
}

func UpdateOrganization(organizationService organization.Service, ctx context.Context, in *pb.OrganizationCrudRequest) (*pb.OrganizationResponse, error) {
	u, svcError := organizationService.UpdateOrganization(toOrganization(in.Organization))
	if svcError.Error != nil {
		return nil, svcError.Error
	}
	return &pb.OrganizationResponse{Organization: toGrpcOrganization(u)}, nil
}

func DeleteOrganization(organizationService organization.Service, ctx context.Context, in *pb.OrganizationCrudRequest) (*pb.OrganizationResponse, error) {
	u, svcError := organizationService.DeleteOrganization(organization.Organization{
		ResourceID: in.Organization.Id,
	})
	if svcError.Error != nil {
		return nil, svcError.Error
	}
	return &pb.OrganizationResponse{Organization: toGrpcOrganization(u)}, nil
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
