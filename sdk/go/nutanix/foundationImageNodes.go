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

type FoundationImageNodes struct {
	pulumi.CustomResourceState

	Blocks FoundationImageNodesBlockArrayOutput `pulumi:"blocks"`
	// - list containing cluster name and cluster urls for created clusters in current session
	// * `cluster_urls.#.cluster_name` :- clusterName
	// * `cluster_urls.#.cluster_url` :- url to access the cluster login
	ClusterUrls FoundationImageNodesClusterUrlArrayOutput `pulumi:"clusterUrls"`
	Clusters    FoundationImageNodesClusterArrayOutput    `pulumi:"clusters"`
	// - (Required) CVM gateway.
	CvmGateway pulumi.StringOutput `pulumi:"cvmGateway"`
	// - (Required) CVM netmask.
	CvmNetmask pulumi.StringOutput `pulumi:"cvmNetmask"`
	// - Contains user data from Eos portal.
	EosMetadata FoundationImageNodesEosMetadataPtrOutput `pulumi:"eosMetadata"`
	// - Foundation Central specific settings.
	FcSettings FoundationImageNodesFcSettingsPtrOutput `pulumi:"fcSettings"`
	// - Hyperv External virtual network adapter name.
	HypervExternalVnic pulumi.StringPtrOutput `pulumi:"hypervExternalVnic"`
	// - Hyperv External vswitch name.
	HypervExternalVswitch pulumi.StringPtrOutput `pulumi:"hypervExternalVswitch"`
	// - Hyperv product key.
	HypervProductKey pulumi.StringPtrOutput `pulumi:"hypervProductKey"`
	// - Hyperv SKU.
	HypervSku pulumi.BoolPtrOutput `pulumi:"hypervSku"`
	// - (Required) Hypervisor gateway.
	HypervisorGateway pulumi.StringOutput `pulumi:"hypervisorGateway"`
	// - Hypervisor ISO.
	HypervisorIso        FoundationImageNodesHypervisorIsoPtrOutput `pulumi:"hypervisorIso"`
	HypervisorNameserver pulumi.StringPtrOutput                     `pulumi:"hypervisorNameserver"`
	// - (Required) Hypervisor netmask.
	HypervisorNetmask pulumi.StringOutput `pulumi:"hypervisorNetmask"`
	// - Hypervisor password.
	HypervisorPassword pulumi.StringPtrOutput `pulumi:"hypervisorPassword"`
	// - install script.
	InstallScript pulumi.StringPtrOutput `pulumi:"installScript"`
	// - (Required incase using IPMI based imaging either here or inside node spec) default IPMI gateway
	IpmiGateway pulumi.StringPtrOutput `pulumi:"ipmiGateway"`
	// - (Required incase using IPMI based imaging either here or inside node spec) default IPMI netmask
	IpmiNetmask pulumi.StringPtrOutput `pulumi:"ipmiNetmask"`
	// - (Required incase using IPMI based imaging either here or inside node spec) IPMI password.
	IpmiPassword pulumi.StringPtrOutput `pulumi:"ipmiPassword"`
	// - (Required incase using IPMI based imaging either here or inside node spec) IPMI username.
	IpmiUser pulumi.StringPtrOutput `pulumi:"ipmiUser"`
	// - Id of the custom layout which needs to be passed to imaging request.
	LayoutEggUuid pulumi.StringPtrOutput `pulumi:"layoutEggUuid"`
	// - (Required) NOS package.
	NosPackage pulumi.StringOutput `pulumi:"nosPackage"`
	// - sessionId of the imaging session
	SessionId pulumi.StringOutput `pulumi:"sessionId"`
	// - If hypervisor installation should be skipped.
	SkipHypervisor pulumi.BoolPtrOutput `pulumi:"skipHypervisor"`
	// - Arguments to be passed to svmRescue for AOS installation. Ensure that the arguments provided are supported by the AOS version used for imaging.
	SvmRescueArgs pulumi.StringArrayOutput `pulumi:"svmRescueArgs"`
	// - Types of tests to be performed.
	Tests FoundationImageNodesTestsPtrOutput `pulumi:"tests"`
	// - UCSM IP address.
	UcsmIp pulumi.StringPtrOutput `pulumi:"ucsmIp"`
	// - UCSM password.
	UcsmPassword pulumi.StringPtrOutput `pulumi:"ucsmPassword"`
	// - UCSM username.
	UcsmUser pulumi.StringPtrOutput `pulumi:"ucsmUser"`
	// - UNC password.
	UncPassword pulumi.StringPtrOutput `pulumi:"uncPassword"`
	// - UNC Path.
	UncPath pulumi.StringPtrOutput `pulumi:"uncPath"`
	// - UNC username.
	UncUsername pulumi.StringPtrOutput `pulumi:"uncUsername"`
	// - xen config types.
	XenConfigType pulumi.StringPtrOutput `pulumi:"xenConfigType"`
	// - xen server master IP address.
	XsMasterIp pulumi.StringPtrOutput `pulumi:"xsMasterIp"`
	// - xen server master label.
	XsMasterLabel pulumi.StringPtrOutput `pulumi:"xsMasterLabel"`
	// - xen server master password.
	XsMasterPassword pulumi.StringPtrOutput `pulumi:"xsMasterPassword"`
	// - xen server master username.
	XsMasterUsername pulumi.StringPtrOutput `pulumi:"xsMasterUsername"`
}

