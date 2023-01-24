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
// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Application Server can be imported by using `<APPLICATION SERVER ID>` or `<APPLICATION SERVER NAME>` as the import ID For example
//
// ```sh
//
//	$ pulumi import zpa:index/zPAApplicationServer:ZPAApplicationServer example <application_server_id>
//
// ```
//
//	or
//
// ```sh
//
//	$ pulumi import zpa:index/zPAApplicationServer:ZPAApplicationServer example <application_server_name>
//
// ```
type ZPAApplicationServer struct {
	pulumi.CustomResourceState

	// Address. The address of the application server to be exported.
	Address pulumi.StringOutput `pulumi:"address"`
	// (Optional) This field defines the list of server group IDs.
	AppServerGroupIds pulumi.StringArrayOutput `pulumi:"appServerGroupIds"`
	// (Optional)
	ConfigSpace pulumi.StringPtrOutput `pulumi:"configSpace"`
	// (Optional) This field defines the description of the server.
	Description pulumi.StringOutput `pulumi:"description"`
	// (Optional) This field defines the status of the server.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Name. The name of the application server to be exported.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewZPAApplicationServer registers a new resource with the given unique name, arguments, and options.
func NewZPAApplicationServer(ctx *pulumi.Context,
	name string, args *ZPAApplicationServerArgs, opts ...pulumi.ResourceOption) (*ZPAApplicationServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ZPAApplicationServer
	err := ctx.RegisterResource("zpa:index/zPAApplicationServer:ZPAApplicationServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZPAApplicationServer gets an existing ZPAApplicationServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZPAApplicationServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZPAApplicationServerState, opts ...pulumi.ResourceOption) (*ZPAApplicationServer, error) {
	var resource ZPAApplicationServer
	err := ctx.ReadResource("zpa:index/zPAApplicationServer:ZPAApplicationServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ZPAApplicationServer resources.
type zpaapplicationServerState struct {
	// Address. The address of the application server to be exported.
	Address *string `pulumi:"address"`
	// (Optional) This field defines the list of server group IDs.
	AppServerGroupIds []string `pulumi:"appServerGroupIds"`
	// (Optional)
	ConfigSpace *string `pulumi:"configSpace"`
	// (Optional) This field defines the description of the server.
	Description *string `pulumi:"description"`
	// (Optional) This field defines the status of the server.
	Enabled *bool `pulumi:"enabled"`
	// Name. The name of the application server to be exported.
	Name *string `pulumi:"name"`
}

type ZPAApplicationServerState struct {
	// Address. The address of the application server to be exported.
	Address pulumi.StringPtrInput
	// (Optional) This field defines the list of server group IDs.
	AppServerGroupIds pulumi.StringArrayInput
	// (Optional)
	ConfigSpace pulumi.StringPtrInput
	// (Optional) This field defines the description of the server.
	Description pulumi.StringPtrInput
	// (Optional) This field defines the status of the server.
	Enabled pulumi.BoolPtrInput
	// Name. The name of the application server to be exported.
	Name pulumi.StringPtrInput
}

func (ZPAApplicationServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*zpaapplicationServerState)(nil)).Elem()
}

type zpaapplicationServerArgs struct {
	// Address. The address of the application server to be exported.
	Address string `pulumi:"address"`
	// (Optional) This field defines the list of server group IDs.
	AppServerGroupIds []string `pulumi:"appServerGroupIds"`
	// (Optional)
	ConfigSpace *string `pulumi:"configSpace"`
	// (Optional) This field defines the description of the server.
	Description *string `pulumi:"description"`
	// (Optional) This field defines the status of the server.
	Enabled *bool `pulumi:"enabled"`
	// Name. The name of the application server to be exported.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ZPAApplicationServer resource.
type ZPAApplicationServerArgs struct {
	// Address. The address of the application server to be exported.
	Address pulumi.StringInput
	// (Optional) This field defines the list of server group IDs.
	AppServerGroupIds pulumi.StringArrayInput
	// (Optional)
	ConfigSpace pulumi.StringPtrInput
	// (Optional) This field defines the description of the server.
	Description pulumi.StringPtrInput
	// (Optional) This field defines the status of the server.
	Enabled pulumi.BoolPtrInput
	// Name. The name of the application server to be exported.
	Name pulumi.StringPtrInput
}

func (ZPAApplicationServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zpaapplicationServerArgs)(nil)).Elem()
}

type ZPAApplicationServerInput interface {
	pulumi.Input

	ToZPAApplicationServerOutput() ZPAApplicationServerOutput
	ToZPAApplicationServerOutputWithContext(ctx context.Context) ZPAApplicationServerOutput
}

func (*ZPAApplicationServer) ElementType() reflect.Type {
	return reflect.TypeOf((**ZPAApplicationServer)(nil)).Elem()
}

func (i *ZPAApplicationServer) ToZPAApplicationServerOutput() ZPAApplicationServerOutput {
	return i.ToZPAApplicationServerOutputWithContext(context.Background())
}

func (i *ZPAApplicationServer) ToZPAApplicationServerOutputWithContext(ctx context.Context) ZPAApplicationServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZPAApplicationServerOutput)
}

