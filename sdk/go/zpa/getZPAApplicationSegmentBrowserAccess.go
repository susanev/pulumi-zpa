// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

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
//			_, err := zpa.LookupZPAApplicationSegmentBrowserAccess(ctx, &zpa.LookupZPAApplicationSegmentBrowserAccessArgs{
//				Name: pulumi.StringRef("example"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
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
//			_, err := zpa.LookupZPAApplicationSegmentBrowserAccess(ctx, &zpa.LookupZPAApplicationSegmentBrowserAccessArgs{
//				Id: pulumi.StringRef("123456789"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupZPAApplicationSegmentBrowserAccess(ctx *pulumi.Context, args *LookupZPAApplicationSegmentBrowserAccessArgs, opts ...pulumi.InvokeOption) (*LookupZPAApplicationSegmentBrowserAccessResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupZPAApplicationSegmentBrowserAccessResult
	err := ctx.Invoke("zpa:index/getZPAApplicationSegmentBrowserAccess:getZPAApplicationSegmentBrowserAccess", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getZPAApplicationSegmentBrowserAccess.
type LookupZPAApplicationSegmentBrowserAccessArgs struct {
	// This field defines the id of the application server.
	Id *string `pulumi:"id"`
	// This field defines the name of the server.
	Name *string `pulumi:"name"`
	// (string) TCP port ranges used to access the app.
	// * `from:`
	// * `to:`
	TcpPortRanges []GetZPAApplicationSegmentBrowserAccessTcpPortRange `pulumi:"tcpPortRanges"`
	// (string) UDP port ranges used to access the app.
	// * `from:`
	// * `to:`
	UdpPortRanges []GetZPAApplicationSegmentBrowserAccessUdpPortRange `pulumi:"udpPortRanges"`
}

// A collection of values returned by getZPAApplicationSegmentBrowserAccess.
type LookupZPAApplicationSegmentBrowserAccessResult struct {
	// (string) Indicates whether users can bypass ZPA to access applications. Default: `NEVER`. Supported values: `ALWAYS`, `NEVER`, `ON_NET`. The value `NEVER` indicates the use of the client forwarding policy.
	BypassType     string                                               `pulumi:"bypassType"`
	ClientlessApps []GetZPAApplicationSegmentBrowserAccessClientlessApp `pulumi:"clientlessApps"`
	// (string)
	ConfigSpace string `pulumi:"configSpace"`
	// (string)
	Description string `pulumi:"description"`
	// List of domains and IPs.
	DomainNames []string `pulumi:"domainNames"`
	// (string) Whether Double Encryption is enabled or disabled for the app. Default: false. Boolean: `true`, `false`.
	DoubleEncrypt bool `pulumi:"doubleEncrypt"`
	// (bool)
	Enabled         bool   `pulumi:"enabled"`
	HealthCheckType string `pulumi:"healthCheckType"`
	// (string)
	HealthReporting string  `pulumi:"healthReporting"`
	Id              *string `pulumi:"id"`
	// (bool)
	IpAnchored bool `pulumi:"ipAnchored"`
	// (bool) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors. Default: true. Boolean: `true`, `false`.
	IsCnameEnabled bool `pulumi:"isCnameEnabled"`
	// (string)
	Name *string `pulumi:"name"`
	// (bool)
	PassiveHealthEnabled bool `pulumi:"passiveHealthEnabled"`
	// (string)
	SegmentGroupId string `pulumi:"segmentGroupId"`
	// (string)
	SegmentGroupName string                                             `pulumi:"segmentGroupName"`
	ServerGroups     []GetZPAApplicationSegmentBrowserAccessServerGroup `pulumi:"serverGroups"`
	// (string) TCP port ranges used to access the app.
	TcpPortRanges []string `pulumi:"tcpPortRanges"`
	// (string) UDP port ranges used to access the app.
	UdpPortRanges []string `pulumi:"udpPortRanges"`
}

func LookupZPAApplicationSegmentBrowserAccessOutput(ctx *pulumi.Context, args LookupZPAApplicationSegmentBrowserAccessOutputArgs, opts ...pulumi.InvokeOption) LookupZPAApplicationSegmentBrowserAccessResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupZPAApplicationSegmentBrowserAccessResult, error) {
			args := v.(LookupZPAApplicationSegmentBrowserAccessArgs)
			r, err := LookupZPAApplicationSegmentBrowserAccess(ctx, &args, opts...)
			var s LookupZPAApplicationSegmentBrowserAccessResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupZPAApplicationSegmentBrowserAccessResultOutput)
}

// A collection of arguments for invoking getZPAApplicationSegmentBrowserAccess.
type LookupZPAApplicationSegmentBrowserAccessOutputArgs struct {
	// This field defines the id of the application server.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// This field defines the name of the server.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// (string) TCP port ranges used to access the app.
	// * `from:`
	// * `to:`
	TcpPortRanges GetZPAApplicationSegmentBrowserAccessTcpPortRangeArrayInput `pulumi:"tcpPortRanges"`
	// (string) UDP port ranges used to access the app.
	// * `from:`
	// * `to:`
	UdpPortRanges GetZPAApplicationSegmentBrowserAccessUdpPortRangeArrayInput `pulumi:"udpPortRanges"`
}

func (LookupZPAApplicationSegmentBrowserAccessOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupZPAApplicationSegmentBrowserAccessArgs)(nil)).Elem()
}

