// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha2

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Patch resources are used to modify existing Kubernetes resources by using
// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
// ResourceClaimTemplate is used to produce ResourceClaim objects.
type ResourceClaimTemplatePatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// Describes the ResourceClaim that is to be generated.
	//
	// This field is immutable. A ResourceClaim will get created by the control plane for a Pod when needed and then not get updated anymore.
	Spec ResourceClaimTemplateSpecPatchPtrOutput `pulumi:"spec"`
}

// NewResourceClaimTemplatePatch registers a new resource with the given unique name, arguments, and options.
func NewResourceClaimTemplatePatch(ctx *pulumi.Context,
	name string, args *ResourceClaimTemplatePatchArgs, opts ...pulumi.ResourceOption) (*ResourceClaimTemplatePatch, error) {
	if args == nil {
		args = &ResourceClaimTemplatePatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("resource.k8s.io/v1alpha2")
	args.Kind = pulumi.StringPtr("ResourceClaimTemplate")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:resource.k8s.io/v1alpha1:ResourceClaimTemplatePatch"),
		},
		{
			Type: pulumi.String("kubernetes:resource.k8s.io/v1alpha3:ResourceClaimTemplatePatch"),
		},
		{
			Type: pulumi.String("kubernetes:resource.k8s.io/v1beta1:ResourceClaimTemplatePatch"),
		},
		{
			Type: pulumi.String("kubernetes:resource.k8s.io/v1beta2:ResourceClaimTemplatePatch"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ResourceClaimTemplatePatch
	err := ctx.RegisterResource("kubernetes:resource.k8s.io/v1alpha2:ResourceClaimTemplatePatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceClaimTemplatePatch gets an existing ResourceClaimTemplatePatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceClaimTemplatePatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceClaimTemplatePatchState, opts ...pulumi.ResourceOption) (*ResourceClaimTemplatePatch, error) {
	var resource ResourceClaimTemplatePatch
	err := ctx.ReadResource("kubernetes:resource.k8s.io/v1alpha2:ResourceClaimTemplatePatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceClaimTemplatePatch resources.
type resourceClaimTemplatePatchState struct {
}

type ResourceClaimTemplatePatchState struct {
}

func (ResourceClaimTemplatePatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceClaimTemplatePatchState)(nil)).Elem()
}

type resourceClaimTemplatePatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// Describes the ResourceClaim that is to be generated.
	//
	// This field is immutable. A ResourceClaim will get created by the control plane for a Pod when needed and then not get updated anymore.
	Spec *ResourceClaimTemplateSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a ResourceClaimTemplatePatch resource.
type ResourceClaimTemplatePatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	// Describes the ResourceClaim that is to be generated.
	//
	// This field is immutable. A ResourceClaim will get created by the control plane for a Pod when needed and then not get updated anymore.
	Spec ResourceClaimTemplateSpecPatchPtrInput
}

func (ResourceClaimTemplatePatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceClaimTemplatePatchArgs)(nil)).Elem()
}

type ResourceClaimTemplatePatchInput interface {
	pulumi.Input

	ToResourceClaimTemplatePatchOutput() ResourceClaimTemplatePatchOutput
	ToResourceClaimTemplatePatchOutputWithContext(ctx context.Context) ResourceClaimTemplatePatchOutput
}

func (*ResourceClaimTemplatePatch) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceClaimTemplatePatch)(nil)).Elem()
}

func (i *ResourceClaimTemplatePatch) ToResourceClaimTemplatePatchOutput() ResourceClaimTemplatePatchOutput {
	return i.ToResourceClaimTemplatePatchOutputWithContext(context.Background())
}

func (i *ResourceClaimTemplatePatch) ToResourceClaimTemplatePatchOutputWithContext(ctx context.Context) ResourceClaimTemplatePatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceClaimTemplatePatchOutput)
}

// ResourceClaimTemplatePatchArrayInput is an input type that accepts ResourceClaimTemplatePatchArray and ResourceClaimTemplatePatchArrayOutput values.
// You can construct a concrete instance of `ResourceClaimTemplatePatchArrayInput` via:
//
//	ResourceClaimTemplatePatchArray{ ResourceClaimTemplatePatchArgs{...} }
type ResourceClaimTemplatePatchArrayInput interface {
	pulumi.Input

	ToResourceClaimTemplatePatchArrayOutput() ResourceClaimTemplatePatchArrayOutput
	ToResourceClaimTemplatePatchArrayOutputWithContext(context.Context) ResourceClaimTemplatePatchArrayOutput
}

type ResourceClaimTemplatePatchArray []ResourceClaimTemplatePatchInput

func (ResourceClaimTemplatePatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceClaimTemplatePatch)(nil)).Elem()
}

func (i ResourceClaimTemplatePatchArray) ToResourceClaimTemplatePatchArrayOutput() ResourceClaimTemplatePatchArrayOutput {
	return i.ToResourceClaimTemplatePatchArrayOutputWithContext(context.Background())
}

func (i ResourceClaimTemplatePatchArray) ToResourceClaimTemplatePatchArrayOutputWithContext(ctx context.Context) ResourceClaimTemplatePatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceClaimTemplatePatchArrayOutput)
}

