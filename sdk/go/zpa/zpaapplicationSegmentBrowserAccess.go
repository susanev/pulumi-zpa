// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/zscaler/pulumi-zpa/sdk/go/zpa"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testCert, err := zpa.GetZPABaCertificate(ctx, &zpa.GetZPABaCertificateArgs{
//				Name: pulumi.StringRef("sales.acme.com"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleZPASegmentGroup, err := zpa.NewZPASegmentGroup(ctx, "exampleZPASegmentGroup", &zpa.ZPASegmentGroupArgs{
//				Description: pulumi.String("Example"),
//				Enabled:     pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			exampleZPAAppConnectorGroup, err := zpa.LookupZPAAppConnectorGroup(ctx, &zpa.LookupZPAAppConnectorGroupArgs{
//				Name: pulumi.StringRef("AWS-Connector"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleZPAServerGroup, err := zpa.NewZPAServerGroup(ctx, "exampleZPAServerGroup", &zpa.ZPAServerGroupArgs{
//				Description:      pulumi.String("Example"),
//				Enabled:          pulumi.Bool(true),
//				DynamicDiscovery: pulumi.Bool(true),
//				AppConnectorGroups: zpa.ZPAServerGroupAppConnectorGroupArray{
//					&zpa.ZPAServerGroupAppConnectorGroupArgs{
//						Ids: pulumi.StringArray{
//							*pulumi.String(exampleZPAAppConnectorGroup.Id),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = zpa.NewZPAApplicationSegmentBrowserAccess(ctx, "browserAccessApps", &zpa.ZPAApplicationSegmentBrowserAccessArgs{
//				Description:     pulumi.String("Browser Access Apps"),
//				Enabled:         pulumi.Bool(true),
//				HealthReporting: pulumi.String("ON_ACCESS"),
//				BypassType:      pulumi.String("NEVER"),
//				TcpPortRanges: pulumi.StringArray{
//					pulumi.String("80"),
//					pulumi.String("80"),
//				},
//				DomainNames: pulumi.StringArray{
//					pulumi.String("sales.acme.com"),
//				},
//				SegmentGroupId: exampleZPASegmentGroup.ID(),
//				ClientlessApps: zpa.ZPAApplicationSegmentBrowserAccessClientlessAppArray{
//					&zpa.ZPAApplicationSegmentBrowserAccessClientlessAppArgs{
//						Name:                pulumi.String("sales.acme.com"),
//						ApplicationProtocol: pulumi.String("HTTP"),
//						ApplicationPort:     pulumi.String("80"),
//						CertificateId:       *pulumi.String(testCert.Id),
//						TrustUntrustedCert:  pulumi.Bool(true),
//						Enabled:             pulumi.Bool(true),
//						Domain:              pulumi.String("sales.acme.com"),
//					},
//				},
//				ServerGroups: zpa.ZPAApplicationSegmentBrowserAccessServerGroupArray{
//					&zpa.ZPAApplicationSegmentBrowserAccessServerGroupArgs{
//						Ids: pulumi.StringArray{
//							exampleZPAServerGroup.ID(),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) **zpa_application_segment_browser_access** Application Segment Browser Access can be imported by using <`BROWSER ACCESS ID`> or `<<BROWSER ACCESS NAME>` as the import ID. For example
//
// ```sh
//
//	$ pulumi import zpa:index/zPAApplicationSegmentBrowserAccess:ZPAApplicationSegmentBrowserAccess example <browser_access_id>.
//
// ```
//
//	or
//
// ```sh
//
//	$ pulumi import zpa:index/zPAApplicationSegmentBrowserAccess:ZPAApplicationSegmentBrowserAccess example <browser_access_name>
//
// ```
type ZPAApplicationSegmentBrowserAccess struct {
	pulumi.CustomResourceState

	// (Optional) Indicates whether users can bypass ZPA to access applications. Default value is: `NEVER` and supported values are: `ALWAYS`, `NEVER` and `ON_NET`. The value `NEVER` indicates the use of the client forwarding policy.
	BypassType     pulumi.StringPtrOutput                                     `pulumi:"bypassType"`
	ClientlessApps ZPAApplicationSegmentBrowserAccessClientlessAppArrayOutput `pulumi:"clientlessApps"`
	// (Optional) Default: `DEFAULT`. Supported values: `DEFAULT`, `SIEM`
	ConfigSpace pulumi.StringPtrOutput `pulumi:"configSpace"`
	// (Optional) Description of the application.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of domains and IPs.
	DomainNames pulumi.StringArrayOutput `pulumi:"domainNames"`
	// (Optional) Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt pulumi.BoolPtrOutput `pulumi:"doubleEncrypt"`
	// (Optional) - Whether this app is enabled or not.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// (Optional) Default: `DEFAULT`. Supported values: `DEFAULT`, `NONE`
	HealthCheckType pulumi.StringOutput `pulumi:"healthCheckType"`
	// (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
	HealthReporting pulumi.StringOutput    `pulumi:"healthReporting"`
	IcmpAccessType  pulumi.StringPtrOutput `pulumi:"icmpAccessType"`
	// (Optional) - If Source IP Anchoring for use with ZIA, is enabled or disabled for the app. Supported values are `true` and `false`
	IpAnchored pulumi.BoolPtrOutput `pulumi:"ipAnchored"`
	// (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
	IsCnameEnabled pulumi.BoolPtrOutput `pulumi:"isCnameEnabled"`
	// Name of the application.
	Name                 pulumi.StringOutput `pulumi:"name"`
	PassiveHealthEnabled pulumi.BoolOutput   `pulumi:"passiveHealthEnabled"`
	// List of Segment Group IDs
	SegmentGroupId   pulumi.StringOutput `pulumi:"segmentGroupId"`
	SegmentGroupName pulumi.StringOutput `pulumi:"segmentGroupName"`
	// List of Server Group IDs
	ServerGroups ZPAApplicationSegmentBrowserAccessServerGroupArrayOutput `pulumi:"serverGroups"`
	// TCP port ranges used to access the app.
	TcpPortRanges pulumi.StringArrayOutput `pulumi:"tcpPortRanges"`
	// UDP port ranges used to access the app.
	UdpPortRanges pulumi.StringArrayOutput `pulumi:"udpPortRanges"`
}

