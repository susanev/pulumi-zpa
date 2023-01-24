// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Application Segment can be imported by using `<APPLICATION SEGMENT ID>` or `<APPLICATION SEGMENT NAME>` as the import ID.
//
// ```sh
//
//	$ pulumi import zpa:index/zPAApplicationSegment:ZPAApplicationSegment example <application_segment_id>
//
// ```
//
//	or
//
// ```sh
//
//	$ pulumi import zpa:index/zPAApplicationSegment:ZPAApplicationSegment example <application_segment_name>
//
// ```
type ZPAApplicationSegment struct {
	pulumi.CustomResourceState

	// (Optional) Indicates whether users can bypass ZPA to access applications.
	BypassType pulumi.StringOutput `pulumi:"bypassType"`
	// (Optional)
	ConfigSpace        pulumi.StringPtrOutput `pulumi:"configSpace"`
	DefaultIdleTimeout pulumi.StringOutput    `pulumi:"defaultIdleTimeout"`
	// (Optional) Description of the application.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of domains and IPs.
	DomainNames pulumi.StringArrayOutput `pulumi:"domainNames"`
	// (Optional) Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt pulumi.BoolPtrOutput `pulumi:"doubleEncrypt"`
	// (Optional) Whether this application is enabled or not.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// (Optional)
	HealthCheckType pulumi.StringPtrOutput `pulumi:"healthCheckType"`
	// (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
	HealthReporting pulumi.StringPtrOutput `pulumi:"healthReporting"`
	// (Optional)
	IcmpAccessType pulumi.StringPtrOutput `pulumi:"icmpAccessType"`
	// (Optional)
	IpAnchored pulumi.BoolPtrOutput `pulumi:"ipAnchored"`
	// (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
	IsCnameEnabled pulumi.BoolOutput `pulumi:"isCnameEnabled"`
	// Name. The name of the App Connector Group to be exported.
	Name pulumi.StringOutput `pulumi:"name"`
	// (Optional)
	PassiveHealthEnabled pulumi.BoolOutput `pulumi:"passiveHealthEnabled"`
	// List of Segment Group IDs
	SegmentGroupId   pulumi.StringOutput `pulumi:"segmentGroupId"`
	SegmentGroupName pulumi.StringOutput `pulumi:"segmentGroupName"`
	// List of Server Group IDs
	ServerGroups ZPAApplicationSegmentServerGroupArrayOutput `pulumi:"serverGroups"`
	// TCP port ranges used to access the app.
	TcpPortRanges pulumi.StringArrayOutput `pulumi:"tcpPortRanges"`
	// UDP port ranges used to access the app.
	UdpPortRanges pulumi.StringArrayOutput `pulumi:"udpPortRanges"`
}

