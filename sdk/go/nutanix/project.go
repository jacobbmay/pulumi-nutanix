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

// Provides a Nutanix Project resource to Create a Project.
type Project struct {
	pulumi.CustomResourceState

	AccountReferenceLists ProjectAccountReferenceListArrayOutput `pulumi:"accountReferenceLists"`
	Acps                  ProjectAcpArrayOutput                  `pulumi:"acps"`
	ApiVersion            pulumi.StringOutput                    `pulumi:"apiVersion"`
	// - (Optional) The category values represented as a dictionary of key > list of values.
	Categories            ProjectCategoryArrayOutput             `pulumi:"categories"`
	ClusterReferenceLists ProjectClusterReferenceListArrayOutput `pulumi:"clusterReferenceLists"`
	// The UUID of cluster. (Required when using projectInternal flag).
	ClusterUuid pulumi.StringPtrOutput `pulumi:"clusterUuid"`
	// Reference to a environment.
	// * `default_environment_reference.kind` - (Optional) The kind name. Default value is `environment`
	// * `default_environment_reference.uuid` - (Required) The UUID of a environment
	// * `default_environment_reference.name` - (Optional/Computed) The name of a environment.
	DefaultEnvironmentReference ProjectDefaultEnvironmentReferenceOutput `pulumi:"defaultEnvironmentReference"`
	// Reference to a subnet.
	// * `default_subnet_reference.kind` - (Optional) The kind name. Default value is `subnet`
	// * `default_subnet_reference.uuid` - (Required) The UUID of a subnet.
	// * `default_subnet_reference.name` - (Optional/Computed) The name of a subnet.
	DefaultSubnetReference ProjectDefaultSubnetReferenceOutput `pulumi:"defaultSubnetReference"`
	// A description for project.
	Description pulumi.StringOutput `pulumi:"description"`
	// flag to allow collaboration of projects. (Use with projectInternal flag)
	EnableCollab                    pulumi.BoolPtrOutput                             `pulumi:"enableCollab"`
	EnvironmentReferenceLists       ProjectEnvironmentReferenceListArrayOutput       `pulumi:"environmentReferenceLists"`
	ExternalNetworkLists            ProjectExternalNetworkListArrayOutput            `pulumi:"externalNetworkLists"`
	ExternalUserGroupReferenceLists ProjectExternalUserGroupReferenceListArrayOutput `pulumi:"externalUserGroupReferenceLists"`
	IsDefault                       pulumi.BoolOutput                                `pulumi:"isDefault"`
	Metadata                        pulumi.StringMapOutput                           `pulumi:"metadata"`
	// The name for the project.
	Name                 pulumi.StringOutput                   `pulumi:"name"`
	OwnerReference       pulumi.StringMapOutput                `pulumi:"ownerReference"`
	ProjectReference     pulumi.StringMapOutput                `pulumi:"projectReference"`
	ResourceDomain       ProjectResourceDomainPtrOutput        `pulumi:"resourceDomain"`
	State                pulumi.StringOutput                   `pulumi:"state"`
	SubnetReferenceLists ProjectSubnetReferenceListArrayOutput `pulumi:"subnetReferenceLists"`
	TunnelReferenceLists ProjectTunnelReferenceListArrayOutput `pulumi:"tunnelReferenceLists"`
	// flag to use project internal for user role mapping
	UseProjectInternal pulumi.BoolPtrOutput            `pulumi:"useProjectInternal"`
	UserGroupLists     ProjectUserGroupListArrayOutput `pulumi:"userGroupLists"`
	UserLists          ProjectUserListArrayOutput      `pulumi:"userLists"`
	// List of Reference of users.
	UserReferenceLists ProjectUserReferenceListArrayOutput `pulumi:"userReferenceLists"`
	VpcReferenceLists  ProjectVpcReferenceListArrayOutput  `pulumi:"vpcReferenceLists"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultSubnetReference == nil {
		return nil, errors.New("invalid value for required argument 'DefaultSubnetReference'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Project
	err := ctx.RegisterResource("nutanix:index/project:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("nutanix:index/project:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
	AccountReferenceLists []ProjectAccountReferenceList `pulumi:"accountReferenceLists"`
	Acps                  []ProjectAcp                  `pulumi:"acps"`
	ApiVersion            *string                       `pulumi:"apiVersion"`
	// - (Optional) The category values represented as a dictionary of key > list of values.
	Categories            []ProjectCategory             `pulumi:"categories"`
	ClusterReferenceLists []ProjectClusterReferenceList `pulumi:"clusterReferenceLists"`
	// The UUID of cluster. (Required when using projectInternal flag).
	ClusterUuid *string `pulumi:"clusterUuid"`
	// Reference to a environment.
	// * `default_environment_reference.kind` - (Optional) The kind name. Default value is `environment`
	// * `default_environment_reference.uuid` - (Required) The UUID of a environment
	// * `default_environment_reference.name` - (Optional/Computed) The name of a environment.
	DefaultEnvironmentReference *ProjectDefaultEnvironmentReference `pulumi:"defaultEnvironmentReference"`
	// Reference to a subnet.
	// * `default_subnet_reference.kind` - (Optional) The kind name. Default value is `subnet`
	// * `default_subnet_reference.uuid` - (Required) The UUID of a subnet.
	// * `default_subnet_reference.name` - (Optional/Computed) The name of a subnet.
	DefaultSubnetReference *ProjectDefaultSubnetReference `pulumi:"defaultSubnetReference"`
	// A description for project.
	Description *string `pulumi:"description"`
	// flag to allow collaboration of projects. (Use with projectInternal flag)
	EnableCollab                    *bool                                   `pulumi:"enableCollab"`
	EnvironmentReferenceLists       []ProjectEnvironmentReferenceList       `pulumi:"environmentReferenceLists"`
	ExternalNetworkLists            []ProjectExternalNetworkList            `pulumi:"externalNetworkLists"`
	ExternalUserGroupReferenceLists []ProjectExternalUserGroupReferenceList `pulumi:"externalUserGroupReferenceLists"`
	IsDefault                       *bool                                   `pulumi:"isDefault"`
	Metadata                        map[string]string                       `pulumi:"metadata"`
	// The name for the project.
	Name                 *string                      `pulumi:"name"`
	OwnerReference       map[string]string            `pulumi:"ownerReference"`
	ProjectReference     map[string]string            `pulumi:"projectReference"`
	ResourceDomain       *ProjectResourceDomain       `pulumi:"resourceDomain"`
	State                *string                      `pulumi:"state"`
	SubnetReferenceLists []ProjectSubnetReferenceList `pulumi:"subnetReferenceLists"`
	TunnelReferenceLists []ProjectTunnelReferenceList `pulumi:"tunnelReferenceLists"`
	// flag to use project internal for user role mapping
	UseProjectInternal *bool                  `pulumi:"useProjectInternal"`
	UserGroupLists     []ProjectUserGroupList `pulumi:"userGroupLists"`
	UserLists          []ProjectUserList      `pulumi:"userLists"`
	// List of Reference of users.
	UserReferenceLists []ProjectUserReferenceList `pulumi:"userReferenceLists"`
	VpcReferenceLists  []ProjectVpcReferenceList  `pulumi:"vpcReferenceLists"`
}

type ProjectState struct {
	AccountReferenceLists ProjectAccountReferenceListArrayInput
	Acps                  ProjectAcpArrayInput
	ApiVersion            pulumi.StringPtrInput
	// - (Optional) The category values represented as a dictionary of key > list of values.
	Categories            ProjectCategoryArrayInput
	ClusterReferenceLists ProjectClusterReferenceListArrayInput
	// The UUID of cluster. (Required when using projectInternal flag).
	ClusterUuid pulumi.StringPtrInput
	// Reference to a environment.
	// * `default_environment_reference.kind` - (Optional) The kind name. Default value is `environment`
	// * `default_environment_reference.uuid` - (Required) The UUID of a environment
	// * `default_environment_reference.name` - (Optional/Computed) The name of a environment.
	DefaultEnvironmentReference ProjectDefaultEnvironmentReferencePtrInput
	// Reference to a subnet.
	// * `default_subnet_reference.kind` - (Optional) The kind name. Default value is `subnet`
	// * `default_subnet_reference.uuid` - (Required) The UUID of a subnet.
	// * `default_subnet_reference.name` - (Optional/Computed) The name of a subnet.
	DefaultSubnetReference ProjectDefaultSubnetReferencePtrInput
	// A description for project.
	Description pulumi.StringPtrInput
	// flag to allow collaboration of projects. (Use with projectInternal flag)
	EnableCollab                    pulumi.BoolPtrInput
	EnvironmentReferenceLists       ProjectEnvironmentReferenceListArrayInput
	ExternalNetworkLists            ProjectExternalNetworkListArrayInput
	ExternalUserGroupReferenceLists ProjectExternalUserGroupReferenceListArrayInput
	IsDefault                       pulumi.BoolPtrInput
	Metadata                        pulumi.StringMapInput
	// The name for the project.
	Name                 pulumi.StringPtrInput
	OwnerReference       pulumi.StringMapInput
	ProjectReference     pulumi.StringMapInput
	ResourceDomain       ProjectResourceDomainPtrInput
	State                pulumi.StringPtrInput
	SubnetReferenceLists ProjectSubnetReferenceListArrayInput
	TunnelReferenceLists ProjectTunnelReferenceListArrayInput
	// flag to use project internal for user role mapping
	UseProjectInternal pulumi.BoolPtrInput
	UserGroupLists     ProjectUserGroupListArrayInput
	UserLists          ProjectUserListArrayInput
	// List of Reference of users.
	UserReferenceLists ProjectUserReferenceListArrayInput
	VpcReferenceLists  ProjectVpcReferenceListArrayInput
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	AccountReferenceLists []ProjectAccountReferenceList `pulumi:"accountReferenceLists"`
	Acps                  []ProjectAcp                  `pulumi:"acps"`
	ApiVersion            *string                       `pulumi:"apiVersion"`
	// - (Optional) The category values represented as a dictionary of key > list of values.
	Categories            []ProjectCategory             `pulumi:"categories"`
	ClusterReferenceLists []ProjectClusterReferenceList `pulumi:"clusterReferenceLists"`
	// The UUID of cluster. (Required when using projectInternal flag).
	ClusterUuid *string `pulumi:"clusterUuid"`
	// Reference to a environment.
	// * `default_environment_reference.kind` - (Optional) The kind name. Default value is `environment`
	// * `default_environment_reference.uuid` - (Required) The UUID of a environment
	// * `default_environment_reference.name` - (Optional/Computed) The name of a environment.
	DefaultEnvironmentReference *ProjectDefaultEnvironmentReference `pulumi:"defaultEnvironmentReference"`
	// Reference to a subnet.
	// * `default_subnet_reference.kind` - (Optional) The kind name. Default value is `subnet`
	// * `default_subnet_reference.uuid` - (Required) The UUID of a subnet.
	// * `default_subnet_reference.name` - (Optional/Computed) The name of a subnet.
	DefaultSubnetReference ProjectDefaultSubnetReference `pulumi:"defaultSubnetReference"`
	// A description for project.
	Description string `pulumi:"description"`
	// flag to allow collaboration of projects. (Use with projectInternal flag)
	EnableCollab                    *bool                                   `pulumi:"enableCollab"`
	EnvironmentReferenceLists       []ProjectEnvironmentReferenceList       `pulumi:"environmentReferenceLists"`
	ExternalNetworkLists            []ProjectExternalNetworkList            `pulumi:"externalNetworkLists"`
	ExternalUserGroupReferenceLists []ProjectExternalUserGroupReferenceList `pulumi:"externalUserGroupReferenceLists"`
	// The name for the project.
	Name                 *string                      `pulumi:"name"`
	OwnerReference       map[string]string            `pulumi:"ownerReference"`
	ProjectReference     map[string]string            `pulumi:"projectReference"`
	ResourceDomain       *ProjectResourceDomain       `pulumi:"resourceDomain"`
	SubnetReferenceLists []ProjectSubnetReferenceList `pulumi:"subnetReferenceLists"`
	TunnelReferenceLists []ProjectTunnelReferenceList `pulumi:"tunnelReferenceLists"`
	// flag to use project internal for user role mapping
	UseProjectInternal *bool                  `pulumi:"useProjectInternal"`
	UserGroupLists     []ProjectUserGroupList `pulumi:"userGroupLists"`
	UserLists          []ProjectUserList      `pulumi:"userLists"`
	// List of Reference of users.
	UserReferenceLists []ProjectUserReferenceList `pulumi:"userReferenceLists"`
	VpcReferenceLists  []ProjectVpcReferenceList  `pulumi:"vpcReferenceLists"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	AccountReferenceLists ProjectAccountReferenceListArrayInput
	Acps                  ProjectAcpArrayInput
	ApiVersion            pulumi.StringPtrInput
	// - (Optional) The category values represented as a dictionary of key > list of values.
	Categories            ProjectCategoryArrayInput
	ClusterReferenceLists ProjectClusterReferenceListArrayInput
	// The UUID of cluster. (Required when using projectInternal flag).
	ClusterUuid pulumi.StringPtrInput
	// Reference to a environment.
	// * `default_environment_reference.kind` - (Optional) The kind name. Default value is `environment`
	// * `default_environment_reference.uuid` - (Required) The UUID of a environment
	// * `default_environment_reference.name` - (Optional/Computed) The name of a environment.
	DefaultEnvironmentReference ProjectDefaultEnvironmentReferencePtrInput
	// Reference to a subnet.
	// * `default_subnet_reference.kind` - (Optional) The kind name. Default value is `subnet`
	// * `default_subnet_reference.uuid` - (Required) The UUID of a subnet.
	// * `default_subnet_reference.name` - (Optional/Computed) The name of a subnet.
	DefaultSubnetReference ProjectDefaultSubnetReferenceInput
	// A description for project.
	Description pulumi.StringInput
	// flag to allow collaboration of projects. (Use with projectInternal flag)
	EnableCollab                    pulumi.BoolPtrInput
	EnvironmentReferenceLists       ProjectEnvironmentReferenceListArrayInput
	ExternalNetworkLists            ProjectExternalNetworkListArrayInput
	ExternalUserGroupReferenceLists ProjectExternalUserGroupReferenceListArrayInput
	// The name for the project.
	Name                 pulumi.StringPtrInput
	OwnerReference       pulumi.StringMapInput
	ProjectReference     pulumi.StringMapInput
	ResourceDomain       ProjectResourceDomainPtrInput
	SubnetReferenceLists ProjectSubnetReferenceListArrayInput
	TunnelReferenceLists ProjectTunnelReferenceListArrayInput
	// flag to use project internal for user role mapping
	UseProjectInternal pulumi.BoolPtrInput
	UserGroupLists     ProjectUserGroupListArrayInput
	UserLists          ProjectUserListArrayInput
	// List of Reference of users.
	UserReferenceLists ProjectUserReferenceListArrayInput
	VpcReferenceLists  ProjectVpcReferenceListArrayInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

type ProjectInput interface {
	pulumi.Input

	ToProjectOutput() ProjectOutput
	ToProjectOutputWithContext(ctx context.Context) ProjectOutput
}

func (*Project) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (i *Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i *Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

func (i *Project) ToOutput(ctx context.Context) pulumix.Output[*Project] {
	return pulumix.Output[*Project]{
		OutputState: i.ToProjectOutputWithContext(ctx).OutputState,
	}
}

// ProjectArrayInput is an input type that accepts ProjectArray and ProjectArrayOutput values.
// You can construct a concrete instance of `ProjectArrayInput` via:
//
//	ProjectArray{ ProjectArgs{...} }
type ProjectArrayInput interface {
	pulumi.Input

	ToProjectArrayOutput() ProjectArrayOutput
	ToProjectArrayOutputWithContext(context.Context) ProjectArrayOutput
}

type ProjectArray []ProjectInput

func (ProjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Project)(nil)).Elem()
}

