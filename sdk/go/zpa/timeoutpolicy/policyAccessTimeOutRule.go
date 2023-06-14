// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package timeoutpolicy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PolicyAccessTimeOutRule struct {
	pulumi.CustomResourceState

	// This is for providing the rule action.
	Action pulumi.StringPtrOutput `pulumi:"action"`
	// This field defines the description of the server.
	ActionId          pulumi.StringPtrOutput `pulumi:"actionId"`
	BypassDefaultRule pulumi.BoolPtrOutput   `pulumi:"bypassDefaultRule"`
	// This is for proviidng the set of conditions for the policy.
	Conditions PolicyAccessTimeOutRuleConditionArrayOutput `pulumi:"conditions"`
	// This is for providing a customer message for the user.
	CustomMsg pulumi.StringPtrOutput `pulumi:"customMsg"`
	// This is for providing a customer message for the user.
	DefaultRule pulumi.BoolPtrOutput `pulumi:"defaultRule"`
	// This is the description of the access policy.
	Description    pulumi.StringPtrOutput `pulumi:"description"`
	LssDefaultRule pulumi.BoolPtrOutput   `pulumi:"lssDefaultRule"`
	// This is the name of the policy.
	Name                   pulumi.StringOutput    `pulumi:"name"`
	Operator               pulumi.StringOutput    `pulumi:"operator"`
	PolicySetId            pulumi.StringOutput    `pulumi:"policySetId"`
	PolicyType             pulumi.StringOutput    `pulumi:"policyType"`
	Priority               pulumi.StringOutput    `pulumi:"priority"`
	ReauthDefaultRule      pulumi.BoolPtrOutput   `pulumi:"reauthDefaultRule"`
	ReauthIdleTimeout      pulumi.StringPtrOutput `pulumi:"reauthIdleTimeout"`
	ReauthTimeout          pulumi.StringPtrOutput `pulumi:"reauthTimeout"`
	RuleOrder              pulumi.StringOutput    `pulumi:"ruleOrder"`
	ZpnInspectionProfileId pulumi.StringPtrOutput `pulumi:"zpnInspectionProfileId"`
}

