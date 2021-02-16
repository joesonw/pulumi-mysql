// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The ``Database`` resource creates and manages a database on a MySQL
// server.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-mysql/sdk/v2/go/mysql"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := mysql.NewDatabase(ctx, "app", nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Databases can be imported using their name, e.g.
//
// ```sh
//  $ pulumi import mysql:index/database:Database example my-example-database
// ```
type Database struct {
	pulumi.CustomResourceState

	// The default character set to use when
	// a table is created without specifying an explicit character set. Defaults
	// to "utf8".
	DefaultCharacterSet pulumi.StringPtrOutput `pulumi:"defaultCharacterSet"`
	// The default collation to use when a table
	// is created without specifying an explicit collation. Defaults to
	// ``utf8GeneralCi``. Each character set has its own set of collations, so
	// changing the character set requires also changing the collation.
	DefaultCollation pulumi.StringPtrOutput `pulumi:"defaultCollation"`
	// The name of the database. This must be unique within
	// a given MySQL server and may or may not be case-sensitive depending on
	// the operating system on which the MySQL server is running.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil {
		args = &DatabaseArgs{}
	}

	var resource Database
	err := ctx.RegisterResource("mysql:index/database:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("mysql:index/database:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
	// The default character set to use when
	// a table is created without specifying an explicit character set. Defaults
	// to "utf8".
	DefaultCharacterSet *string `pulumi:"defaultCharacterSet"`
	// The default collation to use when a table
	// is created without specifying an explicit collation. Defaults to
	// ``utf8GeneralCi``. Each character set has its own set of collations, so
	// changing the character set requires also changing the collation.
	DefaultCollation *string `pulumi:"defaultCollation"`
	// The name of the database. This must be unique within
	// a given MySQL server and may or may not be case-sensitive depending on
	// the operating system on which the MySQL server is running.
	Name *string `pulumi:"name"`
}

type DatabaseState struct {
	// The default character set to use when
	// a table is created without specifying an explicit character set. Defaults
	// to "utf8".
	DefaultCharacterSet pulumi.StringPtrInput
	// The default collation to use when a table
	// is created without specifying an explicit collation. Defaults to
	// ``utf8GeneralCi``. Each character set has its own set of collations, so
	// changing the character set requires also changing the collation.
	DefaultCollation pulumi.StringPtrInput
	// The name of the database. This must be unique within
	// a given MySQL server and may or may not be case-sensitive depending on
	// the operating system on which the MySQL server is running.
	Name pulumi.StringPtrInput
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// The default character set to use when
	// a table is created without specifying an explicit character set. Defaults
	// to "utf8".
	DefaultCharacterSet *string `pulumi:"defaultCharacterSet"`
	// The default collation to use when a table
	// is created without specifying an explicit collation. Defaults to
	// ``utf8GeneralCi``. Each character set has its own set of collations, so
	// changing the character set requires also changing the collation.
	DefaultCollation *string `pulumi:"defaultCollation"`
	// The name of the database. This must be unique within
	// a given MySQL server and may or may not be case-sensitive depending on
	// the operating system on which the MySQL server is running.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// The default character set to use when
	// a table is created without specifying an explicit character set. Defaults
	// to "utf8".
	DefaultCharacterSet pulumi.StringPtrInput
	// The default collation to use when a table
	// is created without specifying an explicit collation. Defaults to
	// ``utf8GeneralCi``. Each character set has its own set of collations, so
	// changing the character set requires also changing the collation.
	DefaultCollation pulumi.StringPtrInput
	// The name of the database. This must be unique within
	// a given MySQL server and may or may not be case-sensitive depending on
	// the operating system on which the MySQL server is running.
	Name pulumi.StringPtrInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}

type DatabaseInput interface {
	pulumi.Input

	ToDatabaseOutput() DatabaseOutput
	ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput
}

func (*Database) ElementType() reflect.Type {
	return reflect.TypeOf((*Database)(nil))
}

func (i *Database) ToDatabaseOutput() DatabaseOutput {
	return i.ToDatabaseOutputWithContext(context.Background())
}

func (i *Database) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseOutput)
}

func (i *Database) ToDatabasePtrOutput() DatabasePtrOutput {
	return i.ToDatabasePtrOutputWithContext(context.Background())
}

func (i *Database) ToDatabasePtrOutputWithContext(ctx context.Context) DatabasePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePtrOutput)
}

