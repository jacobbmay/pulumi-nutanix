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

// Provides a resource to create a role based on the input parameters.
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
//			_, err := nutanix.NewRole(ctx, "test", &nutanix.RoleArgs{
//				Description: pulumi.String("DESCRIPTION"),
//				PermissionReferenceLists: nutanix.RolePermissionReferenceListArray{
//					&nutanix.RolePermissionReferenceListArgs{
//						Kind: pulumi.String("permission"),
//						Uuid: pulumi.String("ID OF PERMISSION"),
//					},
//					&nutanix.RolePermissionReferenceListArgs{
//						Kind: pulumi.String("permission"),
//						Uuid: pulumi.String("ID OF PERMISSION"),
//					},
//					&nutanix.RolePermissionReferenceListArgs{
//						Kind: pulumi.String("permission"),
//						Uuid: pulumi.String("ID OF PERMISSION"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Role struct {
	pulumi.CustomResourceState

	// The version of the API.
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// - (Optional) Categories for the role.
	Categories RoleCategoryArrayOutput `pulumi:"categories"`
	// - (Optional) The description of the role.
	Description pulumi.StringOutput `pulumi:"description"`
	// - The role kind metadata.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// - (Optional) Name of the role.
	Name pulumi.StringOutput `pulumi:"name"`
	// - (Optional) The reference to a user.
	OwnerReference RoleOwnerReferenceOutput `pulumi:"ownerReference"`
	// - (Required) List of permission references.
	PermissionReferenceLists RolePermissionReferenceListArrayOutput `pulumi:"permissionReferenceLists"`
	// - (Optional) The reference to a project.
	ProjectReference RoleProjectReferenceOutput `pulumi:"projectReference"`
	// - The state of the role.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewRole registers a new resource with the given unique name, arguments, and options.
func NewRole(ctx *pulumi.Context,
	name string, args *RoleArgs, opts ...pulumi.ResourceOption) (*Role, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PermissionReferenceLists == nil {
		return nil, errors.New("invalid value for required argument 'PermissionReferenceLists'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Role
	err := ctx.RegisterResource("nutanix:index/role:Role", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRole gets an existing Role resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleState, opts ...pulumi.ResourceOption) (*Role, error) {
	var resource Role
	err := ctx.ReadResource("nutanix:index/role:Role", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Role resources.
type roleState struct {
	// The version of the API.
	ApiVersion *string `pulumi:"apiVersion"`
	// - (Optional) Categories for the role.
	Categories []RoleCategory `pulumi:"categories"`
	// - (Optional) The description of the role.
	Description *string `pulumi:"description"`
	// - The role kind metadata.
	Metadata map[string]string `pulumi:"metadata"`
	// - (Optional) Name of the role.
	Name *string `pulumi:"name"`
	// - (Optional) The reference to a user.
	OwnerReference *RoleOwnerReference `pulumi:"ownerReference"`
	// - (Required) List of permission references.
	PermissionReferenceLists []RolePermissionReferenceList `pulumi:"permissionReferenceLists"`
	// - (Optional) The reference to a project.
	ProjectReference *RoleProjectReference `pulumi:"projectReference"`
	// - The state of the role.
	State *string `pulumi:"state"`
}

type RoleState struct {
	// The version of the API.
	ApiVersion pulumi.StringPtrInput
	// - (Optional) Categories for the role.
	Categories RoleCategoryArrayInput
	// - (Optional) The description of the role.
	Description pulumi.StringPtrInput
	// - The role kind metadata.
	Metadata pulumi.StringMapInput
	// - (Optional) Name of the role.
	Name pulumi.StringPtrInput
	// - (Optional) The reference to a user.
	OwnerReference RoleOwnerReferencePtrInput
	// - (Required) List of permission references.
	PermissionReferenceLists RolePermissionReferenceListArrayInput
	// - (Optional) The reference to a project.
	ProjectReference RoleProjectReferencePtrInput
	// - The state of the role.
	State pulumi.StringPtrInput
}

func (RoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleState)(nil)).Elem()
}

type roleArgs struct {
	// - (Optional) Categories for the role.
	Categories []RoleCategory `pulumi:"categories"`
	// - (Optional) The description of the role.
	Description *string `pulumi:"description"`
	// - (Optional) Name of the role.
	Name *string `pulumi:"name"`
	// - (Optional) The reference to a user.
	OwnerReference *RoleOwnerReference `pulumi:"ownerReference"`
	// - (Required) List of permission references.
	PermissionReferenceLists []RolePermissionReferenceList `pulumi:"permissionReferenceLists"`
	// - (Optional) The reference to a project.
	ProjectReference *RoleProjectReference `pulumi:"projectReference"`
}

// The set of arguments for constructing a Role resource.
type RoleArgs struct {
	// - (Optional) Categories for the role.
	Categories RoleCategoryArrayInput
	// - (Optional) The description of the role.
	Description pulumi.StringPtrInput
	// - (Optional) Name of the role.
	Name pulumi.StringPtrInput
	// - (Optional) The reference to a user.
	OwnerReference RoleOwnerReferencePtrInput
	// - (Required) List of permission references.
	PermissionReferenceLists RolePermissionReferenceListArrayInput
	// - (Optional) The reference to a project.
	ProjectReference RoleProjectReferencePtrInput
}

func (RoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleArgs)(nil)).Elem()
}

type RoleInput interface {
	pulumi.Input

	ToRoleOutput() RoleOutput
	ToRoleOutputWithContext(ctx context.Context) RoleOutput
}

func (*Role) ElementType() reflect.Type {
	return reflect.TypeOf((**Role)(nil)).Elem()
}

func (i *Role) ToRoleOutput() RoleOutput {
	return i.ToRoleOutputWithContext(context.Background())
}

func (i *Role) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleOutput)
}