// NewFoundationImageNodes registers a new resource with the given unique name, arguments, and options.
func NewFoundationImageNodes(ctx *pulumi.Context,
	name string, args *FoundationImageNodesArgs, opts ...pulumi.ResourceOption) (*FoundationImageNodes, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Blocks == nil {
		return nil, errors.New("invalid value for required argument 'Blocks'")
	}
	if args.CvmGateway == nil {
		return nil, errors.New("invalid value for required argument 'CvmGateway'")
	}
	if args.CvmNetmask == nil {
		return nil, errors.New("invalid value for required argument 'CvmNetmask'")
	}
	if args.HypervisorGateway == nil {
		return nil, errors.New("invalid value for required argument 'HypervisorGateway'")
	}
	if args.HypervisorNetmask == nil {
		return nil, errors.New("invalid value for required argument 'HypervisorNetmask'")
	}
	if args.NosPackage == nil {
		return nil, errors.New("invalid value for required argument 'NosPackage'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FoundationImageNodes
	err := ctx.RegisterResource("nutanix:index/foundationImageNodes:FoundationImageNodes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFoundationImageNodes gets an existing FoundationImageNodes resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFoundationImageNodes(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FoundationImageNodesState, opts ...pulumi.ResourceOption) (*FoundationImageNodes, error) {
	var resource FoundationImageNodes
	err := ctx.ReadResource("nutanix:index/foundationImageNodes:FoundationImageNodes", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FoundationImageNodes resources.
type foundationImageNodesState struct {
	Blocks []FoundationImageNodesBlock `pulumi:"blocks"`
	// - list containing cluster name and cluster urls for created clusters in current session
	// * `cluster_urls.#.cluster_name` :- clusterName
	// * `cluster_urls.#.cluster_url` :- url to access the cluster login
	ClusterUrls []FoundationImageNodesClusterUrl `pulumi:"clusterUrls"`
	Clusters    []FoundationImageNodesCluster    `pulumi:"clusters"`
	// - (Required) CVM gateway.
	CvmGateway *string `pulumi:"cvmGateway"`
	// - (Required) CVM netmask.
	CvmNetmask *string `pulumi:"cvmNetmask"`
	// - Contains user data from Eos portal.
	EosMetadata *FoundationImageNodesEosMetadata `pulumi:"eosMetadata"`
	// - Foundation Central specific settings.
	FcSettings *FoundationImageNodesFcSettings `pulumi:"fcSettings"`
	// - Hyperv External virtual network adapter name.
	HypervExternalVnic *string `pulumi:"hypervExternalVnic"`
	// - Hyperv External vswitch name.
	HypervExternalVswitch *string `pulumi:"hypervExternalVswitch"`
	// - Hyperv product key.
	HypervProductKey *string `pulumi:"hypervProductKey"`
	// - Hyperv SKU.
	HypervSku *bool `pulumi:"hypervSku"`
	// - (Required) Hypervisor gateway.
	HypervisorGateway *string `pulumi:"hypervisorGateway"`
	// - Hypervisor ISO.
	HypervisorIso        *FoundationImageNodesHypervisorIso `pulumi:"hypervisorIso"`
	HypervisorNameserver *string                            `pulumi:"hypervisorNameserver"`
	// - (Required) Hypervisor netmask.
	HypervisorNetmask *string `pulumi:"hypervisorNetmask"`
	// - Hypervisor password.
	HypervisorPassword *string `pulumi:"hypervisorPassword"`
	// - install script.
	InstallScript *string `pulumi:"installScript"`
	// - (Required incase using IPMI based imaging either here or inside node spec) default IPMI gateway
	IpmiGateway *string `pulumi:"ipmiGateway"`
	// - (Required incase using IPMI based imaging either here or inside node spec) default IPMI netmask
	IpmiNetmask *string `pulumi:"ipmiNetmask"`
	// - (Required incase using IPMI based imaging either here or inside node spec) IPMI password.
	IpmiPassword *string `pulumi:"ipmiPassword"`
	// - (Required incase using IPMI based imaging either here or inside node spec) IPMI username.
	IpmiUser *string `pulumi:"ipmiUser"`
	// - Id of the custom layout which needs to be passed to imaging request.
	LayoutEggUuid *string `pulumi:"layoutEggUuid"`
	// - (Required) NOS package.
	NosPackage *string `pulumi:"nosPackage"`
	// - sessionId of the imaging session
	SessionId *string `pulumi:"sessionId"`
	// - If hypervisor installation should be skipped.
	SkipHypervisor *bool `pulumi:"skipHypervisor"`
	// - Arguments to be passed to svmRescue for AOS installation. Ensure that the arguments provided are supported by the AOS version used for imaging.
	SvmRescueArgs []string `pulumi:"svmRescueArgs"`
	// - Types of tests to be performed.
	Tests *FoundationImageNodesTests `pulumi:"tests"`
	// - UCSM IP address.
	UcsmIp *string `pulumi:"ucsmIp"`
	// - UCSM password.
	UcsmPassword *string `pulumi:"ucsmPassword"`
	// - UCSM username.
	UcsmUser *string `pulumi:"ucsmUser"`
	// - UNC password.
	UncPassword *string `pulumi:"uncPassword"`
	// - UNC Path.
	UncPath *string `pulumi:"uncPath"`
	// - UNC username.
	UncUsername *string `pulumi:"uncUsername"`
	// - xen config types.
	XenConfigType *string `pulumi:"xenConfigType"`
	// - xen server master IP address.
	XsMasterIp *string `pulumi:"xsMasterIp"`
	// - xen server master label.
	XsMasterLabel *string `pulumi:"xsMasterLabel"`
	// - xen server master password.
	XsMasterPassword *string `pulumi:"xsMasterPassword"`
	// - xen server master username.
	XsMasterUsername *string `pulumi:"xsMasterUsername"`
}

type FoundationImageNodesState struct {
	Blocks FoundationImageNodesBlockArrayInput
	// - list containing cluster name and cluster urls for created clusters in current session
	// * `cluster_urls.#.cluster_name` :- clusterName
	// * `cluster_urls.#.cluster_url` :- url to access the cluster login
	ClusterUrls FoundationImageNodesClusterUrlArrayInput
	Clusters    FoundationImageNodesClusterArrayInput
	// - (Required) CVM gateway.
	CvmGateway pulumi.StringPtrInput
	// - (Required) CVM netmask.
	CvmNetmask pulumi.StringPtrInput
	// - Contains user data from Eos portal.
	EosMetadata FoundationImageNodesEosMetadataPtrInput
	// - Foundation Central specific settings.
	FcSettings FoundationImageNodesFcSettingsPtrInput
	// - Hyperv External virtual network adapter name.
	HypervExternalVnic pulumi.StringPtrInput
	// - Hyperv External vswitch name.
	HypervExternalVswitch pulumi.StringPtrInput
	// - Hyperv product key.
	HypervProductKey pulumi.StringPtrInput
	// - Hyperv SKU.
	HypervSku pulumi.BoolPtrInput
	// - (Required) Hypervisor gateway.
	HypervisorGateway pulumi.StringPtrInput
	// - Hypervisor ISO.
	HypervisorIso        FoundationImageNodesHypervisorIsoPtrInput
	HypervisorNameserver pulumi.StringPtrInput
	// - (Required) Hypervisor netmask.
	HypervisorNetmask pulumi.StringPtrInput
	// - Hypervisor password.
	HypervisorPassword pulumi.StringPtrInput
	// - install script.
	InstallScript pulumi.StringPtrInput
	// - (Required incase using IPMI based imaging either here or inside node spec) default IPMI gateway
	IpmiGateway pulumi.StringPtrInput
	// - (Required incase using IPMI based imaging either here or inside node spec) default IPMI netmask
	IpmiNetmask pulumi.StringPtrInput
	// - (Required incase using IPMI based imaging either here or inside node spec) IPMI password.
	IpmiPassword pulumi.StringPtrInput
	// - (Required incase using IPMI based imaging either here or inside node spec) IPMI username.
	IpmiUser pulumi.StringPtrInput
	// - Id of the custom layout which needs to be passed to imaging request.
	LayoutEggUuid pulumi.StringPtrInput
	// - (Required) NOS package.
	NosPackage pulumi.StringPtrInput
	// - sessionId of the imaging session
	SessionId pulumi.StringPtrInput
	// - If hypervisor installation should be skipped.
	SkipHypervisor pulumi.BoolPtrInput
	// - Arguments to be passed to svmRescue for AOS installation. Ensure that the arguments provided are supported by the AOS version used for imaging.
	SvmRescueArgs pulumi.StringArrayInput
	// - Types of tests to be performed.
	Tests FoundationImageNodesTestsPtrInput
	// - UCSM IP address.
	UcsmIp pulumi.StringPtrInput
	// - UCSM password.
	UcsmPassword pulumi.StringPtrInput
	// - UCSM username.
	UcsmUser pulumi.StringPtrInput
	// - UNC password.
	UncPassword pulumi.StringPtrInput
	// - UNC Path.
	UncPath pulumi.StringPtrInput
	// - UNC username.
	UncUsername pulumi.StringPtrInput
	// - xen config types.
	XenConfigType pulumi.StringPtrInput
	// - xen server master IP address.
	XsMasterIp pulumi.StringPtrInput
	// - xen server master label.
	XsMasterLabel pulumi.StringPtrInput
	// - xen server master password.
	XsMasterPassword pulumi.StringPtrInput
	// - xen server master username.
	XsMasterUsername pulumi.StringPtrInput
}

func (FoundationImageNodesState) ElementType() reflect.Type {
	return reflect.TypeOf((*foundationImageNodesState)(nil)).Elem()
}

type foundationImageNodesArgs struct {
	Blocks   []FoundationImageNodesBlock   `pulumi:"blocks"`
	Clusters []FoundationImageNodesCluster `pulumi:"clusters"`
	// - (Required) CVM gateway.
	CvmGateway string `pulumi:"cvmGateway"`
	// - (Required) CVM netmask.
	CvmNetmask string `pulumi:"cvmNetmask"`
	// - Contains user data from Eos portal.
	EosMetadata *FoundationImageNodesEosMetadata `pulumi:"eosMetadata"`
	// - Foundation Central specific settings.
	FcSettings *FoundationImageNodesFcSettings `pulumi:"fcSettings"`
	// - Hyperv External virtual network adapter name.
	HypervExternalVnic *string `pulumi:"hypervExternalVnic"`
	// - Hyperv External vswitch name.
	HypervExternalVswitch *string `pulumi:"hypervExternalVswitch"`
	// - Hyperv product key.
	HypervProductKey *string `pulumi:"hypervProductKey"`
	// - Hyperv SKU.
	HypervSku *bool `pulumi:"hypervSku"`
	// - (Required) Hypervisor gateway.
	HypervisorGateway string `pulumi:"hypervisorGateway"`
	// - Hypervisor ISO.
	HypervisorIso        *FoundationImageNodesHypervisorIso `pulumi:"hypervisorIso"`
	HypervisorNameserver *string                            `pulumi:"hypervisorNameserver"`
	// - (Required) Hypervisor netmask.
	HypervisorNetmask string `pulumi:"hypervisorNetmask"`
	// - Hypervisor password.
	HypervisorPassword *string `pulumi:"hypervisorPassword"`
	// - install script.
	InstallScript *string `pulumi:"installScript"`
	// - (Required incase using IPMI based imaging either here or inside node spec) default IPMI gateway
	IpmiGateway *string `pulumi:"ipmiGateway"`
	// - (Required incase using IPMI based imaging either here or inside node spec) default IPMI netmask
	IpmiNetmask *string `pulumi:"ipmiNetmask"`
	// - (Required incase using IPMI based imaging either here or inside node spec) IPMI password.
	IpmiPassword *string `pulumi:"ipmiPassword"`
	// - (Required incase using IPMI based imaging either here or inside node spec) IPMI username.
	IpmiUser *string `pulumi:"ipmiUser"`
	// - Id of the custom layout which needs to be passed to imaging request.
	LayoutEggUuid *string `pulumi:"layoutEggUuid"`
	// - (Required) NOS package.
	NosPackage string `pulumi:"nosPackage"`
	// - If hypervisor installation should be skipped.
	SkipHypervisor *bool `pulumi:"skipHypervisor"`
	// - Arguments to be passed to svmRescue for AOS installation. Ensure that the arguments provided are supported by the AOS version used for imaging.
	SvmRescueArgs []string `pulumi:"svmRescueArgs"`
	// - Types of tests to be performed.
	Tests *FoundationImageNodesTests `pulumi:"tests"`
	// - UCSM IP address.
	UcsmIp *string `pulumi:"ucsmIp"`
	// - UCSM password.
	UcsmPassword *string `pulumi:"ucsmPassword"`
	// - UCSM username.
	UcsmUser *string `pulumi:"ucsmUser"`
	// - UNC password.
	UncPassword *string `pulumi:"uncPassword"`
	// - UNC Path.
	UncPath *string `pulumi:"uncPath"`
	// - UNC username.
	UncUsername *string `pulumi:"uncUsername"`
	// - xen config types.
	XenConfigType *string `pulumi:"xenConfigType"`
	// - xen server master IP address.
	XsMasterIp *string `pulumi:"xsMasterIp"`
	// - xen server master label.
	XsMasterLabel *string `pulumi:"xsMasterLabel"`
	// - xen server master password.
	XsMasterPassword *string `pulumi:"xsMasterPassword"`
	// - xen server master username.
	XsMasterUsername *string `pulumi:"xsMasterUsername"`
}

// The set of arguments for constructing a FoundationImageNodes resource.
type FoundationImageNodesArgs struct {
	Blocks   FoundationImageNodesBlockArrayInput
	Clusters FoundationImageNodesClusterArrayInput
	// - (Required) CVM gateway.
	CvmGateway pulumi.StringInput
	// - (Required) CVM netmask.
	CvmNetmask pulumi.StringInput
	// - Contains user data from Eos portal.
	EosMetadata FoundationImageNodesEosMetadataPtrInput
	// - Foundation Central specific settings.
	FcSettings FoundationImageNodesFcSettingsPtrInput
	// - Hyperv External virtual network adapter name.
	HypervExternalVnic pulumi.StringPtrInput
	// - Hyperv External vswitch name.
	HypervExternalVswitch pulumi.StringPtrInput
	// - Hyperv product key.
	HypervProductKey pulumi.StringPtrInput
	// - Hyperv SKU.
	HypervSku pulumi.BoolPtrInput
	// - (Required) Hypervisor gateway.
	HypervisorGateway pulumi.StringInput
	// - Hypervisor ISO.
	HypervisorIso        FoundationImageNodesHypervisorIsoPtrInput
	HypervisorNameserver pulumi.StringPtrInput
	// - (Required) Hypervisor netmask.
	HypervisorNetmask pulumi.StringInput
	// - Hypervisor password.
	HypervisorPassword pulumi.StringPtrInput
	// - install script.
	InstallScript pulumi.StringPtrInput
	// - (Required incase using IPMI based imaging either here or inside node spec) default IPMI gateway
	IpmiGateway pulumi.StringPtrInput
	// - (Required incase using IPMI based imaging either here or inside node spec) default IPMI netmask
	IpmiNetmask pulumi.StringPtrInput
	// - (Required incase using IPMI based imaging either here or inside node spec) IPMI password.
	IpmiPassword pulumi.StringPtrInput
	// - (Required incase using IPMI based imaging either here or inside node spec) IPMI username.
	IpmiUser pulumi.StringPtrInput
	// - Id of the custom layout which needs to be passed to imaging request.
	LayoutEggUuid pulumi.StringPtrInput
	// - (Required) NOS package.
	NosPackage pulumi.StringInput
	// - If hypervisor installation should be skipped.
	SkipHypervisor pulumi.BoolPtrInput
	// - Arguments to be passed to svmRescue for AOS installation. Ensure that the arguments provided are supported by the AOS version used for imaging.
	SvmRescueArgs pulumi.StringArrayInput
	// - Types of tests to be performed.
	Tests FoundationImageNodesTestsPtrInput
	// - UCSM IP address.
	UcsmIp pulumi.StringPtrInput
	// - UCSM password.
	UcsmPassword pulumi.StringPtrInput
	// - UCSM username.
	UcsmUser pulumi.StringPtrInput
	// - UNC password.
	UncPassword pulumi.StringPtrInput
	// - UNC Path.
	UncPath pulumi.StringPtrInput
	// - UNC username.
	UncUsername pulumi.StringPtrInput
	// - xen config types.
	XenConfigType pulumi.StringPtrInput
	// - xen server master IP address.
	XsMasterIp pulumi.StringPtrInput
	// - xen server master label.
	XsMasterLabel pulumi.StringPtrInput
	// - xen server master password.
	XsMasterPassword pulumi.StringPtrInput
	// - xen server master username.
	XsMasterUsername pulumi.StringPtrInput
}

func (FoundationImageNodesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*foundationImageNodesArgs)(nil)).Elem()
}

type FoundationImageNodesInput interface {
	pulumi.Input

	ToFoundationImageNodesOutput() FoundationImageNodesOutput
	ToFoundationImageNodesOutputWithContext(ctx context.Context) FoundationImageNodesOutput
}

func (*FoundationImageNodes) ElementType() reflect.Type {
	return reflect.TypeOf((**FoundationImageNodes)(nil)).Elem()
}

func (i *FoundationImageNodes) ToFoundationImageNodesOutput() FoundationImageNodesOutput {
	return i.ToFoundationImageNodesOutputWithContext(context.Background())
}

func (i *FoundationImageNodes) ToFoundationImageNodesOutputWithContext(ctx context.Context) FoundationImageNodesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FoundationImageNodesOutput)
}

func (i *FoundationImageNodes) ToOutput(ctx context.Context) pulumix.Output[*FoundationImageNodes] {
	return pulumix.Output[*FoundationImageNodes]{
		OutputState: i.ToFoundationImageNodesOutputWithContext(ctx).OutputState,
	}
}

// FoundationImageNodesArrayInput is an input type that accepts FoundationImageNodesArray and FoundationImageNodesArrayOutput values.
// You can construct a concrete instance of `FoundationImageNodesArrayInput` via:
//
//	FoundationImageNodesArray{ FoundationImageNodesArgs{...} }
type FoundationImageNodesArrayInput interface {
	pulumi.Input

	ToFoundationImageNodesArrayOutput() FoundationImageNodesArrayOutput
	ToFoundationImageNodesArrayOutputWithContext(context.Context) FoundationImageNodesArrayOutput
}

type FoundationImageNodesArray []FoundationImageNodesInput

func (FoundationImageNodesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FoundationImageNodes)(nil)).Elem()
}