// NewZPAApplicationSegment registers a new resource with the given unique name, arguments, and options.
func NewZPAApplicationSegment(ctx *pulumi.Context,
	name string, args *ZPAApplicationSegmentArgs, opts ...pulumi.ResourceOption) (*ZPAApplicationSegment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainNames == nil {
		return nil, errors.New("invalid value for required argument 'DomainNames'")
	}
	if args.ServerGroups == nil {
		return nil, errors.New("invalid value for required argument 'ServerGroups'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ZPAApplicationSegment
	err := ctx.RegisterResource("zpa:index/zPAApplicationSegment:ZPAApplicationSegment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZPAApplicationSegment gets an existing ZPAApplicationSegment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZPAApplicationSegment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZPAApplicationSegmentState, opts ...pulumi.ResourceOption) (*ZPAApplicationSegment, error) {
	var resource ZPAApplicationSegment
	err := ctx.ReadResource("zpa:index/zPAApplicationSegment:ZPAApplicationSegment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ZPAApplicationSegment resources.
type zpaapplicationSegmentState struct {
	// (Optional) Indicates whether users can bypass ZPA to access applications.
	BypassType *string `pulumi:"bypassType"`
	// (Optional)
	ConfigSpace        *string `pulumi:"configSpace"`
	DefaultIdleTimeout *string `pulumi:"defaultIdleTimeout"`
	// (Optional) Description of the application.
	Description *string `pulumi:"description"`
	// List of domains and IPs.
	DomainNames []string `pulumi:"domainNames"`
	// (Optional) Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt *bool `pulumi:"doubleEncrypt"`
	// (Optional) Whether this application is enabled or not.
	Enabled *bool `pulumi:"enabled"`
	// (Optional)
	HealthCheckType *string `pulumi:"healthCheckType"`
	// (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
	HealthReporting *string `pulumi:"healthReporting"`
	// (Optional)
	IcmpAccessType *string `pulumi:"icmpAccessType"`
	// (Optional)
	IpAnchored *bool `pulumi:"ipAnchored"`
	// (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
	IsCnameEnabled *bool `pulumi:"isCnameEnabled"`
	// Name. The name of the App Connector Group to be exported.
	Name *string `pulumi:"name"`
	// (Optional)
	PassiveHealthEnabled *bool `pulumi:"passiveHealthEnabled"`
	// List of Segment Group IDs
	SegmentGroupId   *string `pulumi:"segmentGroupId"`
	SegmentGroupName *string `pulumi:"segmentGroupName"`
	// List of Server Group IDs
	ServerGroups []ZPAApplicationSegmentServerGroup `pulumi:"serverGroups"`
	// TCP port ranges used to access the app.
	TcpPortRanges []string `pulumi:"tcpPortRanges"`
	// UDP port ranges used to access the app.
	UdpPortRanges []string `pulumi:"udpPortRanges"`
}

type ZPAApplicationSegmentState struct {
	// (Optional) Indicates whether users can bypass ZPA to access applications.
	BypassType pulumi.StringPtrInput
	// (Optional)
	ConfigSpace        pulumi.StringPtrInput
	DefaultIdleTimeout pulumi.StringPtrInput
	// (Optional) Description of the application.
	Description pulumi.StringPtrInput
	// List of domains and IPs.
	DomainNames pulumi.StringArrayInput
	// (Optional) Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt pulumi.BoolPtrInput
	// (Optional) Whether this application is enabled or not.
	Enabled pulumi.BoolPtrInput
	// (Optional)
	HealthCheckType pulumi.StringPtrInput
	// (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
	HealthReporting pulumi.StringPtrInput
	// (Optional)
	IcmpAccessType pulumi.StringPtrInput
	// (Optional)
	IpAnchored pulumi.BoolPtrInput
	// (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
	IsCnameEnabled pulumi.BoolPtrInput
	// Name. The name of the App Connector Group to be exported.
	Name pulumi.StringPtrInput
	// (Optional)
	PassiveHealthEnabled pulumi.BoolPtrInput
	// List of Segment Group IDs
	SegmentGroupId   pulumi.StringPtrInput
	SegmentGroupName pulumi.StringPtrInput
	// List of Server Group IDs
	ServerGroups ZPAApplicationSegmentServerGroupArrayInput
	// TCP port ranges used to access the app.
	TcpPortRanges pulumi.StringArrayInput
	// UDP port ranges used to access the app.
	UdpPortRanges pulumi.StringArrayInput
}

func (ZPAApplicationSegmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*zpaapplicationSegmentState)(nil)).Elem()
}

type zpaapplicationSegmentArgs struct {
	// (Optional) Indicates whether users can bypass ZPA to access applications.
	BypassType *string `pulumi:"bypassType"`
	// (Optional)
	ConfigSpace        *string `pulumi:"configSpace"`
	DefaultIdleTimeout *string `pulumi:"defaultIdleTimeout"`
	// (Optional) Description of the application.
	Description *string `pulumi:"description"`
	// List of domains and IPs.
	DomainNames []string `pulumi:"domainNames"`
	// (Optional) Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt *bool `pulumi:"doubleEncrypt"`
	// (Optional) Whether this application is enabled or not.
	Enabled *bool `pulumi:"enabled"`
	// (Optional)
	HealthCheckType *string `pulumi:"healthCheckType"`
	// (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
	HealthReporting *string `pulumi:"healthReporting"`
	// (Optional)
	IcmpAccessType *string `pulumi:"icmpAccessType"`
	// (Optional)
	IpAnchored *bool `pulumi:"ipAnchored"`
	// (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
	IsCnameEnabled *bool `pulumi:"isCnameEnabled"`
	// Name. The name of the App Connector Group to be exported.
	Name *string `pulumi:"name"`
	// (Optional)
	PassiveHealthEnabled *bool `pulumi:"passiveHealthEnabled"`
	// List of Segment Group IDs
	SegmentGroupId   *string `pulumi:"segmentGroupId"`
	SegmentGroupName *string `pulumi:"segmentGroupName"`
	// List of Server Group IDs
	ServerGroups []ZPAApplicationSegmentServerGroup `pulumi:"serverGroups"`
	// TCP port ranges used to access the app.
	TcpPortRanges []string `pulumi:"tcpPortRanges"`
	// UDP port ranges used to access the app.
	UdpPortRanges []string `pulumi:"udpPortRanges"`
}

// The set of arguments for constructing a ZPAApplicationSegment resource.
type ZPAApplicationSegmentArgs struct {
	// (Optional) Indicates whether users can bypass ZPA to access applications.
	BypassType pulumi.StringPtrInput
	// (Optional)
	ConfigSpace        pulumi.StringPtrInput
	DefaultIdleTimeout pulumi.StringPtrInput
	// (Optional) Description of the application.
	Description pulumi.StringPtrInput
	// List of domains and IPs.
	DomainNames pulumi.StringArrayInput
	// (Optional) Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt pulumi.BoolPtrInput
	// (Optional) Whether this application is enabled or not.
	Enabled pulumi.BoolPtrInput
	// (Optional)
	HealthCheckType pulumi.StringPtrInput
	// (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
	HealthReporting pulumi.StringPtrInput
	// (Optional)
	IcmpAccessType pulumi.StringPtrInput
	// (Optional)
	IpAnchored pulumi.BoolPtrInput
	// (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
	IsCnameEnabled pulumi.BoolPtrInput
	// Name. The name of the App Connector Group to be exported.
	Name pulumi.StringPtrInput
	// (Optional)
	PassiveHealthEnabled pulumi.BoolPtrInput
	// List of Segment Group IDs
	SegmentGroupId   pulumi.StringPtrInput
	SegmentGroupName pulumi.StringPtrInput
	// List of Server Group IDs
	ServerGroups ZPAApplicationSegmentServerGroupArrayInput
	// TCP port ranges used to access the app.
	TcpPortRanges pulumi.StringArrayInput
	// UDP port ranges used to access the app.
	UdpPortRanges pulumi.StringArrayInput
}

func (ZPAApplicationSegmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zpaapplicationSegmentArgs)(nil)).Elem()
}

type ZPAApplicationSegmentInput interface {
	pulumi.Input

	ToZPAApplicationSegmentOutput() ZPAApplicationSegmentOutput
	ToZPAApplicationSegmentOutputWithContext(ctx context.Context) ZPAApplicationSegmentOutput
}

func (*ZPAApplicationSegment) ElementType() reflect.Type {
	return reflect.TypeOf((**ZPAApplicationSegment)(nil)).Elem()
}

func (i *ZPAApplicationSegment) ToZPAApplicationSegmentOutput() ZPAApplicationSegmentOutput {
	return i.ToZPAApplicationSegmentOutputWithContext(context.Background())
}

func (i *ZPAApplicationSegment) ToZPAApplicationSegmentOutputWithContext(ctx context.Context) ZPAApplicationSegmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZPAApplicationSegmentOutput)
}

