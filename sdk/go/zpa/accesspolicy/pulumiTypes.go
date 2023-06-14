// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package accesspolicy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PolicyAccessRuleAppConnectorGroup struct {
	// (Optional) The ID of a server group resource
	Ids []string `pulumi:"ids"`
}

// PolicyAccessRuleAppConnectorGroupInput is an input type that accepts PolicyAccessRuleAppConnectorGroupArgs and PolicyAccessRuleAppConnectorGroupOutput values.
// You can construct a concrete instance of `PolicyAccessRuleAppConnectorGroupInput` via:
//
//	PolicyAccessRuleAppConnectorGroupArgs{...}
type PolicyAccessRuleAppConnectorGroupInput interface {
	pulumi.Input

	ToPolicyAccessRuleAppConnectorGroupOutput() PolicyAccessRuleAppConnectorGroupOutput
	ToPolicyAccessRuleAppConnectorGroupOutputWithContext(context.Context) PolicyAccessRuleAppConnectorGroupOutput
}

type PolicyAccessRuleAppConnectorGroupArgs struct {
	// (Optional) The ID of a server group resource
	Ids pulumi.StringArrayInput `pulumi:"ids"`
}

func (PolicyAccessRuleAppConnectorGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAccessRuleAppConnectorGroup)(nil)).Elem()
}

func (i PolicyAccessRuleAppConnectorGroupArgs) ToPolicyAccessRuleAppConnectorGroupOutput() PolicyAccessRuleAppConnectorGroupOutput {
	return i.ToPolicyAccessRuleAppConnectorGroupOutputWithContext(context.Background())
}

func (i PolicyAccessRuleAppConnectorGroupArgs) ToPolicyAccessRuleAppConnectorGroupOutputWithContext(ctx context.Context) PolicyAccessRuleAppConnectorGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAccessRuleAppConnectorGroupOutput)
}

// PolicyAccessRuleAppConnectorGroupArrayInput is an input type that accepts PolicyAccessRuleAppConnectorGroupArray and PolicyAccessRuleAppConnectorGroupArrayOutput values.
// You can construct a concrete instance of `PolicyAccessRuleAppConnectorGroupArrayInput` via:
//
//	PolicyAccessRuleAppConnectorGroupArray{ PolicyAccessRuleAppConnectorGroupArgs{...} }
type PolicyAccessRuleAppConnectorGroupArrayInput interface {
	pulumi.Input

	ToPolicyAccessRuleAppConnectorGroupArrayOutput() PolicyAccessRuleAppConnectorGroupArrayOutput
	ToPolicyAccessRuleAppConnectorGroupArrayOutputWithContext(context.Context) PolicyAccessRuleAppConnectorGroupArrayOutput
}

type PolicyAccessRuleAppConnectorGroupArray []PolicyAccessRuleAppConnectorGroupInput

func (PolicyAccessRuleAppConnectorGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyAccessRuleAppConnectorGroup)(nil)).Elem()
}

func (i PolicyAccessRuleAppConnectorGroupArray) ToPolicyAccessRuleAppConnectorGroupArrayOutput() PolicyAccessRuleAppConnectorGroupArrayOutput {
	return i.ToPolicyAccessRuleAppConnectorGroupArrayOutputWithContext(context.Background())
}

func (i PolicyAccessRuleAppConnectorGroupArray) ToPolicyAccessRuleAppConnectorGroupArrayOutputWithContext(ctx context.Context) PolicyAccessRuleAppConnectorGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAccessRuleAppConnectorGroupArrayOutput)
}

type PolicyAccessRuleAppConnectorGroupOutput struct{ *pulumi.OutputState }

func (PolicyAccessRuleAppConnectorGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAccessRuleAppConnectorGroup)(nil)).Elem()
}

func (o PolicyAccessRuleAppConnectorGroupOutput) ToPolicyAccessRuleAppConnectorGroupOutput() PolicyAccessRuleAppConnectorGroupOutput {
	return o
}

func (o PolicyAccessRuleAppConnectorGroupOutput) ToPolicyAccessRuleAppConnectorGroupOutputWithContext(ctx context.Context) PolicyAccessRuleAppConnectorGroupOutput {
	return o
}

// (Optional) The ID of a server group resource
func (o PolicyAccessRuleAppConnectorGroupOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PolicyAccessRuleAppConnectorGroup) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

type PolicyAccessRuleAppConnectorGroupArrayOutput struct{ *pulumi.OutputState }

func (PolicyAccessRuleAppConnectorGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyAccessRuleAppConnectorGroup)(nil)).Elem()
}

func (o PolicyAccessRuleAppConnectorGroupArrayOutput) ToPolicyAccessRuleAppConnectorGroupArrayOutput() PolicyAccessRuleAppConnectorGroupArrayOutput {
	return o
}

func (o PolicyAccessRuleAppConnectorGroupArrayOutput) ToPolicyAccessRuleAppConnectorGroupArrayOutputWithContext(ctx context.Context) PolicyAccessRuleAppConnectorGroupArrayOutput {
	return o
}

func (o PolicyAccessRuleAppConnectorGroupArrayOutput) Index(i pulumi.IntInput) PolicyAccessRuleAppConnectorGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyAccessRuleAppConnectorGroup {
		return vs[0].([]PolicyAccessRuleAppConnectorGroup)[vs[1].(int)]
	}).(PolicyAccessRuleAppConnectorGroupOutput)
}

type PolicyAccessRuleAppServerGroup struct {
	// (Optional) The ID of a server group resource
	Ids []string `pulumi:"ids"`
}

// PolicyAccessRuleAppServerGroupInput is an input type that accepts PolicyAccessRuleAppServerGroupArgs and PolicyAccessRuleAppServerGroupOutput values.
// You can construct a concrete instance of `PolicyAccessRuleAppServerGroupInput` via:
//
//	PolicyAccessRuleAppServerGroupArgs{...}
type PolicyAccessRuleAppServerGroupInput interface {
	pulumi.Input

	ToPolicyAccessRuleAppServerGroupOutput() PolicyAccessRuleAppServerGroupOutput
	ToPolicyAccessRuleAppServerGroupOutputWithContext(context.Context) PolicyAccessRuleAppServerGroupOutput
}

type PolicyAccessRuleAppServerGroupArgs struct {
	// (Optional) The ID of a server group resource
	Ids pulumi.StringArrayInput `pulumi:"ids"`
}

func (PolicyAccessRuleAppServerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAccessRuleAppServerGroup)(nil)).Elem()
}

func (i PolicyAccessRuleAppServerGroupArgs) ToPolicyAccessRuleAppServerGroupOutput() PolicyAccessRuleAppServerGroupOutput {
	return i.ToPolicyAccessRuleAppServerGroupOutputWithContext(context.Background())
}

func (i PolicyAccessRuleAppServerGroupArgs) ToPolicyAccessRuleAppServerGroupOutputWithContext(ctx context.Context) PolicyAccessRuleAppServerGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAccessRuleAppServerGroupOutput)
}

// PolicyAccessRuleAppServerGroupArrayInput is an input type that accepts PolicyAccessRuleAppServerGroupArray and PolicyAccessRuleAppServerGroupArrayOutput values.
// You can construct a concrete instance of `PolicyAccessRuleAppServerGroupArrayInput` via:
//
//	PolicyAccessRuleAppServerGroupArray{ PolicyAccessRuleAppServerGroupArgs{...} }
type PolicyAccessRuleAppServerGroupArrayInput interface {
	pulumi.Input

	ToPolicyAccessRuleAppServerGroupArrayOutput() PolicyAccessRuleAppServerGroupArrayOutput
	ToPolicyAccessRuleAppServerGroupArrayOutputWithContext(context.Context) PolicyAccessRuleAppServerGroupArrayOutput
}

type PolicyAccessRuleAppServerGroupArray []PolicyAccessRuleAppServerGroupInput

func (PolicyAccessRuleAppServerGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyAccessRuleAppServerGroup)(nil)).Elem()
}

func (i PolicyAccessRuleAppServerGroupArray) ToPolicyAccessRuleAppServerGroupArrayOutput() PolicyAccessRuleAppServerGroupArrayOutput {
	return i.ToPolicyAccessRuleAppServerGroupArrayOutputWithContext(context.Background())
}

func (i PolicyAccessRuleAppServerGroupArray) ToPolicyAccessRuleAppServerGroupArrayOutputWithContext(ctx context.Context) PolicyAccessRuleAppServerGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAccessRuleAppServerGroupArrayOutput)
}

type PolicyAccessRuleAppServerGroupOutput struct{ *pulumi.OutputState }

func (PolicyAccessRuleAppServerGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAccessRuleAppServerGroup)(nil)).Elem()
}