func (i FoundationImageNodesArray) ToFoundationImageNodesArrayOutput() FoundationImageNodesArrayOutput {
	return i.ToFoundationImageNodesArrayOutputWithContext(context.Background())
}

func (i FoundationImageNodesArray) ToFoundationImageNodesArrayOutputWithContext(ctx context.Context) FoundationImageNodesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FoundationImageNodesArrayOutput)
}

func (i FoundationImageNodesArray) ToOutput(ctx context.Context) pulumix.Output[[]*FoundationImageNodes] {
	return pulumix.Output[[]*FoundationImageNodes]{
		OutputState: i.ToFoundationImageNodesArrayOutputWithContext(ctx).OutputState,
	}
}

// FoundationImageNodesMapInput is an input type that accepts FoundationImageNodesMap and FoundationImageNodesMapOutput values.
// You can construct a concrete instance of `FoundationImageNodesMapInput` via:
//
//	FoundationImageNodesMap{ "key": FoundationImageNodesArgs{...} }
type FoundationImageNodesMapInput interface {
	pulumi.Input

	ToFoundationImageNodesMapOutput() FoundationImageNodesMapOutput
	ToFoundationImageNodesMapOutputWithContext(context.Context) FoundationImageNodesMapOutput
}

type FoundationImageNodesMap map[string]FoundationImageNodesInput

