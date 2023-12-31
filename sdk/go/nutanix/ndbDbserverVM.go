// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nutanix

import (
	"context"
	"reflect"

	"errors"
	"github.com/jacobbmay/pulumi-nutanix/sdk/go/nutanix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type NdbDbserverVM struct {
	pulumi.CustomResourceState

	ClientId          pulumi.StringOutput                      `pulumi:"clientId"`
	ComputeProfileId  pulumi.StringOutput                      `pulumi:"computeProfileId"`
	Credentials       NdbDbserverVMCredentialArrayOutput       `pulumi:"credentials"`
	DatabaseType      pulumi.StringOutput                      `pulumi:"databaseType"`
	DbserverClusterId pulumi.StringOutput                      `pulumi:"dbserverClusterId"`
	Delete            pulumi.BoolPtrOutput                     `pulumi:"delete"`
	DeleteVgs         pulumi.BoolPtrOutput                     `pulumi:"deleteVgs"`
	DeleteVmSnapshots pulumi.BoolPtrOutput                     `pulumi:"deleteVmSnapshots"`
	Description       pulumi.StringOutput                      `pulumi:"description"`
	EraDriveId        pulumi.StringOutput                      `pulumi:"eraDriveId"`
	EraVersion        pulumi.StringOutput                      `pulumi:"eraVersion"`
	Fqdns             pulumi.StringOutput                      `pulumi:"fqdns"`
	IpAddresses       pulumi.StringArrayOutput                 `pulumi:"ipAddresses"`
	LatestSnapshot    pulumi.BoolPtrOutput                     `pulumi:"latestSnapshot"`
	MacAddresses      pulumi.StringArrayOutput                 `pulumi:"macAddresses"`
	MaintenanceTasks  NdbDbserverVMMaintenanceTasksPtrOutput   `pulumi:"maintenanceTasks"`
	Name              pulumi.StringOutput                      `pulumi:"name"`
	NetworkProfileId  pulumi.StringOutput                      `pulumi:"networkProfileId"`
	NxClusterId       pulumi.StringOutput                      `pulumi:"nxClusterId"`
	PostgresDatabases NdbDbserverVMPostgresDatabaseArrayOutput `pulumi:"postgresDatabases"`
	// List of all the properties
	Properties               NdbDbserverVMPropertyArrayOutput `pulumi:"properties"`
	Remove                   pulumi.BoolPtrOutput             `pulumi:"remove"`
	SnapshotId               pulumi.StringPtrOutput           `pulumi:"snapshotId"`
	SoftRemove               pulumi.BoolPtrOutput             `pulumi:"softRemove"`
	SoftwareProfileId        pulumi.StringPtrOutput           `pulumi:"softwareProfileId"`
	SoftwareProfileVersionId pulumi.StringPtrOutput           `pulumi:"softwareProfileVersionId"`
	Status                   pulumi.StringOutput              `pulumi:"status"`
	Tags                     NdbDbserverVMTagArrayOutput      `pulumi:"tags"`
	TimeMachineId            pulumi.StringPtrOutput           `pulumi:"timeMachineId"`
	Timezone                 pulumi.StringPtrOutput           `pulumi:"timezone"`
	Type                     pulumi.StringOutput              `pulumi:"type"`
	VmClusterName            pulumi.StringOutput              `pulumi:"vmClusterName"`
	VmClusterUuid            pulumi.StringOutput              `pulumi:"vmClusterUuid"`
	VmPassword               pulumi.StringPtrOutput           `pulumi:"vmPassword"`
	VmTimezone               pulumi.StringOutput              `pulumi:"vmTimezone"`
}