func (o PolicyAccessRuleAppServerGroupOutput) ToPolicyAccessRuleAppServerGroupOutput() PolicyAccessRuleAppServerGroupOutput {
	return o
}

func (o PolicyAccessRuleAppServerGroupOutput) ToPolicyAccessRuleAppServerGroupOutputWithContext(ctx context.Context) PolicyAccessRuleAppServerGroupOutput {
	return o
}

// (Optional) The ID of a server group resource
func (o PolicyAccessRuleAppServerGroupOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PolicyAccessRuleAppServerGroup) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

type PolicyAccessRuleAppServerGroupArrayOutput struct{ *pulumi.OutputState }

func (PolicyAccessRuleAppServerGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyAccessRuleAppServerGroup)(nil)).Elem()
}

func (o PolicyAccessRuleAppServerGroupArrayOutput) ToPolicyAccessRuleAppServerGroupArrayOutput() PolicyAccessRuleAppServerGroupArrayOutput {
	return o
}

func (o PolicyAccessRuleAppServerGroupArrayOutput) ToPolicyAccessRuleAppServerGroupArrayOutputWithContext(ctx context.Context) PolicyAccessRuleAppServerGroupArrayOutput {
	return o
}

func (o PolicyAccessRuleAppServerGroupArrayOutput) Index(i pulumi.IntInput) PolicyAccessRuleAppServerGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyAccessRuleAppServerGroup {
		return vs[0].([]PolicyAccessRuleAppServerGroup)[vs[1].(int)]
	}).(PolicyAccessRuleAppServerGroupOutput)
}

type PolicyAccessRuleCondition struct {
	// (Optional) The ID of a server group resource
	Id *string `pulumi:"id"`
	// (Optional) Supported values: ``true`` or ``false``
	Negated *bool `pulumi:"negated"`
	// (Optional) - Operands block must be repeated if multiple per `objectType` conditions are to be added to the rule.
	Operands []PolicyAccessRuleConditionOperand `pulumi:"operands"`
	// (Optional) Supported values: ``AND``, and ``OR``
	Operator string `pulumi:"operator"`
}

// PolicyAccessRuleConditionInput is an input type that accepts PolicyAccessRuleConditionArgs and PolicyAccessRuleConditionOutput values.
// You can construct a concrete instance of `PolicyAccessRuleConditionInput` via:
//
//	PolicyAccessRuleConditionArgs{...}
type PolicyAccessRuleConditionInput interface {
	pulumi.Input

	ToPolicyAccessRuleConditionOutput() PolicyAccessRuleConditionOutput
	ToPolicyAccessRuleConditionOutputWithContext(context.Context) PolicyAccessRuleConditionOutput
}

type PolicyAccessRuleConditionArgs struct {
	// (Optional) The ID of a server group resource
	Id pulumi.StringPtrInput `pulumi:"id"`
	// (Optional) Supported values: ``true`` or ``false``
	Negated pulumi.BoolPtrInput `pulumi:"negated"`
	// (Optional) - Operands block must be repeated if multiple per `objectType` conditions are to be added to the rule.
	Operands PolicyAccessRuleConditionOperandArrayInput `pulumi:"operands"`
	// (Optional) Supported values: ``AND``, and ``OR``
	Operator pulumi.StringInput `pulumi:"operator"`
}

func (PolicyAccessRuleConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAccessRuleCondition)(nil)).Elem()
}

func (i PolicyAccessRuleConditionArgs) ToPolicyAccessRuleConditionOutput() PolicyAccessRuleConditionOutput {
	return i.ToPolicyAccessRuleConditionOutputWithContext(context.Background())
}

func (i PolicyAccessRuleConditionArgs) ToPolicyAccessRuleConditionOutputWithContext(ctx context.Context) PolicyAccessRuleConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAccessRuleConditionOutput)
}

// PolicyAccessRuleConditionArrayInput is an input type that accepts PolicyAccessRuleConditionArray and PolicyAccessRuleConditionArrayOutput values.
// You can construct a concrete instance of `PolicyAccessRuleConditionArrayInput` via:
//
//	PolicyAccessRuleConditionArray{ PolicyAccessRuleConditionArgs{...} }
type PolicyAccessRuleConditionArrayInput interface {
	pulumi.Input

	ToPolicyAccessRuleConditionArrayOutput() PolicyAccessRuleConditionArrayOutput
	ToPolicyAccessRuleConditionArrayOutputWithContext(context.Context) PolicyAccessRuleConditionArrayOutput
}

