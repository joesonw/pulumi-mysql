// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The ``Grant`` resource creates and manages privileges given to
// a user on a MySQL server.
//
// ## Examples
//
// ### Granting Privileges to a User
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-mysql/sdk/v3/go/mysql"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		jdoeUser, err := mysql.NewUser(ctx, "jdoeUser", &mysql.UserArgs{
// 			Host:              pulumi.String("example.com"),
// 			PlaintextPassword: pulumi.String("password"),
// 			User:              pulumi.String("jdoe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = mysql.NewGrant(ctx, "jdoeGrant", &mysql.GrantArgs{
// 			Database: pulumi.String("app"),
// 			Host:     jdoeUser.Host,
// 			Privileges: pulumi.StringArray{
// 				pulumi.String("SELECT"),
// 				pulumi.String("UPDATE"),
// 			},
// 			User: jdoeUser.User,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ### Granting Privileges to a Role
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-mysql/sdk/v3/go/mysql"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		developerRole, err := mysql.NewRole(ctx, "developerRole", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = mysql.NewGrant(ctx, "developerGrant", &mysql.GrantArgs{
// 			Database: pulumi.String("app"),
// 			Privileges: pulumi.StringArray{
// 				pulumi.String("SELECT"),
// 				pulumi.String("UPDATE"),
// 			},
// 			Role: developerRole.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ### Adding a Role to a User
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-mysql/sdk/v3/go/mysql"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		jdoe, err := mysql.NewUser(ctx, "jdoe", &mysql.UserArgs{
// 			Host:              pulumi.String("example.com"),
// 			PlaintextPassword: pulumi.String("password"),
// 			User:              pulumi.String("jdoe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		developerRole, err := mysql.NewRole(ctx, "developerRole", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = mysql.NewGrant(ctx, "developerGrant", &mysql.GrantArgs{
// 			Database: pulumi.String("app"),
// 			Host:     jdoe.Host,
// 			Roles: pulumi.StringArray{
// 				developerRole.Name,
// 			},
// 			User: jdoe.User,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Grant struct {
	pulumi.CustomResourceState

	// The database to grant privileges on.
	Database pulumi.StringOutput `pulumi:"database"`
	// Whether to also give the user privileges to grant the same privileges to other users.
	Grant pulumi.BoolPtrOutput `pulumi:"grant"`
	// The source host of the user. Defaults to "localhost". Conflicts with `role`.
	Host pulumi.StringPtrOutput `pulumi:"host"`
	// A list of privileges to grant to the user. Refer to a list of privileges (such as [here](https://dev.mysql.com/doc/refman/5.5/en/grant.html)) for applicable privileges. Conflicts with `roles`.
	Privileges pulumi.StringArrayOutput `pulumi:"privileges"`
	// The role to grant `privileges` to. Conflicts with `user` and `host`.
	Role pulumi.StringPtrOutput `pulumi:"role"`
	// A list of rols to grant to the user. Conflicts with `privileges`.
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// Which table to grant `privileges` on. Defaults to `*`, which is all tables.
	Table pulumi.StringPtrOutput `pulumi:"table"`
	// An TLS-Option for the `GRANT` statement. The value is suffixed to `REQUIRE`. A value of 'SSL' will generate a `GRANT ... REQUIRE SSL` statement. See the [MYSQL `GRANT` documentation](https://dev.mysql.com/doc/refman/5.7/en/grant.html) for more. Ignored if MySQL version is under 5.7.0.
	TlsOption pulumi.StringPtrOutput `pulumi:"tlsOption"`
	// The name of the user. Conflicts with `role`.
	User pulumi.StringPtrOutput `pulumi:"user"`
}