func (FoundationImageNodesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FoundationImageNodes)(nil)).Elem()
}

func (i FoundationImageNodesMap) ToFoundationImageNodesMapOutput() FoundationImageNodesMapOutput {
	return i.ToFoundationImageNodesMapOutputWithContext(context.Background())
}

func (i FoundationImageNodesMap) ToFoundationImageNodesMapOutputWithContext(ctx context.Context) FoundationImageNodesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FoundationImageNodesMapOutput)
}

func (i FoundationImageNodesMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*FoundationImageNodes] {
	return pulumix.Output[map[string]*FoundationImageNodes]{
		OutputState: i.ToFoundationImageNodesMapOutputWithContext(ctx).OutputState,
	}
}

type FoundationImageNodesOutput struct{ *pulumi.OutputState }

func (FoundationImageNodesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FoundationImageNodes)(nil)).Elem()
}

func (o FoundationImageNodesOutput) ToFoundationImageNodesOutput() FoundationImageNodesOutput {
	return o
}

func (o FoundationImageNodesOutput) ToFoundationImageNodesOutputWithContext(ctx context.Context) FoundationImageNodesOutput {
	return o
}

func (o FoundationImageNodesOutput) ToOutput(ctx context.Context) pulumix.Output[*FoundationImageNodes] {
	return pulumix.Output[*FoundationImageNodes]{
		OutputState: o.OutputState,
	}
}