// NewZPAApplicationSegmentBrowserAccess registers a new resource with the given unique name, arguments, and options.
func NewZPAApplicationSegmentBrowserAccess(ctx *pulumi.Context,
	name string, args *ZPAApplicationSegmentBrowserAccessArgs, opts ...pulumi.ResourceOption) (*ZPAApplicationSegmentBrowserAccess, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientlessApps == nil {
		return nil, errors.New("invalid value for required argument 'ClientlessApps'")
	}
	if args.DomainNames == nil {
		return nil, errors.New("invalid value for required argument 'DomainNames'")
	}
	if args.SegmentGroupId == nil {
		return nil, errors.New("invalid value for required argument 'SegmentGroupId'")
	}
	if args.ServerGroups == nil {
		return nil, errors.New("invalid value for required argument 'ServerGroups'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ZPAApplicationSegmentBrowserAccess
	err := ctx.RegisterResource("zpa:index/zPAApplicationSegmentBrowserAccess:ZPAApplicationSegmentBrowserAccess", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZPAApplicationSegmentBrowserAccess gets an existing ZPAApplicationSegmentBrowserAccess resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZPAApplicationSegmentBrowserAccess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZPAApplicationSegmentBrowserAccessState, opts ...pulumi.ResourceOption) (*ZPAApplicationSegmentBrowserAccess, error) {
	var resource ZPAApplicationSegmentBrowserAccess
	err := ctx.ReadResource("zpa:index/zPAApplicationSegmentBrowserAccess:ZPAApplicationSegmentBrowserAccess", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ZPAApplicationSegmentBrowserAccess resources.
type zpaapplicationSegmentBrowserAccessState struct {
	// (Optional) Indicates whether users can bypass ZPA to access applications. Default value is: `NEVER` and supported values are: `ALWAYS`, `NEVER` and `ON_NET`. The value `NEVER` indicates the use of the client forwarding policy.
	BypassType     *string                                           `pulumi:"bypassType"`
	ClientlessApps []ZPAApplicationSegmentBrowserAccessClientlessApp `pulumi:"clientlessApps"`
	// (Optional) Default: `DEFAULT`. Supported values: `DEFAULT`, `SIEM`
	ConfigSpace *string `pulumi:"configSpace"`
	// (Optional) Description of the application.
	Description *string `pulumi:"description"`
	// List of domains and IPs.
	DomainNames []string `pulumi:"domainNames"`
	// (Optional) Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt *bool `pulumi:"doubleEncrypt"`
	// (Optional) - Whether this app is enabled or not.
	Enabled *bool `pulumi:"enabled"`
	// (Optional) Default: `DEFAULT`. Supported values: `DEFAULT`, `NONE`
	HealthCheckType *string `pulumi:"healthCheckType"`
	// (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
	HealthReporting *string `pulumi:"healthReporting"`
	IcmpAccessType  *string `pulumi:"icmpAccessType"`
	// (Optional) - If Source IP Anchoring for use with ZIA, is enabled or disabled for the app. Supported values are `true` and `false`
	IpAnchored *bool `pulumi:"ipAnchored"`
	// (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
	IsCnameEnabled *bool `pulumi:"isCnameEnabled"`
	// Name of the application.
	Name                 *string `pulumi:"name"`
	PassiveHealthEnabled *bool   `pulumi:"passiveHealthEnabled"`
	// List of Segment Group IDs
	SegmentGroupId   *string `pulumi:"segmentGroupId"`
	SegmentGroupName *string `pulumi:"segmentGroupName"`
	// List of Server Group IDs
	ServerGroups []ZPAApplicationSegmentBrowserAccessServerGroup `pulumi:"serverGroups"`
	// TCP port ranges used to access the app.
	TcpPortRanges []string `pulumi:"tcpPortRanges"`
	// UDP port ranges used to access the app.
	UdpPortRanges []string `pulumi:"udpPortRanges"`
}

type ZPAApplicationSegmentBrowserAccessState struct {
	// (Optional) Indicates whether users can bypass ZPA to access applications. Default value is: `NEVER` and supported values are: `ALWAYS`, `NEVER` and `ON_NET`. The value `NEVER` indicates the use of the client forwarding policy.
	BypassType     pulumi.StringPtrInput
	ClientlessApps ZPAApplicationSegmentBrowserAccessClientlessAppArrayInput
	// (Optional) Default: `DEFAULT`. Supported values: `DEFAULT`, `SIEM`
	ConfigSpace pulumi.StringPtrInput
	// (Optional) Description of the application.
	Description pulumi.StringPtrInput
	// List of domains and IPs.
	DomainNames pulumi.StringArrayInput
	// (Optional) Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt pulumi.BoolPtrInput
	// (Optional) - Whether this app is enabled or not.
	Enabled pulumi.BoolPtrInput
	// (Optional) Default: `DEFAULT`. Supported values: `DEFAULT`, `NONE`
	HealthCheckType pulumi.StringPtrInput
	// (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
	HealthReporting pulumi.StringPtrInput
	IcmpAccessType  pulumi.StringPtrInput
	// (Optional) - If Source IP Anchoring for use with ZIA, is enabled or disabled for the app. Supported values are `true` and `false`
	IpAnchored pulumi.BoolPtrInput
	// (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
	IsCnameEnabled pulumi.BoolPtrInput
	// Name of the application.
	Name                 pulumi.StringPtrInput
	PassiveHealthEnabled pulumi.BoolPtrInput
	// List of Segment Group IDs
	SegmentGroupId   pulumi.StringPtrInput
	SegmentGroupName pulumi.StringPtrInput
	// List of Server Group IDs
	ServerGroups ZPAApplicationSegmentBrowserAccessServerGroupArrayInput
	// TCP port ranges used to access the app.
	TcpPortRanges pulumi.StringArrayInput
	// UDP port ranges used to access the app.
	UdpPortRanges pulumi.StringArrayInput
}

func (ZPAApplicationSegmentBrowserAccessState) ElementType() reflect.Type {
	return reflect.TypeOf((*zpaapplicationSegmentBrowserAccessState)(nil)).Elem()
}

type zpaapplicationSegmentBrowserAccessArgs struct {
	// (Optional) Indicates whether users can bypass ZPA to access applications. Default value is: `NEVER` and supported values are: `ALWAYS`, `NEVER` and `ON_NET`. The value `NEVER` indicates the use of the client forwarding policy.
	BypassType     *string                                           `pulumi:"bypassType"`
	ClientlessApps []ZPAApplicationSegmentBrowserAccessClientlessApp `pulumi:"clientlessApps"`
	// (Optional) Default: `DEFAULT`. Supported values: `DEFAULT`, `SIEM`
	ConfigSpace *string `pulumi:"configSpace"`
	// (Optional) Description of the application.
	Description *string `pulumi:"description"`
	// List of domains and IPs.
	DomainNames []string `pulumi:"domainNames"`
	// (Optional) Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt *bool `pulumi:"doubleEncrypt"`
	// (Optional) - Whether this app is enabled or not.
	Enabled *bool `pulumi:"enabled"`
	// (Optional) Default: `DEFAULT`. Supported values: `DEFAULT`, `NONE`
	HealthCheckType *string `pulumi:"healthCheckType"`
	// (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
	HealthReporting *string `pulumi:"healthReporting"`
	IcmpAccessType  *string `pulumi:"icmpAccessType"`
	// (Optional) - If Source IP Anchoring for use with ZIA, is enabled or disabled for the app. Supported values are `true` and `false`
	IpAnchored *bool `pulumi:"ipAnchored"`
	// (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
	IsCnameEnabled *bool `pulumi:"isCnameEnabled"`
	// Name of the application.
	Name                 *string `pulumi:"name"`
	PassiveHealthEnabled *bool   `pulumi:"passiveHealthEnabled"`
	// List of Segment Group IDs
	SegmentGroupId   string  `pulumi:"segmentGroupId"`
	SegmentGroupName *string `pulumi:"segmentGroupName"`
	// List of Server Group IDs
	ServerGroups []ZPAApplicationSegmentBrowserAccessServerGroup `pulumi:"serverGroups"`
	// TCP port ranges used to access the app.
	TcpPortRanges []string `pulumi:"tcpPortRanges"`
	// UDP port ranges used to access the app.
	UdpPortRanges []string `pulumi:"udpPortRanges"`
}

// The set of arguments for constructing a ZPAApplicationSegmentBrowserAccess resource.
type ZPAApplicationSegmentBrowserAccessArgs struct {
	// (Optional) Indicates whether users can bypass ZPA to access applications. Default value is: `NEVER` and supported values are: `ALWAYS`, `NEVER` and `ON_NET`. The value `NEVER` indicates the use of the client forwarding policy.
	BypassType     pulumi.StringPtrInput
	ClientlessApps ZPAApplicationSegmentBrowserAccessClientlessAppArrayInput
	// (Optional) Default: `DEFAULT`. Supported values: `DEFAULT`, `SIEM`
	ConfigSpace pulumi.StringPtrInput
	// (Optional) Description of the application.
	Description pulumi.StringPtrInput
	// List of domains and IPs.
	DomainNames pulumi.StringArrayInput
	// (Optional) Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt pulumi.BoolPtrInput
	// (Optional) - Whether this app is enabled or not.
	Enabled pulumi.BoolPtrInput
	// (Optional) Default: `DEFAULT`. Supported values: `DEFAULT`, `NONE`
	HealthCheckType pulumi.StringPtrInput
	// (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
	HealthReporting pulumi.StringPtrInput
	IcmpAccessType  pulumi.StringPtrInput
	// (Optional) - If Source IP Anchoring for use with ZIA, is enabled or disabled for the app. Supported values are `true` and `false`
	IpAnchored pulumi.BoolPtrInput
	// (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
	IsCnameEnabled pulumi.BoolPtrInput
	// Name of the application.
	Name                 pulumi.StringPtrInput
	PassiveHealthEnabled pulumi.BoolPtrInput
	// List of Segment Group IDs
	SegmentGroupId   pulumi.StringInput
	SegmentGroupName pulumi.StringPtrInput
	// List of Server Group IDs
	ServerGroups ZPAApplicationSegmentBrowserAccessServerGroupArrayInput
	// TCP port ranges used to access the app.
	TcpPortRanges pulumi.StringArrayInput
	// UDP port ranges used to access the app.
	UdpPortRanges pulumi.StringArrayInput
}

func (ZPAApplicationSegmentBrowserAccessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zpaapplicationSegmentBrowserAccessArgs)(nil)).Elem()
}

type ZPAApplicationSegmentBrowserAccessInput interface {
	pulumi.Input

	ToZPAApplicationSegmentBrowserAccessOutput() ZPAApplicationSegmentBrowserAccessOutput
	ToZPAApplicationSegmentBrowserAccessOutputWithContext(ctx context.Context) ZPAApplicationSegmentBrowserAccessOutput
}

func (*ZPAApplicationSegmentBrowserAccess) ElementType() reflect.Type {
	return reflect.TypeOf((**ZPAApplicationSegmentBrowserAccess)(nil)).Elem()
}

func (i *ZPAApplicationSegmentBrowserAccess) ToZPAApplicationSegmentBrowserAccessOutput() ZPAApplicationSegmentBrowserAccessOutput {
	return i.ToZPAApplicationSegmentBrowserAccessOutputWithContext(context.Background())
}

func (i *ZPAApplicationSegmentBrowserAccess) ToZPAApplicationSegmentBrowserAccessOutputWithContext(ctx context.Context) ZPAApplicationSegmentBrowserAccessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZPAApplicationSegmentBrowserAccessOutput)
}

// ZPAApplicationSegmentBrowserAccessArrayInput is an input type that accepts ZPAApplicationSegmentBrowserAccessArray and ZPAApplicationSegmentBrowserAccessArrayOutput values.
// You can construct a concrete instance of `ZPAApplicationSegmentBrowserAccessArrayInput` via:
//
//	ZPAApplicationSegmentBrowserAccessArray{ ZPAApplicationSegmentBrowserAccessArgs{...} }
type ZPAApplicationSegmentBrowserAccessArrayInput interface {
	pulumi.Input

	ToZPAApplicationSegmentBrowserAccessArrayOutput() ZPAApplicationSegmentBrowserAccessArrayOutput
	ToZPAApplicationSegmentBrowserAccessArrayOutputWithContext(context.Context) ZPAApplicationSegmentBrowserAccessArrayOutput
}

type ZPAApplicationSegmentBrowserAccessArray []ZPAApplicationSegmentBrowserAccessInput

func (ZPAApplicationSegmentBrowserAccessArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ZPAApplicationSegmentBrowserAccess)(nil)).Elem()
}

