// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appconnectorgroup

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
//	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/AppConnectorGroup"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := AppConnectorGroup.GetAppConnectorGroup(ctx, &appconnectorgroup.GetAppConnectorGroupArgs{
//				Name: pulumi.StringRef("DataCenter"),
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
//	"github.com/zscaler/pulumi-zpa/sdk/go/zpa/AppConnectorGroup"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := AppConnectorGroup.GetAppConnectorGroup(ctx, &appconnectorgroup.GetAppConnectorGroupArgs{
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
func GetAppConnectorGroup(ctx *pulumi.Context, args *GetAppConnectorGroupArgs, opts ...pulumi.InvokeOption) (*GetAppConnectorGroupResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetAppConnectorGroupResult
	err := ctx.Invoke("zpa:AppConnectorGroup/getAppConnectorGroup:getAppConnectorGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppConnectorGroup.
type GetAppConnectorGroupArgs struct {
	// ID of the App Connector Group.
	Id *string `pulumi:"id"`
	// Name of the App Connector Group.
	Name *string `pulumi:"name"`
	// (bool) Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
	OverrideVersionProfile *bool `pulumi:"overrideVersionProfile"`
}

// A collection of values returned by getAppConnectorGroup.
type GetAppConnectorGroupResult struct {
	// (String) Whether Double Encryption is enabled or disabled for the app.
	CityCountry string                          `pulumi:"cityCountry"`
	Connectors  []GetAppConnectorGroupConnector `pulumi:"connectors"`
	// (String)
	CountryCode  string `pulumi:"countryCode"`
	CreationTime string `pulumi:"creationTime"`
	// (String) Description of the App Connector Group.
	Description string `pulumi:"description"`
	// (String)
	DnsQueryType string `pulumi:"dnsQueryType"`
	// (String) Whether this App Connector Group is enabled or not. Default value: `true`. Supported values: `true`, `false`
	Enabled bool `pulumi:"enabled"`
	// (String)
	GeoLocationId string  `pulumi:"geoLocationId"`
	Id            *string `pulumi:"id"`
	// (String) Latitude of the App Connector Group. Integer or decimal. With values in the range of `-90` to `90`
	Latitude string `pulumi:"latitude"`
	// (String) Location of the App Connector Group.
	Location string `pulumi:"location"`
	// (String) Longitude of the App Connector Group. Integer or decimal. With values in the range of `-180` to `180`
	Longitude            string  `pulumi:"longitude"`
	LssAppConnectorGroup bool    `pulumi:"lssAppConnectorGroup"`
	ModifiedTime         string  `pulumi:"modifiedTime"`
	Modifiedby           string  `pulumi:"modifiedby"`
	Name                 *string `pulumi:"name"`
	// (bool) Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
	OverrideVersionProfile   *bool                             `pulumi:"overrideVersionProfile"`
	ServerGroups             []GetAppConnectorGroupServerGroup `pulumi:"serverGroups"`
	TcpQuickAckApp           bool                              `pulumi:"tcpQuickAckApp"`
	TcpQuickAckAssistant     bool                              `pulumi:"tcpQuickAckAssistant"`
	TcpQuickAckReadAssistant bool                              `pulumi:"tcpQuickAckReadAssistant"`
	// (String) App Connectors in this group will attempt to update to a newer version of the software during this specified day
	UpgradeDay string `pulumi:"upgradeDay"`
	// (String) App Connectors in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600`. Integer in seconds (i.e., `-66600`). The integer should be greater than or equal to `0` and less than `86400`, in `15` minute intervals
	UpgradeTimeInSecs string `pulumi:"upgradeTimeInSecs"`
	UseInDrMode       bool   `pulumi:"useInDrMode"`
	// (String) ID of the version profile.
	// Exported values are:
	VersionProfileId string `pulumi:"versionProfileId"`
	// (String)
	// Exported values are:
	VersionProfileName string `pulumi:"versionProfileName"`
	// (String)
	// Exported values are:
	VersionProfileVisibilityScope string `pulumi:"versionProfileVisibilityScope"`
}

func GetAppConnectorGroupOutput(ctx *pulumi.Context, args GetAppConnectorGroupOutputArgs, opts ...pulumi.InvokeOption) GetAppConnectorGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppConnectorGroupResult, error) {
			args := v.(GetAppConnectorGroupArgs)
			r, err := GetAppConnectorGroup(ctx, &args, opts...)
			var s GetAppConnectorGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAppConnectorGroupResultOutput)
}

// A collection of arguments for invoking getAppConnectorGroup.
type GetAppConnectorGroupOutputArgs struct {
	// ID of the App Connector Group.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Name of the App Connector Group.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// (bool) Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
	OverrideVersionProfile pulumi.BoolPtrInput `pulumi:"overrideVersionProfile"`
}

func (GetAppConnectorGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppConnectorGroupArgs)(nil)).Elem()
}

// A collection of values returned by getAppConnectorGroup.
type GetAppConnectorGroupResultOutput struct{ *pulumi.OutputState }

