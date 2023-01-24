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
//			_, err := zpa.LookupZPASegmentGroup(ctx, &zpa.LookupZPASegmentGroupArgs{
//				Name: pulumi.StringRef("segment_group_name"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupZPASegmentGroup(ctx *pulumi.Context, args *LookupZPASegmentGroupArgs, opts ...pulumi.InvokeOption) (*LookupZPASegmentGroupResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupZPASegmentGroupResult
	err := ctx.Invoke("zpa:index/getZPASegmentGroup:getZPASegmentGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getZPASegmentGroup.
type LookupZPASegmentGroupArgs struct {
	// The ID of the segment group to be exported.
	Id *string `pulumi:"id"`
	// The name of the segment group to be exported.
	Name *string `pulumi:"name"`
	// (bool)
	PolicyMigrated *bool `pulumi:"policyMigrated"`
}

// A collection of values returned by getZPASegmentGroup.
type LookupZPASegmentGroupResult struct {
	// (Computed)
	Applications []GetZPASegmentGroupApplication `pulumi:"applications"`
	// (string)
	ConfigSpace string `pulumi:"configSpace"`
	// (string)
	CreationTime string `pulumi:"creationTime"`
	// (string)
	Description string `pulumi:"description"`
	// (bool)
	Enabled bool `pulumi:"enabled"`
	// (string)
	Id *string `pulumi:"id"`
	// (string)
	ModifiedBy string `pulumi:"modifiedBy"`
	// (string)
	ModifiedTime string `pulumi:"modifiedTime"`
	// (string)
	Name *string `pulumi:"name"`
	// (bool)
	PolicyMigrated bool `pulumi:"policyMigrated"`
	// (string)
	TcpKeepAliveEnabled string `pulumi:"tcpKeepAliveEnabled"`
}

func LookupZPASegmentGroupOutput(ctx *pulumi.Context, args LookupZPASegmentGroupOutputArgs, opts ...pulumi.InvokeOption) LookupZPASegmentGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupZPASegmentGroupResult, error) {
			args := v.(LookupZPASegmentGroupArgs)
			r, err := LookupZPASegmentGroup(ctx, &args, opts...)
			var s LookupZPASegmentGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupZPASegmentGroupResultOutput)
}

// A collection of arguments for invoking getZPASegmentGroup.
type LookupZPASegmentGroupOutputArgs struct {
	// The ID of the segment group to be exported.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of the segment group to be exported.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// (bool)
	PolicyMigrated pulumi.BoolPtrInput `pulumi:"policyMigrated"`
}

func (LookupZPASegmentGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupZPASegmentGroupArgs)(nil)).Elem()
}

// A collection of values returned by getZPASegmentGroup.
type LookupZPASegmentGroupResultOutput struct{ *pulumi.OutputState }

func (LookupZPASegmentGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupZPASegmentGroupResult)(nil)).Elem()
}

func (o LookupZPASegmentGroupResultOutput) ToLookupZPASegmentGroupResultOutput() LookupZPASegmentGroupResultOutput {
	return o
}

func (o LookupZPASegmentGroupResultOutput) ToLookupZPASegmentGroupResultOutputWithContext(ctx context.Context) LookupZPASegmentGroupResultOutput {
	return o
}

// (Computed)
func (o LookupZPASegmentGroupResultOutput) Applications() GetZPASegmentGroupApplicationArrayOutput {
	return o.ApplyT(func(v LookupZPASegmentGroupResult) []GetZPASegmentGroupApplication { return v.Applications }).(GetZPASegmentGroupApplicationArrayOutput)
}

// (string)
func (o LookupZPASegmentGroupResultOutput) ConfigSpace() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZPASegmentGroupResult) string { return v.ConfigSpace }).(pulumi.StringOutput)
}

// (string)
func (o LookupZPASegmentGroupResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZPASegmentGroupResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// (string)
func (o LookupZPASegmentGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZPASegmentGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

// (bool)
func (o LookupZPASegmentGroupResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupZPASegmentGroupResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// (string)
func (o LookupZPASegmentGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupZPASegmentGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// (string)
func (o LookupZPASegmentGroupResultOutput) ModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZPASegmentGroupResult) string { return v.ModifiedBy }).(pulumi.StringOutput)
}

// (string)
func (o LookupZPASegmentGroupResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZPASegmentGroupResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

// (string)
func (o LookupZPASegmentGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupZPASegmentGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// (bool)
func (o LookupZPASegmentGroupResultOutput) PolicyMigrated() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupZPASegmentGroupResult) bool { return v.PolicyMigrated }).(pulumi.BoolOutput)
}

// (string)
func (o LookupZPASegmentGroupResultOutput) TcpKeepAliveEnabled() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZPASegmentGroupResult) string { return v.TcpKeepAliveEnabled }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupZPASegmentGroupResultOutput{})
}