func (i ProjectArray) ToProjectArrayOutput() ProjectArrayOutput {
	return i.ToProjectArrayOutputWithContext(context.Background())
}

func (i ProjectArray) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectArrayOutput)
}

func (i ProjectArray) ToOutput(ctx context.Context) pulumix.Output[[]*Project] {
	return pulumix.Output[[]*Project]{
		OutputState: i.ToProjectArrayOutputWithContext(ctx).OutputState,
	}
}

// ProjectMapInput is an input type that accepts ProjectMap and ProjectMapOutput values.
// You can construct a concrete instance of `ProjectMapInput` via:
//
//	ProjectMap{ "key": ProjectArgs{...} }
type ProjectMapInput interface {
	pulumi.Input

	ToProjectMapOutput() ProjectMapOutput
	ToProjectMapOutputWithContext(context.Context) ProjectMapOutput
}

type ProjectMap map[string]ProjectInput

func (ProjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Project)(nil)).Elem()
}

func (i ProjectMap) ToProjectMapOutput() ProjectMapOutput {
	return i.ToProjectMapOutputWithContext(context.Background())
}

func (i ProjectMap) ToProjectMapOutputWithContext(ctx context.Context) ProjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMapOutput)
}

func (i ProjectMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Project] {
	return pulumix.Output[map[string]*Project]{
		OutputState: i.ToProjectMapOutputWithContext(ctx).OutputState,
	}
}