func (GetAppConnectorGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppConnectorGroupResult)(nil)).Elem()
}

func (o GetAppConnectorGroupResultOutput) ToGetAppConnectorGroupResultOutput() GetAppConnectorGroupResultOutput {
	return o
}

func (o GetAppConnectorGroupResultOutput) ToGetAppConnectorGroupResultOutputWithContext(ctx context.Context) GetAppConnectorGroupResultOutput {
	return o
}

// (String) Whether Double Encryption is enabled or disabled for the app.
func (o GetAppConnectorGroupResultOutput) CityCountry() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) string { return v.CityCountry }).(pulumi.StringOutput)
}

func (o GetAppConnectorGroupResultOutput) Connectors() GetAppConnectorGroupConnectorArrayOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) []GetAppConnectorGroupConnector { return v.Connectors }).(GetAppConnectorGroupConnectorArrayOutput)
}

// (String)
func (o GetAppConnectorGroupResultOutput) CountryCode() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) string { return v.CountryCode }).(pulumi.StringOutput)
}

func (o GetAppConnectorGroupResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// (String) Description of the App Connector Group.
func (o GetAppConnectorGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

// (String)
func (o GetAppConnectorGroupResultOutput) DnsQueryType() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) string { return v.DnsQueryType }).(pulumi.StringOutput)
}

// (String) Whether this App Connector Group is enabled or not. Default value: `true`. Supported values: `true`, `false`
func (o GetAppConnectorGroupResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// (String)
func (o GetAppConnectorGroupResultOutput) GeoLocationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) string { return v.GeoLocationId }).(pulumi.StringOutput)
}

func (o GetAppConnectorGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// (String) Latitude of the App Connector Group. Integer or decimal. With values in the range of `-90` to `90`
func (o GetAppConnectorGroupResultOutput) Latitude() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) string { return v.Latitude }).(pulumi.StringOutput)
}

// (String) Location of the App Connector Group.
func (o GetAppConnectorGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

// (String) Longitude of the App Connector Group. Integer or decimal. With values in the range of `-180` to `180`
func (o GetAppConnectorGroupResultOutput) Longitude() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) string { return v.Longitude }).(pulumi.StringOutput)
}

func (o GetAppConnectorGroupResultOutput) LssAppConnectorGroup() pulumi.BoolOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) bool { return v.LssAppConnectorGroup }).(pulumi.BoolOutput)
}

func (o GetAppConnectorGroupResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

func (o GetAppConnectorGroupResultOutput) Modifiedby() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) string { return v.Modifiedby }).(pulumi.StringOutput)
}

func (o GetAppConnectorGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// (bool) Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
func (o GetAppConnectorGroupResultOutput) OverrideVersionProfile() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) *bool { return v.OverrideVersionProfile }).(pulumi.BoolPtrOutput)
}

func (o GetAppConnectorGroupResultOutput) ServerGroups() GetAppConnectorGroupServerGroupArrayOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) []GetAppConnectorGroupServerGroup { return v.ServerGroups }).(GetAppConnectorGroupServerGroupArrayOutput)
}

func (o GetAppConnectorGroupResultOutput) TcpQuickAckApp() pulumi.BoolOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) bool { return v.TcpQuickAckApp }).(pulumi.BoolOutput)
}

func (o GetAppConnectorGroupResultOutput) TcpQuickAckAssistant() pulumi.BoolOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) bool { return v.TcpQuickAckAssistant }).(pulumi.BoolOutput)
}

func (o GetAppConnectorGroupResultOutput) TcpQuickAckReadAssistant() pulumi.BoolOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) bool { return v.TcpQuickAckReadAssistant }).(pulumi.BoolOutput)
}

// (String) App Connectors in this group will attempt to update to a newer version of the software during this specified day
func (o GetAppConnectorGroupResultOutput) UpgradeDay() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) string { return v.UpgradeDay }).(pulumi.StringOutput)
}

// (String) App Connectors in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600`. Integer in seconds (i.e., `-66600`). The integer should be greater than or equal to `0` and less than `86400`, in `15` minute intervals
func (o GetAppConnectorGroupResultOutput) UpgradeTimeInSecs() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) string { return v.UpgradeTimeInSecs }).(pulumi.StringOutput)
}

func (o GetAppConnectorGroupResultOutput) UseInDrMode() pulumi.BoolOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) bool { return v.UseInDrMode }).(pulumi.BoolOutput)
}

// (String) ID of the version profile.
// Exported values are:
func (o GetAppConnectorGroupResultOutput) VersionProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) string { return v.VersionProfileId }).(pulumi.StringOutput)
}

// (String)
// Exported values are:
func (o GetAppConnectorGroupResultOutput) VersionProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) string { return v.VersionProfileName }).(pulumi.StringOutput)
}

// (String)
// Exported values are:
func (o GetAppConnectorGroupResultOutput) VersionProfileVisibilityScope() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppConnectorGroupResult) string { return v.VersionProfileVisibilityScope }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppConnectorGroupResultOutput{})
}