// NewPolicyAccessTimeOutRule registers a new resource with the given unique name, arguments, and options.
func NewPolicyAccessTimeOutRule(ctx *pulumi.Context,
	name string, args *PolicyAccessTimeOutRuleArgs, opts ...pulumi.ResourceOption) (*PolicyAccessTimeOutRule, error) {
	if args == nil {
		args = &PolicyAccessTimeOutRuleArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource PolicyAccessTimeOutRule
	err := ctx.RegisterResource("zpa:TimeoutPolicy/policyAccessTimeOutRule:PolicyAccessTimeOutRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyAccessTimeOutRule gets an existing PolicyAccessTimeOutRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyAccessTimeOutRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyAccessTimeOutRuleState, opts ...pulumi.ResourceOption) (*PolicyAccessTimeOutRule, error) {
	var resource PolicyAccessTimeOutRule
	err := ctx.ReadResource("zpa:TimeoutPolicy/policyAccessTimeOutRule:PolicyAccessTimeOutRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyAccessTimeOutRule resources.
type policyAccessTimeOutRuleState struct {
	// This is for providing the rule action.
	Action *string `pulumi:"action"`
	// This field defines the description of the server.
	ActionId          *string `pulumi:"actionId"`
	BypassDefaultRule *bool   `pulumi:"bypassDefaultRule"`
	// This is for proviidng the set of conditions for the policy.
	Conditions []PolicyAccessTimeOutRuleCondition `pulumi:"conditions"`
	// This is for providing a customer message for the user.
	CustomMsg *string `pulumi:"customMsg"`
	// This is for providing a customer message for the user.
	DefaultRule *bool `pulumi:"defaultRule"`
	// This is the description of the access policy.
	Description    *string `pulumi:"description"`
	LssDefaultRule *bool   `pulumi:"lssDefaultRule"`
	// This is the name of the policy.
	Name                   *string `pulumi:"name"`
	Operator               *string `pulumi:"operator"`
	PolicySetId            *string `pulumi:"policySetId"`
	PolicyType             *string `pulumi:"policyType"`
	Priority               *string `pulumi:"priority"`
	ReauthDefaultRule      *bool   `pulumi:"reauthDefaultRule"`
	ReauthIdleTimeout      *string `pulumi:"reauthIdleTimeout"`
	ReauthTimeout          *string `pulumi:"reauthTimeout"`
	RuleOrder              *string `pulumi:"ruleOrder"`
	ZpnInspectionProfileId *string `pulumi:"zpnInspectionProfileId"`
}

type PolicyAccessTimeOutRuleState struct {
	// This is for providing the rule action.
	Action pulumi.StringPtrInput
	// This field defines the description of the server.
	ActionId          pulumi.StringPtrInput
	BypassDefaultRule pulumi.BoolPtrInput
	// This is for proviidng the set of conditions for the policy.
	Conditions PolicyAccessTimeOutRuleConditionArrayInput
	// This is for providing a customer message for the user.
	CustomMsg pulumi.StringPtrInput
	// This is for providing a customer message for the user.
	DefaultRule pulumi.BoolPtrInput
	// This is the description of the access policy.
	Description    pulumi.StringPtrInput
	LssDefaultRule pulumi.BoolPtrInput
	// This is the name of the policy.
	Name                   pulumi.StringPtrInput
	Operator               pulumi.StringPtrInput
	PolicySetId            pulumi.StringPtrInput
	PolicyType             pulumi.StringPtrInput
	Priority               pulumi.StringPtrInput
	ReauthDefaultRule      pulumi.BoolPtrInput
	ReauthIdleTimeout      pulumi.StringPtrInput
	ReauthTimeout          pulumi.StringPtrInput
	RuleOrder              pulumi.StringPtrInput
	ZpnInspectionProfileId pulumi.StringPtrInput
}

func (PolicyAccessTimeOutRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAccessTimeOutRuleState)(nil)).Elem()
}

type policyAccessTimeOutRuleArgs struct {
	// This is for providing the rule action.
	Action *string `pulumi:"action"`
	// This field defines the description of the server.
	ActionId          *string `pulumi:"actionId"`
	BypassDefaultRule *bool   `pulumi:"bypassDefaultRule"`
	// This is for proviidng the set of conditions for the policy.
	Conditions []PolicyAccessTimeOutRuleCondition `pulumi:"conditions"`
	// This is for providing a customer message for the user.
	CustomMsg *string `pulumi:"customMsg"`
	// This is for providing a customer message for the user.
	DefaultRule *bool `pulumi:"defaultRule"`
	// This is the description of the access policy.
	Description    *string `pulumi:"description"`
	LssDefaultRule *bool   `pulumi:"lssDefaultRule"`
	// This is the name of the policy.
	Name                   *string `pulumi:"name"`
	Operator               *string `pulumi:"operator"`
	PolicySetId            *string `pulumi:"policySetId"`
	PolicyType             *string `pulumi:"policyType"`
	Priority               *string `pulumi:"priority"`
	ReauthDefaultRule      *bool   `pulumi:"reauthDefaultRule"`
	ReauthIdleTimeout      *string `pulumi:"reauthIdleTimeout"`
	ReauthTimeout          *string `pulumi:"reauthTimeout"`
	RuleOrder              *string `pulumi:"ruleOrder"`
	ZpnInspectionProfileId *string `pulumi:"zpnInspectionProfileId"`
}

// The set of arguments for constructing a PolicyAccessTimeOutRule resource.
type PolicyAccessTimeOutRuleArgs struct {
	// This is for providing the rule action.
	Action pulumi.StringPtrInput
	// This field defines the description of the server.
	ActionId          pulumi.StringPtrInput
	BypassDefaultRule pulumi.BoolPtrInput
	// This is for proviidng the set of conditions for the policy.
	Conditions PolicyAccessTimeOutRuleConditionArrayInput
	// This is for providing a customer message for the user.
	CustomMsg pulumi.StringPtrInput
	// This is for providing a customer message for the user.
	DefaultRule pulumi.BoolPtrInput
	// This is the description of the access policy.
	Description    pulumi.StringPtrInput
	LssDefaultRule pulumi.BoolPtrInput
	// This is the name of the policy.
	Name                   pulumi.StringPtrInput
	Operator               pulumi.StringPtrInput
	PolicySetId            pulumi.StringPtrInput
	PolicyType             pulumi.StringPtrInput
	Priority               pulumi.StringPtrInput
	ReauthDefaultRule      pulumi.BoolPtrInput
	ReauthIdleTimeout      pulumi.StringPtrInput
	ReauthTimeout          pulumi.StringPtrInput
	RuleOrder              pulumi.StringPtrInput
	ZpnInspectionProfileId pulumi.StringPtrInput
}

func (PolicyAccessTimeOutRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAccessTimeOutRuleArgs)(nil)).Elem()
}

