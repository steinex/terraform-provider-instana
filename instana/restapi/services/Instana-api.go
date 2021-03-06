package services

import "github.com/gessnerfl/terraform-provider-instana/instana/restapi"
import "github.com/gessnerfl/terraform-provider-instana/instana/restapi/resources"

//NewInstanaAPI creates a new instance of the instana API
func NewInstanaAPI(apiToken string, endpoint string) restapi.InstanaAPI {
	client := NewClient(apiToken, endpoint)
	return &baseInstanaAPI{client: client}
}

type baseInstanaAPI struct {
	client restapi.RestClient
}

//CustomEventSpecifications implementation of InstanaAPI interface
func (api baseInstanaAPI) CustomEventSpecifications() restapi.CustomEventSpecificationResource {
	return resources.NewCustomEventSpecificationResource(api.client)
}

//UserRoles implementation of InstanaAPI interface
func (api baseInstanaAPI) UserRoles() restapi.UserRoleResource {
	return resources.NewUserRoleResource(api.client)
}

//ApplicationConfigs implementation of InstanaAPI interface
func (api baseInstanaAPI) ApplicationConfigs() restapi.ApplicationConfigResource {
	return resources.NewApplicationConfigResource(api.client)
}