// NewGrant registers a new resource with the given unique name, arguments, and options.
func NewGrant(ctx *pulumi.Context,
	name string, args *GrantArgs, opts ...pulumi.ResourceOption) (*Grant, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Database == nil {
		return nil, errors.New("invalid value for required argument 'Database'")
	}
	var resource Grant
	err := ctx.RegisterResource("mysql:index/grant:Grant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGrant gets an existing Grant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGrant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GrantState, opts ...pulumi.ResourceOption) (*Grant, error) {
	var resource Grant
	err := ctx.ReadResource("mysql:index/grant:Grant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Grant resources.
type grantState struct {
	// The database to grant privileges on.
	Database *string `pulumi:"database"`
	// Whether to also give the user privileges to grant the same privileges to other users.
	Grant *bool `pulumi:"grant"`
	// The source host of the user. Defaults to "localhost". Conflicts with `role`.
	Host *string `pulumi:"host"`
	// A list of privileges to grant to the user. Refer to a list of privileges (such as [here](https://dev.mysql.com/doc/refman/5.5/en/grant.html)) for applicable privileges. Conflicts with `roles`.
	Privileges []string `pulumi:"privileges"`
	// The role to grant `privileges` to. Conflicts with `user` and `host`.
	Role *string `pulumi:"role"`
	// A list of rols to grant to the user. Conflicts with `privileges`.
	Roles []string `pulumi:"roles"`
	// Which table to grant `privileges` on. Defaults to `*`, which is all tables.
	Table *string `pulumi:"table"`
	// An TLS-Option for the `GRANT` statement. The value is suffixed to `REQUIRE`. A value of 'SSL' will generate a `GRANT ... REQUIRE SSL` statement. See the [MYSQL `GRANT` documentation](https://dev.mysql.com/doc/refman/5.7/en/grant.html) for more. Ignored if MySQL version is under 5.7.0.
	TlsOption *string `pulumi:"tlsOption"`
	// The name of the user. Conflicts with `role`.
	User *string `pulumi:"user"`
}

type GrantState struct {
	// The database to grant privileges on.
	Database pulumi.StringPtrInput
	// Whether to also give the user privileges to grant the same privileges to other users.
	Grant pulumi.BoolPtrInput
	// The source host of the user. Defaults to "localhost". Conflicts with `role`.
	Host pulumi.StringPtrInput
	// A list of privileges to grant to the user. Refer to a list of privileges (such as [here](https://dev.mysql.com/doc/refman/5.5/en/grant.html)) for applicable privileges. Conflicts with `roles`.
	Privileges pulumi.StringArrayInput
	// The role to grant `privileges` to. Conflicts with `user` and `host`.
	Role pulumi.StringPtrInput
	// A list of rols to grant to the user. Conflicts with `privileges`.
	Roles pulumi.StringArrayInput
	// Which table to grant `privileges` on. Defaults to `*`, which is all tables.
	Table pulumi.StringPtrInput
	// An TLS-Option for the `GRANT` statement. The value is suffixed to `REQUIRE`. A value of 'SSL' will generate a `GRANT ... REQUIRE SSL` statement. See the [MYSQL `GRANT` documentation](https://dev.mysql.com/doc/refman/5.7/en/grant.html) for more. Ignored if MySQL version is under 5.7.0.
	TlsOption pulumi.StringPtrInput
	// The name of the user. Conflicts with `role`.
	User pulumi.StringPtrInput
}

func (GrantState) ElementType() reflect.Type {
	return reflect.TypeOf((*grantState)(nil)).Elem()
}

type grantArgs struct {
	// The database to grant privileges on.
	Database string `pulumi:"database"`
	// Whether to also give the user privileges to grant the same privileges to other users.
	Grant *bool `pulumi:"grant"`
	// The source host of the user. Defaults to "localhost". Conflicts with `role`.
	Host *string `pulumi:"host"`
	// A list of privileges to grant to the user. Refer to a list of privileges (such as [here](https://dev.mysql.com/doc/refman/5.5/en/grant.html)) for applicable privileges. Conflicts with `roles`.
	Privileges []string `pulumi:"privileges"`
	// The role to grant `privileges` to. Conflicts with `user` and `host`.
	Role *string `pulumi:"role"`
	// A list of rols to grant to the user. Conflicts with `privileges`.
	Roles []string `pulumi:"roles"`
	// Which table to grant `privileges` on. Defaults to `*`, which is all tables.
	Table *string `pulumi:"table"`
	// An TLS-Option for the `GRANT` statement. The value is suffixed to `REQUIRE`. A value of 'SSL' will generate a `GRANT ... REQUIRE SSL` statement. See the [MYSQL `GRANT` documentation](https://dev.mysql.com/doc/refman/5.7/en/grant.html) for more. Ignored if MySQL version is under 5.7.0.
	TlsOption *string `pulumi:"tlsOption"`
	// The name of the user. Conflicts with `role`.
	User *string `pulumi:"user"`
}

// The set of arguments for constructing a Grant resource.
type GrantArgs struct {
	// The database to grant privileges on.
	Database pulumi.StringInput
	// Whether to also give the user privileges to grant the same privileges to other users.
	Grant pulumi.BoolPtrInput
	// The source host of the user. Defaults to "localhost". Conflicts with `role`.
	Host pulumi.StringPtrInput
	// A list of privileges to grant to the user. Refer to a list of privileges (such as [here](https://dev.mysql.com/doc/refman/5.5/en/grant.html)) for applicable privileges. Conflicts with `roles`.
	Privileges pulumi.StringArrayInput
	// The role to grant `privileges` to. Conflicts with `user` and `host`.
	Role pulumi.StringPtrInput
	// A list of rols to grant to the user. Conflicts with `privileges`.
	Roles pulumi.StringArrayInput
	// Which table to grant `privileges` on. Defaults to `*`, which is all tables.
	Table pulumi.StringPtrInput
	// An TLS-Option for the `GRANT` statement. The value is suffixed to `REQUIRE`. A value of 'SSL' will generate a `GRANT ... REQUIRE SSL` statement. See the [MYSQL `GRANT` documentation](https://dev.mysql.com/doc/refman/5.7/en/grant.html) for more. Ignored if MySQL version is under 5.7.0.
	TlsOption pulumi.StringPtrInput
	// The name of the user. Conflicts with `role`.
	User pulumi.StringPtrInput
}

func (GrantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*grantArgs)(nil)).Elem()
}

type GrantInput interface {
	pulumi.Input

	ToGrantOutput() GrantOutput
	ToGrantOutputWithContext(ctx context.Context) GrantOutput
}

func (*Grant) ElementType() reflect.Type {
	return reflect.TypeOf((*Grant)(nil))
}

func (i *Grant) ToGrantOutput() GrantOutput {
	return i.ToGrantOutputWithContext(context.Background())
}

func (i *Grant) ToGrantOutputWithContext(ctx context.Context) GrantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrantOutput)
}

