// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package segmentgroup

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
//	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/SegmentGroup"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := SegmentGroup.GetSegmentGroup(ctx, &segmentgroup.GetSegmentGroupArgs{
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
func LookupSegmentGroup(ctx *pulumi.Context, args *LookupSegmentGroupArgs, opts ...pulumi.InvokeOption) (*LookupSegmentGroupResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSegmentGroupResult
	err := ctx.Invoke("zpa:SegmentGroup/getSegmentGroup:getSegmentGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSegmentGroup.
type LookupSegmentGroupArgs struct {
	// The ID of the segment group to be exported.
	Id *string `pulumi:"id"`
	// The name of the segment group to be exported.
	Name *string `pulumi:"name"`
	// (bool)
	PolicyMigrated *bool `pulumi:"policyMigrated"`
}

// A collection of values returned by getSegmentGroup.
type LookupSegmentGroupResult struct {
	// (Computed)
	Applications []GetSegmentGroupApplication `pulumi:"applications"`
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

func LookupSegmentGroupOutput(ctx *pulumi.Context, args LookupSegmentGroupOutputArgs, opts ...pulumi.InvokeOption) LookupSegmentGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSegmentGroupResult, error) {
			args := v.(LookupSegmentGroupArgs)
			r, err := LookupSegmentGroup(ctx, &args, opts...)
			var s LookupSegmentGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSegmentGroupResultOutput)
}

// A collection of arguments for invoking getSegmentGroup.
type LookupSegmentGroupOutputArgs struct {
	// The ID of the segment group to be exported.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of the segment group to be exported.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// (bool)
	PolicyMigrated pulumi.BoolPtrInput `pulumi:"policyMigrated"`
}

func (LookupSegmentGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSegmentGroupArgs)(nil)).Elem()
}

// A collection of values returned by getSegmentGroup.
type LookupSegmentGroupResultOutput struct{ *pulumi.OutputState }

func (LookupSegmentGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSegmentGroupResult)(nil)).Elem()
}

func (o LookupSegmentGroupResultOutput) ToLookupSegmentGroupResultOutput() LookupSegmentGroupResultOutput {
	return o
}

func (o LookupSegmentGroupResultOutput) ToLookupSegmentGroupResultOutputWithContext(ctx context.Context) LookupSegmentGroupResultOutput {
	return o
}

// (Computed)
func (o LookupSegmentGroupResultOutput) Applications() GetSegmentGroupApplicationArrayOutput {
	return o.ApplyT(func(v LookupSegmentGroupResult) []GetSegmentGroupApplication { return v.Applications }).(GetSegmentGroupApplicationArrayOutput)
}

// (string)
func (o LookupSegmentGroupResultOutput) ConfigSpace() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSegmentGroupResult) string { return v.ConfigSpace }).(pulumi.StringOutput)
}

// (string)
func (o LookupSegmentGroupResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSegmentGroupResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// (string)
func (o LookupSegmentGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSegmentGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

// (bool)
func (o LookupSegmentGroupResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSegmentGroupResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// (string)
func (o LookupSegmentGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSegmentGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// (string)
func (o LookupSegmentGroupResultOutput) ModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSegmentGroupResult) string { return v.ModifiedBy }).(pulumi.StringOutput)
}

// (string)
func (o LookupSegmentGroupResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSegmentGroupResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

// (string)
func (o LookupSegmentGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSegmentGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// (bool)
func (o LookupSegmentGroupResultOutput) PolicyMigrated() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSegmentGroupResult) bool { return v.PolicyMigrated }).(pulumi.BoolOutput)
}

// (string)
func (o LookupSegmentGroupResultOutput) TcpKeepAliveEnabled() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSegmentGroupResult) string { return v.TcpKeepAliveEnabled }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSegmentGroupResultOutput{})
}