type ProjectOutput struct{ *pulumi.OutputState }

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

func (o ProjectOutput) ToOutput(ctx context.Context) pulumix.Output[*Project] {
	return pulumix.Output[*Project]{
		OutputState: o.OutputState,
	}
}

func (o ProjectOutput) AccountReferenceLists() ProjectAccountReferenceListArrayOutput {
	return o.ApplyT(func(v *Project) ProjectAccountReferenceListArrayOutput { return v.AccountReferenceLists }).(ProjectAccountReferenceListArrayOutput)
}

func (o ProjectOutput) Acps() ProjectAcpArrayOutput {
	return o.ApplyT(func(v *Project) ProjectAcpArrayOutput { return v.Acps }).(ProjectAcpArrayOutput)
}

func (o ProjectOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// - (Optional) The category values represented as a dictionary of key > list of values.
func (o ProjectOutput) Categories() ProjectCategoryArrayOutput {
	return o.ApplyT(func(v *Project) ProjectCategoryArrayOutput { return v.Categories }).(ProjectCategoryArrayOutput)
}

func (o ProjectOutput) ClusterReferenceLists() ProjectClusterReferenceListArrayOutput {
	return o.ApplyT(func(v *Project) ProjectClusterReferenceListArrayOutput { return v.ClusterReferenceLists }).(ProjectClusterReferenceListArrayOutput)
}

// The UUID of cluster. (Required when using projectInternal flag).
func (o ProjectOutput) ClusterUuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.ClusterUuid }).(pulumi.StringPtrOutput)
}

