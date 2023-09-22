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

// Describes a Karbon Cluster
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
//			_, err := nutanix.LookupKarbonCluster(ctx, &nutanix.LookupKarbonClusterArgs{
//				KarbonClusterId: pulumi.StringRef("<YOUR-CLUSTER-ID>"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupKarbonCluster(ctx *pulumi.Context, args *LookupKarbonClusterArgs, opts ...pulumi.InvokeOption) (*LookupKarbonClusterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupKarbonClusterResult
	err := ctx.Invoke("nutanix:index/getKarbonCluster:getKarbonCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKarbonCluster.
type LookupKarbonClusterArgs struct {
	// Represents karbon cluster uuid
	KarbonClusterId *string `pulumi:"karbonClusterId"`
	// Represents the name of karbon cluster
	KarbonClusterName *string `pulumi:"karbonClusterName"`
}

// A collection of values returned by getKarbonCluster.
type LookupKarbonClusterResult struct {
	DeploymentType string `pulumi:"deploymentType"`
	// - Configuration of the node pools that the nodes in the etcd cluster belong to. The etcd nodes require a minimum of 8,192 MiB memory and 409,60 MiB disk space.
	EtcdNodePools []GetKarbonClusterEtcdNodePool `pulumi:"etcdNodePools"`
	// The provider-assigned unique ID for this managed resource.
	Id                       string  `pulumi:"id"`
	KarbonClusterId          *string `pulumi:"karbonClusterId"`
	KarbonClusterName        *string `pulumi:"karbonClusterName"`
	KubeapiServerIpv4Address string  `pulumi:"kubeapiServerIpv4Address"`
	// - Configuration of the master node pools.
	MasterNodePools []GetKarbonClusterMasterNodePool `pulumi:"masterNodePools"`
	// - Unique name of the node pool.
	Name   string `pulumi:"name"`
	Status string `pulumi:"status"`
	Uuid   string `pulumi:"uuid"`
	// - K8s version of the cluster.
	Version         string                           `pulumi:"version"`
	WorkerNodePools []GetKarbonClusterWorkerNodePool `pulumi:"workerNodePools"`
}

func LookupKarbonClusterOutput(ctx *pulumi.Context, args LookupKarbonClusterOutputArgs, opts ...pulumi.InvokeOption) LookupKarbonClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKarbonClusterResult, error) {
			args := v.(LookupKarbonClusterArgs)
			r, err := LookupKarbonCluster(ctx, &args, opts...)
			var s LookupKarbonClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKarbonClusterResultOutput)
}

// A collection of arguments for invoking getKarbonCluster.
type LookupKarbonClusterOutputArgs struct {
	// Represents karbon cluster uuid
	KarbonClusterId pulumi.StringPtrInput `pulumi:"karbonClusterId"`
	// Represents the name of karbon cluster
	KarbonClusterName pulumi.StringPtrInput `pulumi:"karbonClusterName"`
}

func (LookupKarbonClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKarbonClusterArgs)(nil)).Elem()
}

// A collection of values returned by getKarbonCluster.
type LookupKarbonClusterResultOutput struct{ *pulumi.OutputState }

func (LookupKarbonClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKarbonClusterResult)(nil)).Elem()
}

func (o LookupKarbonClusterResultOutput) ToLookupKarbonClusterResultOutput() LookupKarbonClusterResultOutput {
	return o
}

func (o LookupKarbonClusterResultOutput) ToLookupKarbonClusterResultOutputWithContext(ctx context.Context) LookupKarbonClusterResultOutput {
	return o
}

func (o LookupKarbonClusterResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupKarbonClusterResult] {
	return pulumix.Output[LookupKarbonClusterResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupKarbonClusterResultOutput) DeploymentType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKarbonClusterResult) string { return v.DeploymentType }).(pulumi.StringOutput)
}

// - Configuration of the node pools that the nodes in the etcd cluster belong to. The etcd nodes require a minimum of 8,192 MiB memory and 409,60 MiB disk space.
func (o LookupKarbonClusterResultOutput) EtcdNodePools() GetKarbonClusterEtcdNodePoolArrayOutput {
	return o.ApplyT(func(v LookupKarbonClusterResult) []GetKarbonClusterEtcdNodePool { return v.EtcdNodePools }).(GetKarbonClusterEtcdNodePoolArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupKarbonClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKarbonClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupKarbonClusterResultOutput) KarbonClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKarbonClusterResult) *string { return v.KarbonClusterId }).(pulumi.StringPtrOutput)
}

func (o LookupKarbonClusterResultOutput) KarbonClusterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKarbonClusterResult) *string { return v.KarbonClusterName }).(pulumi.StringPtrOutput)
}

func (o LookupKarbonClusterResultOutput) KubeapiServerIpv4Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKarbonClusterResult) string { return v.KubeapiServerIpv4Address }).(pulumi.StringOutput)
}

// - Configuration of the master node pools.
func (o LookupKarbonClusterResultOutput) MasterNodePools() GetKarbonClusterMasterNodePoolArrayOutput {
	return o.ApplyT(func(v LookupKarbonClusterResult) []GetKarbonClusterMasterNodePool { return v.MasterNodePools }).(GetKarbonClusterMasterNodePoolArrayOutput)
}

// - Unique name of the node pool.
func (o LookupKarbonClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKarbonClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupKarbonClusterResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKarbonClusterResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupKarbonClusterResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKarbonClusterResult) string { return v.Uuid }).(pulumi.StringOutput)
}

// - K8s version of the cluster.
func (o LookupKarbonClusterResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKarbonClusterResult) string { return v.Version }).(pulumi.StringOutput)
}

func (o LookupKarbonClusterResultOutput) WorkerNodePools() GetKarbonClusterWorkerNodePoolArrayOutput {
	return o.ApplyT(func(v LookupKarbonClusterResult) []GetKarbonClusterWorkerNodePool { return v.WorkerNodePools }).(GetKarbonClusterWorkerNodePoolArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKarbonClusterResultOutput{})
}