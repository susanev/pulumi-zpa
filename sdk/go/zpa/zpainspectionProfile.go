// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The  **zpa_inspection_profile** resource creates an inspection profile in the Zscaler Private Access cloud. This resource can then be referenced in an inspection custom control resource.
//
// ## Import
//
// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer)
type ZPAInspectionProfile struct {
	pulumi.CustomResourceState

	AssociateAllControls pulumi.BoolPtrOutput `pulumi:"associateAllControls"`
	// (Optional)
	CommonGlobalOverrideActionsConfig pulumi.StringMapOutput `pulumi:"commonGlobalOverrideActionsConfig"`
	// (Optional) Types for custom controls
	ControlsInfos ZPAInspectionProfileControlsInfoArrayOutput `pulumi:"controlsInfos"`
	// (Optional) Types for custom controls
	CustomControls ZPAInspectionProfileCustomControlArrayOutput `pulumi:"customControls"`
	// Description of the inspection profile.
	Description          pulumi.StringOutput      `pulumi:"description"`
	GlobalControlActions pulumi.StringArrayOutput `pulumi:"globalControlActions"`
	IncarnationNumber    pulumi.StringOutput      `pulumi:"incarnationNumber"`
	// The name of the inspection profile.
	Name pulumi.StringOutput `pulumi:"name"`
	// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
	ParanoiaLevel pulumi.StringOutput `pulumi:"paranoiaLevel"`
	// The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefinedControls` can be set by using the data source `dataSourceZpaPredefinedControls` or by group using the data source `getZPAInspectionAllPredefinedControls`.
	PredefinedControls        ZPAInspectionProfilePredefinedControlArrayOutput `pulumi:"predefinedControls"`
	PredefinedControlsVersion pulumi.StringOutput                              `pulumi:"predefinedControlsVersion"`
}