func (i ZPAApplicationSegmentBrowserAccessArray) ToZPAApplicationSegmentBrowserAccessArrayOutput() ZPAApplicationSegmentBrowserAccessArrayOutput {
	return i.ToZPAApplicationSegmentBrowserAccessArrayOutputWithContext(context.Background())
}

func (i ZPAApplicationSegmentBrowserAccessArray) ToZPAApplicationSegmentBrowserAccessArrayOutputWithContext(ctx context.Context) ZPAApplicationSegmentBrowserAccessArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZPAApplicationSegmentBrowserAccessArrayOutput)
}

// ZPAApplicationSegmentBrowserAccessMapInput is an input type that accepts ZPAApplicationSegmentBrowserAccessMap and ZPAApplicationSegmentBrowserAccessMapOutput values.
// You can construct a concrete instance of `ZPAApplicationSegmentBrowserAccessMapInput` via:
//
//	ZPAApplicationSegmentBrowserAccessMap{ "key": ZPAApplicationSegmentBrowserAccessArgs{...} }
type ZPAApplicationSegmentBrowserAccessMapInput interface {
	pulumi.Input

	ToZPAApplicationSegmentBrowserAccessMapOutput() ZPAApplicationSegmentBrowserAccessMapOutput
	ToZPAApplicationSegmentBrowserAccessMapOutputWithContext(context.Context) ZPAApplicationSegmentBrowserAccessMapOutput
}