type PolicyAccessRuleConditionArray []PolicyAccessRuleConditionInput

func (PolicyAccessRuleConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyAccessRuleCondition)(nil)).Elem()
}

func (i PolicyAccessRuleConditionArray) ToPolicyAccessRuleConditionArrayOutput() PolicyAccessRuleConditionArrayOutput {
	return i.ToPolicyAccessRuleConditionArrayOutputWithContext(context.Background())
}

func (i PolicyAccessRuleConditionArray) ToPolicyAccessRuleConditionArrayOutputWithContext(ctx context.Context) PolicyAccessRuleConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAccessRuleConditionArrayOutput)
}

type PolicyAccessRuleConditionOutput struct{ *pulumi.OutputState }

func (PolicyAccessRuleConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAccessRuleCondition)(nil)).Elem()
}

func (o PolicyAccessRuleConditionOutput) ToPolicyAccessRuleConditionOutput() PolicyAccessRuleConditionOutput {
	return o
}

func (o PolicyAccessRuleConditionOutput) ToPolicyAccessRuleConditionOutputWithContext(ctx context.Context) PolicyAccessRuleConditionOutput {
	return o
}

// (Optional) The ID of a server group resource
func (o PolicyAccessRuleConditionOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyAccessRuleCondition) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// (Optional) Supported values: “true“ or “false“
func (o PolicyAccessRuleConditionOutput) Negated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PolicyAccessRuleCondition) *bool { return v.Negated }).(pulumi.BoolPtrOutput)
}

// (Optional) - Operands block must be repeated if multiple per `objectType` conditions are to be added to the rule.
func (o PolicyAccessRuleConditionOutput) Operands() PolicyAccessRuleConditionOperandArrayOutput {
	return o.ApplyT(func(v PolicyAccessRuleCondition) []PolicyAccessRuleConditionOperand { return v.Operands }).(PolicyAccessRuleConditionOperandArrayOutput)
}

// (Optional) Supported values: “AND“, and “OR“
func (o PolicyAccessRuleConditionOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyAccessRuleCondition) string { return v.Operator }).(pulumi.StringOutput)
}

type PolicyAccessRuleConditionArrayOutput struct{ *pulumi.OutputState }

func (PolicyAccessRuleConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyAccessRuleCondition)(nil)).Elem()
}

func (o PolicyAccessRuleConditionArrayOutput) ToPolicyAccessRuleConditionArrayOutput() PolicyAccessRuleConditionArrayOutput {
	return o
}

func (o PolicyAccessRuleConditionArrayOutput) ToPolicyAccessRuleConditionArrayOutputWithContext(ctx context.Context) PolicyAccessRuleConditionArrayOutput {
	return o
}

func (o PolicyAccessRuleConditionArrayOutput) Index(i pulumi.IntInput) PolicyAccessRuleConditionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyAccessRuleCondition {
		return vs[0].([]PolicyAccessRuleCondition)[vs[1].(int)]
	}).(PolicyAccessRuleConditionOutput)
}

type PolicyAccessRuleConditionOperand struct {
	// (Optional) The ID of a server group resource
	Id *string `pulumi:"id"`
	// (Optional)
	IdpId *string `pulumi:"idpId"`
	// (Optional) LHS must always carry the string value ``id`` or the attribute ID of the resource being associated with the rule.
	Lhs string `pulumi:"lhs"`
	// (Optional)
	Name *string `pulumi:"name"`
	// (Optional) This is for specifying the policy critiera. Supported values: `APP`, `APP_GROUP`, `SAML`, `IDP`, `CLIENT_TYPE`, `TRUSTED_NETWORK`, `POSTURE`, `SCIM`, `SCIM_GROUP`, and `CLOUD_CONNECTOR_GROUP`. `TRUSTED_NETWORK`, and `CLIENT_TYPE`.
	ObjectType string `pulumi:"objectType"`
	// (Optional) RHS is either the ID attribute of a resource or fixed string value. Refer to the chart below for further details.
	Rhs      *string  `pulumi:"rhs"`
	RhsLists []string `pulumi:"rhsLists"`
}

