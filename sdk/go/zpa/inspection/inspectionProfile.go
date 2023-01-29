// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package inspection

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
type InspectionProfile struct {
	pulumi.CustomResourceState

	AssociateAllControls pulumi.BoolPtrOutput `pulumi:"associateAllControls"`
	// (Optional)
	CommonGlobalOverrideActionsConfig pulumi.StringMapOutput `pulumi:"commonGlobalOverrideActionsConfig"`
	// (Optional) Types for custom controls
	ControlsInfos InspectionProfileControlsInfoArrayOutput `pulumi:"controlsInfos"`
	// (Optional) Types for custom controls
	CustomControls InspectionProfileCustomControlArrayOutput `pulumi:"customControls"`
	// Description of the inspection profile.
	Description          pulumi.StringOutput      `pulumi:"description"`
	GlobalControlActions pulumi.StringArrayOutput `pulumi:"globalControlActions"`
	IncarnationNumber    pulumi.StringOutput      `pulumi:"incarnationNumber"`
	// The name of the inspection profile.
	Name pulumi.StringOutput `pulumi:"name"`
	// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
	ParanoiaLevel pulumi.StringOutput `pulumi:"paranoiaLevel"`
	// The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefinedControls` can be set by using the data source `dataSourceZpaPredefinedControls` or by group using the data source `Inspection.getInspectionAllPredefinedControls`.
	PredefinedControls        InspectionProfilePredefinedControlArrayOutput `pulumi:"predefinedControls"`
	PredefinedControlsVersion pulumi.StringOutput                           `pulumi:"predefinedControlsVersion"`
	// (string)
	WebSocketControls InspectionProfileWebSocketControlArrayOutput `pulumi:"webSocketControls"`
}

// NewInspectionProfile registers a new resource with the given unique name, arguments, and options.
func NewInspectionProfile(ctx *pulumi.Context,
	name string, args *InspectionProfileArgs, opts ...pulumi.ResourceOption) (*InspectionProfile, error) {
	if args == nil {
		args = &InspectionProfileArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource InspectionProfile
	err := ctx.RegisterResource("zpa:Inspection/inspectionProfile:InspectionProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInspectionProfile gets an existing InspectionProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInspectionProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InspectionProfileState, opts ...pulumi.ResourceOption) (*InspectionProfile, error) {
	var resource InspectionProfile
	err := ctx.ReadResource("zpa:Inspection/inspectionProfile:InspectionProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InspectionProfile resources.
type inspectionProfileState struct {
	AssociateAllControls *bool `pulumi:"associateAllControls"`
	// (Optional)
	CommonGlobalOverrideActionsConfig map[string]string `pulumi:"commonGlobalOverrideActionsConfig"`
	// (Optional) Types for custom controls
	ControlsInfos []InspectionProfileControlsInfo `pulumi:"controlsInfos"`
	// (Optional) Types for custom controls
	CustomControls []InspectionProfileCustomControl `pulumi:"customControls"`
	// Description of the inspection profile.
	Description          *string  `pulumi:"description"`
	GlobalControlActions []string `pulumi:"globalControlActions"`
	IncarnationNumber    *string  `pulumi:"incarnationNumber"`
	// The name of the inspection profile.
	Name *string `pulumi:"name"`
	// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
	ParanoiaLevel *string `pulumi:"paranoiaLevel"`
	// The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefinedControls` can be set by using the data source `dataSourceZpaPredefinedControls` or by group using the data source `Inspection.getInspectionAllPredefinedControls`.
	PredefinedControls        []InspectionProfilePredefinedControl `pulumi:"predefinedControls"`
	PredefinedControlsVersion *string                              `pulumi:"predefinedControlsVersion"`
	// (string)
	WebSocketControls []InspectionProfileWebSocketControl `pulumi:"webSocketControls"`
}

type InspectionProfileState struct {
	AssociateAllControls pulumi.BoolPtrInput
	// (Optional)
	CommonGlobalOverrideActionsConfig pulumi.StringMapInput
	// (Optional) Types for custom controls
	ControlsInfos InspectionProfileControlsInfoArrayInput
	// (Optional) Types for custom controls
	CustomControls InspectionProfileCustomControlArrayInput
	// Description of the inspection profile.
	Description          pulumi.StringPtrInput
	GlobalControlActions pulumi.StringArrayInput
	IncarnationNumber    pulumi.StringPtrInput
	// The name of the inspection profile.
	Name pulumi.StringPtrInput
	// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
	ParanoiaLevel pulumi.StringPtrInput
	// The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefinedControls` can be set by using the data source `dataSourceZpaPredefinedControls` or by group using the data source `Inspection.getInspectionAllPredefinedControls`.
	PredefinedControls        InspectionProfilePredefinedControlArrayInput
	PredefinedControlsVersion pulumi.StringPtrInput
	// (string)
	WebSocketControls InspectionProfileWebSocketControlArrayInput
}

func (InspectionProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*inspectionProfileState)(nil)).Elem()
}

type inspectionProfileArgs struct {
	AssociateAllControls *bool `pulumi:"associateAllControls"`
	// (Optional)
	CommonGlobalOverrideActionsConfig map[string]string `pulumi:"commonGlobalOverrideActionsConfig"`
	// (Optional) Types for custom controls
	ControlsInfos []InspectionProfileControlsInfo `pulumi:"controlsInfos"`
	// (Optional) Types for custom controls
	CustomControls []InspectionProfileCustomControl `pulumi:"customControls"`
	// Description of the inspection profile.
	Description          *string  `pulumi:"description"`
	GlobalControlActions []string `pulumi:"globalControlActions"`
	IncarnationNumber    *string  `pulumi:"incarnationNumber"`
	// The name of the inspection profile.
	Name *string `pulumi:"name"`
	// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
	ParanoiaLevel *string `pulumi:"paranoiaLevel"`
	// The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefinedControls` can be set by using the data source `dataSourceZpaPredefinedControls` or by group using the data source `Inspection.getInspectionAllPredefinedControls`.
	PredefinedControls        []InspectionProfilePredefinedControl `pulumi:"predefinedControls"`
	PredefinedControlsVersion *string                              `pulumi:"predefinedControlsVersion"`
	// (string)
	WebSocketControls []InspectionProfileWebSocketControl `pulumi:"webSocketControls"`
}

// The set of arguments for constructing a InspectionProfile resource.
type InspectionProfileArgs struct {
	AssociateAllControls pulumi.BoolPtrInput
	// (Optional)
	CommonGlobalOverrideActionsConfig pulumi.StringMapInput
	// (Optional) Types for custom controls
	ControlsInfos InspectionProfileControlsInfoArrayInput
	// (Optional) Types for custom controls
	CustomControls InspectionProfileCustomControlArrayInput
	// Description of the inspection profile.
	Description          pulumi.StringPtrInput
	GlobalControlActions pulumi.StringArrayInput
	IncarnationNumber    pulumi.StringPtrInput
	// The name of the inspection profile.
	Name pulumi.StringPtrInput
	// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
	ParanoiaLevel pulumi.StringPtrInput
	// The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefinedControls` can be set by using the data source `dataSourceZpaPredefinedControls` or by group using the data source `Inspection.getInspectionAllPredefinedControls`.
	PredefinedControls        InspectionProfilePredefinedControlArrayInput
	PredefinedControlsVersion pulumi.StringPtrInput
	// (string)
	WebSocketControls InspectionProfileWebSocketControlArrayInput
}

func (InspectionProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inspectionProfileArgs)(nil)).Elem()
}

