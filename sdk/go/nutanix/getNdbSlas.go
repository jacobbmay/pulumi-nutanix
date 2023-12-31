// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nutanix

import (
	"github.com/jacobbmay/pulumi-nutanix/sdk/go/nutanix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists all SLAs in Nutanix Database Service
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/jacobbmay/pulumi-nutanix/sdk/go/nutanix"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			slas, err := nutanix.GetNdbSlas(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("sla", slas)
//			return nil
//		})
//	}
//
// ```
func GetNdbSlas(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetNdbSlasResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetNdbSlasResult
	err := ctx.Invoke("nutanix:index/getNdbSlas:getNdbSlas", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getNdbSlas.
type GetNdbSlasResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// - list of slas
	Slas []GetNdbSlasSla `pulumi:"slas"`
}