type ZPAApplicationSegmentBrowserAccessMap map[string]ZPAApplicationSegmentBrowserAccessInput

func (ZPAApplicationSegmentBrowserAccessMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ZPAApplicationSegmentBrowserAccess)(nil)).Elem()
}

func (i ZPAApplicationSegmentBrowserAccessMap) ToZPAApplicationSegmentBrowserAccessMapOutput() ZPAApplicationSegmentBrowserAccessMapOutput {
	return i.ToZPAApplicationSegmentBrowserAccessMapOutputWithContext(context.Background())
}

func (i ZPAApplicationSegmentBrowserAccessMap) ToZPAApplicationSegmentBrowserAccessMapOutputWithContext(ctx context.Context) ZPAApplicationSegmentBrowserAccessMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZPAApplicationSegmentBrowserAccessMapOutput)
}

type ZPAApplicationSegmentBrowserAccessOutput struct{ *pulumi.OutputState }

func (ZPAApplicationSegmentBrowserAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ZPAApplicationSegmentBrowserAccess)(nil)).Elem()
}

func (o ZPAApplicationSegmentBrowserAccessOutput) ToZPAApplicationSegmentBrowserAccessOutput() ZPAApplicationSegmentBrowserAccessOutput {
	return o
}

func (o ZPAApplicationSegmentBrowserAccessOutput) ToZPAApplicationSegmentBrowserAccessOutputWithContext(ctx context.Context) ZPAApplicationSegmentBrowserAccessOutput {
	return o
}