// NewNdbDbserverVM registers a new resource with the given unique name, arguments, and options.
func NewNdbDbserverVM(ctx *pulumi.Context,
	name string, args *NdbDbserverVMArgs, opts ...pulumi.ResourceOption) (*NdbDbserverVM, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComputeProfileId == nil {
		return nil, errors.New("invalid value for required argument 'ComputeProfileId'")
	}
	if args.DatabaseType == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseType'")
	}
	if args.NetworkProfileId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkProfileId'")
	}
	if args.NxClusterId == nil {
		return nil, errors.New("invalid value for required argument 'NxClusterId'")
	}
	if args.VmPassword != nil {
		args.VmPassword = pulumi.ToSecret(args.VmPassword).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"vmPassword",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NdbDbserverVM
	err := ctx.RegisterResource("nutanix:index/ndbDbserverVM:NdbDbserverVM", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNdbDbserverVM gets an existing NdbDbserverVM resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNdbDbserverVM(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NdbDbserverVMState, opts ...pulumi.ResourceOption) (*NdbDbserverVM, error) {
	var resource NdbDbserverVM
	err := ctx.ReadResource("nutanix:index/ndbDbserverVM:NdbDbserverVM", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NdbDbserverVM resources.
type ndbDbserverVMState struct {
	ClientId          *string                         `pulumi:"clientId"`
	ComputeProfileId  *string                         `pulumi:"computeProfileId"`
	Credentials       []NdbDbserverVMCredential       `pulumi:"credentials"`
	DatabaseType      *string                         `pulumi:"databaseType"`
	DbserverClusterId *string                         `pulumi:"dbserverClusterId"`
	Delete            *bool                           `pulumi:"delete"`
	DeleteVgs         *bool                           `pulumi:"deleteVgs"`
	DeleteVmSnapshots *bool                           `pulumi:"deleteVmSnapshots"`
	Description       *string                         `pulumi:"description"`
	EraDriveId        *string                         `pulumi:"eraDriveId"`
	EraVersion        *string                         `pulumi:"eraVersion"`
	Fqdns             *string                         `pulumi:"fqdns"`
	IpAddresses       []string                        `pulumi:"ipAddresses"`
	LatestSnapshot    *bool                           `pulumi:"latestSnapshot"`
	MacAddresses      []string                        `pulumi:"macAddresses"`
	MaintenanceTasks  *NdbDbserverVMMaintenanceTasks  `pulumi:"maintenanceTasks"`
	Name              *string                         `pulumi:"name"`
	NetworkProfileId  *string                         `pulumi:"networkProfileId"`
	NxClusterId       *string                         `pulumi:"nxClusterId"`
	PostgresDatabases []NdbDbserverVMPostgresDatabase `pulumi:"postgresDatabases"`
	// List of all the properties
	Properties               []NdbDbserverVMProperty `pulumi:"properties"`
	Remove                   *bool                   `pulumi:"remove"`
	SnapshotId               *string                 `pulumi:"snapshotId"`
	SoftRemove               *bool                   `pulumi:"softRemove"`
	SoftwareProfileId        *string                 `pulumi:"softwareProfileId"`
	SoftwareProfileVersionId *string                 `pulumi:"softwareProfileVersionId"`
	Status                   *string                 `pulumi:"status"`
	Tags                     []NdbDbserverVMTag      `pulumi:"tags"`
	TimeMachineId            *string                 `pulumi:"timeMachineId"`
	Timezone                 *string                 `pulumi:"timezone"`
	Type                     *string                 `pulumi:"type"`
	VmClusterName            *string                 `pulumi:"vmClusterName"`
	VmClusterUuid            *string                 `pulumi:"vmClusterUuid"`
	VmPassword               *string                 `pulumi:"vmPassword"`
	VmTimezone               *string                 `pulumi:"vmTimezone"`
}

type NdbDbserverVMState struct {
	ClientId          pulumi.StringPtrInput
	ComputeProfileId  pulumi.StringPtrInput
	Credentials       NdbDbserverVMCredentialArrayInput
	DatabaseType      pulumi.StringPtrInput
	DbserverClusterId pulumi.StringPtrInput
	Delete            pulumi.BoolPtrInput
	DeleteVgs         pulumi.BoolPtrInput
	DeleteVmSnapshots pulumi.BoolPtrInput
	Description       pulumi.StringPtrInput
	EraDriveId        pulumi.StringPtrInput
	EraVersion        pulumi.StringPtrInput
	Fqdns             pulumi.StringPtrInput
	IpAddresses       pulumi.StringArrayInput
	LatestSnapshot    pulumi.BoolPtrInput
	MacAddresses      pulumi.StringArrayInput
	MaintenanceTasks  NdbDbserverVMMaintenanceTasksPtrInput
	Name              pulumi.StringPtrInput
	NetworkProfileId  pulumi.StringPtrInput
	NxClusterId       pulumi.StringPtrInput
	PostgresDatabases NdbDbserverVMPostgresDatabaseArrayInput
	// List of all the properties
	Properties               NdbDbserverVMPropertyArrayInput
	Remove                   pulumi.BoolPtrInput
	SnapshotId               pulumi.StringPtrInput
	SoftRemove               pulumi.BoolPtrInput
	SoftwareProfileId        pulumi.StringPtrInput
	SoftwareProfileVersionId pulumi.StringPtrInput
	Status                   pulumi.StringPtrInput
	Tags                     NdbDbserverVMTagArrayInput
	TimeMachineId            pulumi.StringPtrInput
	Timezone                 pulumi.StringPtrInput
	Type                     pulumi.StringPtrInput
	VmClusterName            pulumi.StringPtrInput
	VmClusterUuid            pulumi.StringPtrInput
	VmPassword               pulumi.StringPtrInput
	VmTimezone               pulumi.StringPtrInput
}

func (NdbDbserverVMState) ElementType() reflect.Type {
	return reflect.TypeOf((*ndbDbserverVMState)(nil)).Elem()
}

type ndbDbserverVMArgs struct {
	ComputeProfileId         string                          `pulumi:"computeProfileId"`
	Credentials              []NdbDbserverVMCredential       `pulumi:"credentials"`
	DatabaseType             string                          `pulumi:"databaseType"`
	Delete                   *bool                           `pulumi:"delete"`
	DeleteVgs                *bool                           `pulumi:"deleteVgs"`
	DeleteVmSnapshots        *bool                           `pulumi:"deleteVmSnapshots"`
	Description              *string                         `pulumi:"description"`
	LatestSnapshot           *bool                           `pulumi:"latestSnapshot"`
	MaintenanceTasks         *NdbDbserverVMMaintenanceTasks  `pulumi:"maintenanceTasks"`
	NetworkProfileId         string                          `pulumi:"networkProfileId"`
	NxClusterId              string                          `pulumi:"nxClusterId"`
	PostgresDatabases        []NdbDbserverVMPostgresDatabase `pulumi:"postgresDatabases"`
	Remove                   *bool                           `pulumi:"remove"`
	SnapshotId               *string                         `pulumi:"snapshotId"`
	SoftRemove               *bool                           `pulumi:"softRemove"`
	SoftwareProfileId        *string                         `pulumi:"softwareProfileId"`
	SoftwareProfileVersionId *string                         `pulumi:"softwareProfileVersionId"`
	Tags                     []NdbDbserverVMTag              `pulumi:"tags"`
	TimeMachineId            *string                         `pulumi:"timeMachineId"`
	Timezone                 *string                         `pulumi:"timezone"`
	VmPassword               *string                         `pulumi:"vmPassword"`
}

// The set of arguments for constructing a NdbDbserverVM resource.
type NdbDbserverVMArgs struct {
	ComputeProfileId         pulumi.StringInput
	Credentials              NdbDbserverVMCredentialArrayInput
	DatabaseType             pulumi.StringInput
	Delete                   pulumi.BoolPtrInput
	DeleteVgs                pulumi.BoolPtrInput
	DeleteVmSnapshots        pulumi.BoolPtrInput
	Description              pulumi.StringPtrInput
	LatestSnapshot           pulumi.BoolPtrInput
	MaintenanceTasks         NdbDbserverVMMaintenanceTasksPtrInput
	NetworkProfileId         pulumi.StringInput
	NxClusterId              pulumi.StringInput
	PostgresDatabases        NdbDbserverVMPostgresDatabaseArrayInput
	Remove                   pulumi.BoolPtrInput
	SnapshotId               pulumi.StringPtrInput
	SoftRemove               pulumi.BoolPtrInput
	SoftwareProfileId        pulumi.StringPtrInput
	SoftwareProfileVersionId pulumi.StringPtrInput
	Tags                     NdbDbserverVMTagArrayInput
	TimeMachineId            pulumi.StringPtrInput
	Timezone                 pulumi.StringPtrInput
	VmPassword               pulumi.StringPtrInput
}

func (NdbDbserverVMArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ndbDbserverVMArgs)(nil)).Elem()
}

type NdbDbserverVMInput interface {
	pulumi.Input

	ToNdbDbserverVMOutput() NdbDbserverVMOutput
	ToNdbDbserverVMOutputWithContext(ctx context.Context) NdbDbserverVMOutput
}

func (*NdbDbserverVM) ElementType() reflect.Type {
	return reflect.TypeOf((**NdbDbserverVM)(nil)).Elem()
}

func (i *NdbDbserverVM) ToNdbDbserverVMOutput() NdbDbserverVMOutput {
	return i.ToNdbDbserverVMOutputWithContext(context.Background())
}

func (i *NdbDbserverVM) ToNdbDbserverVMOutputWithContext(ctx context.Context) NdbDbserverVMOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NdbDbserverVMOutput)
}

