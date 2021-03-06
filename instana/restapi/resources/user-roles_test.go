package resources_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/gessnerfl/terraform-provider-instana/instana/restapi"
	. "github.com/gessnerfl/terraform-provider-instana/instana/restapi/resources"
	mocks "github.com/gessnerfl/terraform-provider-instana/mocks"
	"github.com/gessnerfl/terraform-provider-instana/testutils"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

const (
	userRoleID   = "test-user-role-id"
	userRoleName = "Test User Role"
)

func TestSuccessfulGetOneUserRole(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewUserRoleResource(client)
	userRole := makeTestUserRole()
	serializedJSON, _ := json.Marshal(userRole)

	client.EXPECT().GetOne(gomock.Eq(userRole.ID), gomock.Eq(restapi.UserRolesResourcePath)).Return(serializedJSON, nil)

	data, err := sut.GetOne(userRole.ID)

	if err != nil {
		t.Fatalf(testutils.ExpectedNoErrorButGotMessage, err)
	}

	if !cmp.Equal(userRole, data) {
		t.Fatalf(testutils.ExpectedUnmarshalledJSONWithStruct, userRole, data, cmp.Diff(userRole, data))
	}
}

func TestFailedGetOneUserRoleBecauseOfErrorFromRestClient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewUserRoleResource(client)

	client.EXPECT().GetOne(gomock.Eq(userRoleID), gomock.Eq(restapi.UserRolesResourcePath)).Return(nil, errors.New("error during test"))

	_, err := sut.GetOne(userRoleID)

	if err == nil {
		t.Fatalf(testutils.ExpectedError)
	}
}

func TestFailedGetOneUserRoleBecauseOfInvalidJsonArray(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewUserRoleResource(client)

	client.EXPECT().GetOne(gomock.Eq(userRoleID), gomock.Eq(restapi.UserRolesResourcePath)).Return([]byte("[{ \"invalid\" : \"data\" }]"), nil)

	_, err := sut.GetOne(userRoleID)

	if err == nil {
		t.Fatalf(testutils.ExpectedError)
	}
}

func TestFailedGetOneUserRoleBecauseOfInvalidJsonObject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewUserRoleResource(client)

	client.EXPECT().GetOne(gomock.Eq(userRoleID), gomock.Eq(restapi.UserRolesResourcePath)).Return([]byte("{ \"invalid\" : \"data\" }"), nil)

	_, err := sut.GetOne(userRoleID)

	if err == nil {
		t.Fatalf(testutils.ExpectedError)
	}
}

func TestFailedGetOneUserRoleBecauseResponseIsNotAValidJsonDocument(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewUserRoleResource(client)

	client.EXPECT().GetOne(gomock.Eq(userRoleID), gomock.Eq(restapi.UserRolesResourcePath)).Return([]byte("Invalid Data"), nil)

	_, err := sut.GetOne(userRoleID)

	if err == nil {
		t.Fatalf(testutils.ExpectedError)
	}
}

func TestSuccessfulUpsertOfUserRole(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewUserRoleResource(client)
	userRole := makeTestUserRole()
	serializedJSON, _ := json.Marshal(userRole)

	client.EXPECT().Put(gomock.Eq(userRole), gomock.Eq(restapi.UserRolesResourcePath)).Return(serializedJSON, nil)

	result, err := sut.Upsert(userRole)

	if err != nil {
		t.Fatalf(testutils.ExpectedNoErrorButGotMessage, err)
	}

	if !cmp.Equal(userRole, result) {
		t.Fatalf(testutils.ExpectedUnmarshalledJSONWithStruct, userRole, result, cmp.Diff(result, result))
	}
}

func TestFailedUpsertOfUserRoleBecauseOfClientError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewUserRoleResource(client)
	userRole := makeTestUserRole()

	client.EXPECT().Put(gomock.Eq(userRole), gomock.Eq(restapi.UserRolesResourcePath)).Return(nil, errors.New("Error during test"))

	_, err := sut.Upsert(userRole)

	if err == nil {
		t.Fatal(testutils.ExpectedError)
	}
}

func TestFailedUpsertOfUserRoleBecauseOfInvalidResponseMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewUserRoleResource(client)
	userRole := makeTestUserRole()

	client.EXPECT().Put(gomock.Eq(userRole), gomock.Eq(restapi.UserRolesResourcePath)).Return([]byte("invalid response"), nil)

	_, err := sut.Upsert(userRole)

	if err == nil {
		t.Fatalf(testutils.ExpectedError)
	}
}

func TestFailedUpsertOfUserRoleBecauseOfInvalidUserRoleInResponse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewUserRoleResource(client)
	userRole := makeTestUserRole()

	client.EXPECT().Put(gomock.Eq(userRole), gomock.Eq(restapi.UserRolesResourcePath)).Return([]byte("{ \"invalid\" : \"userRole\" }"), nil)

	_, err := sut.Upsert(userRole)

	if err == nil {
		t.Fatalf(testutils.ExpectedError)
	}
}

func TestFailedUpsertOfUserRoleBecauseOfInvalidUserRoleProvided(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewUserRoleResource(client)
	userRole := restapi.UserRole{
		Name: userRoleName,
	}

	client.EXPECT().Put(gomock.Eq(userRole), gomock.Eq(restapi.UserRolesResourcePath)).Times(0)

	_, err := sut.Upsert(userRole)

	if err == nil {
		t.Fatal(testutils.ExpectedError)
	}
}

func TestSuccessfulDeleteOfUserRoleByUserRole(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewUserRoleResource(client)
	userRole := makeTestUserRole()

	client.EXPECT().Delete(gomock.Eq(userRoleID), gomock.Eq(restapi.UserRolesResourcePath)).Return(nil)

	err := sut.Delete(userRole)

	if err != nil {
		t.Fatalf("Expected no error got %s", err)
	}
}

func TestFailedDeleteOfUserRoleByUserRole(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewUserRoleResource(client)
	userRole := makeTestUserRole()

	client.EXPECT().Delete(gomock.Eq(userRoleID), gomock.Eq(restapi.UserRolesResourcePath)).Return(errors.New("Error during test"))

	err := sut.Delete(userRole)

	if err == nil {
		t.Fatal(testutils.ExpectedError)
	}
}

func makeTestUserRole() restapi.UserRole {
	return restapi.UserRole{
		ID:                                userRoleID,
		Name:                              userRoleName,
		ImplicitViewFilter:                "Test view filter",
		CanConfigureServiceMapping:        true,
		CanConfigureEumApplications:       true,
		CanConfigureUsers:                 true,
		CanInstallNewAgents:               true,
		CanSeeUsageInformation:            true,
		CanConfigureIntegrations:          true,
		CanSeeOnPremiseLicenseInformation: true,
		CanConfigureRoles:                 true,
		CanConfigureCustomAlerts:          true,
		CanConfigureAPITokens:             true,
		CanConfigureAgentRunMode:          true,
		CanViewAuditLog:                   true,
		CanConfigureObjectives:            true,
		CanConfigureAgents:                true,
		CanConfigureAuthenticationMethods: true,
		CanConfigureApplications:          true,
	}
}
