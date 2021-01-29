// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `UserPassword` resource sets and manages a password for a given
// user on a MySQL server.
//
// > **NOTE on MySQL Passwords:** This resource conflicts with the `password`
//    argument for `User`. This resource uses PGP encryption to avoid
//    storing unencrypted passwords in the provider state.
//
// > **NOTE on How Passwords are Created:** This resource **automatically**
//    generates a **random** password. The password will be a random UUID.
type UserPassword struct {
	pulumi.CustomResourceState

	// The encrypted password, base64 encoded.
	EncryptedPassword pulumi.StringOutput `pulumi:"encryptedPassword"`
	// The source host of the user. Defaults to `localhost`.
	Host pulumi.StringPtrOutput `pulumi:"host"`
	// The fingerprint of the PGP key used to encrypt the password
	KeyFingerprint pulumi.StringOutput `pulumi:"keyFingerprint"`
	// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:some_person_that_exists`.
	PgpKey pulumi.StringOutput `pulumi:"pgpKey"`
	// The IAM user to associate with this access key.
	User pulumi.StringOutput `pulumi:"user"`
}

// NewUserPassword registers a new resource with the given unique name, arguments, and options.
func NewUserPassword(ctx *pulumi.Context,
	name string, args *UserPasswordArgs, opts ...pulumi.ResourceOption) (*UserPassword, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PgpKey == nil {
		return nil, errors.New("invalid value for required argument 'PgpKey'")
	}
	if args.User == nil {
		return nil, errors.New("invalid value for required argument 'User'")
	}
	var resource UserPassword
	err := ctx.RegisterResource("mysql:index/userPassword:UserPassword", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserPassword gets an existing UserPassword resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserPassword(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserPasswordState, opts ...pulumi.ResourceOption) (*UserPassword, error) {
	var resource UserPassword
	err := ctx.ReadResource("mysql:index/userPassword:UserPassword", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserPassword resources.
type userPasswordState struct {
	// The encrypted password, base64 encoded.
	EncryptedPassword *string `pulumi:"encryptedPassword"`
	// The source host of the user. Defaults to `localhost`.
	Host *string `pulumi:"host"`
	// The fingerprint of the PGP key used to encrypt the password
	KeyFingerprint *string `pulumi:"keyFingerprint"`
	// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:some_person_that_exists`.
	PgpKey *string `pulumi:"pgpKey"`
	// The IAM user to associate with this access key.
	User *string `pulumi:"user"`
}

type UserPasswordState struct {
	// The encrypted password, base64 encoded.
	EncryptedPassword pulumi.StringPtrInput
	// The source host of the user. Defaults to `localhost`.
	Host pulumi.StringPtrInput
	// The fingerprint of the PGP key used to encrypt the password
	KeyFingerprint pulumi.StringPtrInput
	// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:some_person_that_exists`.
	PgpKey pulumi.StringPtrInput
	// The IAM user to associate with this access key.
	User pulumi.StringPtrInput
}

func (UserPasswordState) ElementType() reflect.Type {
	return reflect.TypeOf((*userPasswordState)(nil)).Elem()
}

type userPasswordArgs struct {
	// The source host of the user. Defaults to `localhost`.
	Host *string `pulumi:"host"`
	// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:some_person_that_exists`.
	PgpKey string `pulumi:"pgpKey"`
	// The IAM user to associate with this access key.
	User string `pulumi:"user"`
}

// The set of arguments for constructing a UserPassword resource.
type UserPasswordArgs struct {
	// The source host of the user. Defaults to `localhost`.
	Host pulumi.StringPtrInput
	// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:some_person_that_exists`.
	PgpKey pulumi.StringInput
	// The IAM user to associate with this access key.
	User pulumi.StringInput
}

func (UserPasswordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userPasswordArgs)(nil)).Elem()
}

type UserPasswordInput interface {
	pulumi.Input

	ToUserPasswordOutput() UserPasswordOutput
	ToUserPasswordOutputWithContext(ctx context.Context) UserPasswordOutput
}

func (*UserPassword) ElementType() reflect.Type {
	return reflect.TypeOf((*UserPassword)(nil))
}

func (i *UserPassword) ToUserPasswordOutput() UserPasswordOutput {
	return i.ToUserPasswordOutputWithContext(context.Background())
}

func (i *UserPassword) ToUserPasswordOutputWithContext(ctx context.Context) UserPasswordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPasswordOutput)
}

type UserPasswordOutput struct {
	*pulumi.OutputState
}

func (UserPasswordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserPassword)(nil))
}

func (o UserPasswordOutput) ToUserPasswordOutput() UserPasswordOutput {
	return o
}

func (o UserPasswordOutput) ToUserPasswordOutputWithContext(ctx context.Context) UserPasswordOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(UserPasswordOutput{})
}