func (i *NdbDbserverVM) ToOutput(ctx context.Context) pulumix.Output[*NdbDbserverVM] {
	return pulumix.Output[*NdbDbserverVM]{
		OutputState: i.ToNdbDbserverVMOutputWithContext(ctx).OutputState,
	}
}

// NdbDbserverVMArrayInput is an input type that accepts NdbDbserverVMArray and NdbDbserverVMArrayOutput values.
// You can construct a concrete instance of `NdbDbserverVMArrayInput` via:
//
//	NdbDbserverVMArray{ NdbDbserverVMArgs{...} }
type NdbDbserverVMArrayInput interface {
	pulumi.Input

	ToNdbDbserverVMArrayOutput() NdbDbserverVMArrayOutput
	ToNdbDbserverVMArrayOutputWithContext(context.Context) NdbDbserverVMArrayOutput
}

type NdbDbserverVMArray []NdbDbserverVMInput

func (NdbDbserverVMArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NdbDbserverVM)(nil)).Elem()
}

func (i NdbDbserverVMArray) ToNdbDbserverVMArrayOutput() NdbDbserverVMArrayOutput {
	return i.ToNdbDbserverVMArrayOutputWithContext(context.Background())
}

func (i NdbDbserverVMArray) ToNdbDbserverVMArrayOutputWithContext(ctx context.Context) NdbDbserverVMArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NdbDbserverVMArrayOutput)
}

