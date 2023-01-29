// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package segment

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
//	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/Segment"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Segment.GetApplicationSegment(ctx, &segment.GetApplicationSegmentArgs{
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
//	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/Segment"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Segment.GetApplicationSegment(ctx, &segment.GetApplicationSegmentArgs{
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
func LookupApplicationSegment(ctx *pulumi.Context, args *LookupApplicationSegmentArgs, opts ...pulumi.InvokeOption) (*LookupApplicationSegmentResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupApplicationSegmentResult
	err := ctx.Invoke("zpa:Segment/getApplicationSegment:getApplicationSegment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApplicationSegment.
type LookupApplicationSegmentArgs struct {
	Id *string `pulumi:"id"`
	// Name of the application.
	Name          *string                             `pulumi:"name"`
	TcpPortRanges []GetApplicationSegmentTcpPortRange `pulumi:"tcpPortRanges"`
	UdpPortRanges []GetApplicationSegmentUdpPortRange `pulumi:"udpPortRanges"`
}

// A collection of values returned by getApplicationSegment.
type LookupApplicationSegmentResult struct {
	BypassType                string                             `pulumi:"bypassType"`
	ConfigSpace               string                             `pulumi:"configSpace"`
	CreationTime              string                             `pulumi:"creationTime"`
	DefaultIdleTimeout        string                             `pulumi:"defaultIdleTimeout"`
	DefaultMaxAge             string                             `pulumi:"defaultMaxAge"`
	Description               string                             `pulumi:"description"`
	DomainNames               []string                           `pulumi:"domainNames"`
	DoubleEncrypt             bool                               `pulumi:"doubleEncrypt"`
	Enabled                   bool                               `pulumi:"enabled"`
	HealthCheckType           string                             `pulumi:"healthCheckType"`
	HealthReporting           string                             `pulumi:"healthReporting"`
	Id                        *string                            `pulumi:"id"`
	IpAnchored                bool                               `pulumi:"ipAnchored"`
	IsCnameEnabled            bool                               `pulumi:"isCnameEnabled"`
	ModifiedTime              string                             `pulumi:"modifiedTime"`
	Modifiedby                string                             `pulumi:"modifiedby"`
	Name                      *string                            `pulumi:"name"`
	PassiveHealthEnabled      bool                               `pulumi:"passiveHealthEnabled"`
	SegmentGroupId            string                             `pulumi:"segmentGroupId"`
	SegmentGroupName          string                             `pulumi:"segmentGroupName"`
	SelectConnectorCloseToApp bool                               `pulumi:"selectConnectorCloseToApp"`
	ServerGroups              []GetApplicationSegmentServerGroup `pulumi:"serverGroups"`
	TcpPortRanges             []string                           `pulumi:"tcpPortRanges"`
	UdpPortRanges             []string                           `pulumi:"udpPortRanges"`
}

func LookupApplicationSegmentOutput(ctx *pulumi.Context, args LookupApplicationSegmentOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationSegmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationSegmentResult, error) {
			args := v.(LookupApplicationSegmentArgs)
			r, err := LookupApplicationSegment(ctx, &args, opts...)
			var s LookupApplicationSegmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationSegmentResultOutput)
}

// A collection of arguments for invoking getApplicationSegment.
type LookupApplicationSegmentOutputArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Name of the application.
	Name          pulumi.StringPtrInput                       `pulumi:"name"`
	TcpPortRanges GetApplicationSegmentTcpPortRangeArrayInput `pulumi:"tcpPortRanges"`
	UdpPortRanges GetApplicationSegmentUdpPortRangeArrayInput `pulumi:"udpPortRanges"`
}

func (LookupApplicationSegmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationSegmentArgs)(nil)).Elem()
}

// A collection of values returned by getApplicationSegment.
type LookupApplicationSegmentResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationSegmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationSegmentResult)(nil)).Elem()
}

func (o LookupApplicationSegmentResultOutput) ToLookupApplicationSegmentResultOutput() LookupApplicationSegmentResultOutput {
	return o
}

func (o LookupApplicationSegmentResultOutput) ToLookupApplicationSegmentResultOutputWithContext(ctx context.Context) LookupApplicationSegmentResultOutput {
	return o
}

func (o LookupApplicationSegmentResultOutput) BypassType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentResult) string { return v.BypassType }).(pulumi.StringOutput)
}

func (o LookupApplicationSegmentResultOutput) ConfigSpace() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentResult) string { return v.ConfigSpace }).(pulumi.StringOutput)
}

func (o LookupApplicationSegmentResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupApplicationSegmentResultOutput) DefaultIdleTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentResult) string { return v.DefaultIdleTimeout }).(pulumi.StringOutput)
}

func (o LookupApplicationSegmentResultOutput) DefaultMaxAge() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentResult) string { return v.DefaultMaxAge }).(pulumi.StringOutput)
}

func (o LookupApplicationSegmentResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupApplicationSegmentResultOutput) DomainNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApplicationSegmentResult) []string { return v.DomainNames }).(pulumi.StringArrayOutput)
}

func (o LookupApplicationSegmentResultOutput) DoubleEncrypt() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApplicationSegmentResult) bool { return v.DoubleEncrypt }).(pulumi.BoolOutput)
}

func (o LookupApplicationSegmentResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApplicationSegmentResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupApplicationSegmentResultOutput) HealthCheckType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentResult) string { return v.HealthCheckType }).(pulumi.StringOutput)
}

func (o LookupApplicationSegmentResultOutput) HealthReporting() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentResult) string { return v.HealthReporting }).(pulumi.StringOutput)
}

func (o LookupApplicationSegmentResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationSegmentResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationSegmentResultOutput) IpAnchored() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApplicationSegmentResult) bool { return v.IpAnchored }).(pulumi.BoolOutput)
}

func (o LookupApplicationSegmentResultOutput) IsCnameEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApplicationSegmentResult) bool { return v.IsCnameEnabled }).(pulumi.BoolOutput)
}

func (o LookupApplicationSegmentResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

func (o LookupApplicationSegmentResultOutput) Modifiedby() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentResult) string { return v.Modifiedby }).(pulumi.StringOutput)
}

func (o LookupApplicationSegmentResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationSegmentResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationSegmentResultOutput) PassiveHealthEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApplicationSegmentResult) bool { return v.PassiveHealthEnabled }).(pulumi.BoolOutput)
}

func (o LookupApplicationSegmentResultOutput) SegmentGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentResult) string { return v.SegmentGroupId }).(pulumi.StringOutput)
}

func (o LookupApplicationSegmentResultOutput) SegmentGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSegmentResult) string { return v.SegmentGroupName }).(pulumi.StringOutput)
}

func (o LookupApplicationSegmentResultOutput) SelectConnectorCloseToApp() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApplicationSegmentResult) bool { return v.SelectConnectorCloseToApp }).(pulumi.BoolOutput)
}

func (o LookupApplicationSegmentResultOutput) ServerGroups() GetApplicationSegmentServerGroupArrayOutput {
	return o.ApplyT(func(v LookupApplicationSegmentResult) []GetApplicationSegmentServerGroup { return v.ServerGroups }).(GetApplicationSegmentServerGroupArrayOutput)
}

func (o LookupApplicationSegmentResultOutput) TcpPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApplicationSegmentResult) []string { return v.TcpPortRanges }).(pulumi.StringArrayOutput)
}

func (o LookupApplicationSegmentResultOutput) UdpPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApplicationSegmentResult) []string { return v.UdpPortRanges }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationSegmentResultOutput{})
}