func (i *Role) ToOutput(ctx context.Context) pulumix.Output[*Role] {
	return pulumix.Output[*Role]{
		OutputState: i.ToRoleOutputWithContext(ctx).OutputState,
	}
}

// RoleArrayInput is an input type that accepts RoleArray and RoleArrayOutput values.
// You can construct a concrete instance of `RoleArrayInput` via:
//
//	RoleArray{ RoleArgs{...} }
type RoleArrayInput interface {
	pulumi.Input

	ToRoleArrayOutput() RoleArrayOutput
	ToRoleArrayOutputWithContext(context.Context) RoleArrayOutput
}

type RoleArray []RoleInput

func (RoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Role)(nil)).Elem()
}

func (i RoleArray) ToRoleArrayOutput() RoleArrayOutput {
	return i.ToRoleArrayOutputWithContext(context.Background())
}

func (i RoleArray) ToRoleArrayOutputWithContext(ctx context.Context) RoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleArrayOutput)
}

func (i RoleArray) ToOutput(ctx context.Context) pulumix.Output[[]*Role] {
	return pulumix.Output[[]*Role]{
		OutputState: i.ToRoleArrayOutputWithContext(ctx).OutputState,
	}
}

// RoleMapInput is an input type that accepts RoleMap and RoleMapOutput values.
// You can construct a concrete instance of `RoleMapInput` via:
//
//	RoleMap{ "key": RoleArgs{...} }
type RoleMapInput interface {
	pulumi.Input

	ToRoleMapOutput() RoleMapOutput
	ToRoleMapOutputWithContext(context.Context) RoleMapOutput
}

type RoleMap map[string]RoleInput

func (RoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Role)(nil)).Elem()
}

func (i RoleMap) ToRoleMapOutput() RoleMapOutput {
	return i.ToRoleMapOutputWithContext(context.Background())
}

func (i RoleMap) ToRoleMapOutputWithContext(ctx context.Context) RoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleMapOutput)
}

func (i RoleMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Role] {
	return pulumix.Output[map[string]*Role]{
		OutputState: i.ToRoleMapOutputWithContext(ctx).OutputState,
	}
}

type RoleOutput struct{ *pulumi.OutputState }

func (RoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Role)(nil)).Elem()
}

func (o RoleOutput) ToRoleOutput() RoleOutput {
	return o
}

func (o RoleOutput) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return o
}

func (o RoleOutput) ToOutput(ctx context.Context) pulumix.Output[*Role] {
	return pulumix.Output[*Role]{
		OutputState: o.OutputState,
	}
}

// The version of the API.
func (o RoleOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// - (Optional) Categories for the role.
func (o RoleOutput) Categories() RoleCategoryArrayOutput {
	return o.ApplyT(func(v *Role) RoleCategoryArrayOutput { return v.Categories }).(RoleCategoryArrayOutput)
}

// - (Optional) The description of the role.
func (o RoleOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// - The role kind metadata.
func (o RoleOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Role) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// - (Optional) Name of the role.
func (o RoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// - (Optional) The reference to a user.
func (o RoleOutput) OwnerReference() RoleOwnerReferenceOutput {
	return o.ApplyT(func(v *Role) RoleOwnerReferenceOutput { return v.OwnerReference }).(RoleOwnerReferenceOutput)
}

// - (Required) List of permission references.
func (o RoleOutput) PermissionReferenceLists() RolePermissionReferenceListArrayOutput {
	return o.ApplyT(func(v *Role) RolePermissionReferenceListArrayOutput { return v.PermissionReferenceLists }).(RolePermissionReferenceListArrayOutput)
}

// - (Optional) The reference to a project.
func (o RoleOutput) ProjectReference() RoleProjectReferenceOutput {
	return o.ApplyT(func(v *Role) RoleProjectReferenceOutput { return v.ProjectReference }).(RoleProjectReferenceOutput)
}

// - The state of the role.
func (o RoleOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

type RoleArrayOutput struct{ *pulumi.OutputState }

func (RoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Role)(nil)).Elem()
}

func (o RoleArrayOutput) ToRoleArrayOutput() RoleArrayOutput {
	return o
}

func (o RoleArrayOutput) ToRoleArrayOutputWithContext(ctx context.Context) RoleArrayOutput {
	return o
}

func (o RoleArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Role] {
	return pulumix.Output[[]*Role]{
		OutputState: o.OutputState,
	}
}

func (o RoleArrayOutput) Index(i pulumi.IntInput) RoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Role {
		return vs[0].([]*Role)[vs[1].(int)]
	}).(RoleOutput)
}

type RoleMapOutput struct{ *pulumi.OutputState }

func (RoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Role)(nil)).Elem()
}

func (o RoleMapOutput) ToRoleMapOutput() RoleMapOutput {
	return o
}

func (o RoleMapOutput) ToRoleMapOutputWithContext(ctx context.Context) RoleMapOutput {
	return o
}

func (o RoleMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Role] {
	return pulumix.Output[map[string]*Role]{
		OutputState: o.OutputState,
	}
}

func (o RoleMapOutput) MapIndex(k pulumi.StringInput) RoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Role {
		return vs[0].(map[string]*Role)[vs[1].(string)]
	}).(RoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoleInput)(nil)).Elem(), &Role{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleArrayInput)(nil)).Elem(), RoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleMapInput)(nil)).Elem(), RoleMap{})
	pulumi.RegisterOutputType(RoleOutput{})
	pulumi.RegisterOutputType(RoleArrayOutput{})
	pulumi.RegisterOutputType(RoleMapOutput{})
}