// NewZPAInspectionProfile registers a new resource with the given unique name, arguments, and options.
func NewZPAInspectionProfile(ctx *pulumi.Context,
	name string, args *ZPAInspectionProfileArgs, opts ...pulumi.ResourceOption) (*ZPAInspectionProfile, error) {
	if args == nil {
		args = &ZPAInspectionProfileArgs{}
	}

	var resource ZPAInspectionProfile
	err := ctx.RegisterResource("zpa:index/zPAInspectionProfile:ZPAInspectionProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZPAInspectionProfile gets an existing ZPAInspectionProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZPAInspectionProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZPAInspectionProfileState, opts ...pulumi.ResourceOption) (*ZPAInspectionProfile, error) {
	var resource ZPAInspectionProfile
	err := ctx.ReadResource("zpa:index/zPAInspectionProfile:ZPAInspectionProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ZPAInspectionProfile resources.
type zpainspectionProfileState struct {
	AssociateAllControls *bool `pulumi:"associateAllControls"`
	// (Optional)
	CommonGlobalOverrideActionsConfig map[string]string `pulumi:"commonGlobalOverrideActionsConfig"`
	// (Optional) Types for custom controls
	ControlsInfos []ZPAInspectionProfileControlsInfo `pulumi:"controlsInfos"`
	// (Optional) Types for custom controls
	CustomControls []ZPAInspectionProfileCustomControl `pulumi:"customControls"`
	// Description of the inspection profile.
	Description          *string  `pulumi:"description"`
	GlobalControlActions []string `pulumi:"globalControlActions"`
	IncarnationNumber    *string  `pulumi:"incarnationNumber"`
	// The name of the inspection profile.
	Name *string `pulumi:"name"`
	// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
	ParanoiaLevel *string `pulumi:"paranoiaLevel"`
	// The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefinedControls` can be set by using the data source `dataSourceZpaPredefinedControls` or by group using the data source `getZPAInspectionAllPredefinedControls`.
	PredefinedControls        []ZPAInspectionProfilePredefinedControl `pulumi:"predefinedControls"`
	PredefinedControlsVersion *string                                 `pulumi:"predefinedControlsVersion"`
}

type ZPAInspectionProfileState struct {
	AssociateAllControls pulumi.BoolPtrInput
	// (Optional)
	CommonGlobalOverrideActionsConfig pulumi.StringMapInput
	// (Optional) Types for custom controls
	ControlsInfos ZPAInspectionProfileControlsInfoArrayInput
	// (Optional) Types for custom controls
	CustomControls ZPAInspectionProfileCustomControlArrayInput
	// Description of the inspection profile.
	Description          pulumi.StringPtrInput
	GlobalControlActions pulumi.StringArrayInput
	IncarnationNumber    pulumi.StringPtrInput
	// The name of the inspection profile.
	Name pulumi.StringPtrInput
	// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
	ParanoiaLevel pulumi.StringPtrInput
	// The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefinedControls` can be set by using the data source `dataSourceZpaPredefinedControls` or by group using the data source `getZPAInspectionAllPredefinedControls`.
	PredefinedControls        ZPAInspectionProfilePredefinedControlArrayInput
	PredefinedControlsVersion pulumi.StringPtrInput
}

func (ZPAInspectionProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*zpainspectionProfileState)(nil)).Elem()
}

type zpainspectionProfileArgs struct {
	AssociateAllControls *bool `pulumi:"associateAllControls"`
	// (Optional)
	CommonGlobalOverrideActionsConfig map[string]string `pulumi:"commonGlobalOverrideActionsConfig"`
	// (Optional) Types for custom controls
	ControlsInfos []ZPAInspectionProfileControlsInfo `pulumi:"controlsInfos"`
	// (Optional) Types for custom controls
	CustomControls []ZPAInspectionProfileCustomControl `pulumi:"customControls"`
	// Description of the inspection profile.
	Description          *string  `pulumi:"description"`
	GlobalControlActions []string `pulumi:"globalControlActions"`
	IncarnationNumber    *string  `pulumi:"incarnationNumber"`
	// The name of the inspection profile.
	Name *string `pulumi:"name"`
	// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
	ParanoiaLevel *string `pulumi:"paranoiaLevel"`
	// The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefinedControls` can be set by using the data source `dataSourceZpaPredefinedControls` or by group using the data source `getZPAInspectionAllPredefinedControls`.
	PredefinedControls        []ZPAInspectionProfilePredefinedControl `pulumi:"predefinedControls"`
	PredefinedControlsVersion *string                                 `pulumi:"predefinedControlsVersion"`
}

// The set of arguments for constructing a ZPAInspectionProfile resource.
type ZPAInspectionProfileArgs struct {
	AssociateAllControls pulumi.BoolPtrInput
	// (Optional)
	CommonGlobalOverrideActionsConfig pulumi.StringMapInput
	// (Optional) Types for custom controls
	ControlsInfos ZPAInspectionProfileControlsInfoArrayInput
	// (Optional) Types for custom controls
	CustomControls ZPAInspectionProfileCustomControlArrayInput
	// Description of the inspection profile.
	Description          pulumi.StringPtrInput
	GlobalControlActions pulumi.StringArrayInput
	IncarnationNumber    pulumi.StringPtrInput
	// The name of the inspection profile.
	Name pulumi.StringPtrInput
	// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
	ParanoiaLevel pulumi.StringPtrInput
	// The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefinedControls` can be set by using the data source `dataSourceZpaPredefinedControls` or by group using the data source `getZPAInspectionAllPredefinedControls`.
	PredefinedControls        ZPAInspectionProfilePredefinedControlArrayInput
	PredefinedControlsVersion pulumi.StringPtrInput
}

func (ZPAInspectionProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zpainspectionProfileArgs)(nil)).Elem()
}

type ZPAInspectionProfileInput interface {
	pulumi.Input

	ToZPAInspectionProfileOutput() ZPAInspectionProfileOutput
	ToZPAInspectionProfileOutputWithContext(ctx context.Context) ZPAInspectionProfileOutput
}

func (*ZPAInspectionProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**ZPAInspectionProfile)(nil)).Elem()
}

func (i *ZPAInspectionProfile) ToZPAInspectionProfileOutput() ZPAInspectionProfileOutput {
	return i.ToZPAInspectionProfileOutputWithContext(context.Background())
}

