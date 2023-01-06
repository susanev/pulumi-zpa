// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the **zpa_application_segment_inspection** data source to get information about an inspection application segment in the Zscaler Private Access cloud. This resource can then be referenced in a ZPA access inspection policy. This resource supports ZPA Inspection for both `HTTP` and `HTTPS`.
//
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
//			_, err := zpa.LookupZPAApplicationSegmentInspection(ctx, &zpa.LookupZPAApplicationSegmentInspectionArgs{
//				Name: pulumi.StringRef("ZPA_Inspection_Example"),
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
//			_, err := zpa.LookupZPAApplicationSegmentInspection(ctx, &zpa.LookupZPAApplicationSegmentInspectionArgs{
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
func LookupZPAApplicationSegmentInspection(ctx *pulumi.Context, args *LookupZPAApplicationSegmentInspectionArgs, opts ...pulumi.InvokeOption) (*LookupZPAApplicationSegmentInspectionResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupZPAApplicationSegmentInspectionResult
	err := ctx.Invoke("zpa:index/getZPAApplicationSegmentInspection:getZPAApplicationSegmentInspection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getZPAApplicationSegmentInspection.
type LookupZPAApplicationSegmentInspectionArgs struct {
	// The ID of the Inspection Application Segment to be exported.
	Id *string `pulumi:"id"`
	// The name of the Inspection Application Segment to be exported.
	Name *string `pulumi:"name"`
	// (string) TCP port ranges used to access the app.
	// * `from:`
	// * `to:`
	TcpPortRanges []GetZPAApplicationSegmentInspectionTcpPortRange `pulumi:"tcpPortRanges"`
	// (string) UDP port ranges used to access the app.
	// * `from:`
	// * `to:`
	UdpPortRanges []GetZPAApplicationSegmentInspectionUdpPortRange `pulumi:"udpPortRanges"`
}

// A collection of values returned by getZPAApplicationSegmentInspection.
type LookupZPAApplicationSegmentInspectionResult struct {
	// (string) Indicates whether users can bypass ZPA to access applications.
	BypassType string `pulumi:"bypassType"`
	// (String)
	CreationTime string `pulumi:"creationTime"`
	// (string) Description of the application.
	Description string `pulumi:"description"`
	// (string) List of domains and IPs.
	DomainNames []string `pulumi:"domainNames"`
	// (bool) Whether Double Encryption is enabled or disabled for the app.
	DoubleEncrypt bool `pulumi:"doubleEncrypt"`
	// (bool) Whether this application is enabled or not
	Enabled         bool   `pulumi:"enabled"`
	HealthCheckType string `pulumi:"healthCheckType"`
	// (string) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
	// * `healthCheckType` (string)
	HealthReporting string `pulumi:"healthReporting"`
	// (string)
	IcmpAccessType string  `pulumi:"icmpAccessType"`
	Id             *string `pulumi:"id"`
	// (string) TCP port ranges used to access the app.
	// * `app_id:` - (string)
	// * `name:` - (string) Name of the Inspection Application
	// * `description:` - (string) Description of the Inspection Application
	// * `domain:` - (string) Domain name of the inspection application
	InspectionApps []GetZPAApplicationSegmentInspectionInspectionApp `pulumi:"inspectionApps"`
	// (bool)
	IpAnchored bool `pulumi:"ipAnchored"`
	// (bool) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
	IsCnameEnabled bool   `pulumi:"isCnameEnabled"`
	ModifiedBy     string `pulumi:"modifiedBy"`
	// (String)
	ModifiedTime string  `pulumi:"modifiedTime"`
	Name         *string `pulumi:"name"`
	// (bool)
	PassiveHealthEnabled bool `pulumi:"passiveHealthEnabled"`
	// (String) Segment Group IDs
	SegmentGroupId            string `pulumi:"segmentGroupId"`
	SegmentGroupName          string `pulumi:"segmentGroupName"`
	SelectConnectorCloseToApp bool   `pulumi:"selectConnectorCloseToApp"`
	// (string) List of Server Group IDs
	// * `id:` - (string) List of Server Group IDs
	ServerGroups []GetZPAApplicationSegmentInspectionServerGroup `pulumi:"serverGroups"`
	// (string) TCP port ranges used to access the app.
	TcpPortRanges []string `pulumi:"tcpPortRanges"`
	// (string) UDP port ranges used to access the app.
	UdpPortRanges []string `pulumi:"udpPortRanges"`
}

func LookupZPAApplicationSegmentInspectionOutput(ctx *pulumi.Context, args LookupZPAApplicationSegmentInspectionOutputArgs, opts ...pulumi.InvokeOption) LookupZPAApplicationSegmentInspectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupZPAApplicationSegmentInspectionResult, error) {
			args := v.(LookupZPAApplicationSegmentInspectionArgs)
			r, err := LookupZPAApplicationSegmentInspection(ctx, &args, opts...)
			var s LookupZPAApplicationSegmentInspectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupZPAApplicationSegmentInspectionResultOutput)
}

// A collection of arguments for invoking getZPAApplicationSegmentInspection.
type LookupZPAApplicationSegmentInspectionOutputArgs struct {
	// The ID of the Inspection Application Segment to be exported.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of the Inspection Application Segment to be exported.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// (string) TCP port ranges used to access the app.
	// * `from:`
	// * `to:`
	TcpPortRanges GetZPAApplicationSegmentInspectionTcpPortRangeArrayInput `pulumi:"tcpPortRanges"`
	// (string) UDP port ranges used to access the app.
	// * `from:`
	// * `to:`
	UdpPortRanges GetZPAApplicationSegmentInspectionUdpPortRangeArrayInput `pulumi:"udpPortRanges"`
}

func (LookupZPAApplicationSegmentInspectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupZPAApplicationSegmentInspectionArgs)(nil)).Elem()
}

// A collection of values returned by getZPAApplicationSegmentInspection.
type LookupZPAApplicationSegmentInspectionResultOutput struct{ *pulumi.OutputState }

func (LookupZPAApplicationSegmentInspectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupZPAApplicationSegmentInspectionResult)(nil)).Elem()
}

func (o LookupZPAApplicationSegmentInspectionResultOutput) ToLookupZPAApplicationSegmentInspectionResultOutput() LookupZPAApplicationSegmentInspectionResultOutput {
	return o
}

func (o LookupZPAApplicationSegmentInspectionResultOutput) ToLookupZPAApplicationSegmentInspectionResultOutputWithContext(ctx context.Context) LookupZPAApplicationSegmentInspectionResultOutput {
	return o
}

// (string) Indicates whether users can bypass ZPA to access applications.
func (o LookupZPAApplicationSegmentInspectionResultOutput) BypassType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentInspectionResult) string { return v.BypassType }).(pulumi.StringOutput)
}

// (String)
func (o LookupZPAApplicationSegmentInspectionResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentInspectionResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// (string) Description of the application.
func (o LookupZPAApplicationSegmentInspectionResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentInspectionResult) string { return v.Description }).(pulumi.StringOutput)
}

// (string) List of domains and IPs.
func (o LookupZPAApplicationSegmentInspectionResultOutput) DomainNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentInspectionResult) []string { return v.DomainNames }).(pulumi.StringArrayOutput)
}

// (bool) Whether Double Encryption is enabled or disabled for the app.
func (o LookupZPAApplicationSegmentInspectionResultOutput) DoubleEncrypt() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentInspectionResult) bool { return v.DoubleEncrypt }).(pulumi.BoolOutput)
}

// (bool) Whether this application is enabled or not
func (o LookupZPAApplicationSegmentInspectionResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentInspectionResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupZPAApplicationSegmentInspectionResultOutput) HealthCheckType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentInspectionResult) string { return v.HealthCheckType }).(pulumi.StringOutput)
}

// (string) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
// * `healthCheckType` (string)
func (o LookupZPAApplicationSegmentInspectionResultOutput) HealthReporting() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentInspectionResult) string { return v.HealthReporting }).(pulumi.StringOutput)
}

// (string)
func (o LookupZPAApplicationSegmentInspectionResultOutput) IcmpAccessType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentInspectionResult) string { return v.IcmpAccessType }).(pulumi.StringOutput)
}

func (o LookupZPAApplicationSegmentInspectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentInspectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// (string) TCP port ranges used to access the app.
// * `app_id:` - (string)
// * `name:` - (string) Name of the Inspection Application
// * `description:` - (string) Description of the Inspection Application
// * `domain:` - (string) Domain name of the inspection application
func (o LookupZPAApplicationSegmentInspectionResultOutput) InspectionApps() GetZPAApplicationSegmentInspectionInspectionAppArrayOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentInspectionResult) []GetZPAApplicationSegmentInspectionInspectionApp {
		return v.InspectionApps
	}).(GetZPAApplicationSegmentInspectionInspectionAppArrayOutput)
}

// (bool)
func (o LookupZPAApplicationSegmentInspectionResultOutput) IpAnchored() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentInspectionResult) bool { return v.IpAnchored }).(pulumi.BoolOutput)
}

// (bool) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
func (o LookupZPAApplicationSegmentInspectionResultOutput) IsCnameEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentInspectionResult) bool { return v.IsCnameEnabled }).(pulumi.BoolOutput)
}

func (o LookupZPAApplicationSegmentInspectionResultOutput) ModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentInspectionResult) string { return v.ModifiedBy }).(pulumi.StringOutput)
}

// (String)
func (o LookupZPAApplicationSegmentInspectionResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentInspectionResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

func (o LookupZPAApplicationSegmentInspectionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentInspectionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// (bool)
func (o LookupZPAApplicationSegmentInspectionResultOutput) PassiveHealthEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentInspectionResult) bool { return v.PassiveHealthEnabled }).(pulumi.BoolOutput)
}

// (String) Segment Group IDs
func (o LookupZPAApplicationSegmentInspectionResultOutput) SegmentGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentInspectionResult) string { return v.SegmentGroupId }).(pulumi.StringOutput)
}

func (o LookupZPAApplicationSegmentInspectionResultOutput) SegmentGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentInspectionResult) string { return v.SegmentGroupName }).(pulumi.StringOutput)
}

func (o LookupZPAApplicationSegmentInspectionResultOutput) SelectConnectorCloseToApp() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentInspectionResult) bool { return v.SelectConnectorCloseToApp }).(pulumi.BoolOutput)
}

// (string) List of Server Group IDs
// * `id:` - (string) List of Server Group IDs
func (o LookupZPAApplicationSegmentInspectionResultOutput) ServerGroups() GetZPAApplicationSegmentInspectionServerGroupArrayOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentInspectionResult) []GetZPAApplicationSegmentInspectionServerGroup {
		return v.ServerGroups
	}).(GetZPAApplicationSegmentInspectionServerGroupArrayOutput)
}

// (string) TCP port ranges used to access the app.
func (o LookupZPAApplicationSegmentInspectionResultOutput) TcpPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentInspectionResult) []string { return v.TcpPortRanges }).(pulumi.StringArrayOutput)
}

// (string) UDP port ranges used to access the app.
func (o LookupZPAApplicationSegmentInspectionResultOutput) UdpPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupZPAApplicationSegmentInspectionResult) []string { return v.UdpPortRanges }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupZPAApplicationSegmentInspectionResultOutput{})
}