func (i NdbDbserverVMArray) ToOutput(ctx context.Context) pulumix.Output[[]*NdbDbserverVM] {
	return pulumix.Output[[]*NdbDbserverVM]{
		OutputState: i.ToNdbDbserverVMArrayOutputWithContext(ctx).OutputState,
	}
}

// NdbDbserverVMMapInput is an input type that accepts NdbDbserverVMMap and NdbDbserverVMMapOutput values.
// You can construct a concrete instance of `NdbDbserverVMMapInput` via:
//
//	NdbDbserverVMMap{ "key": NdbDbserverVMArgs{...} }
type NdbDbserverVMMapInput interface {
	pulumi.Input

	ToNdbDbserverVMMapOutput() NdbDbserverVMMapOutput
	ToNdbDbserverVMMapOutputWithContext(context.Context) NdbDbserverVMMapOutput
}

type NdbDbserverVMMap map[string]NdbDbserverVMInput

func (NdbDbserverVMMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NdbDbserverVM)(nil)).Elem()
}

func (i NdbDbserverVMMap) ToNdbDbserverVMMapOutput() NdbDbserverVMMapOutput {
	return i.ToNdbDbserverVMMapOutputWithContext(context.Background())
}

func (i NdbDbserverVMMap) ToNdbDbserverVMMapOutputWithContext(ctx context.Context) NdbDbserverVMMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NdbDbserverVMMapOutput)
}

func (i NdbDbserverVMMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*NdbDbserverVM] {
	return pulumix.Output[map[string]*NdbDbserverVM]{
		OutputState: i.ToNdbDbserverVMMapOutputWithContext(ctx).OutputState,
	}
}

type NdbDbserverVMOutput struct{ *pulumi.OutputState }

func (NdbDbserverVMOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NdbDbserverVM)(nil)).Elem()
}

func (o NdbDbserverVMOutput) ToNdbDbserverVMOutput() NdbDbserverVMOutput {
	return o
}

func (o NdbDbserverVMOutput) ToNdbDbserverVMOutputWithContext(ctx context.Context) NdbDbserverVMOutput {
	return o
}