func (i *ZPAInspectionProfile) ToZPAInspectionProfileOutputWithContext(ctx context.Context) ZPAInspectionProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZPAInspectionProfileOutput)
}

// ZPAInspectionProfileArrayInput is an input type that accepts ZPAInspectionProfileArray and ZPAInspectionProfileArrayOutput values.
// You can construct a concrete instance of `ZPAInspectionProfileArrayInput` via:
//
//	ZPAInspectionProfileArray{ ZPAInspectionProfileArgs{...} }
type ZPAInspectionProfileArrayInput interface {
	pulumi.Input

	ToZPAInspectionProfileArrayOutput() ZPAInspectionProfileArrayOutput
	ToZPAInspectionProfileArrayOutputWithContext(context.Context) ZPAInspectionProfileArrayOutput
}

type ZPAInspectionProfileArray []ZPAInspectionProfileInput

func (ZPAInspectionProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ZPAInspectionProfile)(nil)).Elem()
}

func (i ZPAInspectionProfileArray) ToZPAInspectionProfileArrayOutput() ZPAInspectionProfileArrayOutput {
	return i.ToZPAInspectionProfileArrayOutputWithContext(context.Background())
}

func (i ZPAInspectionProfileArray) ToZPAInspectionProfileArrayOutputWithContext(ctx context.Context) ZPAInspectionProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZPAInspectionProfileArrayOutput)
}

// ZPAInspectionProfileMapInput is an input type that accepts ZPAInspectionProfileMap and ZPAInspectionProfileMapOutput values.
// You can construct a concrete instance of `ZPAInspectionProfileMapInput` via:
//
//	ZPAInspectionProfileMap{ "key": ZPAInspectionProfileArgs{...} }
type ZPAInspectionProfileMapInput interface {
	pulumi.Input

	ToZPAInspectionProfileMapOutput() ZPAInspectionProfileMapOutput
	ToZPAInspectionProfileMapOutputWithContext(context.Context) ZPAInspectionProfileMapOutput
}

type ZPAInspectionProfileMap map[string]ZPAInspectionProfileInput

func (ZPAInspectionProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ZPAInspectionProfile)(nil)).Elem()
}

func (i ZPAInspectionProfileMap) ToZPAInspectionProfileMapOutput() ZPAInspectionProfileMapOutput {
	return i.ToZPAInspectionProfileMapOutputWithContext(context.Background())
}

func (i ZPAInspectionProfileMap) ToZPAInspectionProfileMapOutputWithContext(ctx context.Context) ZPAInspectionProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZPAInspectionProfileMapOutput)
}

type ZPAInspectionProfileOutput struct{ *pulumi.OutputState }

func (ZPAInspectionProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ZPAInspectionProfile)(nil)).Elem()
}

func (o ZPAInspectionProfileOutput) ToZPAInspectionProfileOutput() ZPAInspectionProfileOutput {
	return o
}

func (o ZPAInspectionProfileOutput) ToZPAInspectionProfileOutputWithContext(ctx context.Context) ZPAInspectionProfileOutput {
	return o
}

func (o ZPAInspectionProfileOutput) AssociateAllControls() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ZPAInspectionProfile) pulumi.BoolPtrOutput { return v.AssociateAllControls }).(pulumi.BoolPtrOutput)
}

// (Optional)
func (o ZPAInspectionProfileOutput) CommonGlobalOverrideActionsConfig() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ZPAInspectionProfile) pulumi.StringMapOutput { return v.CommonGlobalOverrideActionsConfig }).(pulumi.StringMapOutput)
}

// (Optional) Types for custom controls
func (o ZPAInspectionProfileOutput) ControlsInfos() ZPAInspectionProfileControlsInfoArrayOutput {
	return o.ApplyT(func(v *ZPAInspectionProfile) ZPAInspectionProfileControlsInfoArrayOutput { return v.ControlsInfos }).(ZPAInspectionProfileControlsInfoArrayOutput)
}

