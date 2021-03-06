// *** WARNING: this file was generated by crd2pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AzurePodIdentityException contains the pod selectors for all pods that don't require NMI to process and request token on their behalf.
type AzurePodIdentityException struct {
	pulumi.CustomResourceState

	ApiVersion pulumi.StringPtrOutput     `pulumi:"apiVersion"`
	Kind       pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata   metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// AzurePodIdentityExceptionSpec matches pods with the selector defined. If request originates from a pod that matches the selector, nmi will proxy the request and send response back without any validation.
	Spec AzurePodIdentityExceptionSpecPtrOutput `pulumi:"spec"`
	// AzurePodIdentityExceptionStatus contains the status of an AzurePodIdentityException.
	Status AzurePodIdentityExceptionStatusPtrOutput `pulumi:"status"`
}

// NewAzurePodIdentityException registers a new resource with the given unique name, arguments, and options.
func NewAzurePodIdentityException(ctx *pulumi.Context,
	name string, args *AzurePodIdentityExceptionArgs, opts ...pulumi.ResourceOption) (*AzurePodIdentityException, error) {
	if args == nil {
		args = &AzurePodIdentityExceptionArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("aadpodidentity.k8s.io/v1")
	args.Kind = pulumi.StringPtr("AzurePodIdentityException")
	var resource AzurePodIdentityException
	err := ctx.RegisterResource("kubernetes:aadpodidentity.k8s.io/v1:AzurePodIdentityException", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAzurePodIdentityException gets an existing AzurePodIdentityException resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAzurePodIdentityException(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzurePodIdentityExceptionState, opts ...pulumi.ResourceOption) (*AzurePodIdentityException, error) {
	var resource AzurePodIdentityException
	err := ctx.ReadResource("kubernetes:aadpodidentity.k8s.io/v1:AzurePodIdentityException", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AzurePodIdentityException resources.
type azurePodIdentityExceptionState struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// AzurePodIdentityExceptionSpec matches pods with the selector defined. If request originates from a pod that matches the selector, nmi will proxy the request and send response back without any validation.
	Spec *AzurePodIdentityExceptionSpec `pulumi:"spec"`
	// AzurePodIdentityExceptionStatus contains the status of an AzurePodIdentityException.
	Status *AzurePodIdentityExceptionStatus `pulumi:"status"`
}

type AzurePodIdentityExceptionState struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	// AzurePodIdentityExceptionSpec matches pods with the selector defined. If request originates from a pod that matches the selector, nmi will proxy the request and send response back without any validation.
	Spec AzurePodIdentityExceptionSpecPtrInput
	// AzurePodIdentityExceptionStatus contains the status of an AzurePodIdentityException.
	Status AzurePodIdentityExceptionStatusPtrInput
}

func (AzurePodIdentityExceptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*azurePodIdentityExceptionState)(nil)).Elem()
}

type azurePodIdentityExceptionArgs struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// AzurePodIdentityExceptionSpec matches pods with the selector defined. If request originates from a pod that matches the selector, nmi will proxy the request and send response back without any validation.
	Spec *AzurePodIdentityExceptionSpec `pulumi:"spec"`
	// AzurePodIdentityExceptionStatus contains the status of an AzurePodIdentityException.
	Status *AzurePodIdentityExceptionStatus `pulumi:"status"`
}

// The set of arguments for constructing a AzurePodIdentityException resource.
type AzurePodIdentityExceptionArgs struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	// AzurePodIdentityExceptionSpec matches pods with the selector defined. If request originates from a pod that matches the selector, nmi will proxy the request and send response back without any validation.
	Spec AzurePodIdentityExceptionSpecPtrInput
	// AzurePodIdentityExceptionStatus contains the status of an AzurePodIdentityException.
	Status AzurePodIdentityExceptionStatusPtrInput
}

func (AzurePodIdentityExceptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azurePodIdentityExceptionArgs)(nil)).Elem()
}

type AzurePodIdentityExceptionInput interface {
	pulumi.Input

	ToAzurePodIdentityExceptionOutput() AzurePodIdentityExceptionOutput
	ToAzurePodIdentityExceptionOutputWithContext(ctx context.Context) AzurePodIdentityExceptionOutput
}

func (*AzurePodIdentityException) ElementType() reflect.Type {
	return reflect.TypeOf((*AzurePodIdentityException)(nil))
}

func (i *AzurePodIdentityException) ToAzurePodIdentityExceptionOutput() AzurePodIdentityExceptionOutput {
	return i.ToAzurePodIdentityExceptionOutputWithContext(context.Background())
}

func (i *AzurePodIdentityException) ToAzurePodIdentityExceptionOutputWithContext(ctx context.Context) AzurePodIdentityExceptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzurePodIdentityExceptionOutput)
}

type AzurePodIdentityExceptionOutput struct {
	*pulumi.OutputState
}

func (AzurePodIdentityExceptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzurePodIdentityException)(nil))
}

func (o AzurePodIdentityExceptionOutput) ToAzurePodIdentityExceptionOutput() AzurePodIdentityExceptionOutput {
	return o
}

func (o AzurePodIdentityExceptionOutput) ToAzurePodIdentityExceptionOutputWithContext(ctx context.Context) AzurePodIdentityExceptionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AzurePodIdentityExceptionOutput{})
}