func (o FoundationImageNodesOutput) Blocks() FoundationImageNodesBlockArrayOutput {
	return o.ApplyT(func(v *FoundationImageNodes) FoundationImageNodesBlockArrayOutput { return v.Blocks }).(FoundationImageNodesBlockArrayOutput)
}

// - list containing cluster name and cluster urls for created clusters in current session
// * `cluster_urls.#.cluster_name` :- clusterName
// * `cluster_urls.#.cluster_url` :- url to access the cluster login
func (o FoundationImageNodesOutput) ClusterUrls() FoundationImageNodesClusterUrlArrayOutput {
	return o.ApplyT(func(v *FoundationImageNodes) FoundationImageNodesClusterUrlArrayOutput { return v.ClusterUrls }).(FoundationImageNodesClusterUrlArrayOutput)
}

func (o FoundationImageNodesOutput) Clusters() FoundationImageNodesClusterArrayOutput {
	return o.ApplyT(func(v *FoundationImageNodes) FoundationImageNodesClusterArrayOutput { return v.Clusters }).(FoundationImageNodesClusterArrayOutput)
}

// - (Required) CVM gateway.
func (o FoundationImageNodesOutput) CvmGateway() pulumi.StringOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringOutput { return v.CvmGateway }).(pulumi.StringOutput)
}