// PolicyAccessRuleConditionOperandInput is an input type that accepts PolicyAccessRuleConditionOperandArgs and PolicyAccessRuleConditionOperandOutput values.
// You can construct a concrete instance of `PolicyAccessRuleConditionOperandInput` via:
//
//	PolicyAccessRuleConditionOperandArgs{...}
type PolicyAccessRuleConditionOperandInput interface {
	pulumi.Input

	ToPolicyAccessRuleConditionOperandOutput() PolicyAccessRuleConditionOperandOutput
	ToPolicyAccessRuleConditionOperandOutputWithContext(context.Context) PolicyAccessRuleConditionOperandOutput
}

type PolicyAccessRuleConditionOperandArgs struct {
	// (Optional) The ID of a server group resource
	Id pulumi.StringPtrInput `pulumi:"id"`
	// (Optional)
	IdpId pulumi.StringPtrInput `pulumi:"idpId"`
	// (Optional) LHS must always carry the string value ``id`` or the attribute ID of the resource being associated with the rule.
	Lhs pulumi.StringInput `pulumi:"lhs"`
	// (Optional)
	Name pulumi.StringPtrInput `pulumi:"name"`
	// (Optional) This is for specifying the policy critiera. Supported values: `APP`, `APP_GROUP`, `SAML`, `IDP`, `CLIENT_TYPE`, `TRUSTED_NETWORK`, `POSTURE`, `SCIM`, `SCIM_GROUP`, and `CLOUD_CONNECTOR_GROUP`. `TRUSTED_NETWORK`, and `CLIENT_TYPE`.
	ObjectType pulumi.StringInput `pulumi:"objectType"`
	// (Optional) RHS is either the ID attribute of a resource or fixed string value. Refer to the chart below for further details.
	Rhs      pulumi.StringPtrInput   `pulumi:"rhs"`
	RhsLists pulumi.StringArrayInput `pulumi:"rhsLists"`
}

func (PolicyAccessRuleConditionOperandArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAccessRuleConditionOperand)(nil)).Elem()
}

func (i PolicyAccessRuleConditionOperandArgs) ToPolicyAccessRuleConditionOperandOutput() PolicyAccessRuleConditionOperandOutput {
	return i.ToPolicyAccessRuleConditionOperandOutputWithContext(context.Background())
}

func (i PolicyAccessRuleConditionOperandArgs) ToPolicyAccessRuleConditionOperandOutputWithContext(ctx context.Context) PolicyAccessRuleConditionOperandOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAccessRuleConditionOperandOutput)
}

// PolicyAccessRuleConditionOperandArrayInput is an input type that accepts PolicyAccessRuleConditionOperandArray and PolicyAccessRuleConditionOperandArrayOutput values.
// You can construct a concrete instance of `PolicyAccessRuleConditionOperandArrayInput` via:
//
//	PolicyAccessRuleConditionOperandArray{ PolicyAccessRuleConditionOperandArgs{...} }
type PolicyAccessRuleConditionOperandArrayInput interface {
	pulumi.Input

	ToPolicyAccessRuleConditionOperandArrayOutput() PolicyAccessRuleConditionOperandArrayOutput
	ToPolicyAccessRuleConditionOperandArrayOutputWithContext(context.Context) PolicyAccessRuleConditionOperandArrayOutput
}

type PolicyAccessRuleConditionOperandArray []PolicyAccessRuleConditionOperandInput

func (PolicyAccessRuleConditionOperandArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyAccessRuleConditionOperand)(nil)).Elem()
}

func (i PolicyAccessRuleConditionOperandArray) ToPolicyAccessRuleConditionOperandArrayOutput() PolicyAccessRuleConditionOperandArrayOutput {
	return i.ToPolicyAccessRuleConditionOperandArrayOutputWithContext(context.Background())
}

func (i PolicyAccessRuleConditionOperandArray) ToPolicyAccessRuleConditionOperandArrayOutputWithContext(ctx context.Context) PolicyAccessRuleConditionOperandArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAccessRuleConditionOperandArrayOutput)
}

type PolicyAccessRuleConditionOperandOutput struct{ *pulumi.OutputState }

func (PolicyAccessRuleConditionOperandOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAccessRuleConditionOperand)(nil)).Elem()
}

func (o PolicyAccessRuleConditionOperandOutput) ToPolicyAccessRuleConditionOperandOutput() PolicyAccessRuleConditionOperandOutput {
	return o
}

func (o PolicyAccessRuleConditionOperandOutput) ToPolicyAccessRuleConditionOperandOutputWithContext(ctx context.Context) PolicyAccessRuleConditionOperandOutput {
	return o
}