type PolicyAccessTimeOutRuleInput interface {
	pulumi.Input

	ToPolicyAccessTimeOutRuleOutput() PolicyAccessTimeOutRuleOutput
	ToPolicyAccessTimeOutRuleOutputWithContext(ctx context.Context) PolicyAccessTimeOutRuleOutput
}

func (*PolicyAccessTimeOutRule) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAccessTimeOutRule)(nil)).Elem()
}

func (i *PolicyAccessTimeOutRule) ToPolicyAccessTimeOutRuleOutput() PolicyAccessTimeOutRuleOutput {
	return i.ToPolicyAccessTimeOutRuleOutputWithContext(context.Background())
}

func (i *PolicyAccessTimeOutRule) ToPolicyAccessTimeOutRuleOutputWithContext(ctx context.Context) PolicyAccessTimeOutRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAccessTimeOutRuleOutput)
}

// PolicyAccessTimeOutRuleArrayInput is an input type that accepts PolicyAccessTimeOutRuleArray and PolicyAccessTimeOutRuleArrayOutput values.
// You can construct a concrete instance of `PolicyAccessTimeOutRuleArrayInput` via:
//
//	PolicyAccessTimeOutRuleArray{ PolicyAccessTimeOutRuleArgs{...} }
type PolicyAccessTimeOutRuleArrayInput interface {
	pulumi.Input

	ToPolicyAccessTimeOutRuleArrayOutput() PolicyAccessTimeOutRuleArrayOutput
	ToPolicyAccessTimeOutRuleArrayOutputWithContext(context.Context) PolicyAccessTimeOutRuleArrayOutput
}

type PolicyAccessTimeOutRuleArray []PolicyAccessTimeOutRuleInput

func (PolicyAccessTimeOutRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyAccessTimeOutRule)(nil)).Elem()
}

func (i PolicyAccessTimeOutRuleArray) ToPolicyAccessTimeOutRuleArrayOutput() PolicyAccessTimeOutRuleArrayOutput {
	return i.ToPolicyAccessTimeOutRuleArrayOutputWithContext(context.Background())
}

func (i PolicyAccessTimeOutRuleArray) ToPolicyAccessTimeOutRuleArrayOutputWithContext(ctx context.Context) PolicyAccessTimeOutRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAccessTimeOutRuleArrayOutput)
}

// PolicyAccessTimeOutRuleMapInput is an input type that accepts PolicyAccessTimeOutRuleMap and PolicyAccessTimeOutRuleMapOutput values.
// You can construct a concrete instance of `PolicyAccessTimeOutRuleMapInput` via:
//
//	PolicyAccessTimeOutRuleMap{ "key": PolicyAccessTimeOutRuleArgs{...} }
type PolicyAccessTimeOutRuleMapInput interface {
	pulumi.Input

	ToPolicyAccessTimeOutRuleMapOutput() PolicyAccessTimeOutRuleMapOutput
	ToPolicyAccessTimeOutRuleMapOutputWithContext(context.Context) PolicyAccessTimeOutRuleMapOutput
}