// Reference to a environment.
// * `default_environment_reference.kind` - (Optional) The kind name. Default value is `environment`
// * `default_environment_reference.uuid` - (Required) The UUID of a environment
// * `default_environment_reference.name` - (Optional/Computed) The name of a environment.
func (o ProjectOutput) DefaultEnvironmentReference() ProjectDefaultEnvironmentReferenceOutput {
	return o.ApplyT(func(v *Project) ProjectDefaultEnvironmentReferenceOutput { return v.DefaultEnvironmentReference }).(ProjectDefaultEnvironmentReferenceOutput)
}

// Reference to a subnet.
// * `default_subnet_reference.kind` - (Optional) The kind name. Default value is `subnet`
// * `default_subnet_reference.uuid` - (Required) The UUID of a subnet.
// * `default_subnet_reference.name` - (Optional/Computed) The name of a subnet.
func (o ProjectOutput) DefaultSubnetReference() ProjectDefaultSubnetReferenceOutput {
	return o.ApplyT(func(v *Project) ProjectDefaultSubnetReferenceOutput { return v.DefaultSubnetReference }).(ProjectDefaultSubnetReferenceOutput)
}

// A description for project.
func (o ProjectOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// flag to allow collaboration of projects. (Use with projectInternal flag)
func (o ProjectOutput) EnableCollab() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.BoolPtrOutput { return v.EnableCollab }).(pulumi.BoolPtrOutput)
}

