package instana

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
)

//NewTestHelper creates a new instance of TestHelper
func NewTestHelper(t *testing.T) TestHelper {
	return &testHelperImpl{t: t}
}

//TestHelper definition of the test helper utility
type TestHelper interface {
	CreateEmptyUserRoleResourceData() *schema.ResourceData
	CreateUserRoleResourceData(data map[string]interface{}) *schema.ResourceData
	CreateEmptyApplicationConfigResourceData() *schema.ResourceData
	CreateApplicationConfigResourceData(data map[string]interface{}) *schema.ResourceData
	CreateEmptyCustomEventSpecificationWithSystemRuleResourceData() *schema.ResourceData
	CreateCustomEventSpecificationWithSystemRuleResourceData(data map[string]interface{}) *schema.ResourceData
	CreateEmptyCustomEventSpecificationWithThresholdRuleResourceData() *schema.ResourceData
	CreateCustomEventSpecificationWithThresholdRuleResourceData(data map[string]interface{}) *schema.ResourceData
}

type testHelperImpl struct {
	t *testing.T
}

func (inst *testHelperImpl) CreateEmptyUserRoleResourceData() *schema.ResourceData {
	data := make(map[string]interface{})
	return inst.CreateUserRoleResourceData(data)
}

func (inst *testHelperImpl) CreateUserRoleResourceData(data map[string]interface{}) *schema.ResourceData {
	schemaMap := CreateResourceUserRole().Schema
	return schema.TestResourceDataRaw(inst.t, schemaMap, data)
}

func (inst *testHelperImpl) CreateEmptyApplicationConfigResourceData() *schema.ResourceData {
	data := make(map[string]interface{})
	return inst.CreateApplicationConfigResourceData(data)
}

func (inst *testHelperImpl) CreateApplicationConfigResourceData(data map[string]interface{}) *schema.ResourceData {
	schemaMap := CreateResourceApplicationConfig().Schema
	return schema.TestResourceDataRaw(inst.t, schemaMap, data)
}

func (inst *testHelperImpl) CreateEmptyCustomEventSpecificationWithSystemRuleResourceData() *schema.ResourceData {
	data := make(map[string]interface{})
	return inst.CreateCustomEventSpecificationWithSystemRuleResourceData(data)
}

func (inst *testHelperImpl) CreateCustomEventSpecificationWithSystemRuleResourceData(data map[string]interface{}) *schema.ResourceData {
	schemaMap := CreateResourceCustomEventSpecificationWithSystemRule().Schema
	return schema.TestResourceDataRaw(inst.t, schemaMap, data)
}

func (inst *testHelperImpl) CreateEmptyCustomEventSpecificationWithThresholdRuleResourceData() *schema.ResourceData {
	data := make(map[string]interface{})
	return inst.CreateCustomEventSpecificationWithThresholdRuleResourceData(data)
}

func (inst *testHelperImpl) CreateCustomEventSpecificationWithThresholdRuleResourceData(data map[string]interface{}) *schema.ResourceData {
	schemaMap := CreateResourceCustomEventSpecificationWithThresholdRule().Schema
	return schema.TestResourceDataRaw(inst.t, schemaMap, data)
}
