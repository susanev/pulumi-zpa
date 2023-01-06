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
//			_, err := zpa.NewZPASegmentGroup(ctx, "testSegmentGroup", &zpa.ZPASegmentGroupArgs{
//				Description:         pulumi.String("test1-segment-group"),
//				Enabled:             pulumi.Bool(true),
//				TcpKeepAliveEnabled: pulumi.String("1"),
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
// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) **segment_group** can be imported by using `<SEGMENT GROUP ID>` or `<SEGMENT GROUP NAME>` as the import ID. For example
//
// ```sh
//
//	$ pulumi import zpa:index/zPASegmentGroup:ZPASegmentGroup example <segment_group_id>
//
// ```
//
//	or
//
// ```sh
//
//	$ pulumi import zpa:index/zPASegmentGroup:ZPASegmentGroup example <segment_group_name>
//
// ```
type ZPASegmentGroup struct {
	pulumi.CustomResourceState

	Applications ZPASegmentGroupApplicationArrayOutput `pulumi:"applications"`
	// (Optional)
	ConfigSpace pulumi.StringPtrOutput `pulumi:"configSpace"`
	// (Optional) Description of the segment group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// (Optional) Whether this segment group is enabled or not.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// (Required) Name of the segment group.
	Name           pulumi.StringOutput  `pulumi:"name"`
	PolicyMigrated pulumi.BoolPtrOutput `pulumi:"policyMigrated"`
	// (Optional)
	TcpKeepAliveEnabled pulumi.StringPtrOutput `pulumi:"tcpKeepAliveEnabled"`
}

// NewZPASegmentGroup registers a new resource with the given unique name, arguments, and options.
func NewZPASegmentGroup(ctx *pulumi.Context,
	name string, args *ZPASegmentGroupArgs, opts ...pulumi.ResourceOption) (*ZPASegmentGroup, error) {
	if args == nil {
		args = &ZPASegmentGroupArgs{}
	}

	var resource ZPASegmentGroup
	err := ctx.RegisterResource("zpa:index/zPASegmentGroup:ZPASegmentGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZPASegmentGroup gets an existing ZPASegmentGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZPASegmentGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZPASegmentGroupState, opts ...pulumi.ResourceOption) (*ZPASegmentGroup, error) {
	var resource ZPASegmentGroup
	err := ctx.ReadResource("zpa:index/zPASegmentGroup:ZPASegmentGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ZPASegmentGroup resources.
type zpasegmentGroupState struct {
	Applications []ZPASegmentGroupApplication `pulumi:"applications"`
	// (Optional)
	ConfigSpace *string `pulumi:"configSpace"`
	// (Optional) Description of the segment group.
	Description *string `pulumi:"description"`
	// (Optional) Whether this segment group is enabled or not.
	Enabled *bool `pulumi:"enabled"`
	// (Required) Name of the segment group.
	Name           *string `pulumi:"name"`
	PolicyMigrated *bool   `pulumi:"policyMigrated"`
	// (Optional)
	TcpKeepAliveEnabled *string `pulumi:"tcpKeepAliveEnabled"`
}

type ZPASegmentGroupState struct {
	Applications ZPASegmentGroupApplicationArrayInput
	// (Optional)
	ConfigSpace pulumi.StringPtrInput
	// (Optional) Description of the segment group.
	Description pulumi.StringPtrInput
	// (Optional) Whether this segment group is enabled or not.
	Enabled pulumi.BoolPtrInput
	// (Required) Name of the segment group.
	Name           pulumi.StringPtrInput
	PolicyMigrated pulumi.BoolPtrInput
	// (Optional)
	TcpKeepAliveEnabled pulumi.StringPtrInput
}

func (ZPASegmentGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*zpasegmentGroupState)(nil)).Elem()
}

type zpasegmentGroupArgs struct {
	Applications []ZPASegmentGroupApplication `pulumi:"applications"`
	// (Optional)
	ConfigSpace *string `pulumi:"configSpace"`
	// (Optional) Description of the segment group.
	Description *string `pulumi:"description"`
	// (Optional) Whether this segment group is enabled or not.
	Enabled *bool `pulumi:"enabled"`
	// (Required) Name of the segment group.
	Name           *string `pulumi:"name"`
	PolicyMigrated *bool   `pulumi:"policyMigrated"`
	// (Optional)
	TcpKeepAliveEnabled *string `pulumi:"tcpKeepAliveEnabled"`
}

// The set of arguments for constructing a ZPASegmentGroup resource.
type ZPASegmentGroupArgs struct {
	Applications ZPASegmentGroupApplicationArrayInput
	// (Optional)
	ConfigSpace pulumi.StringPtrInput
	// (Optional) Description of the segment group.
	Description pulumi.StringPtrInput
	// (Optional) Whether this segment group is enabled or not.
	Enabled pulumi.BoolPtrInput
	// (Required) Name of the segment group.
	Name           pulumi.StringPtrInput
	PolicyMigrated pulumi.BoolPtrInput
	// (Optional)
	TcpKeepAliveEnabled pulumi.StringPtrInput
}

func (ZPASegmentGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zpasegmentGroupArgs)(nil)).Elem()
}