func (o NdbDbserverVMOutput) ToOutput(ctx context.Context) pulumix.Output[*NdbDbserverVM] {
	return pulumix.Output[*NdbDbserverVM]{
		OutputState: o.OutputState,
	}
}

func (o NdbDbserverVMOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

func (o NdbDbserverVMOutput) ComputeProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.StringOutput { return v.ComputeProfileId }).(pulumi.StringOutput)
}

func (o NdbDbserverVMOutput) Credentials() NdbDbserverVMCredentialArrayOutput {
	return o.ApplyT(func(v *NdbDbserverVM) NdbDbserverVMCredentialArrayOutput { return v.Credentials }).(NdbDbserverVMCredentialArrayOutput)
}

func (o NdbDbserverVMOutput) DatabaseType() pulumi.StringOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.StringOutput { return v.DatabaseType }).(pulumi.StringOutput)
}

func (o NdbDbserverVMOutput) DbserverClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.StringOutput { return v.DbserverClusterId }).(pulumi.StringOutput)
}

func (o NdbDbserverVMOutput) Delete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.BoolPtrOutput { return v.Delete }).(pulumi.BoolPtrOutput)
}

func (o NdbDbserverVMOutput) DeleteVgs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.BoolPtrOutput { return v.DeleteVgs }).(pulumi.BoolPtrOutput)
}

func (o NdbDbserverVMOutput) DeleteVmSnapshots() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.BoolPtrOutput { return v.DeleteVmSnapshots }).(pulumi.BoolPtrOutput)
}

func (o NdbDbserverVMOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o NdbDbserverVMOutput) EraDriveId() pulumi.StringOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.StringOutput { return v.EraDriveId }).(pulumi.StringOutput)
}

func (o NdbDbserverVMOutput) EraVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.StringOutput { return v.EraVersion }).(pulumi.StringOutput)
}

func (o NdbDbserverVMOutput) Fqdns() pulumi.StringOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.StringOutput { return v.Fqdns }).(pulumi.StringOutput)
}

func (o NdbDbserverVMOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.StringArrayOutput { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

func (o NdbDbserverVMOutput) LatestSnapshot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.BoolPtrOutput { return v.LatestSnapshot }).(pulumi.BoolPtrOutput)
}

func (o NdbDbserverVMOutput) MacAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.StringArrayOutput { return v.MacAddresses }).(pulumi.StringArrayOutput)
}

func (o NdbDbserverVMOutput) MaintenanceTasks() NdbDbserverVMMaintenanceTasksPtrOutput {
	return o.ApplyT(func(v *NdbDbserverVM) NdbDbserverVMMaintenanceTasksPtrOutput { return v.MaintenanceTasks }).(NdbDbserverVMMaintenanceTasksPtrOutput)
}

func (o NdbDbserverVMOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NdbDbserverVMOutput) NetworkProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.StringOutput { return v.NetworkProfileId }).(pulumi.StringOutput)
}

func (o NdbDbserverVMOutput) NxClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.StringOutput { return v.NxClusterId }).(pulumi.StringOutput)
}

func (o NdbDbserverVMOutput) PostgresDatabases() NdbDbserverVMPostgresDatabaseArrayOutput {
	return o.ApplyT(func(v *NdbDbserverVM) NdbDbserverVMPostgresDatabaseArrayOutput { return v.PostgresDatabases }).(NdbDbserverVMPostgresDatabaseArrayOutput)
}

// List of all the properties
func (o NdbDbserverVMOutput) Properties() NdbDbserverVMPropertyArrayOutput {
	return o.ApplyT(func(v *NdbDbserverVM) NdbDbserverVMPropertyArrayOutput { return v.Properties }).(NdbDbserverVMPropertyArrayOutput)
}

func (o NdbDbserverVMOutput) Remove() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.BoolPtrOutput { return v.Remove }).(pulumi.BoolPtrOutput)
}

func (o NdbDbserverVMOutput) SnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.StringPtrOutput { return v.SnapshotId }).(pulumi.StringPtrOutput)
}

func (o NdbDbserverVMOutput) SoftRemove() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.BoolPtrOutput { return v.SoftRemove }).(pulumi.BoolPtrOutput)
}