type InspectionProfileInput interface {
	pulumi.Input

	ToInspectionProfileOutput() InspectionProfileOutput
	ToInspectionProfileOutputWithContext(ctx context.Context) InspectionProfileOutput
}

func (*InspectionProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**InspectionProfile)(nil)).Elem()
}

func (i *InspectionProfile) ToInspectionProfileOutput() InspectionProfileOutput {
	return i.ToInspectionProfileOutputWithContext(context.Background())
}

func (i *InspectionProfile) ToInspectionProfileOutputWithContext(ctx context.Context) InspectionProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InspectionProfileOutput)
}

// InspectionProfileArrayInput is an input type that accepts InspectionProfileArray and InspectionProfileArrayOutput values.
// You can construct a concrete instance of `InspectionProfileArrayInput` via:
//
//	InspectionProfileArray{ InspectionProfileArgs{...} }
type InspectionProfileArrayInput interface {
	pulumi.Input

	ToInspectionProfileArrayOutput() InspectionProfileArrayOutput
	ToInspectionProfileArrayOutputWithContext(context.Context) InspectionProfileArrayOutput
}

type InspectionProfileArray []InspectionProfileInput

func (InspectionProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InspectionProfile)(nil)).Elem()
}

func (i InspectionProfileArray) ToInspectionProfileArrayOutput() InspectionProfileArrayOutput {
	return i.ToInspectionProfileArrayOutputWithContext(context.Background())
}

func (i InspectionProfileArray) ToInspectionProfileArrayOutputWithContext(ctx context.Context) InspectionProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InspectionProfileArrayOutput)
}

// InspectionProfileMapInput is an input type that accepts InspectionProfileMap and InspectionProfileMapOutput values.
// You can construct a concrete instance of `InspectionProfileMapInput` via:
//
//	InspectionProfileMap{ "key": InspectionProfileArgs{...} }
type InspectionProfileMapInput interface {
	pulumi.Input

	ToInspectionProfileMapOutput() InspectionProfileMapOutput
	ToInspectionProfileMapOutputWithContext(context.Context) InspectionProfileMapOutput
}

type InspectionProfileMap map[string]InspectionProfileInput

func (InspectionProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InspectionProfile)(nil)).Elem()
}

func (i InspectionProfileMap) ToInspectionProfileMapOutput() InspectionProfileMapOutput {
	return i.ToInspectionProfileMapOutputWithContext(context.Background())
}

func (i InspectionProfileMap) ToInspectionProfileMapOutputWithContext(ctx context.Context) InspectionProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InspectionProfileMapOutput)
}