// (Optional) Types for custom controls
func (o ZPAInspectionProfileOutput) CustomControls() ZPAInspectionProfileCustomControlArrayOutput {
	return o.ApplyT(func(v *ZPAInspectionProfile) ZPAInspectionProfileCustomControlArrayOutput { return v.CustomControls }).(ZPAInspectionProfileCustomControlArrayOutput)
}

// Description of the inspection profile.
func (o ZPAInspectionProfileOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ZPAInspectionProfile) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o ZPAInspectionProfileOutput) GlobalControlActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ZPAInspectionProfile) pulumi.StringArrayOutput { return v.GlobalControlActions }).(pulumi.StringArrayOutput)
}

func (o ZPAInspectionProfileOutput) IncarnationNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *ZPAInspectionProfile) pulumi.StringOutput { return v.IncarnationNumber }).(pulumi.StringOutput)
}

// The name of the inspection profile.
func (o ZPAInspectionProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ZPAInspectionProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
func (o ZPAInspectionProfileOutput) ParanoiaLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *ZPAInspectionProfile) pulumi.StringOutput { return v.ParanoiaLevel }).(pulumi.StringOutput)
}

// The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefinedControls` can be set by using the data source `dataSourceZpaPredefinedControls` or by group using the data source `getZPAInspectionAllPredefinedControls`.
func (o ZPAInspectionProfileOutput) PredefinedControls() ZPAInspectionProfilePredefinedControlArrayOutput {
	return o.ApplyT(func(v *ZPAInspectionProfile) ZPAInspectionProfilePredefinedControlArrayOutput {
		return v.PredefinedControls
	}).(ZPAInspectionProfilePredefinedControlArrayOutput)
}

func (o ZPAInspectionProfileOutput) PredefinedControlsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ZPAInspectionProfile) pulumi.StringOutput { return v.PredefinedControlsVersion }).(pulumi.StringOutput)
}

type ZPAInspectionProfileArrayOutput struct{ *pulumi.OutputState }

func (ZPAInspectionProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ZPAInspectionProfile)(nil)).Elem()
}

func (o ZPAInspectionProfileArrayOutput) ToZPAInspectionProfileArrayOutput() ZPAInspectionProfileArrayOutput {
	return o
}

func (o ZPAInspectionProfileArrayOutput) ToZPAInspectionProfileArrayOutputWithContext(ctx context.Context) ZPAInspectionProfileArrayOutput {
	return o
}

func (o ZPAInspectionProfileArrayOutput) Index(i pulumi.IntInput) ZPAInspectionProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ZPAInspectionProfile {
		return vs[0].([]*ZPAInspectionProfile)[vs[1].(int)]
	}).(ZPAInspectionProfileOutput)
}

type ZPAInspectionProfileMapOutput struct{ *pulumi.OutputState }

func (ZPAInspectionProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ZPAInspectionProfile)(nil)).Elem()
}

func (o ZPAInspectionProfileMapOutput) ToZPAInspectionProfileMapOutput() ZPAInspectionProfileMapOutput {
	return o
}

func (o ZPAInspectionProfileMapOutput) ToZPAInspectionProfileMapOutputWithContext(ctx context.Context) ZPAInspectionProfileMapOutput {
	return o
}

func (o ZPAInspectionProfileMapOutput) MapIndex(k pulumi.StringInput) ZPAInspectionProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ZPAInspectionProfile {
		return vs[0].(map[string]*ZPAInspectionProfile)[vs[1].(string)]
	}).(ZPAInspectionProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ZPAInspectionProfileInput)(nil)).Elem(), &ZPAInspectionProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZPAInspectionProfileArrayInput)(nil)).Elem(), ZPAInspectionProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZPAInspectionProfileMapInput)(nil)).Elem(), ZPAInspectionProfileMap{})
	pulumi.RegisterOutputType(ZPAInspectionProfileOutput{})
	pulumi.RegisterOutputType(ZPAInspectionProfileArrayOutput{})
	pulumi.RegisterOutputType(ZPAInspectionProfileMapOutput{})
}