// (Optional) Indicates whether users can bypass ZPA to access applications. Default value is: `NEVER` and supported values are: `ALWAYS`, `NEVER` and `ON_NET`. The value `NEVER` indicates the use of the client forwarding policy.
func (o ZPAApplicationSegmentBrowserAccessOutput) BypassType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ZPAApplicationSegmentBrowserAccess) pulumi.StringPtrOutput { return v.BypassType }).(pulumi.StringPtrOutput)
}

func (o ZPAApplicationSegmentBrowserAccessOutput) ClientlessApps() ZPAApplicationSegmentBrowserAccessClientlessAppArrayOutput {
	return o.ApplyT(func(v *ZPAApplicationSegmentBrowserAccess) ZPAApplicationSegmentBrowserAccessClientlessAppArrayOutput {
		return v.ClientlessApps
	}).(ZPAApplicationSegmentBrowserAccessClientlessAppArrayOutput)
}

// (Optional) Default: `DEFAULT`. Supported values: `DEFAULT`, `SIEM`
func (o ZPAApplicationSegmentBrowserAccessOutput) ConfigSpace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ZPAApplicationSegmentBrowserAccess) pulumi.StringPtrOutput { return v.ConfigSpace }).(pulumi.StringPtrOutput)
}

// (Optional) Description of the application.
func (o ZPAApplicationSegmentBrowserAccessOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ZPAApplicationSegmentBrowserAccess) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of domains and IPs.
func (o ZPAApplicationSegmentBrowserAccessOutput) DomainNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ZPAApplicationSegmentBrowserAccess) pulumi.StringArrayOutput { return v.DomainNames }).(pulumi.StringArrayOutput)
}