// A collection of values returned by getZPAApplicationSegmentBrowserAccess.
type LookupZPAApplicationSegmentBrowserAccessResultOutput struct{ *pulumi.OutputState }

func (LookupZPAApplicationSegmentBrowserAccessResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupZPAApplicationSegmentBrowserAccessResult)(nil)).Elem()
}

func (o LookupZPAApplicationSegmentBrowserAccessResultOutput) ToLookupZPAApplicationSegmentBrowserAccessResultOutput() LookupZPAApplicationSegmentBrowserAccessResultOutput {
	return o
}

func (o LookupZPAApplicationSegmentBrowserAccessResultOutput) ToLookupZPAApplicationSegmentBrowserAccessResultOutputWithContext(ctx context.Context) LookupZPAApplicationSegmentBrowserAccessResultOutput {
	return o
}

// (string) Indicates whether users can bypass ZPA to access applications. Default: `NEVER`. Supported values: `ALWAYS`, `NEVER`, `ON_NET`. The value `NEVER` indicates the use of the client forwarding policy.
func (o LookupZPAApplicationSegmentBrowserAccessResultOutput) BypassType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentBrowserAccessResult) string { return v.BypassType }).(pulumi.StringOutput)
}

func (o LookupZPAApplicationSegmentBrowserAccessResultOutput) ClientlessApps() GetZPAApplicationSegmentBrowserAccessClientlessAppArrayOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentBrowserAccessResult) []GetZPAApplicationSegmentBrowserAccessClientlessApp {
		return v.ClientlessApps
	}).(GetZPAApplicationSegmentBrowserAccessClientlessAppArrayOutput)
}

// (string)
func (o LookupZPAApplicationSegmentBrowserAccessResultOutput) ConfigSpace() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentBrowserAccessResult) string { return v.ConfigSpace }).(pulumi.StringOutput)
}

// (string)
func (o LookupZPAApplicationSegmentBrowserAccessResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentBrowserAccessResult) string { return v.Description }).(pulumi.StringOutput)
}

// List of domains and IPs.
func (o LookupZPAApplicationSegmentBrowserAccessResultOutput) DomainNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentBrowserAccessResult) []string { return v.DomainNames }).(pulumi.StringArrayOutput)
}

// (string) Whether Double Encryption is enabled or disabled for the app. Default: false. Boolean: `true`, `false`.
func (o LookupZPAApplicationSegmentBrowserAccessResultOutput) DoubleEncrypt() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentBrowserAccessResult) bool { return v.DoubleEncrypt }).(pulumi.BoolOutput)
}

// (bool)
func (o LookupZPAApplicationSegmentBrowserAccessResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentBrowserAccessResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupZPAApplicationSegmentBrowserAccessResultOutput) HealthCheckType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentBrowserAccessResult) string { return v.HealthCheckType }).(pulumi.StringOutput)
}

// (string)
func (o LookupZPAApplicationSegmentBrowserAccessResultOutput) HealthReporting() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentBrowserAccessResult) string { return v.HealthReporting }).(pulumi.StringOutput)
}

func (o LookupZPAApplicationSegmentBrowserAccessResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentBrowserAccessResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// (bool)
func (o LookupZPAApplicationSegmentBrowserAccessResultOutput) IpAnchored() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentBrowserAccessResult) bool { return v.IpAnchored }).(pulumi.BoolOutput)
}

// (bool) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors. Default: true. Boolean: `true`, `false`.
func (o LookupZPAApplicationSegmentBrowserAccessResultOutput) IsCnameEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentBrowserAccessResult) bool { return v.IsCnameEnabled }).(pulumi.BoolOutput)
}

// (string)
func (o LookupZPAApplicationSegmentBrowserAccessResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentBrowserAccessResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// (bool)
func (o LookupZPAApplicationSegmentBrowserAccessResultOutput) PassiveHealthEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentBrowserAccessResult) bool { return v.PassiveHealthEnabled }).(pulumi.BoolOutput)
}

// (string)
func (o LookupZPAApplicationSegmentBrowserAccessResultOutput) SegmentGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentBrowserAccessResult) string { return v.SegmentGroupId }).(pulumi.StringOutput)
}

// (string)
func (o LookupZPAApplicationSegmentBrowserAccessResultOutput) SegmentGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentBrowserAccessResult) string { return v.SegmentGroupName }).(pulumi.StringOutput)
}

func (o LookupZPAApplicationSegmentBrowserAccessResultOutput) ServerGroups() GetZPAApplicationSegmentBrowserAccessServerGroupArrayOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentBrowserAccessResult) []GetZPAApplicationSegmentBrowserAccessServerGroup {
		return v.ServerGroups
	}).(GetZPAApplicationSegmentBrowserAccessServerGroupArrayOutput)
}

// (string) TCP port ranges used to access the app.
func (o LookupZPAApplicationSegmentBrowserAccessResultOutput) TcpPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentBrowserAccessResult) []string { return v.TcpPortRanges }).(pulumi.StringArrayOutput)
}

// (string) UDP port ranges used to access the app.
func (o LookupZPAApplicationSegmentBrowserAccessResultOutput) UdpPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentBrowserAccessResult) []string { return v.UdpPortRanges }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupZPAApplicationSegmentBrowserAccessResultOutput{})
}
