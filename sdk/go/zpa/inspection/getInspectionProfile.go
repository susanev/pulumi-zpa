// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package inspection

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the **zpa_inspection_profile** data source to get information about an inspection profile in the Zscaler Private Access cloud. This resource can then be referenced in an inspection custom control resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/Inspection"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Inspection.GetInspectionProfile(ctx, &inspection.GetInspectionProfileArgs{
//				Name: pulumi.StringRef("Example"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupInspectionProfile(ctx *pulumi.Context, args *LookupInspectionProfileArgs, opts ...pulumi.InvokeOption) (*LookupInspectionProfileResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupInspectionProfileResult
	err := ctx.Invoke("zpa:Inspection/getInspectionProfile:getInspectionProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInspectionProfile.
type LookupInspectionProfileArgs struct {
	// This field defines the id of the application server.
	Id *string `pulumi:"id"`
	// This field defines the name of the server.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getInspectionProfile.
type LookupInspectionProfileResult struct {
	// (string)
	CommonGlobalOverrideActionsConfig map[string]string `pulumi:"commonGlobalOverrideActionsConfig"`
	// (string) Types for custom controls
	ControlsInfos []GetInspectionProfileControlsInfo `pulumi:"controlsInfos"`
	CreationTime  string                             `pulumi:"creationTime"`
	// (string) Types for custom controls
	CustomControls []GetInspectionProfileCustomControl `pulumi:"customControls"`
	// (string) Description of the inspection profile.
	Description          string   `pulumi:"description"`
	GlobalControlActions []string `pulumi:"globalControlActions"`
	// (string) ID of the predefined control
	Id                string `pulumi:"id"`
	IncarnationNumber string `pulumi:"incarnationNumber"`
	ModifiedBy        string `pulumi:"modifiedBy"`
	ModifiedTime      string `pulumi:"modifiedTime"`
	// (string)
	Name string `pulumi:"name"`
	// (string) OWASP Predefined Paranoia Level. Range: [1-4], inclusive
	ParanoiaLevel string `pulumi:"paranoiaLevel"`
	// (string) The predefined controls
	PredefinedControls        []GetInspectionProfilePredefinedControl `pulumi:"predefinedControls"`
	PredefinedControlsVersion string                                  `pulumi:"predefinedControlsVersion"`
	// (string)
	WebSocketControls []GetInspectionProfileWebSocketControl `pulumi:"webSocketControls"`
}

func LookupInspectionProfileOutput(ctx *pulumi.Context, args LookupInspectionProfileOutputArgs, opts ...pulumi.InvokeOption) LookupInspectionProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInspectionProfileResult, error) {
			args := v.(LookupInspectionProfileArgs)
			r, err := LookupInspectionProfile(ctx, &args, opts...)
			var s LookupInspectionProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInspectionProfileResultOutput)
}

// A collection of arguments for invoking getInspectionProfile.
type LookupInspectionProfileOutputArgs struct {
	// This field defines the id of the application server.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// This field defines the name of the server.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupInspectionProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInspectionProfileArgs)(nil)).Elem()
}

// A collection of values returned by getInspectionProfile.
type LookupInspectionProfileResultOutput struct{ *pulumi.OutputState }

func (LookupInspectionProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInspectionProfileResult)(nil)).Elem()
}

func (o LookupInspectionProfileResultOutput) ToLookupInspectionProfileResultOutput() LookupInspectionProfileResultOutput {
	return o
}

func (o LookupInspectionProfileResultOutput) ToLookupInspectionProfileResultOutputWithContext(ctx context.Context) LookupInspectionProfileResultOutput {
	return o
}

// (string)
func (o LookupInspectionProfileResultOutput) CommonGlobalOverrideActionsConfig() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInspectionProfileResult) map[string]string { return v.CommonGlobalOverrideActionsConfig }).(pulumi.StringMapOutput)
}

// (string) Types for custom controls
func (o LookupInspectionProfileResultOutput) ControlsInfos() GetInspectionProfileControlsInfoArrayOutput {
	return o.ApplyT(func(v LookupInspectionProfileResult) []GetInspectionProfileControlsInfo { return v.ControlsInfos }).(GetInspectionProfileControlsInfoArrayOutput)
}

func (o LookupInspectionProfileResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInspectionProfileResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// (string) Types for custom controls
func (o LookupInspectionProfileResultOutput) CustomControls() GetInspectionProfileCustomControlArrayOutput {
	return o.ApplyT(func(v LookupInspectionProfileResult) []GetInspectionProfileCustomControl { return v.CustomControls }).(GetInspectionProfileCustomControlArrayOutput)
}

// (string) Description of the inspection profile.
func (o LookupInspectionProfileResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInspectionProfileResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupInspectionProfileResultOutput) GlobalControlActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInspectionProfileResult) []string { return v.GlobalControlActions }).(pulumi.StringArrayOutput)
}

// (string) ID of the predefined control
func (o LookupInspectionProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInspectionProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupInspectionProfileResultOutput) IncarnationNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInspectionProfileResult) string { return v.IncarnationNumber }).(pulumi.StringOutput)
}

func (o LookupInspectionProfileResultOutput) ModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInspectionProfileResult) string { return v.ModifiedBy }).(pulumi.StringOutput)
}

func (o LookupInspectionProfileResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInspectionProfileResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

// (string)
func (o LookupInspectionProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInspectionProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

// (string) OWASP Predefined Paranoia Level. Range: [1-4], inclusive
func (o LookupInspectionProfileResultOutput) ParanoiaLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInspectionProfileResult) string { return v.ParanoiaLevel }).(pulumi.StringOutput)
}

// (string) The predefined controls
func (o LookupInspectionProfileResultOutput) PredefinedControls() GetInspectionProfilePredefinedControlArrayOutput {
	return o.ApplyT(func(v LookupInspectionProfileResult) []GetInspectionProfilePredefinedControl {
		return v.PredefinedControls
	}).(GetInspectionProfilePredefinedControlArrayOutput)
}

func (o LookupInspectionProfileResultOutput) PredefinedControlsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInspectionProfileResult) string { return v.PredefinedControlsVersion }).(pulumi.StringOutput)
}

// (string)
func (o LookupInspectionProfileResultOutput) WebSocketControls() GetInspectionProfileWebSocketControlArrayOutput {
	return o.ApplyT(func(v LookupInspectionProfileResult) []GetInspectionProfileWebSocketControl {
		return v.WebSocketControls
	}).(GetInspectionProfileWebSocketControlArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInspectionProfileResultOutput{})
}