func (o ProjectOutput) EnvironmentReferenceLists() ProjectEnvironmentReferenceListArrayOutput {
	return o.ApplyT(func(v *Project) ProjectEnvironmentReferenceListArrayOutput { return v.EnvironmentReferenceLists }).(ProjectEnvironmentReferenceListArrayOutput)
}

func (o ProjectOutput) ExternalNetworkLists() ProjectExternalNetworkListArrayOutput {
	return o.ApplyT(func(v *Project) ProjectExternalNetworkListArrayOutput { return v.ExternalNetworkLists }).(ProjectExternalNetworkListArrayOutput)
}

func (o ProjectOutput) ExternalUserGroupReferenceLists() ProjectExternalUserGroupReferenceListArrayOutput {
	return o.ApplyT(func(v *Project) ProjectExternalUserGroupReferenceListArrayOutput {
		return v.ExternalUserGroupReferenceLists
	}).(ProjectExternalUserGroupReferenceListArrayOutput)
}

func (o ProjectOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v *Project) pulumi.BoolOutput { return v.IsDefault }).(pulumi.BoolOutput)
}

func (o ProjectOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Project) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// The name for the project.
func (o ProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ProjectOutput) OwnerReference() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Project) pulumi.StringMapOutput { return v.OwnerReference }).(pulumi.StringMapOutput)
}