type ZPASegmentGroupInput interface {
	pulumi.Input

	ToZPASegmentGroupOutput() ZPASegmentGroupOutput
	ToZPASegmentGroupOutputWithContext(ctx context.Context) ZPASegmentGroupOutput
}

func (*ZPASegmentGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ZPASegmentGroup)(nil)).Elem()
}

func (i *ZPASegmentGroup) ToZPASegmentGroupOutput() ZPASegmentGroupOutput {
	return i.ToZPASegmentGroupOutputWithContext(context.Background())
}

func (i *ZPASegmentGroup) ToZPASegmentGroupOutputWithContext(ctx context.Context) ZPASegmentGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZPASegmentGroupOutput)
}

// ZPASegmentGroupArrayInput is an input type that accepts ZPASegmentGroupArray and ZPASegmentGroupArrayOutput values.
// You can construct a concrete instance of `ZPASegmentGroupArrayInput` via:
//
//	ZPASegmentGroupArray{ ZPASegmentGroupArgs{...} }
type ZPASegmentGroupArrayInput interface {
	pulumi.Input

	ToZPASegmentGroupArrayOutput() ZPASegmentGroupArrayOutput
	ToZPASegmentGroupArrayOutputWithContext(context.Context) ZPASegmentGroupArrayOutput
}

type ZPASegmentGroupArray []ZPASegmentGroupInput

func (ZPASegmentGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ZPASegmentGroup)(nil)).Elem()
}

func (i ZPASegmentGroupArray) ToZPASegmentGroupArrayOutput() ZPASegmentGroupArrayOutput {
	return i.ToZPASegmentGroupArrayOutputWithContext(context.Background())
}

func (i ZPASegmentGroupArray) ToZPASegmentGroupArrayOutputWithContext(ctx context.Context) ZPASegmentGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZPASegmentGroupArrayOutput)
}

// ZPASegmentGroupMapInput is an input type that accepts ZPASegmentGroupMap and ZPASegmentGroupMapOutput values.
// You can construct a concrete instance of `ZPASegmentGroupMapInput` via:
//
//	ZPASegmentGroupMap{ "key": ZPASegmentGroupArgs{...} }
type ZPASegmentGroupMapInput interface {
	pulumi.Input

	ToZPASegmentGroupMapOutput() ZPASegmentGroupMapOutput
	ToZPASegmentGroupMapOutputWithContext(context.Context) ZPASegmentGroupMapOutput
}

type ZPASegmentGroupMap map[string]ZPASegmentGroupInput

func (ZPASegmentGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ZPASegmentGroup)(nil)).Elem()
}

func (i ZPASegmentGroupMap) ToZPASegmentGroupMapOutput() ZPASegmentGroupMapOutput {
	return i.ToZPASegmentGroupMapOutputWithContext(context.Background())
}