type PolicyAccessTimeOutRuleMap map[string]PolicyAccessTimeOutRuleInput

func (PolicyAccessTimeOutRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyAccessTimeOutRule)(nil)).Elem()
}

func (i PolicyAccessTimeOutRuleMap) ToPolicyAccessTimeOutRuleMapOutput() PolicyAccessTimeOutRuleMapOutput {
	return i.ToPolicyAccessTimeOutRuleMapOutputWithContext(context.Background())
}

func (i PolicyAccessTimeOutRuleMap) ToPolicyAccessTimeOutRuleMapOutputWithContext(ctx context.Context) PolicyAccessTimeOutRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAccessTimeOutRuleMapOutput)
}

type PolicyAccessTimeOutRuleOutput struct{ *pulumi.OutputState }

func (PolicyAccessTimeOutRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAccessTimeOutRule)(nil)).Elem()
}

func (o PolicyAccessTimeOutRuleOutput) ToPolicyAccessTimeOutRuleOutput() PolicyAccessTimeOutRuleOutput {
	return o
}

func (o PolicyAccessTimeOutRuleOutput) ToPolicyAccessTimeOutRuleOutputWithContext(ctx context.Context) PolicyAccessTimeOutRuleOutput {
	return o
}

// This is for providing the rule action.
func (o PolicyAccessTimeOutRuleOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAccessTimeOutRule) pulumi.StringPtrOutput { return v.Action }).(pulumi.StringPtrOutput)
}

// This field defines the description of the server.
func (o PolicyAccessTimeOutRuleOutput) ActionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAccessTimeOutRule) pulumi.StringPtrOutput { return v.ActionId }).(pulumi.StringPtrOutput)
}

func (o PolicyAccessTimeOutRuleOutput) BypassDefaultRule() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PolicyAccessTimeOutRule) pulumi.BoolPtrOutput { return v.BypassDefaultRule }).(pulumi.BoolPtrOutput)
}

// This is for proviidng the set of conditions for the policy.
func (o PolicyAccessTimeOutRuleOutput) Conditions() PolicyAccessTimeOutRuleConditionArrayOutput {
	return o.ApplyT(func(v *PolicyAccessTimeOutRule) PolicyAccessTimeOutRuleConditionArrayOutput { return v.Conditions }).(PolicyAccessTimeOutRuleConditionArrayOutput)
}

// This is for providing a customer message for the user.
func (o PolicyAccessTimeOutRuleOutput) CustomMsg() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAccessTimeOutRule) pulumi.StringPtrOutput { return v.CustomMsg }).(pulumi.StringPtrOutput)
}

// This is for providing a customer message for the user.
func (o PolicyAccessTimeOutRuleOutput) DefaultRule() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PolicyAccessTimeOutRule) pulumi.BoolPtrOutput { return v.DefaultRule }).(pulumi.BoolPtrOutput)
}

// This is the description of the access policy.
func (o PolicyAccessTimeOutRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAccessTimeOutRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PolicyAccessTimeOutRuleOutput) LssDefaultRule() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PolicyAccessTimeOutRule) pulumi.BoolPtrOutput { return v.LssDefaultRule }).(pulumi.BoolPtrOutput)
}

// This is the name of the policy.
func (o PolicyAccessTimeOutRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAccessTimeOutRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PolicyAccessTimeOutRuleOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAccessTimeOutRule) pulumi.StringOutput { return v.Operator }).(pulumi.StringOutput)
}

func (o PolicyAccessTimeOutRuleOutput) PolicySetId() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAccessTimeOutRule) pulumi.StringOutput { return v.PolicySetId }).(pulumi.StringOutput)
}