// - (Required) CVM netmask.
func (o FoundationImageNodesOutput) CvmNetmask() pulumi.StringOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringOutput { return v.CvmNetmask }).(pulumi.StringOutput)
}

// - Contains user data from Eos portal.
func (o FoundationImageNodesOutput) EosMetadata() FoundationImageNodesEosMetadataPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) FoundationImageNodesEosMetadataPtrOutput { return v.EosMetadata }).(FoundationImageNodesEosMetadataPtrOutput)
}

// - Foundation Central specific settings.
func (o FoundationImageNodesOutput) FcSettings() FoundationImageNodesFcSettingsPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) FoundationImageNodesFcSettingsPtrOutput { return v.FcSettings }).(FoundationImageNodesFcSettingsPtrOutput)
}

// - Hyperv External virtual network adapter name.
func (o FoundationImageNodesOutput) HypervExternalVnic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringPtrOutput { return v.HypervExternalVnic }).(pulumi.StringPtrOutput)
}

// - Hyperv External vswitch name.
func (o FoundationImageNodesOutput) HypervExternalVswitch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringPtrOutput { return v.HypervExternalVswitch }).(pulumi.StringPtrOutput)
}

// - Hyperv product key.
func (o FoundationImageNodesOutput) HypervProductKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringPtrOutput { return v.HypervProductKey }).(pulumi.StringPtrOutput)
}

