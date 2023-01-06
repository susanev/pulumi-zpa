// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the **zpa_machine_group** data source to get information about a machine group created in the Zscaler Private Access cloud. This data source can then be referenced in an Access Policy, Timeout policy, Forwarding Policy, Inspection Policy or Isolation Policy.
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
//			_, err := zpa.GetZPAMachineGroup(ctx, &zpa.GetZPAMachineGroupArgs{
//				Name: pulumi.StringRef("MGR01"),
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
//			_, err := zpa.GetZPAMachineGroup(ctx, &zpa.GetZPAMachineGroupArgs{
//				Id: pulumi.StringRef("1234567890"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetZPAMachineGroup(ctx *pulumi.Context, args *GetZPAMachineGroupArgs, opts ...pulumi.InvokeOption) (*GetZPAMachineGroupResult, error) {
	var rv GetZPAMachineGroupResult
	err := ctx.Invoke("zpa:index/getZPAMachineGroup:getZPAMachineGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getZPAMachineGroup.
type GetZPAMachineGroupArgs struct {
	// The ID of the machine group to be exported.
	Id *string `pulumi:"id"`
	// The name of the machine group to be exported.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getZPAMachineGroup.
type GetZPAMachineGroupResult struct {
	// (string)
	CreationTime string `pulumi:"creationTime"`
	// (string)
	Description string `pulumi:"description"`
	// (bool)
	Enabled bool `pulumi:"enabled"`
	// (string)
	Id *string `pulumi:"id"`
	// (string)
	Machines []GetZPAMachineGroupMachine `pulumi:"machines"`
	// (string)
	ModifiedBy string `pulumi:"modifiedBy"`
	// (string)
	ModifiedTime string `pulumi:"modifiedTime"`
	// (string)
	Name *string `pulumi:"name"`
}

func GetZPAMachineGroupOutput(ctx *pulumi.Context, args GetZPAMachineGroupOutputArgs, opts ...pulumi.InvokeOption) GetZPAMachineGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetZPAMachineGroupResult, error) {
			args := v.(GetZPAMachineGroupArgs)
			r, err := GetZPAMachineGroup(ctx, &args, opts...)
			var s GetZPAMachineGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetZPAMachineGroupResultOutput)
}

// A collection of arguments for invoking getZPAMachineGroup.
type GetZPAMachineGroupOutputArgs struct {
	// The ID of the machine group to be exported.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of the machine group to be exported.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetZPAMachineGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZPAMachineGroupArgs)(nil)).Elem()
}

// A collection of values returned by getZPAMachineGroup.
type GetZPAMachineGroupResultOutput struct{ *pulumi.OutputState }

func (GetZPAMachineGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZPAMachineGroupResult)(nil)).Elem()
}

func (o GetZPAMachineGroupResultOutput) ToGetZPAMachineGroupResultOutput() GetZPAMachineGroupResultOutput {
	return o
}

func (o GetZPAMachineGroupResultOutput) ToGetZPAMachineGroupResultOutputWithContext(ctx context.Context) GetZPAMachineGroupResultOutput {
	return o
}

// (string)
func (o GetZPAMachineGroupResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetZPAMachineGroupResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// (string)
func (o GetZPAMachineGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetZPAMachineGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

// (bool)
func (o GetZPAMachineGroupResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetZPAMachineGroupResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// (string)
func (o GetZPAMachineGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZPAMachineGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// (string)
func (o GetZPAMachineGroupResultOutput) Machines() GetZPAMachineGroupMachineArrayOutput {
	return o.ApplyT(func(v GetZPAMachineGroupResult) []GetZPAMachineGroupMachine { return v.Machines }).(GetZPAMachineGroupMachineArrayOutput)
}

// (string)
func (o GetZPAMachineGroupResultOutput) ModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v GetZPAMachineGroupResult) string { return v.ModifiedBy }).(pulumi.StringOutput)
}

// (string)
func (o GetZPAMachineGroupResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetZPAMachineGroupResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

// (string)
func (o GetZPAMachineGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZPAMachineGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetZPAMachineGroupResultOutput{})
}