// (Optional) Whether Double Encryption is enabled or disabled for the app.
func (o ZPAApplicationSegmentBrowserAccessOutput) DoubleEncrypt() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ZPAApplicationSegmentBrowserAccess) pulumi.BoolPtrOutput { return v.DoubleEncrypt }).(pulumi.BoolPtrOutput)
}

// (Optional) - Whether this app is enabled or not.
func (o ZPAApplicationSegmentBrowserAccessOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ZPAApplicationSegmentBrowserAccess) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// (Optional) Default: `DEFAULT`. Supported values: `DEFAULT`, `NONE`
func (o ZPAApplicationSegmentBrowserAccessOutput) HealthCheckType() pulumi.StringOutput {
	return o.ApplyT(func(v *ZPAApplicationSegmentBrowserAccess) pulumi.StringOutput { return v.HealthCheckType }).(pulumi.StringOutput)
}

// (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
func (o ZPAApplicationSegmentBrowserAccessOutput) HealthReporting() pulumi.StringOutput {
	return o.ApplyT(func(v *ZPAApplicationSegmentBrowserAccess) pulumi.StringOutput { return v.HealthReporting }).(pulumi.StringOutput)
}

func (o ZPAApplicationSegmentBrowserAccessOutput) IcmpAccessType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ZPAApplicationSegmentBrowserAccess) pulumi.StringPtrOutput { return v.IcmpAccessType }).(pulumi.StringPtrOutput)
}

// (Optional) - If Source IP Anchoring for use with ZIA, is enabled or disabled for the app. Supported values are `true` and `false`
func (o ZPAApplicationSegmentBrowserAccessOutput) IpAnchored() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ZPAApplicationSegmentBrowserAccess) pulumi.BoolPtrOutput { return v.IpAnchored }).(pulumi.BoolPtrOutput)
}

// (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
func (o ZPAApplicationSegmentBrowserAccessOutput) IsCnameEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ZPAApplicationSegmentBrowserAccess) pulumi.BoolPtrOutput { return v.IsCnameEnabled }).(pulumi.BoolPtrOutput)
}

// Name of the application.
func (o ZPAApplicationSegmentBrowserAccessOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ZPAApplicationSegmentBrowserAccess) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ZPAApplicationSegmentBrowserAccessOutput) PassiveHealthEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ZPAApplicationSegmentBrowserAccess) pulumi.BoolOutput { return v.PassiveHealthEnabled }).(pulumi.BoolOutput)
}