type DatabasePtrInput interface {
	pulumi.Input

	ToDatabasePtrOutput() DatabasePtrOutput
	ToDatabasePtrOutputWithContext(ctx context.Context) DatabasePtrOutput
}

type databasePtrType DatabaseArgs

func (*databasePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil))
}

func (i *databasePtrType) ToDatabasePtrOutput() DatabasePtrOutput {
	return i.ToDatabasePtrOutputWithContext(context.Background())
}

func (i *databasePtrType) ToDatabasePtrOutputWithContext(ctx context.Context) DatabasePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePtrOutput)
}

// DatabaseArrayInput is an input type that accepts DatabaseArray and DatabaseArrayOutput values.
// You can construct a concrete instance of `DatabaseArrayInput` via:
//
//          DatabaseArray{ DatabaseArgs{...} }
type DatabaseArrayInput interface {
	pulumi.Input

	ToDatabaseArrayOutput() DatabaseArrayOutput
	ToDatabaseArrayOutputWithContext(context.Context) DatabaseArrayOutput
}

type DatabaseArray []DatabaseInput

func (DatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Database)(nil))
}

func (i DatabaseArray) ToDatabaseArrayOutput() DatabaseArrayOutput {
	return i.ToDatabaseArrayOutputWithContext(context.Background())
}

func (i DatabaseArray) ToDatabaseArrayOutputWithContext(ctx context.Context) DatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseArrayOutput)
}

// DatabaseMapInput is an input type that accepts DatabaseMap and DatabaseMapOutput values.
// You can construct a concrete instance of `DatabaseMapInput` via:
//
//          DatabaseMap{ "key": DatabaseArgs{...} }
type DatabaseMapInput interface {
	pulumi.Input

	ToDatabaseMapOutput() DatabaseMapOutput
	ToDatabaseMapOutputWithContext(context.Context) DatabaseMapOutput
}

type DatabaseMap map[string]DatabaseInput

func (DatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Database)(nil))
}

func (i DatabaseMap) ToDatabaseMapOutput() DatabaseMapOutput {
	return i.ToDatabaseMapOutputWithContext(context.Background())
}

func (i DatabaseMap) ToDatabaseMapOutputWithContext(ctx context.Context) DatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseMapOutput)
}

type DatabaseOutput struct {
	*pulumi.OutputState
}

func (DatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Database)(nil))
}

func (o DatabaseOutput) ToDatabaseOutput() DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabasePtrOutput() DatabasePtrOutput {
	return o.ToDatabasePtrOutputWithContext(context.Background())
}

func (o DatabaseOutput) ToDatabasePtrOutputWithContext(ctx context.Context) DatabasePtrOutput {
	return o.ApplyT(func(v Database) *Database {
		return &v
	}).(DatabasePtrOutput)
}

type DatabasePtrOutput struct {
	*pulumi.OutputState
}

func (DatabasePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil))
}

func (o DatabasePtrOutput) ToDatabasePtrOutput() DatabasePtrOutput {
	return o
}

func (o DatabasePtrOutput) ToDatabasePtrOutputWithContext(ctx context.Context) DatabasePtrOutput {
	return o
}

type DatabaseArrayOutput struct{ *pulumi.OutputState }

func (DatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Database)(nil))
}

func (o DatabaseArrayOutput) ToDatabaseArrayOutput() DatabaseArrayOutput {
	return o
}

func (o DatabaseArrayOutput) ToDatabaseArrayOutputWithContext(ctx context.Context) DatabaseArrayOutput {
	return o
}

func (o DatabaseArrayOutput) Index(i pulumi.IntInput) DatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Database {
		return vs[0].([]Database)[vs[1].(int)]
	}).(DatabaseOutput)
}

type DatabaseMapOutput struct{ *pulumi.OutputState }

func (DatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Database)(nil))
}

func (o DatabaseMapOutput) ToDatabaseMapOutput() DatabaseMapOutput {
	return o
}

func (o DatabaseMapOutput) ToDatabaseMapOutputWithContext(ctx context.Context) DatabaseMapOutput {
	return o
}

func (o DatabaseMapOutput) MapIndex(k pulumi.StringInput) DatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Database {
		return vs[0].(map[string]Database)[vs[1].(string)]
	}).(DatabaseOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseOutput{})
	pulumi.RegisterOutputType(DatabasePtrOutput{})
	pulumi.RegisterOutputType(DatabaseArrayOutput{})
	pulumi.RegisterOutputType(DatabaseMapOutput{})
}