func (i *Grant) ToGrantPtrOutput() GrantPtrOutput {
	return i.ToGrantPtrOutputWithContext(context.Background())
}

func (i *Grant) ToGrantPtrOutputWithContext(ctx context.Context) GrantPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrantPtrOutput)
}

type GrantPtrInput interface {
	pulumi.Input

	ToGrantPtrOutput() GrantPtrOutput
	ToGrantPtrOutputWithContext(ctx context.Context) GrantPtrOutput
}

type grantPtrType GrantArgs

func (*grantPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Grant)(nil))
}

func (i *grantPtrType) ToGrantPtrOutput() GrantPtrOutput {
	return i.ToGrantPtrOutputWithContext(context.Background())
}

func (i *grantPtrType) ToGrantPtrOutputWithContext(ctx context.Context) GrantPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrantPtrOutput)
}

// GrantArrayInput is an input type that accepts GrantArray and GrantArrayOutput values.
// You can construct a concrete instance of `GrantArrayInput` via:
//
//          GrantArray{ GrantArgs{...} }
type GrantArrayInput interface {
	pulumi.Input

	ToGrantArrayOutput() GrantArrayOutput
	ToGrantArrayOutputWithContext(context.Context) GrantArrayOutput
}

type GrantArray []GrantInput

func (GrantArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Grant)(nil))
}

func (i GrantArray) ToGrantArrayOutput() GrantArrayOutput {
	return i.ToGrantArrayOutputWithContext(context.Background())
}

func (i GrantArray) ToGrantArrayOutputWithContext(ctx context.Context) GrantArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrantArrayOutput)
}

// GrantMapInput is an input type that accepts GrantMap and GrantMapOutput values.
// You can construct a concrete instance of `GrantMapInput` via:
//
//          GrantMap{ "key": GrantArgs{...} }
type GrantMapInput interface {
	pulumi.Input

	ToGrantMapOutput() GrantMapOutput
	ToGrantMapOutputWithContext(context.Context) GrantMapOutput
}

type GrantMap map[string]GrantInput

func (GrantMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Grant)(nil))
}

func (i GrantMap) ToGrantMapOutput() GrantMapOutput {
	return i.ToGrantMapOutputWithContext(context.Background())
}

func (i GrantMap) ToGrantMapOutputWithContext(ctx context.Context) GrantMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrantMapOutput)
}

type GrantOutput struct {
	*pulumi.OutputState
}

func (GrantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Grant)(nil))
}

func (o GrantOutput) ToGrantOutput() GrantOutput {
	return o
}

func (o GrantOutput) ToGrantOutputWithContext(ctx context.Context) GrantOutput {
	return o
}

func (o GrantOutput) ToGrantPtrOutput() GrantPtrOutput {
	return o.ToGrantPtrOutputWithContext(context.Background())
}

func (o GrantOutput) ToGrantPtrOutputWithContext(ctx context.Context) GrantPtrOutput {
	return o.ApplyT(func(v Grant) *Grant {
		return &v
	}).(GrantPtrOutput)
}

type GrantPtrOutput struct {
	*pulumi.OutputState
}

func (GrantPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Grant)(nil))
}

func (o GrantPtrOutput) ToGrantPtrOutput() GrantPtrOutput {
	return o
}

func (o GrantPtrOutput) ToGrantPtrOutputWithContext(ctx context.Context) GrantPtrOutput {
	return o
}

type GrantArrayOutput struct{ *pulumi.OutputState }

func (GrantArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Grant)(nil))
}

func (o GrantArrayOutput) ToGrantArrayOutput() GrantArrayOutput {
	return o
}

func (o GrantArrayOutput) ToGrantArrayOutputWithContext(ctx context.Context) GrantArrayOutput {
	return o
}

func (o GrantArrayOutput) Index(i pulumi.IntInput) GrantOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Grant {
		return vs[0].([]Grant)[vs[1].(int)]
	}).(GrantOutput)
}

type GrantMapOutput struct{ *pulumi.OutputState }

func (GrantMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Grant)(nil))
}

func (o GrantMapOutput) ToGrantMapOutput() GrantMapOutput {
	return o
}

func (o GrantMapOutput) ToGrantMapOutputWithContext(ctx context.Context) GrantMapOutput {
	return o
}

func (o GrantMapOutput) MapIndex(k pulumi.StringInput) GrantOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Grant {
		return vs[0].(map[string]Grant)[vs[1].(string)]
	}).(GrantOutput)
}

func init() {
	pulumi.RegisterOutputType(GrantOutput{})
	pulumi.RegisterOutputType(GrantPtrOutput{})
	pulumi.RegisterOutputType(GrantArrayOutput{})
	pulumi.RegisterOutputType(GrantMapOutput{})
}