// ZPAApplicationServerArrayInput is an input type that accepts ZPAApplicationServerArray and ZPAApplicationServerArrayOutput values.
// You can construct a concrete instance of `ZPAApplicationServerArrayInput` via:
//
//	ZPAApplicationServerArray{ ZPAApplicationServerArgs{...} }
type ZPAApplicationServerArrayInput interface {
	pulumi.Input

	ToZPAApplicationServerArrayOutput() ZPAApplicationServerArrayOutput
	ToZPAApplicationServerArrayOutputWithContext(context.Context) ZPAApplicationServerArrayOutput
}

type ZPAApplicationServerArray []ZPAApplicationServerInput

func (ZPAApplicationServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ZPAApplicationServer)(nil)).Elem()
}

func (i ZPAApplicationServerArray) ToZPAApplicationServerArrayOutput() ZPAApplicationServerArrayOutput {
	return i.ToZPAApplicationServerArrayOutputWithContext(context.Background())
}

func (i ZPAApplicationServerArray) ToZPAApplicationServerArrayOutputWithContext(ctx context.Context) ZPAApplicationServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZPAApplicationServerArrayOutput)
}

// ZPAApplicationServerMapInput is an input type that accepts ZPAApplicationServerMap and ZPAApplicationServerMapOutput values.
// You can construct a concrete instance of `ZPAApplicationServerMapInput` via:
//
//	ZPAApplicationServerMap{ "key": ZPAApplicationServerArgs{...} }
type ZPAApplicationServerMapInput interface {
	pulumi.Input

	ToZPAApplicationServerMapOutput() ZPAApplicationServerMapOutput
	ToZPAApplicationServerMapOutputWithContext(context.Context) ZPAApplicationServerMapOutput
}

type ZPAApplicationServerMap map[string]ZPAApplicationServerInput

func (ZPAApplicationServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ZPAApplicationServer)(nil)).Elem()
}

func (i ZPAApplicationServerMap) ToZPAApplicationServerMapOutput() ZPAApplicationServerMapOutput {
	return i.ToZPAApplicationServerMapOutputWithContext(context.Background())
}

func (i ZPAApplicationServerMap) ToZPAApplicationServerMapOutputWithContext(ctx context.Context) ZPAApplicationServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZPAApplicationServerMapOutput)
}

type ZPAApplicationServerOutput struct{ *pulumi.OutputState }

func (ZPAApplicationServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ZPAApplicationServer)(nil)).Elem()
}

func (o ZPAApplicationServerOutput) ToZPAApplicationServerOutput() ZPAApplicationServerOutput {
	return o
}

func (o ZPAApplicationServerOutput) ToZPAApplicationServerOutputWithContext(ctx context.Context) ZPAApplicationServerOutput {
	return o
}

// Address. The address of the application server to be exported.
func (o ZPAApplicationServerOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *ZPAApplicationServer) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// (Optional) This field defines the list of server group IDs.
func (o ZPAApplicationServerOutput) AppServerGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ZPAApplicationServer) pulumi.StringArrayOutput { return v.AppServerGroupIds }).(pulumi.StringArrayOutput)
}

// (Optional)
func (o ZPAApplicationServerOutput) ConfigSpace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ZPAApplicationServer) pulumi.StringPtrOutput { return v.ConfigSpace }).(pulumi.StringPtrOutput)
}

// (Optional) This field defines the description of the server.
func (o ZPAApplicationServerOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ZPAApplicationServer) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// (Optional) This field defines the status of the server.
func (o ZPAApplicationServerOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ZPAApplicationServer) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// Name. The name of the application server to be exported.
func (o ZPAApplicationServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ZPAApplicationServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type ZPAApplicationServerArrayOutput struct{ *pulumi.OutputState }

func (ZPAApplicationServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ZPAApplicationServer)(nil)).Elem()
}

func (o ZPAApplicationServerArrayOutput) ToZPAApplicationServerArrayOutput() ZPAApplicationServerArrayOutput {
	return o
}

func (o ZPAApplicationServerArrayOutput) ToZPAApplicationServerArrayOutputWithContext(ctx context.Context) ZPAApplicationServerArrayOutput {
	return o
}

func (o ZPAApplicationServerArrayOutput) Index(i pulumi.IntInput) ZPAApplicationServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ZPAApplicationServer {
		return vs[0].([]*ZPAApplicationServer)[vs[1].(int)]
	}).(ZPAApplicationServerOutput)
}

type ZPAApplicationServerMapOutput struct{ *pulumi.OutputState }

func (ZPAApplicationServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ZPAApplicationServer)(nil)).Elem()
}

func (o ZPAApplicationServerMapOutput) ToZPAApplicationServerMapOutput() ZPAApplicationServerMapOutput {
	return o
}

func (o ZPAApplicationServerMapOutput) ToZPAApplicationServerMapOutputWithContext(ctx context.Context) ZPAApplicationServerMapOutput {
	return o
}

func (o ZPAApplicationServerMapOutput) MapIndex(k pulumi.StringInput) ZPAApplicationServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ZPAApplicationServer {
		return vs[0].(map[string]*ZPAApplicationServer)[vs[1].(string)]
	}).(ZPAApplicationServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ZPAApplicationServerInput)(nil)).Elem(), &ZPAApplicationServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZPAApplicationServerArrayInput)(nil)).Elem(), ZPAApplicationServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZPAApplicationServerMapInput)(nil)).Elem(), ZPAApplicationServerMap{})
	pulumi.RegisterOutputType(ZPAApplicationServerOutput{})
	pulumi.RegisterOutputType(ZPAApplicationServerArrayOutput{})
	pulumi.RegisterOutputType(ZPAApplicationServerMapOutput{})
}