func (o ProjectOutput) ProjectReference() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Project) pulumi.StringMapOutput { return v.ProjectReference }).(pulumi.StringMapOutput)
}

func (o ProjectOutput) ResourceDomain() ProjectResourceDomainPtrOutput {
	return o.ApplyT(func(v *Project) ProjectResourceDomainPtrOutput { return v.ResourceDomain }).(ProjectResourceDomainPtrOutput)
}

func (o ProjectOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o ProjectOutput) SubnetReferenceLists() ProjectSubnetReferenceListArrayOutput {
	return o.ApplyT(func(v *Project) ProjectSubnetReferenceListArrayOutput { return v.SubnetReferenceLists }).(ProjectSubnetReferenceListArrayOutput)
}

func (o ProjectOutput) TunnelReferenceLists() ProjectTunnelReferenceListArrayOutput {
	return o.ApplyT(func(v *Project) ProjectTunnelReferenceListArrayOutput { return v.TunnelReferenceLists }).(ProjectTunnelReferenceListArrayOutput)
}

// flag to use project internal for user role mapping
func (o ProjectOutput) UseProjectInternal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.BoolPtrOutput { return v.UseProjectInternal }).(pulumi.BoolPtrOutput)
}

func (o ProjectOutput) UserGroupLists() ProjectUserGroupListArrayOutput {
	return o.ApplyT(func(v *Project) ProjectUserGroupListArrayOutput { return v.UserGroupLists }).(ProjectUserGroupListArrayOutput)
}

func (o ProjectOutput) UserLists() ProjectUserListArrayOutput {
	return o.ApplyT(func(v *Project) ProjectUserListArrayOutput { return v.UserLists }).(ProjectUserListArrayOutput)
}

// List of Reference of users.
func (o ProjectOutput) UserReferenceLists() ProjectUserReferenceListArrayOutput {
	return o.ApplyT(func(v *Project) ProjectUserReferenceListArrayOutput { return v.UserReferenceLists }).(ProjectUserReferenceListArrayOutput)
}

func (o ProjectOutput) VpcReferenceLists() ProjectVpcReferenceListArrayOutput {
	return o.ApplyT(func(v *Project) ProjectVpcReferenceListArrayOutput { return v.VpcReferenceLists }).(ProjectVpcReferenceListArrayOutput)
}

type ProjectArrayOutput struct{ *pulumi.OutputState }

func (ProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Project)(nil)).Elem()
}

func (o ProjectArrayOutput) ToProjectArrayOutput() ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Project] {
	return pulumix.Output[[]*Project]{
		OutputState: o.OutputState,
	}
}

func (o ProjectArrayOutput) Index(i pulumi.IntInput) ProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Project {
		return vs[0].([]*Project)[vs[1].(int)]
	}).(ProjectOutput)
}

type ProjectMapOutput struct{ *pulumi.OutputState }

func (ProjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Project)(nil)).Elem()
}

func (o ProjectMapOutput) ToProjectMapOutput() ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) ToProjectMapOutputWithContext(ctx context.Context) ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Project] {
	return pulumix.Output[map[string]*Project]{
		OutputState: o.OutputState,
	}
}

func (o ProjectMapOutput) MapIndex(k pulumi.StringInput) ProjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Project {
		return vs[0].(map[string]*Project)[vs[1].(string)]
	}).(ProjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectInput)(nil)).Elem(), &Project{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectArrayInput)(nil)).Elem(), ProjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMapInput)(nil)).Elem(), ProjectMap{})
	pulumi.RegisterOutputType(ProjectOutput{})
	pulumi.RegisterOutputType(ProjectArrayOutput{})
	pulumi.RegisterOutputType(ProjectMapOutput{})
}
