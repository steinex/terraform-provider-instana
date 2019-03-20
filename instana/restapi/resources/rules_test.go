package resources_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/gessnerfl/terraform-provider-instana/instana/restapi"
	. "github.com/gessnerfl/terraform-provider-instana/instana/restapi/resources"
	mocks "github.com/gessnerfl/terraform-provider-instana/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

const rulePath = "/events/settings/rules"

func TestSuccessfulGetOneRule(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewRuleResource(client)
	rule := makeTestRule()
	serializedJSON, _ := json.Marshal(rule)

	client.EXPECT().GetOne(gomock.Eq(rule.ID), gomock.Eq(rulePath)).Return(serializedJSON, nil)

	data, err := sut.GetOne(rule.ID)

	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	if !cmp.Equal(rule, data) {
		t.Fatalf("Expected json to be unmarshalled to %v but got %v; diff %s", rule, data, cmp.Diff(rule, data))
	}
}

func TestFailedGetOneRuleBecauseOfErrorFromRestClient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewRuleResource(client)
	ruleID := "test-rule-id"

	client.EXPECT().GetOne(gomock.Eq(ruleID), gomock.Eq(rulePath)).Return(nil, errors.New("error during test"))

	_, err := sut.GetOne(ruleID)

	if err == nil {
		t.Fatalf("Expected to get error")
	}
}

func TestFailedGetOneRuleBecauseOfInvalidJsonArray(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewRuleResource(client)
	ruleID := "test-rule-id"

	client.EXPECT().GetOne(gomock.Eq(ruleID), gomock.Eq(rulePath)).Return([]byte("[{ \"invalid\" : \"data\" }]"), nil)

	_, err := sut.GetOne(ruleID)

	if err == nil {
		t.Fatalf("Expected to get error")
	}
}

func TestFailedGetOneRuleBecauseOfInvalidJsonObject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewRuleResource(client)
	ruleID := "test-rule-id"

	client.EXPECT().GetOne(gomock.Eq(ruleID), gomock.Eq(rulePath)).Return([]byte("{ \"invalid\" : \"data\" }"), nil)

	_, err := sut.GetOne(ruleID)

	if err == nil {
		t.Fatalf("Expected to get error")
	}
}

func TestFailedGetOneRuleBecauseOfNoJsonAsResponse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewRuleResource(client)
	ruleID := "test-rule-id"

	client.EXPECT().GetOne(gomock.Eq(ruleID), gomock.Eq(rulePath)).Return([]byte("Invalid Data"), nil)

	_, err := sut.GetOne(ruleID)

	if err == nil {
		t.Fatalf("Expected to get error")
	}
}

func TestSuccessfulGetAllRules(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewRuleResource(client)
	rule1 := makeTestRuleWithCounter(1)
	rule2 := makeTestRuleWithCounter(2)
	rules := []restapi.Rule{rule1, rule2}
	serializedJSON, _ := json.Marshal(rules)

	client.EXPECT().GetAll(gomock.Eq(rulePath)).Return(serializedJSON, nil)

	data, err := sut.GetAll()

	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	if !cmp.Equal(rules, data) {
		t.Fatalf("Expected json to be unmarshalled to %v but got %v; diff %s", rules, data, cmp.Diff(rules, data))
	}
}

func TestFailedGetAllRulesWithErrorFromRestClient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewRuleResource(client)

	client.EXPECT().GetAll(gomock.Eq(rulePath)).Return(nil, errors.New("error during test"))

	_, err := sut.GetAll()

	if err == nil {
		t.Fatalf("Expected to get error")
	}
}

func TestFailedGetAllRulesWithInvalidJsonArray(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewRuleResource(client)

	client.EXPECT().GetAll(gomock.Eq(rulePath)).Return([]byte("[{ \"invalid\" : \"data\" }]"), nil)

	_, err := sut.GetAll()

	if err == nil {
		t.Fatalf("Expected to get error")
	}
}

func TestFailedGetAllRulesWithInvalidJsonObject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewRuleResource(client)

	client.EXPECT().GetAll(gomock.Eq(rulePath)).Return([]byte("{ \"invalid\" : \"data\" }"), nil)

	_, err := sut.GetAll()

	if err == nil {
		t.Fatalf("Expected to get error")
	}
}

