// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nutanix

import (
	"context"
	"reflect"

	"github.com/jacobbmay/pulumi-nutanix/sdk/go/nutanix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a datasource to retrieve VPC with vpcUuid or vpcName .
func LookupVpc(ctx *pulumi.Context, args *LookupVpcArgs, opts ...pulumi.InvokeOption) (*LookupVpcResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcResult
	err := ctx.Invoke("nutanix:index/getVpc:getVpc", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpc.
type LookupVpcArgs struct {
	// vpc Name
	VpcName *string `pulumi:"vpcName"`
	// vpc UUID
	VpcUuid *string `pulumi:"vpcUuid"`
}

// A collection of values returned by getVpc.
type LookupVpcResult struct {
	// The version of the API.
	ApiVersion string `pulumi:"apiVersion"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// - The vpc kind metadata.
	Metadata map[string]string `pulumi:"metadata"`
	// VPC input spec
	Specs []GetVpcSpec `pulumi:"specs"`
	// VPC output status
	Statuses []GetVpcStatus `pulumi:"statuses"`
	VpcName  *string        `pulumi:"vpcName"`
	VpcUuid  *string        `pulumi:"vpcUuid"`
}

func LookupVpcOutput(ctx *pulumi.Context, args LookupVpcOutputArgs, opts ...pulumi.InvokeOption) LookupVpcResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpcResult, error) {
			args := v.(LookupVpcArgs)
			r, err := LookupVpc(ctx, &args, opts...)
			var s LookupVpcResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpcResultOutput)
}

// A collection of arguments for invoking getVpc.
type LookupVpcOutputArgs struct {
	// vpc Name
	VpcName pulumi.StringPtrInput `pulumi:"vpcName"`
	// vpc UUID
	VpcUuid pulumi.StringPtrInput `pulumi:"vpcUuid"`
}

func (LookupVpcOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcArgs)(nil)).Elem()
}

// A collection of values returned by getVpc.
type LookupVpcResultOutput struct{ *pulumi.OutputState }

func (LookupVpcResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcResult)(nil)).Elem()
}

func (o LookupVpcResultOutput) ToLookupVpcResultOutput() LookupVpcResultOutput {
	return o
}

func (o LookupVpcResultOutput) ToLookupVpcResultOutputWithContext(ctx context.Context) LookupVpcResultOutput {
	return o
}

func (o LookupVpcResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVpcResult] {
	return pulumix.Output[LookupVpcResult]{
		OutputState: o.OutputState,
	}
}

// The version of the API.
func (o LookupVpcResultOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcResult) string { return v.ApiVersion }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVpcResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcResult) string { return v.Id }).(pulumi.StringOutput)
}

// - The vpc kind metadata.
func (o LookupVpcResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVpcResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

// VPC input spec
func (o LookupVpcResultOutput) Specs() GetVpcSpecArrayOutput {
	return o.ApplyT(func(v LookupVpcResult) []GetVpcSpec { return v.Specs }).(GetVpcSpecArrayOutput)
}

// VPC output status
func (o LookupVpcResultOutput) Statuses() GetVpcStatusArrayOutput {
	return o.ApplyT(func(v LookupVpcResult) []GetVpcStatus { return v.Statuses }).(GetVpcStatusArrayOutput)
}

func (o LookupVpcResultOutput) VpcName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcResult) *string { return v.VpcName }).(pulumi.StringPtrOutput)
}

func (o LookupVpcResultOutput) VpcUuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcResult) *string { return v.VpcUuid }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcResultOutput{})
}