// List of Segment Group IDs
func (o ZPAApplicationSegmentBrowserAccessOutput) SegmentGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ZPAApplicationSegmentBrowserAccess) pulumi.StringOutput { return v.SegmentGroupId }).(pulumi.StringOutput)
}

func (o ZPAApplicationSegmentBrowserAccessOutput) SegmentGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *ZPAApplicationSegmentBrowserAccess) pulumi.StringOutput { return v.SegmentGroupName }).(pulumi.StringOutput)
}

// List of Server Group IDs
func (o ZPAApplicationSegmentBrowserAccessOutput) ServerGroups() ZPAApplicationSegmentBrowserAccessServerGroupArrayOutput {
	return o.ApplyT(func(v *ZPAApplicationSegmentBrowserAccess) ZPAApplicationSegmentBrowserAccessServerGroupArrayOutput {
		return v.ServerGroups
	}).(ZPAApplicationSegmentBrowserAccessServerGroupArrayOutput)
}

// TCP port ranges used to access the app.
func (o ZPAApplicationSegmentBrowserAccessOutput) TcpPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ZPAApplicationSegmentBrowserAccess) pulumi.StringArrayOutput { return v.TcpPortRanges }).(pulumi.StringArrayOutput)
}

// UDP port ranges used to access the app.
func (o ZPAApplicationSegmentBrowserAccessOutput) UdpPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ZPAApplicationSegmentBrowserAccess) pulumi.StringArrayOutput { return v.UdpPortRanges }).(pulumi.StringArrayOutput)
}

type ZPAApplicationSegmentBrowserAccessArrayOutput struct{ *pulumi.OutputState }

func (ZPAApplicationSegmentBrowserAccessArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ZPAApplicationSegmentBrowserAccess)(nil)).Elem()
}

func (o ZPAApplicationSegmentBrowserAccessArrayOutput) ToZPAApplicationSegmentBrowserAccessArrayOutput() ZPAApplicationSegmentBrowserAccessArrayOutput {
	return o
}

func (o ZPAApplicationSegmentBrowserAccessArrayOutput) ToZPAApplicationSegmentBrowserAccessArrayOutputWithContext(ctx context.Context) ZPAApplicationSegmentBrowserAccessArrayOutput {
	return o
}

func (o ZPAApplicationSegmentBrowserAccessArrayOutput) Index(i pulumi.IntInput) ZPAApplicationSegmentBrowserAccessOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ZPAApplicationSegmentBrowserAccess {
		return vs[0].([]*ZPAApplicationSegmentBrowserAccess)[vs[1].(int)]
	}).(ZPAApplicationSegmentBrowserAccessOutput)
}

type ZPAApplicationSegmentBrowserAccessMapOutput struct{ *pulumi.OutputState }

func (ZPAApplicationSegmentBrowserAccessMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ZPAApplicationSegmentBrowserAccess)(nil)).Elem()
}

func (o ZPAApplicationSegmentBrowserAccessMapOutput) ToZPAApplicationSegmentBrowserAccessMapOutput() ZPAApplicationSegmentBrowserAccessMapOutput {
	return o
}

func (o ZPAApplicationSegmentBrowserAccessMapOutput) ToZPAApplicationSegmentBrowserAccessMapOutputWithContext(ctx context.Context) ZPAApplicationSegmentBrowserAccessMapOutput {
	return o
}

func (o ZPAApplicationSegmentBrowserAccessMapOutput) MapIndex(k pulumi.StringInput) ZPAApplicationSegmentBrowserAccessOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ZPAApplicationSegmentBrowserAccess {
		return vs[0].(map[string]*ZPAApplicationSegmentBrowserAccess)[vs[1].(string)]
	}).(ZPAApplicationSegmentBrowserAccessOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ZPAApplicationSegmentBrowserAccessInput)(nil)).Elem(), &ZPAApplicationSegmentBrowserAccess{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZPAApplicationSegmentBrowserAccessArrayInput)(nil)).Elem(), ZPAApplicationSegmentBrowserAccessArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZPAApplicationSegmentBrowserAccessMapInput)(nil)).Elem(), ZPAApplicationSegmentBrowserAccessMap{})
	pulumi.RegisterOutputType(ZPAApplicationSegmentBrowserAccessOutput{})
	pulumi.RegisterOutputType(ZPAApplicationSegmentBrowserAccessArrayOutput{})
	pulumi.RegisterOutputType(ZPAApplicationSegmentBrowserAccessMapOutput{})
}