// ResourceClaimTemplatePatchMapInput is an input type that accepts ResourceClaimTemplatePatchMap and ResourceClaimTemplatePatchMapOutput values.
// You can construct a concrete instance of `ResourceClaimTemplatePatchMapInput` via:
//
//	ResourceClaimTemplatePatchMap{ "key": ResourceClaimTemplatePatchArgs{...} }
type ResourceClaimTemplatePatchMapInput interface {
	pulumi.Input

	ToResourceClaimTemplatePatchMapOutput() ResourceClaimTemplatePatchMapOutput
	ToResourceClaimTemplatePatchMapOutputWithContext(context.Context) ResourceClaimTemplatePatchMapOutput
}

type ResourceClaimTemplatePatchMap map[string]ResourceClaimTemplatePatchInput

func (ResourceClaimTemplatePatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceClaimTemplatePatch)(nil)).Elem()
}

func (i ResourceClaimTemplatePatchMap) ToResourceClaimTemplatePatchMapOutput() ResourceClaimTemplatePatchMapOutput {
	return i.ToResourceClaimTemplatePatchMapOutputWithContext(context.Background())
}

func (i ResourceClaimTemplatePatchMap) ToResourceClaimTemplatePatchMapOutputWithContext(ctx context.Context) ResourceClaimTemplatePatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceClaimTemplatePatchMapOutput)
}

type ResourceClaimTemplatePatchOutput struct{ *pulumi.OutputState }

func (ResourceClaimTemplatePatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceClaimTemplatePatch)(nil)).Elem()
}

func (o ResourceClaimTemplatePatchOutput) ToResourceClaimTemplatePatchOutput() ResourceClaimTemplatePatchOutput {
	return o
}

func (o ResourceClaimTemplatePatchOutput) ToResourceClaimTemplatePatchOutputWithContext(ctx context.Context) ResourceClaimTemplatePatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ResourceClaimTemplatePatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceClaimTemplatePatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ResourceClaimTemplatePatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceClaimTemplatePatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object metadata
func (o ResourceClaimTemplatePatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *ResourceClaimTemplatePatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// Describes the ResourceClaim that is to be generated.
//
// This field is immutable. A ResourceClaim will get created by the control plane for a Pod when needed and then not get updated anymore.
func (o ResourceClaimTemplatePatchOutput) Spec() ResourceClaimTemplateSpecPatchPtrOutput {
	return o.ApplyT(func(v *ResourceClaimTemplatePatch) ResourceClaimTemplateSpecPatchPtrOutput { return v.Spec }).(ResourceClaimTemplateSpecPatchPtrOutput)
}

type ResourceClaimTemplatePatchArrayOutput struct{ *pulumi.OutputState }

func (ResourceClaimTemplatePatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceClaimTemplatePatch)(nil)).Elem()
}

func (o ResourceClaimTemplatePatchArrayOutput) ToResourceClaimTemplatePatchArrayOutput() ResourceClaimTemplatePatchArrayOutput {
	return o
}

func (o ResourceClaimTemplatePatchArrayOutput) ToResourceClaimTemplatePatchArrayOutputWithContext(ctx context.Context) ResourceClaimTemplatePatchArrayOutput {
	return o
}

func (o ResourceClaimTemplatePatchArrayOutput) Index(i pulumi.IntInput) ResourceClaimTemplatePatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourceClaimTemplatePatch {
		return vs[0].([]*ResourceClaimTemplatePatch)[vs[1].(int)]
	}).(ResourceClaimTemplatePatchOutput)
}

type ResourceClaimTemplatePatchMapOutput struct{ *pulumi.OutputState }

func (ResourceClaimTemplatePatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceClaimTemplatePatch)(nil)).Elem()
}

func (o ResourceClaimTemplatePatchMapOutput) ToResourceClaimTemplatePatchMapOutput() ResourceClaimTemplatePatchMapOutput {
	return o
}

func (o ResourceClaimTemplatePatchMapOutput) ToResourceClaimTemplatePatchMapOutputWithContext(ctx context.Context) ResourceClaimTemplatePatchMapOutput {
	return o
}

func (o ResourceClaimTemplatePatchMapOutput) MapIndex(k pulumi.StringInput) ResourceClaimTemplatePatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourceClaimTemplatePatch {
		return vs[0].(map[string]*ResourceClaimTemplatePatch)[vs[1].(string)]
	}).(ResourceClaimTemplatePatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceClaimTemplatePatchInput)(nil)).Elem(), &ResourceClaimTemplatePatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceClaimTemplatePatchArrayInput)(nil)).Elem(), ResourceClaimTemplatePatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceClaimTemplatePatchMapInput)(nil)).Elem(), ResourceClaimTemplatePatchMap{})
	pulumi.RegisterOutputType(ResourceClaimTemplatePatchOutput{})
	pulumi.RegisterOutputType(ResourceClaimTemplatePatchArrayOutput{})
	pulumi.RegisterOutputType(ResourceClaimTemplatePatchMapOutput{})
}