func (i ZPASegmentGroupMap) ToZPASegmentGroupMapOutputWithContext(ctx context.Context) ZPASegmentGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZPASegmentGroupMapOutput)
}

type ZPASegmentGroupOutput struct{ *pulumi.OutputState }

func (ZPASegmentGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ZPASegmentGroup)(nil)).Elem()
}

func (o ZPASegmentGroupOutput) ToZPASegmentGroupOutput() ZPASegmentGroupOutput {
	return o
}

func (o ZPASegmentGroupOutput) ToZPASegmentGroupOutputWithContext(ctx context.Context) ZPASegmentGroupOutput {
	return o
}

func (o ZPASegmentGroupOutput) Applications() ZPASegmentGroupApplicationArrayOutput {
	return o.ApplyT(func(v *ZPASegmentGroup) ZPASegmentGroupApplicationArrayOutput { return v.Applications }).(ZPASegmentGroupApplicationArrayOutput)
}

// (Optional)
func (o ZPASegmentGroupOutput) ConfigSpace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ZPASegmentGroup) pulumi.StringPtrOutput { return v.ConfigSpace }).(pulumi.StringPtrOutput)
}

// (Optional) Description of the segment group.
func (o ZPASegmentGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ZPASegmentGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// (Optional) Whether this segment group is enabled or not.
func (o ZPASegmentGroupOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ZPASegmentGroup) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// (Required) Name of the segment group.
func (o ZPASegmentGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ZPASegmentGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ZPASegmentGroupOutput) PolicyMigrated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ZPASegmentGroup) pulumi.BoolPtrOutput { return v.PolicyMigrated }).(pulumi.BoolPtrOutput)
}

// (Optional)
func (o ZPASegmentGroupOutput) TcpKeepAliveEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ZPASegmentGroup) pulumi.StringPtrOutput { return v.TcpKeepAliveEnabled }).(pulumi.StringPtrOutput)
}

type ZPASegmentGroupArrayOutput struct{ *pulumi.OutputState }

func (ZPASegmentGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ZPASegmentGroup)(nil)).Elem()
}

func (o ZPASegmentGroupArrayOutput) ToZPASegmentGroupArrayOutput() ZPASegmentGroupArrayOutput {
	return o
}

func (o ZPASegmentGroupArrayOutput) ToZPASegmentGroupArrayOutputWithContext(ctx context.Context) ZPASegmentGroupArrayOutput {
	return o
}

func (o ZPASegmentGroupArrayOutput) Index(i pulumi.IntInput) ZPASegmentGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ZPASegmentGroup {
		return vs[0].([]*ZPASegmentGroup)[vs[1].(int)]
	}).(ZPASegmentGroupOutput)
}

type ZPASegmentGroupMapOutput struct{ *pulumi.OutputState }

func (ZPASegmentGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ZPASegmentGroup)(nil)).Elem()
}

func (o ZPASegmentGroupMapOutput) ToZPASegmentGroupMapOutput() ZPASegmentGroupMapOutput {
	return o
}

func (o ZPASegmentGroupMapOutput) ToZPASegmentGroupMapOutputWithContext(ctx context.Context) ZPASegmentGroupMapOutput {
	return o
}

func (o ZPASegmentGroupMapOutput) MapIndex(k pulumi.StringInput) ZPASegmentGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ZPASegmentGroup {
		return vs[0].(map[string]*ZPASegmentGroup)[vs[1].(string)]
	}).(ZPASegmentGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ZPASegmentGroupInput)(nil)).Elem(), &ZPASegmentGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZPASegmentGroupArrayInput)(nil)).Elem(), ZPASegmentGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZPASegmentGroupMapInput)(nil)).Elem(), ZPASegmentGroupMap{})
	pulumi.RegisterOutputType(ZPASegmentGroupOutput{})
	pulumi.RegisterOutputType(ZPASegmentGroupArrayOutput{})
	pulumi.RegisterOutputType(ZPASegmentGroupMapOutput{})
}