// ZPAApplicationSegmentArrayInput is an input type that accepts ZPAApplicationSegmentArray and ZPAApplicationSegmentArrayOutput values.
// You can construct a concrete instance of `ZPAApplicationSegmentArrayInput` via:
//
//	ZPAApplicationSegmentArray{ ZPAApplicationSegmentArgs{...} }
type ZPAApplicationSegmentArrayInput interface {
	pulumi.Input

	ToZPAApplicationSegmentArrayOutput() ZPAApplicationSegmentArrayOutput
	ToZPAApplicationSegmentArrayOutputWithContext(context.Context) ZPAApplicationSegmentArrayOutput
}

type ZPAApplicationSegmentArray []ZPAApplicationSegmentInput

func (ZPAApplicationSegmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ZPAApplicationSegment)(nil)).Elem()
}

func (i ZPAApplicationSegmentArray) ToZPAApplicationSegmentArrayOutput() ZPAApplicationSegmentArrayOutput {
	return i.ToZPAApplicationSegmentArrayOutputWithContext(context.Background())
}

func (i ZPAApplicationSegmentArray) ToZPAApplicationSegmentArrayOutputWithContext(ctx context.Context) ZPAApplicationSegmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZPAApplicationSegmentArrayOutput)
}

// ZPAApplicationSegmentMapInput is an input type that accepts ZPAApplicationSegmentMap and ZPAApplicationSegmentMapOutput values.
// You can construct a concrete instance of `ZPAApplicationSegmentMapInput` via:
//
//	ZPAApplicationSegmentMap{ "key": ZPAApplicationSegmentArgs{...} }
type ZPAApplicationSegmentMapInput interface {
	pulumi.Input

	ToZPAApplicationSegmentMapOutput() ZPAApplicationSegmentMapOutput
	ToZPAApplicationSegmentMapOutputWithContext(context.Context) ZPAApplicationSegmentMapOutput
}