func TestFailedGetAllRulesWithNoJsonAsResponse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewRuleResource(client)

	client.EXPECT().GetAll(gomock.Eq(rulePath)).Return([]byte("Invalid Data"), nil)

	_, err := sut.GetAll()

	if err == nil {
		t.Fatalf("Expected to get error")
	}
}

func TestSuccessfulUpsertOfRule(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewRuleResource(client)
	rule := makeTestRule()
	serializedJSON, _ := json.Marshal(rule)

	client.EXPECT().Put(gomock.Eq(rule), gomock.Eq(rulePath)).Return(serializedJSON, nil)

	result, err := sut.Upsert(rule)

	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	if !cmp.Equal(rule, result) {
		t.Fatalf("Expected json to be unmarshalled to %v but got %v; diff %s", rule, result, cmp.Diff(result, result))
	}
}

func TestFailedUpsertOfRuleBecauseOfClientError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewRuleResource(client)
	rule := makeTestRule()

	client.EXPECT().Put(gomock.Eq(rule), gomock.Eq(rulePath)).Return(nil, errors.New("Error during test"))

	_, err := sut.Upsert(rule)

	if err == nil {
		t.Fatal("Expected to get error")
	}
}

func TestFailedUpsertOfRuleBecauseOfInvalidResponseMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewRuleResource(client)
	rule := makeTestRule()

	client.EXPECT().Put(gomock.Eq(rule), gomock.Eq(rulePath)).Return([]byte("invalid response"), nil)

	_, err := sut.Upsert(rule)

	if err == nil {
		t.Fatalf("Expected to get error")
	}
}

func TestFailedUpsertOfRuleBecauseOfInvalidRuleInResponse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewRuleResource(client)
	rule := makeTestRule()

	client.EXPECT().Put(gomock.Eq(rule), gomock.Eq(rulePath)).Return([]byte("{ \"invalid\" : \"rule\" }"), nil)

	_, err := sut.Upsert(rule)

	if err == nil {
		t.Fatalf("Expected to get error")
	}
}

func TestFailedUpsertOfRuleBecauseOfInvalidRuleProvided(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewRuleResource(client)
	rule := restapi.Rule{
		Name:              "Test Rule",
		EntityType:        "test",
		MetricName:        "test.metric",
		Rollup:            0,
		Window:            300000,
		Aggregation:       "sum",
		ConditionValue:    0,
		ConditionOperator: ">",
	}

	client.EXPECT().Put(gomock.Eq(rule), gomock.Eq(rulePath)).Times(0)

	_, err := sut.Upsert(rule)

	if err == nil {
		t.Fatal("Expected to get error")
	}
}

func TestSuccessfulDeleteOfRuleByRule(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewRuleResource(client)
	rule := makeTestRule()

	client.EXPECT().Delete(gomock.Eq("test-rule-id-1"), gomock.Eq(rulePath)).Return(nil)

	err := sut.Delete(rule)

	if err != nil {
		t.Fatalf("Expected no error got %s", err)
	}
}

func TestFailedDeleteOfRuleByRule(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mocks.NewMockRestClient(ctrl)

	sut := NewRuleResource(client)
	rule := makeTestRule()

	client.EXPECT().Delete(gomock.Eq("test-rule-id-1"), gomock.Eq(rulePath)).Return(errors.New("Error during test"))

	err := sut.Delete(rule)

	if err == nil {
		t.Fatal("Expected to get error")
	}
}

func makeTestRule() restapi.Rule {
	return makeTestRuleWithCounter(1)
}

func makeTestRuleWithCounter(counter int) restapi.Rule {
	id := fmt.Sprintf("test-rule-id-%d", counter)
	name := fmt.Sprintf("Test Rule %d", counter)
	return restapi.Rule{
		ID:                id,
		Name:              name,
		EntityType:        "test",
		MetricName:        "test.metric",
		Rollup:            0,
		Window:            300000,
		Aggregation:       "sum",
		ConditionOperator: ">",
		ConditionValue:    0,
	}
}
