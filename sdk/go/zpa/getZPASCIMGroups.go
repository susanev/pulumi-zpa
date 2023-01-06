// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zpa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the **zpa_scim_groups** data source to get information about a SCIM Group from an Identity Provider (IdP). This data source can then be referenced in an Access Policy, Timeout policy, Forwarding Policy, Inspection Policy or Isolation Policy.
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
//			_, err := zpa.GetZPASCIMGroups(ctx, &zpa.GetZPASCIMGroupsArgs{
//				IdpName: pulumi.StringRef("idp_name"),
//				Name:    pulumi.StringRef("Engineering"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetZPASCIMGroups(ctx *pulumi.Context, args *GetZPASCIMGroupsArgs, opts ...pulumi.InvokeOption) (*GetZPASCIMGroupsResult, error) {
	var rv GetZPASCIMGroupsResult
	err := ctx.Invoke("zpa:index/getZPASCIMGroups:getZPASCIMGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getZPASCIMGroups.
type GetZPASCIMGroupsArgs struct {
	Id *string `pulumi:"id"`
	// (string) The ID of the IdP corresponding to the SAML attribute.
	IdpId *int `pulumi:"idpId"`
	// Name. The name of the IdP where the scim group must be exported from.
	IdpName *string `pulumi:"idpName"`
	// Name. The name of the scim group to be exported.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getZPASCIMGroups.
type GetZPASCIMGroupsResult struct {
	// (string)
	CreationTime int     `pulumi:"creationTime"`
	Id           *string `pulumi:"id"`
	// (string)
	IdpGroupId string `pulumi:"idpGroupId"`
	// (string) The ID of the IdP corresponding to the SAML attribute.
	IdpId   *int    `pulumi:"idpId"`
	IdpName *string `pulumi:"idpName"`
	// (string)
	ModifiedTime int     `pulumi:"modifiedTime"`
	Name         *string `pulumi:"name"`
}

func GetZPASCIMGroupsOutput(ctx *pulumi.Context, args GetZPASCIMGroupsOutputArgs, opts ...pulumi.InvokeOption) GetZPASCIMGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetZPASCIMGroupsResult, error) {
			args := v.(GetZPASCIMGroupsArgs)
			r, err := GetZPASCIMGroups(ctx, &args, opts...)
			var s GetZPASCIMGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetZPASCIMGroupsResultOutput)
}

// A collection of arguments for invoking getZPASCIMGroups.
type GetZPASCIMGroupsOutputArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
	// (string) The ID of the IdP corresponding to the SAML attribute.
	IdpId pulumi.IntPtrInput `pulumi:"idpId"`
	// Name. The name of the IdP where the scim group must be exported from.
	IdpName pulumi.StringPtrInput `pulumi:"idpName"`
	// Name. The name of the scim group to be exported.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetZPASCIMGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZPASCIMGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getZPASCIMGroups.
type GetZPASCIMGroupsResultOutput struct{ *pulumi.OutputState }

func (GetZPASCIMGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZPASCIMGroupsResult)(nil)).Elem()
}

func (o GetZPASCIMGroupsResultOutput) ToGetZPASCIMGroupsResultOutput() GetZPASCIMGroupsResultOutput {
	return o
}

func (o GetZPASCIMGroupsResultOutput) ToGetZPASCIMGroupsResultOutputWithContext(ctx context.Context) GetZPASCIMGroupsResultOutput {
	return o
}

// (string)
func (o GetZPASCIMGroupsResultOutput) CreationTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetZPASCIMGroupsResult) int { return v.CreationTime }).(pulumi.IntOutput)
}

func (o GetZPASCIMGroupsResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZPASCIMGroupsResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// (string)
func (o GetZPASCIMGroupsResultOutput) IdpGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v GetZPASCIMGroupsResult) string { return v.IdpGroupId }).(pulumi.StringOutput)
}

// (string) The ID of the IdP corresponding to the SAML attribute.
func (o GetZPASCIMGroupsResultOutput) IdpId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetZPASCIMGroupsResult) *int { return v.IdpId }).(pulumi.IntPtrOutput)
}

func (o GetZPASCIMGroupsResultOutput) IdpName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZPASCIMGroupsResult) *string { return v.IdpName }).(pulumi.StringPtrOutput)
}

// (string)
func (o GetZPASCIMGroupsResultOutput) ModifiedTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetZPASCIMGroupsResult) int { return v.ModifiedTime }).(pulumi.IntOutput)
}

func (o GetZPASCIMGroupsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZPASCIMGroupsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetZPASCIMGroupsResultOutput{})
}