type ZPAApplicationSegmentMap map[string]ZPAApplicationSegmentInput

func (ZPAApplicationSegmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ZPAApplicationSegment)(nil)).Elem()
}

func (i ZPAApplicationSegmentMap) ToZPAApplicationSegmentMapOutput() ZPAApplicationSegmentMapOutput {
	return i.ToZPAApplicationSegmentMapOutputWithContext(context.Background())
}

func (i ZPAApplicationSegmentMap) ToZPAApplicationSegmentMapOutputWithContext(ctx context.Context) ZPAApplicationSegmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZPAApplicationSegmentMapOutput)
}

type ZPAApplicationSegmentOutput struct{ *pulumi.OutputState }

func (ZPAApplicationSegmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ZPAApplicationSegment)(nil)).Elem()
}

func (o ZPAApplicationSegmentOutput) ToZPAApplicationSegmentOutput() ZPAApplicationSegmentOutput {
	return o
}

func (o ZPAApplicationSegmentOutput) ToZPAApplicationSegmentOutputWithContext(ctx context.Context) ZPAApplicationSegmentOutput {
	return o
}

// (Optional) Indicates whether users can bypass ZPA to access applications.
func (o ZPAApplicationSegmentOutput) BypassType() pulumi.StringOutput {
	return o.ApplyT(func(v *ZPAApplicationSegment) pulumi.StringOutput { return v.BypassType }).(pulumi.StringOutput)
}

// (Optional)
func (o ZPAApplicationSegmentOutput) ConfigSpace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ZPAApplicationSegment) pulumi.StringPtrOutput { return v.ConfigSpace }).(pulumi.StringPtrOutput)
}

func (o ZPAApplicationSegmentOutput) DefaultIdleTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v *ZPAApplicationSegment) pulumi.StringOutput { return v.DefaultIdleTimeout }).(pulumi.StringOutput)
}

// (Optional) Description of the application.
func (o ZPAApplicationSegmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ZPAApplicationSegment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of domains and IPs.
func (o ZPAApplicationSegmentOutput) DomainNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ZPAApplicationSegment) pulumi.StringArrayOutput { return v.DomainNames }).(pulumi.StringArrayOutput)
}

// (Optional) Whether Double Encryption is enabled or disabled for the app.
func (o ZPAApplicationSegmentOutput) DoubleEncrypt() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ZPAApplicationSegment) pulumi.BoolPtrOutput { return v.DoubleEncrypt }).(pulumi.BoolPtrOutput)
}

// (Optional) Whether this application is enabled or not.
func (o ZPAApplicationSegmentOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ZPAApplicationSegment) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// (Optional)
func (o ZPAApplicationSegmentOutput) HealthCheckType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ZPAApplicationSegment) pulumi.StringPtrOutput { return v.HealthCheckType }).(pulumi.StringPtrOutput)
}

// (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
func (o ZPAApplicationSegmentOutput) HealthReporting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ZPAApplicationSegment) pulumi.StringPtrOutput { return v.HealthReporting }).(pulumi.StringPtrOutput)
}

// (Optional)
func (o ZPAApplicationSegmentOutput) IcmpAccessType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ZPAApplicationSegment) pulumi.StringPtrOutput { return v.IcmpAccessType }).(pulumi.StringPtrOutput)
}

