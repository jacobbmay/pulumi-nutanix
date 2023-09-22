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

// Describes the clone present in Nutanix Database Service
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
//			_, err := nutanix.LookupNdbClone(ctx, &nutanix.LookupNdbCloneArgs{
//				CloneName: pulumi.StringRef("test-inst-tf-check"),
//				Filters: []nutanix.GetNdbCloneFilter{
//					{
//						Detailed: pulumi.StringRef("true"),
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupNdbClone(ctx *pulumi.Context, args *LookupNdbCloneArgs, opts ...pulumi.InvokeOption) (*LookupNdbCloneResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNdbCloneResult
	err := ctx.Invoke("nutanix:index/getNdbClone:getNdbClone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNdbClone.
type LookupNdbCloneArgs struct {
	// Clone id
	CloneId *string `pulumi:"cloneId"`
	// Clone Name
	CloneName *string `pulumi:"cloneName"`
	// Fetches info based on filter
	Filters []GetNdbCloneFilter `pulumi:"filters"`
	// allows you to assign metadata to entities (clones, time machines, databases, and database servers) by using tags.
	Tags []GetNdbCloneTag `pulumi:"tags"`
}

// A collection of values returned by getNdbClone.
type LookupNdbCloneResult struct {
	// clone or not
	Clone     bool    `pulumi:"clone"`
	CloneId   *string `pulumi:"cloneId"`
	CloneName *string `pulumi:"cloneName"`
	// clustered or not
	Clustered bool `pulumi:"clustered"`
	// database cluster type
	DatabaseClusterType string `pulumi:"databaseClusterType"`
	// database name
	DatabaseName string `pulumi:"databaseName"`
	// database nodes associated with database instance
	DatabaseNodes []GetNdbCloneDatabaseNode `pulumi:"databaseNodes"`
	// database status
	DatabaseStatus string `pulumi:"databaseStatus"`
	// database for a cloned instance
	Databases map[string]string `pulumi:"databases"`
	// date created for clone
	DateCreated string `pulumi:"dateCreated"`
	// last modified date for clone
	DateModified string `pulumi:"dateModified"`
	// dbserver logical cluster
	DbserverLogicalCluster map[string]string `pulumi:"dbserverLogicalCluster"`
	// dbserver logical cluster id
	DbserverLogicalClusterId string `pulumi:"dbserverLogicalClusterId"`
	// cloned description
	Description string              `pulumi:"description"`
	Filters     []GetNdbCloneFilter `pulumi:"filters"`
	// cloned id
	Id string `pulumi:"id"`
	// cloned info
	Infos []GetNdbCloneInfo `pulumi:"infos"`
	// LCM Config
	LcmConfigs []GetNdbCloneLcmConfig `pulumi:"lcmConfigs"`
	// linked databases within database instance
	LinkedDatabases []GetNdbCloneLinkedDatabase `pulumi:"linkedDatabases"`
	// Metric of clone
	Metric map[string]string `pulumi:"metric"`
	// cloned name
	Name string `pulumi:"name"`
	// parent database id
	ParentDatabaseId string `pulumi:"parentDatabaseId"`
	// parent source database id
	ParentSourceDatabaseId string `pulumi:"parentSourceDatabaseId"`
	// parent time machine id
	ParentTimeMachineId string `pulumi:"parentTimeMachineId"`
	// properties of clone
	Properties []GetNdbCloneProperty `pulumi:"properties"`
	// status of clone
	Status string `pulumi:"status"`
	// allows you to assign metadata to entities (clones, time machines, databases, and database servers) by using tags.
	Tags []GetNdbCloneTag `pulumi:"tags"`
	// time machine id
	TimeMachineId string `pulumi:"timeMachineId"`
	// Time machine info
	TimeMachines []GetNdbCloneTimeMachine `pulumi:"timeMachines"`
	// time zone
	TimeZone string `pulumi:"timeZone"`
	// type
	Type string `pulumi:"type"`
}

func LookupNdbCloneOutput(ctx *pulumi.Context, args LookupNdbCloneOutputArgs, opts ...pulumi.InvokeOption) LookupNdbCloneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNdbCloneResult, error) {
			args := v.(LookupNdbCloneArgs)
			r, err := LookupNdbClone(ctx, &args, opts...)
			var s LookupNdbCloneResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNdbCloneResultOutput)
}

// A collection of arguments for invoking getNdbClone.
type LookupNdbCloneOutputArgs struct {
	// Clone id
	CloneId pulumi.StringPtrInput `pulumi:"cloneId"`
	// Clone Name
	CloneName pulumi.StringPtrInput `pulumi:"cloneName"`
	// Fetches info based on filter
	Filters GetNdbCloneFilterArrayInput `pulumi:"filters"`
	// allows you to assign metadata to entities (clones, time machines, databases, and database servers) by using tags.
	Tags GetNdbCloneTagArrayInput `pulumi:"tags"`
}

func (LookupNdbCloneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNdbCloneArgs)(nil)).Elem()
}

// A collection of values returned by getNdbClone.
type LookupNdbCloneResultOutput struct{ *pulumi.OutputState }

func (LookupNdbCloneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNdbCloneResult)(nil)).Elem()
}

func (o LookupNdbCloneResultOutput) ToLookupNdbCloneResultOutput() LookupNdbCloneResultOutput {
	return o
}

func (o LookupNdbCloneResultOutput) ToLookupNdbCloneResultOutputWithContext(ctx context.Context) LookupNdbCloneResultOutput {
	return o
}