// - Hyperv SKU.
func (o FoundationImageNodesOutput) HypervSku() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.BoolPtrOutput { return v.HypervSku }).(pulumi.BoolPtrOutput)
}

// - (Required) Hypervisor gateway.
func (o FoundationImageNodesOutput) HypervisorGateway() pulumi.StringOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringOutput { return v.HypervisorGateway }).(pulumi.StringOutput)
}

// - Hypervisor ISO.
func (o FoundationImageNodesOutput) HypervisorIso() FoundationImageNodesHypervisorIsoPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) FoundationImageNodesHypervisorIsoPtrOutput { return v.HypervisorIso }).(FoundationImageNodesHypervisorIsoPtrOutput)
}

func (o FoundationImageNodesOutput) HypervisorNameserver() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringPtrOutput { return v.HypervisorNameserver }).(pulumi.StringPtrOutput)
}

// - (Required) Hypervisor netmask.
func (o FoundationImageNodesOutput) HypervisorNetmask() pulumi.StringOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringOutput { return v.HypervisorNetmask }).(pulumi.StringOutput)
}

// - Hypervisor password.
func (o FoundationImageNodesOutput) HypervisorPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringPtrOutput { return v.HypervisorPassword }).(pulumi.StringPtrOutput)
}

// - install script.
func (o FoundationImageNodesOutput) InstallScript() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringPtrOutput { return v.InstallScript }).(pulumi.StringPtrOutput)
}

// - (Required incase using IPMI based imaging either here or inside node spec) default IPMI gateway
func (o FoundationImageNodesOutput) IpmiGateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringPtrOutput { return v.IpmiGateway }).(pulumi.StringPtrOutput)
}

// - (Required incase using IPMI based imaging either here or inside node spec) default IPMI netmask
func (o FoundationImageNodesOutput) IpmiNetmask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringPtrOutput { return v.IpmiNetmask }).(pulumi.StringPtrOutput)
}

// - (Required incase using IPMI based imaging either here or inside node spec) IPMI password.
func (o FoundationImageNodesOutput) IpmiPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringPtrOutput { return v.IpmiPassword }).(pulumi.StringPtrOutput)
}

// - (Required incase using IPMI based imaging either here or inside node spec) IPMI username.
func (o FoundationImageNodesOutput) IpmiUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringPtrOutput { return v.IpmiUser }).(pulumi.StringPtrOutput)
}

// - Id of the custom layout which needs to be passed to imaging request.
func (o FoundationImageNodesOutput) LayoutEggUuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringPtrOutput { return v.LayoutEggUuid }).(pulumi.StringPtrOutput)
}

// - (Required) NOS package.
func (o FoundationImageNodesOutput) NosPackage() pulumi.StringOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringOutput { return v.NosPackage }).(pulumi.StringOutput)
}

// - sessionId of the imaging session
func (o FoundationImageNodesOutput) SessionId() pulumi.StringOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringOutput { return v.SessionId }).(pulumi.StringOutput)
}

// - If hypervisor installation should be skipped.
func (o FoundationImageNodesOutput) SkipHypervisor() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.BoolPtrOutput { return v.SkipHypervisor }).(pulumi.BoolPtrOutput)
}

// - Arguments to be passed to svmRescue for AOS installation. Ensure that the arguments provided are supported by the AOS version used for imaging.
func (o FoundationImageNodesOutput) SvmRescueArgs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringArrayOutput { return v.SvmRescueArgs }).(pulumi.StringArrayOutput)
}