func (o PolicyAccessTimeOutRuleOutput) PolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAccessTimeOutRule) pulumi.StringOutput { return v.PolicyType }).(pulumi.StringOutput)
}

func (o PolicyAccessTimeOutRuleOutput) Priority() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAccessTimeOutRule) pulumi.StringOutput { return v.Priority }).(pulumi.StringOutput)
}

func (o PolicyAccessTimeOutRuleOutput) ReauthDefaultRule() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PolicyAccessTimeOutRule) pulumi.BoolPtrOutput { return v.ReauthDefaultRule }).(pulumi.BoolPtrOutput)
}

func (o PolicyAccessTimeOutRuleOutput) ReauthIdleTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAccessTimeOutRule) pulumi.StringPtrOutput { return v.ReauthIdleTimeout }).(pulumi.StringPtrOutput)
}

func (o PolicyAccessTimeOutRuleOutput) ReauthTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAccessTimeOutRule) pulumi.StringPtrOutput { return v.ReauthTimeout }).(pulumi.StringPtrOutput)
}

func (o PolicyAccessTimeOutRuleOutput) RuleOrder() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAccessTimeOutRule) pulumi.StringOutput { return v.RuleOrder }).(pulumi.StringOutput)
}

func (o PolicyAccessTimeOutRuleOutput) ZpnInspectionProfileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAccessTimeOutRule) pulumi.StringPtrOutput { return v.ZpnInspectionProfileId }).(pulumi.StringPtrOutput)
}

type PolicyAccessTimeOutRuleArrayOutput struct{ *pulumi.OutputState }

func (PolicyAccessTimeOutRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyAccessTimeOutRule)(nil)).Elem()
}

func (o PolicyAccessTimeOutRuleArrayOutput) ToPolicyAccessTimeOutRuleArrayOutput() PolicyAccessTimeOutRuleArrayOutput {
	return o
}

func (o PolicyAccessTimeOutRuleArrayOutput) ToPolicyAccessTimeOutRuleArrayOutputWithContext(ctx context.Context) PolicyAccessTimeOutRuleArrayOutput {
	return o
}

func (o PolicyAccessTimeOutRuleArrayOutput) Index(i pulumi.IntInput) PolicyAccessTimeOutRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PolicyAccessTimeOutRule {
		return vs[0].([]*PolicyAccessTimeOutRule)[vs[1].(int)]
	}).(PolicyAccessTimeOutRuleOutput)
}

type PolicyAccessTimeOutRuleMapOutput struct{ *pulumi.OutputState }

func (PolicyAccessTimeOutRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyAccessTimeOutRule)(nil)).Elem()
}

func (o PolicyAccessTimeOutRuleMapOutput) ToPolicyAccessTimeOutRuleMapOutput() PolicyAccessTimeOutRuleMapOutput {
	return o
}

func (o PolicyAccessTimeOutRuleMapOutput) ToPolicyAccessTimeOutRuleMapOutputWithContext(ctx context.Context) PolicyAccessTimeOutRuleMapOutput {
	return o
}

func (o PolicyAccessTimeOutRuleMapOutput) MapIndex(k pulumi.StringInput) PolicyAccessTimeOutRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PolicyAccessTimeOutRule {
		return vs[0].(map[string]*PolicyAccessTimeOutRule)[vs[1].(string)]
	}).(PolicyAccessTimeOutRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAccessTimeOutRuleInput)(nil)).Elem(), &PolicyAccessTimeOutRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAccessTimeOutRuleArrayInput)(nil)).Elem(), PolicyAccessTimeOutRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAccessTimeOutRuleMapInput)(nil)).Elem(), PolicyAccessTimeOutRuleMap{})
	pulumi.RegisterOutputType(PolicyAccessTimeOutRuleOutput{})
	pulumi.RegisterOutputType(PolicyAccessTimeOutRuleArrayOutput{})
	pulumi.RegisterOutputType(PolicyAccessTimeOutRuleMapOutput{})
}