func (o NdbDbserverVMOutput) SoftwareProfileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.StringPtrOutput { return v.SoftwareProfileId }).(pulumi.StringPtrOutput)
}

func (o NdbDbserverVMOutput) SoftwareProfileVersionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.StringPtrOutput { return v.SoftwareProfileVersionId }).(pulumi.StringPtrOutput)
}

func (o NdbDbserverVMOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o NdbDbserverVMOutput) Tags() NdbDbserverVMTagArrayOutput {
	return o.ApplyT(func(v *NdbDbserverVM) NdbDbserverVMTagArrayOutput { return v.Tags }).(NdbDbserverVMTagArrayOutput)
}

func (o NdbDbserverVMOutput) TimeMachineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.StringPtrOutput { return v.TimeMachineId }).(pulumi.StringPtrOutput)
}

func (o NdbDbserverVMOutput) Timezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.StringPtrOutput { return v.Timezone }).(pulumi.StringPtrOutput)
}

func (o NdbDbserverVMOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o NdbDbserverVMOutput) VmClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.StringOutput { return v.VmClusterName }).(pulumi.StringOutput)
}

func (o NdbDbserverVMOutput) VmClusterUuid() pulumi.StringOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.StringOutput { return v.VmClusterUuid }).(pulumi.StringOutput)
}

func (o NdbDbserverVMOutput) VmPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.StringPtrOutput { return v.VmPassword }).(pulumi.StringPtrOutput)
}

func (o NdbDbserverVMOutput) VmTimezone() pulumi.StringOutput {
	return o.ApplyT(func(v *NdbDbserverVM) pulumi.StringOutput { return v.VmTimezone }).(pulumi.StringOutput)
}

type NdbDbserverVMArrayOutput struct{ *pulumi.OutputState }

func (NdbDbserverVMArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NdbDbserverVM)(nil)).Elem()
}

func (o NdbDbserverVMArrayOutput) ToNdbDbserverVMArrayOutput() NdbDbserverVMArrayOutput {
	return o
}

func (o NdbDbserverVMArrayOutput) ToNdbDbserverVMArrayOutputWithContext(ctx context.Context) NdbDbserverVMArrayOutput {
	return o
}

func (o NdbDbserverVMArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*NdbDbserverVM] {
	return pulumix.Output[[]*NdbDbserverVM]{
		OutputState: o.OutputState,
	}
}

func (o NdbDbserverVMArrayOutput) Index(i pulumi.IntInput) NdbDbserverVMOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NdbDbserverVM {
		return vs[0].([]*NdbDbserverVM)[vs[1].(int)]
	}).(NdbDbserverVMOutput)
}

type NdbDbserverVMMapOutput struct{ *pulumi.OutputState }

func (NdbDbserverVMMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NdbDbserverVM)(nil)).Elem()
}

func (o NdbDbserverVMMapOutput) ToNdbDbserverVMMapOutput() NdbDbserverVMMapOutput {
	return o
}

func (o NdbDbserverVMMapOutput) ToNdbDbserverVMMapOutputWithContext(ctx context.Context) NdbDbserverVMMapOutput {
	return o
}

func (o NdbDbserverVMMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*NdbDbserverVM] {
	return pulumix.Output[map[string]*NdbDbserverVM]{
		OutputState: o.OutputState,
	}
}

func (o NdbDbserverVMMapOutput) MapIndex(k pulumi.StringInput) NdbDbserverVMOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NdbDbserverVM {
		return vs[0].(map[string]*NdbDbserverVM)[vs[1].(string)]
	}).(NdbDbserverVMOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NdbDbserverVMInput)(nil)).Elem(), &NdbDbserverVM{})
	pulumi.RegisterInputType(reflect.TypeOf((*NdbDbserverVMArrayInput)(nil)).Elem(), NdbDbserverVMArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NdbDbserverVMMapInput)(nil)).Elem(), NdbDbserverVMMap{})
	pulumi.RegisterOutputType(NdbDbserverVMOutput{})
	pulumi.RegisterOutputType(NdbDbserverVMArrayOutput{})
	pulumi.RegisterOutputType(NdbDbserverVMMapOutput{})
}