// - Types of tests to be performed.
func (o FoundationImageNodesOutput) Tests() FoundationImageNodesTestsPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) FoundationImageNodesTestsPtrOutput { return v.Tests }).(FoundationImageNodesTestsPtrOutput)
}

// - UCSM IP address.
func (o FoundationImageNodesOutput) UcsmIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringPtrOutput { return v.UcsmIp }).(pulumi.StringPtrOutput)
}

// - UCSM password.
func (o FoundationImageNodesOutput) UcsmPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringPtrOutput { return v.UcsmPassword }).(pulumi.StringPtrOutput)
}

// - UCSM username.
func (o FoundationImageNodesOutput) UcsmUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringPtrOutput { return v.UcsmUser }).(pulumi.StringPtrOutput)
}

// - UNC password.
func (o FoundationImageNodesOutput) UncPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringPtrOutput { return v.UncPassword }).(pulumi.StringPtrOutput)
}

// - UNC Path.
func (o FoundationImageNodesOutput) UncPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringPtrOutput { return v.UncPath }).(pulumi.StringPtrOutput)
}

// - UNC username.
func (o FoundationImageNodesOutput) UncUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringPtrOutput { return v.UncUsername }).(pulumi.StringPtrOutput)
}

// - xen config types.
func (o FoundationImageNodesOutput) XenConfigType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringPtrOutput { return v.XenConfigType }).(pulumi.StringPtrOutput)
}

// - xen server master IP address.
func (o FoundationImageNodesOutput) XsMasterIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringPtrOutput { return v.XsMasterIp }).(pulumi.StringPtrOutput)
}

// - xen server master label.
func (o FoundationImageNodesOutput) XsMasterLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringPtrOutput { return v.XsMasterLabel }).(pulumi.StringPtrOutput)
}

// - xen server master password.
func (o FoundationImageNodesOutput) XsMasterPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringPtrOutput { return v.XsMasterPassword }).(pulumi.StringPtrOutput)
}

// - xen server master username.
func (o FoundationImageNodesOutput) XsMasterUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FoundationImageNodes) pulumi.StringPtrOutput { return v.XsMasterUsername }).(pulumi.StringPtrOutput)
}

type FoundationImageNodesArrayOutput struct{ *pulumi.OutputState }

func (FoundationImageNodesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FoundationImageNodes)(nil)).Elem()
}

func (o FoundationImageNodesArrayOutput) ToFoundationImageNodesArrayOutput() FoundationImageNodesArrayOutput {
	return o
}

func (o FoundationImageNodesArrayOutput) ToFoundationImageNodesArrayOutputWithContext(ctx context.Context) FoundationImageNodesArrayOutput {
	return o
}

func (o FoundationImageNodesArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*FoundationImageNodes] {
	return pulumix.Output[[]*FoundationImageNodes]{
		OutputState: o.OutputState,
	}
}

func (o FoundationImageNodesArrayOutput) Index(i pulumi.IntInput) FoundationImageNodesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FoundationImageNodes {
		return vs[0].([]*FoundationImageNodes)[vs[1].(int)]
	}).(FoundationImageNodesOutput)
}

type FoundationImageNodesMapOutput struct{ *pulumi.OutputState }

func (FoundationImageNodesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FoundationImageNodes)(nil)).Elem()
}

func (o FoundationImageNodesMapOutput) ToFoundationImageNodesMapOutput() FoundationImageNodesMapOutput {
	return o
}

func (o FoundationImageNodesMapOutput) ToFoundationImageNodesMapOutputWithContext(ctx context.Context) FoundationImageNodesMapOutput {
	return o
}

func (o FoundationImageNodesMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*FoundationImageNodes] {
	return pulumix.Output[map[string]*FoundationImageNodes]{
		OutputState: o.OutputState,
	}
}

func (o FoundationImageNodesMapOutput) MapIndex(k pulumi.StringInput) FoundationImageNodesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FoundationImageNodes {
		return vs[0].(map[string]*FoundationImageNodes)[vs[1].(string)]
	}).(FoundationImageNodesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FoundationImageNodesInput)(nil)).Elem(), &FoundationImageNodes{})
	pulumi.RegisterInputType(reflect.TypeOf((*FoundationImageNodesArrayInput)(nil)).Elem(), FoundationImageNodesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FoundationImageNodesMapInput)(nil)).Elem(), FoundationImageNodesMap{})
	pulumi.RegisterOutputType(FoundationImageNodesOutput{})
	pulumi.RegisterOutputType(FoundationImageNodesArrayOutput{})
	pulumi.RegisterOutputType(FoundationImageNodesMapOutput{})
}