type InspectionProfileOutput struct{ *pulumi.OutputState }

func (InspectionProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InspectionProfile)(nil)).Elem()
}

func (o InspectionProfileOutput) ToInspectionProfileOutput() InspectionProfileOutput {
	return o
}

func (o InspectionProfileOutput) ToInspectionProfileOutputWithContext(ctx context.Context) InspectionProfileOutput {
	return o
}

func (o InspectionProfileOutput) AssociateAllControls() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InspectionProfile) pulumi.BoolPtrOutput { return v.AssociateAllControls }).(pulumi.BoolPtrOutput)
}

// (Optional)
func (o InspectionProfileOutput) CommonGlobalOverrideActionsConfig() pulumi.StringMapOutput {
	return o.ApplyT(func(v *InspectionProfile) pulumi.StringMapOutput { return v.CommonGlobalOverrideActionsConfig }).(pulumi.StringMapOutput)
}

// (Optional) Types for custom controls
func (o InspectionProfileOutput) ControlsInfos() InspectionProfileControlsInfoArrayOutput {
	return o.ApplyT(func(v *InspectionProfile) InspectionProfileControlsInfoArrayOutput { return v.ControlsInfos }).(InspectionProfileControlsInfoArrayOutput)
}

// (Optional) Types for custom controls
func (o InspectionProfileOutput) CustomControls() InspectionProfileCustomControlArrayOutput {
	return o.ApplyT(func(v *InspectionProfile) InspectionProfileCustomControlArrayOutput { return v.CustomControls }).(InspectionProfileCustomControlArrayOutput)
}

// Description of the inspection profile.
func (o InspectionProfileOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *InspectionProfile) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o InspectionProfileOutput) GlobalControlActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InspectionProfile) pulumi.StringArrayOutput { return v.GlobalControlActions }).(pulumi.StringArrayOutput)
}

func (o InspectionProfileOutput) IncarnationNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *InspectionProfile) pulumi.StringOutput { return v.IncarnationNumber }).(pulumi.StringOutput)
}

// The name of the inspection profile.
func (o InspectionProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InspectionProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
func (o InspectionProfileOutput) ParanoiaLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *InspectionProfile) pulumi.StringOutput { return v.ParanoiaLevel }).(pulumi.StringOutput)
}

// The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefinedControls` can be set by using the data source `dataSourceZpaPredefinedControls` or by group using the data source `Inspection.getInspectionAllPredefinedControls`.
func (o InspectionProfileOutput) PredefinedControls() InspectionProfilePredefinedControlArrayOutput {
	return o.ApplyT(func(v *InspectionProfile) InspectionProfilePredefinedControlArrayOutput { return v.PredefinedControls }).(InspectionProfilePredefinedControlArrayOutput)
}

func (o InspectionProfileOutput) PredefinedControlsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *InspectionProfile) pulumi.StringOutput { return v.PredefinedControlsVersion }).(pulumi.StringOutput)
}

// (string)
func (o InspectionProfileOutput) WebSocketControls() InspectionProfileWebSocketControlArrayOutput {
	return o.ApplyT(func(v *InspectionProfile) InspectionProfileWebSocketControlArrayOutput { return v.WebSocketControls }).(InspectionProfileWebSocketControlArrayOutput)
}

type InspectionProfileArrayOutput struct{ *pulumi.OutputState }

func (InspectionProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InspectionProfile)(nil)).Elem()
}

func (o InspectionProfileArrayOutput) ToInspectionProfileArrayOutput() InspectionProfileArrayOutput {
	return o
}

func (o InspectionProfileArrayOutput) ToInspectionProfileArrayOutputWithContext(ctx context.Context) InspectionProfileArrayOutput {
	return o
}

func (o InspectionProfileArrayOutput) Index(i pulumi.IntInput) InspectionProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InspectionProfile {
		return vs[0].([]*InspectionProfile)[vs[1].(int)]
	}).(InspectionProfileOutput)
}

type InspectionProfileMapOutput struct{ *pulumi.OutputState }

func (InspectionProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InspectionProfile)(nil)).Elem()
}

func (o InspectionProfileMapOutput) ToInspectionProfileMapOutput() InspectionProfileMapOutput {
	return o
}

func (o InspectionProfileMapOutput) ToInspectionProfileMapOutputWithContext(ctx context.Context) InspectionProfileMapOutput {
	return o
}

func (o InspectionProfileMapOutput) MapIndex(k pulumi.StringInput) InspectionProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InspectionProfile {
		return vs[0].(map[string]*InspectionProfile)[vs[1].(string)]
	}).(InspectionProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InspectionProfileInput)(nil)).Elem(), &InspectionProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*InspectionProfileArrayInput)(nil)).Elem(), InspectionProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InspectionProfileMapInput)(nil)).Elem(), InspectionProfileMap{})
	pulumi.RegisterOutputType(InspectionProfileOutput{})
	pulumi.RegisterOutputType(InspectionProfileArrayOutput{})
	pulumi.RegisterOutputType(InspectionProfileMapOutput{})
}