// (Optional)
func (o ZPAApplicationSegmentOutput) IpAnchored() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ZPAApplicationSegment) pulumi.BoolPtrOutput { return v.IpAnchored }).(pulumi.BoolPtrOutput)
}

// (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
func (o ZPAApplicationSegmentOutput) IsCnameEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ZPAApplicationSegment) pulumi.BoolOutput { return v.IsCnameEnabled }).(pulumi.BoolOutput)
}

// Name. The name of the App Connector Group to be exported.
func (o ZPAApplicationSegmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ZPAApplicationSegment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// (Optional)
func (o ZPAApplicationSegmentOutput) PassiveHealthEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ZPAApplicationSegment) pulumi.BoolOutput { return v.PassiveHealthEnabled }).(pulumi.BoolOutput)
}

// List of Segment Group IDs
func (o ZPAApplicationSegmentOutput) SegmentGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ZPAApplicationSegment) pulumi.StringOutput { return v.SegmentGroupId }).(pulumi.StringOutput)
}

func (o ZPAApplicationSegmentOutput) SegmentGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *ZPAApplicationSegment) pulumi.StringOutput { return v.SegmentGroupName }).(pulumi.StringOutput)
}

// List of Server Group IDs
func (o ZPAApplicationSegmentOutput) ServerGroups() ZPAApplicationSegmentServerGroupArrayOutput {
	return o.ApplyT(func(v *ZPAApplicationSegment) ZPAApplicationSegmentServerGroupArrayOutput { return v.ServerGroups }).(ZPAApplicationSegmentServerGroupArrayOutput)
}

// TCP port ranges used to access the app.
func (o ZPAApplicationSegmentOutput) TcpPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ZPAApplicationSegment) pulumi.StringArrayOutput { return v.TcpPortRanges }).(pulumi.StringArrayOutput)
}

// UDP port ranges used to access the app.
func (o ZPAApplicationSegmentOutput) UdpPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ZPAApplicationSegment) pulumi.StringArrayOutput { return v.UdpPortRanges }).(pulumi.StringArrayOutput)
}

type ZPAApplicationSegmentArrayOutput struct{ *pulumi.OutputState }

func (ZPAApplicationSegmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ZPAApplicationSegment)(nil)).Elem()
}

func (o ZPAApplicationSegmentArrayOutput) ToZPAApplicationSegmentArrayOutput() ZPAApplicationSegmentArrayOutput {
	return o
}

func (o ZPAApplicationSegmentArrayOutput) ToZPAApplicationSegmentArrayOutputWithContext(ctx context.Context) ZPAApplicationSegmentArrayOutput {
	return o
}

func (o ZPAApplicationSegmentArrayOutput) Index(i pulumi.IntInput) ZPAApplicationSegmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ZPAApplicationSegment {
		return vs[0].([]*ZPAApplicationSegment)[vs[1].(int)]
	}).(ZPAApplicationSegmentOutput)
}

type ZPAApplicationSegmentMapOutput struct{ *pulumi.OutputState }

func (ZPAApplicationSegmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ZPAApplicationSegment)(nil)).Elem()
}

func (o ZPAApplicationSegmentMapOutput) ToZPAApplicationSegmentMapOutput() ZPAApplicationSegmentMapOutput {
	return o
}

func (o ZPAApplicationSegmentMapOutput) ToZPAApplicationSegmentMapOutputWithContext(ctx context.Context) ZPAApplicationSegmentMapOutput {
	return o
}

func (o ZPAApplicationSegmentMapOutput) MapIndex(k pulumi.StringInput) ZPAApplicationSegmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ZPAApplicationSegment {
		return vs[0].(map[string]*ZPAApplicationSegment)[vs[1].(string)]
	}).(ZPAApplicationSegmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ZPAApplicationSegmentInput)(nil)).Elem(), &ZPAApplicationSegment{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZPAApplicationSegmentArrayInput)(nil)).Elem(), ZPAApplicationSegmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZPAApplicationSegmentMapInput)(nil)).Elem(), ZPAApplicationSegmentMap{})
	pulumi.RegisterOutputType(ZPAApplicationSegmentOutput{})
	pulumi.RegisterOutputType(ZPAApplicationSegmentArrayOutput{})
	pulumi.RegisterOutputType(ZPAApplicationSegmentMapOutput{})
}
