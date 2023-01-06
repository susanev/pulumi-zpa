// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the **zpa_lss_config_log_type_formats** data source to get information about all LSS log type formats in the Zscaler Private Access cloud. This data source is required when creating an LSS Config Controller resource.
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
//			_, err := zpa.GetZPALSSLogTypeFormats(ctx, &zpa.GetZPALSSLogTypeFormatsArgs{
//				LogType: "zpn_trans_log",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = zpa.GetZPALSSLogTypeFormats(ctx, &zpa.GetZPALSSLogTypeFormatsArgs{
//				LogType: "zpn_auth_log",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = zpa.GetZPALSSLogTypeFormats(ctx, &zpa.GetZPALSSLogTypeFormatsArgs{
//				LogType: "zpn_ast_auth_log",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = zpa.GetZPALSSLogTypeFormats(ctx, &zpa.GetZPALSSLogTypeFormatsArgs{
//				LogType: "zpn_http_trans_log",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = zpa.GetZPALSSLogTypeFormats(ctx, &zpa.GetZPALSSLogTypeFormatsArgs{
//				LogType: "zpn_audit_log",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = zpa.GetZPALSSLogTypeFormats(ctx, &zpa.GetZPALSSLogTypeFormatsArgs{
//				LogType: "zpn_sys_auth_log",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = zpa.GetZPALSSLogTypeFormats(ctx, &zpa.GetZPALSSLogTypeFormatsArgs{
//				LogType: "zpn_ast_comprehensive_stats",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = zpa.GetZPALSSLogTypeFormats(ctx, &zpa.GetZPALSSLogTypeFormatsArgs{
//				LogType: "zpn_waf_http_exchanges_log",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetZPALSSLogTypeFormats(ctx *pulumi.Context, args *GetZPALSSLogTypeFormatsArgs, opts ...pulumi.InvokeOption) (*GetZPALSSLogTypeFormatsResult, error) {
	var rv GetZPALSSLogTypeFormatsResult
	err := ctx.Invoke("zpa:index/getZPALSSLogTypeFormats:getZPALSSLogTypeFormats", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getZPALSSLogTypeFormats.
type GetZPALSSLogTypeFormatsArgs struct {
	// The type of log to be exported.
	LogType string `pulumi:"logType"`
}

// A collection of values returned by getZPALSSLogTypeFormats.
type GetZPALSSLogTypeFormatsResult struct {
	Csv string `pulumi:"csv"`
	// The provider-assigned unique ID for this managed resource.
	Id      string `pulumi:"id"`
	Json    string `pulumi:"json"`
	LogType string `pulumi:"logType"`
	Tsv     string `pulumi:"tsv"`
}

func GetZPALSSLogTypeFormatsOutput(ctx *pulumi.Context, args GetZPALSSLogTypeFormatsOutputArgs, opts ...pulumi.InvokeOption) GetZPALSSLogTypeFormatsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetZPALSSLogTypeFormatsResult, error) {
			args := v.(GetZPALSSLogTypeFormatsArgs)
			r, err := GetZPALSSLogTypeFormats(ctx, &args, opts...)
			var s GetZPALSSLogTypeFormatsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetZPALSSLogTypeFormatsResultOutput)
}

// A collection of arguments for invoking getZPALSSLogTypeFormats.
type GetZPALSSLogTypeFormatsOutputArgs struct {
	// The type of log to be exported.
	LogType pulumi.StringInput `pulumi:"logType"`
}

func (GetZPALSSLogTypeFormatsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZPALSSLogTypeFormatsArgs)(nil)).Elem()
}

// A collection of values returned by getZPALSSLogTypeFormats.
type GetZPALSSLogTypeFormatsResultOutput struct{ *pulumi.OutputState }

func (GetZPALSSLogTypeFormatsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZPALSSLogTypeFormatsResult)(nil)).Elem()
}

func (o GetZPALSSLogTypeFormatsResultOutput) ToGetZPALSSLogTypeFormatsResultOutput() GetZPALSSLogTypeFormatsResultOutput {
	return o
}

func (o GetZPALSSLogTypeFormatsResultOutput) ToGetZPALSSLogTypeFormatsResultOutputWithContext(ctx context.Context) GetZPALSSLogTypeFormatsResultOutput {
	return o
}

func (o GetZPALSSLogTypeFormatsResultOutput) Csv() pulumi.StringOutput {
	return o.ApplyT(func(v GetZPALSSLogTypeFormatsResult) string { return v.Csv }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetZPALSSLogTypeFormatsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetZPALSSLogTypeFormatsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetZPALSSLogTypeFormatsResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetZPALSSLogTypeFormatsResult) string { return v.Json }).(pulumi.StringOutput)
}

func (o GetZPALSSLogTypeFormatsResultOutput) LogType() pulumi.StringOutput {
	return o.ApplyT(func(v GetZPALSSLogTypeFormatsResult) string { return v.LogType }).(pulumi.StringOutput)
}

func (o GetZPALSSLogTypeFormatsResultOutput) Tsv() pulumi.StringOutput {
	return o.ApplyT(func(v GetZPALSSLogTypeFormatsResult) string { return v.Tsv }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetZPALSSLogTypeFormatsResultOutput{})
}