func (o LookupNdbCloneResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupNdbCloneResult] {
	return pulumix.Output[LookupNdbCloneResult]{
		OutputState: o.OutputState,
	}
}

// clone or not
func (o LookupNdbCloneResultOutput) Clone() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) bool { return v.Clone }).(pulumi.BoolOutput)
}

func (o LookupNdbCloneResultOutput) CloneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) *string { return v.CloneId }).(pulumi.StringPtrOutput)
}

func (o LookupNdbCloneResultOutput) CloneName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) *string { return v.CloneName }).(pulumi.StringPtrOutput)
}

// clustered or not
func (o LookupNdbCloneResultOutput) Clustered() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) bool { return v.Clustered }).(pulumi.BoolOutput)
}

// database cluster type
func (o LookupNdbCloneResultOutput) DatabaseClusterType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) string { return v.DatabaseClusterType }).(pulumi.StringOutput)
}

// database name
func (o LookupNdbCloneResultOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) string { return v.DatabaseName }).(pulumi.StringOutput)
}

// database nodes associated with database instance
func (o LookupNdbCloneResultOutput) DatabaseNodes() GetNdbCloneDatabaseNodeArrayOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) []GetNdbCloneDatabaseNode { return v.DatabaseNodes }).(GetNdbCloneDatabaseNodeArrayOutput)
}

// database status
func (o LookupNdbCloneResultOutput) DatabaseStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) string { return v.DatabaseStatus }).(pulumi.StringOutput)
}

// database for a cloned instance
func (o LookupNdbCloneResultOutput) Databases() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) map[string]string { return v.Databases }).(pulumi.StringMapOutput)
}

// date created for clone
func (o LookupNdbCloneResultOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) string { return v.DateCreated }).(pulumi.StringOutput)
}

// last modified date for clone
func (o LookupNdbCloneResultOutput) DateModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) string { return v.DateModified }).(pulumi.StringOutput)
}

// dbserver logical cluster
func (o LookupNdbCloneResultOutput) DbserverLogicalCluster() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) map[string]string { return v.DbserverLogicalCluster }).(pulumi.StringMapOutput)
}

// dbserver logical cluster id
func (o LookupNdbCloneResultOutput) DbserverLogicalClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) string { return v.DbserverLogicalClusterId }).(pulumi.StringOutput)
}

// cloned description
func (o LookupNdbCloneResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupNdbCloneResultOutput) Filters() GetNdbCloneFilterArrayOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) []GetNdbCloneFilter { return v.Filters }).(GetNdbCloneFilterArrayOutput)
}

// cloned id
func (o LookupNdbCloneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) string { return v.Id }).(pulumi.StringOutput)
}

// cloned info
func (o LookupNdbCloneResultOutput) Infos() GetNdbCloneInfoArrayOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) []GetNdbCloneInfo { return v.Infos }).(GetNdbCloneInfoArrayOutput)
}

// LCM Config
func (o LookupNdbCloneResultOutput) LcmConfigs() GetNdbCloneLcmConfigArrayOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) []GetNdbCloneLcmConfig { return v.LcmConfigs }).(GetNdbCloneLcmConfigArrayOutput)
}

// linked databases within database instance
func (o LookupNdbCloneResultOutput) LinkedDatabases() GetNdbCloneLinkedDatabaseArrayOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) []GetNdbCloneLinkedDatabase { return v.LinkedDatabases }).(GetNdbCloneLinkedDatabaseArrayOutput)
}

// Metric of clone
func (o LookupNdbCloneResultOutput) Metric() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) map[string]string { return v.Metric }).(pulumi.StringMapOutput)
}

// cloned name
func (o LookupNdbCloneResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) string { return v.Name }).(pulumi.StringOutput)
}

// parent database id
func (o LookupNdbCloneResultOutput) ParentDatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) string { return v.ParentDatabaseId }).(pulumi.StringOutput)
}

// parent source database id
func (o LookupNdbCloneResultOutput) ParentSourceDatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) string { return v.ParentSourceDatabaseId }).(pulumi.StringOutput)
}

// parent time machine id
func (o LookupNdbCloneResultOutput) ParentTimeMachineId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) string { return v.ParentTimeMachineId }).(pulumi.StringOutput)
}

// properties of clone
func (o LookupNdbCloneResultOutput) Properties() GetNdbClonePropertyArrayOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) []GetNdbCloneProperty { return v.Properties }).(GetNdbClonePropertyArrayOutput)
}

// status of clone
func (o LookupNdbCloneResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) string { return v.Status }).(pulumi.StringOutput)
}

// allows you to assign metadata to entities (clones, time machines, databases, and database servers) by using tags.
func (o LookupNdbCloneResultOutput) Tags() GetNdbCloneTagArrayOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) []GetNdbCloneTag { return v.Tags }).(GetNdbCloneTagArrayOutput)
}

// time machine id
func (o LookupNdbCloneResultOutput) TimeMachineId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) string { return v.TimeMachineId }).(pulumi.StringOutput)
}

// Time machine info
func (o LookupNdbCloneResultOutput) TimeMachines() GetNdbCloneTimeMachineArrayOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) []GetNdbCloneTimeMachine { return v.TimeMachines }).(GetNdbCloneTimeMachineArrayOutput)
}

// time zone
func (o LookupNdbCloneResultOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) string { return v.TimeZone }).(pulumi.StringOutput)
}

// type
func (o LookupNdbCloneResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNdbCloneResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNdbCloneResultOutput{})
}