// (Optional) The ID of a server group resource
func (o PolicyAccessRuleConditionOperandOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyAccessRuleConditionOperand) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// (Optional)
func (o PolicyAccessRuleConditionOperandOutput) IdpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyAccessRuleConditionOperand) *string { return v.IdpId }).(pulumi.StringPtrOutput)
}

// (Optional) LHS must always carry the string value “id“ or the attribute ID of the resource being associated with the rule.
func (o PolicyAccessRuleConditionOperandOutput) Lhs() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyAccessRuleConditionOperand) string { return v.Lhs }).(pulumi.StringOutput)
}

// (Optional)
func (o PolicyAccessRuleConditionOperandOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyAccessRuleConditionOperand) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// (Optional) This is for specifying the policy critiera. Supported values: `APP`, `APP_GROUP`, `SAML`, `IDP`, `CLIENT_TYPE`, `TRUSTED_NETWORK`, `POSTURE`, `SCIM`, `SCIM_GROUP`, and `CLOUD_CONNECTOR_GROUP`. `TRUSTED_NETWORK`, and `CLIENT_TYPE`.
func (o PolicyAccessRuleConditionOperandOutput) ObjectType() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyAccessRuleConditionOperand) string { return v.ObjectType }).(pulumi.StringOutput)
}

// (Optional) RHS is either the ID attribute of a resource or fixed string value. Refer to the chart below for further details.
func (o PolicyAccessRuleConditionOperandOutput) Rhs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyAccessRuleConditionOperand) *string { return v.Rhs }).(pulumi.StringPtrOutput)
}

func (o PolicyAccessRuleConditionOperandOutput) RhsLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PolicyAccessRuleConditionOperand) []string { return v.RhsLists }).(pulumi.StringArrayOutput)
}

type PolicyAccessRuleConditionOperandArrayOutput struct{ *pulumi.OutputState }

func (PolicyAccessRuleConditionOperandArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyAccessRuleConditionOperand)(nil)).Elem()
}

func (o PolicyAccessRuleConditionOperandArrayOutput) ToPolicyAccessRuleConditionOperandArrayOutput() PolicyAccessRuleConditionOperandArrayOutput {
	return o
}

func (o PolicyAccessRuleConditionOperandArrayOutput) ToPolicyAccessRuleConditionOperandArrayOutputWithContext(ctx context.Context) PolicyAccessRuleConditionOperandArrayOutput {
	return o
}

func (o PolicyAccessRuleConditionOperandArrayOutput) Index(i pulumi.IntInput) PolicyAccessRuleConditionOperandOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyAccessRuleConditionOperand {
		return vs[0].([]PolicyAccessRuleConditionOperand)[vs[1].(int)]
	}).(PolicyAccessRuleConditionOperandOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAccessRuleAppConnectorGroupInput)(nil)).Elem(), PolicyAccessRuleAppConnectorGroupArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAccessRuleAppConnectorGroupArrayInput)(nil)).Elem(), PolicyAccessRuleAppConnectorGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAccessRuleAppServerGroupInput)(nil)).Elem(), PolicyAccessRuleAppServerGroupArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAccessRuleAppServerGroupArrayInput)(nil)).Elem(), PolicyAccessRuleAppServerGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAccessRuleConditionInput)(nil)).Elem(), PolicyAccessRuleConditionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAccessRuleConditionArrayInput)(nil)).Elem(), PolicyAccessRuleConditionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAccessRuleConditionOperandInput)(nil)).Elem(), PolicyAccessRuleConditionOperandArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAccessRuleConditionOperandArrayInput)(nil)).Elem(), PolicyAccessRuleConditionOperandArray{})
	pulumi.RegisterOutputType(PolicyAccessRuleAppConnectorGroupOutput{})
	pulumi.RegisterOutputType(PolicyAccessRuleAppConnectorGroupArrayOutput{})
	pulumi.RegisterOutputType(PolicyAccessRuleAppServerGroupOutput{})
	pulumi.RegisterOutputType(PolicyAccessRuleAppServerGroupArrayOutput{})
	pulumi.RegisterOutputType(PolicyAccessRuleConditionOutput{})
	pulumi.RegisterOutputType(PolicyAccessRuleConditionArrayOutput{})
	pulumi.RegisterOutputType(PolicyAccessRuleConditionOperandOutput{})
	pulumi.RegisterOutputType(PolicyAccessRuleConditionOperandArrayOutput{})
}