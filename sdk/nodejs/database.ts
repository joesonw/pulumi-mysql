// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The ``mysql_database`` resource creates and manages a database on a MySQL
 * server.
 * 
 * > **Caution:** The ``mysql_database`` resource can completely delete your
 * database just as easily as it can create it. To avoid costly accidents,
 * consider setting
 * [``prevent_destroy``](https://www.terraform.io/docs/configuration/resources.html#prevent_destroy)
 * on your database resources as an extra safety measure.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mysql from "@pulumi/mysql";
 * 
 * const app = new mysql.Database("app", {});
 * ```
 */
export class Database extends pulumi.CustomResource {
    /**
     * Get an existing Database resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseState, opts?: pulumi.CustomResourceOptions): Database {
        return new Database(name, <any>state, { ...opts, id: id });
    }

    /**
     * The default character set to use when
     * a table is created without specifying an explicit character set. Defaults
     * to "utf8".
     */
    public readonly defaultCharacterSet: pulumi.Output<string | undefined>;
    /**
     * The default collation to use when a table
     * is created without specifying an explicit collation. Defaults to
     * ``utf8_general_ci``. Each character set has its own set of collations, so
     * changing the character set requires also changing the collation.
     */
    public readonly defaultCollation: pulumi.Output<string | undefined>;
    /**
     * The name of the database. This must be unique within
     * a given MySQL server and may or may not be case-sensitive depending on
     * the operating system on which the MySQL server is running.
     */
    public readonly name: pulumi.Output<string>;

    /**
     * Create a Database resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DatabaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseArgs | DatabaseState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: DatabaseState = argsOrState as DatabaseState | undefined;
            inputs["defaultCharacterSet"] = state ? state.defaultCharacterSet : undefined;
            inputs["defaultCollation"] = state ? state.defaultCollation : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as DatabaseArgs | undefined;
            inputs["defaultCharacterSet"] = args ? args.defaultCharacterSet : undefined;
            inputs["defaultCollation"] = args ? args.defaultCollation : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        super("mysql:index/database:Database", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Database resources.
 */
export interface DatabaseState {
    /**
     * The default character set to use when
     * a table is created without specifying an explicit character set. Defaults
     * to "utf8".
     */
    readonly defaultCharacterSet?: pulumi.Input<string>;
    /**
     * The default collation to use when a table
     * is created without specifying an explicit collation. Defaults to
     * ``utf8_general_ci``. Each character set has its own set of collations, so
     * changing the character set requires also changing the collation.
     */
    readonly defaultCollation?: pulumi.Input<string>;
    /**
     * The name of the database. This must be unique within
     * a given MySQL server and may or may not be case-sensitive depending on
     * the operating system on which the MySQL server is running.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Database resource.
 */
export interface DatabaseArgs {
    /**
     * The default character set to use when
     * a table is created without specifying an explicit character set. Defaults
     * to "utf8".
     */
    readonly defaultCharacterSet?: pulumi.Input<string>;
    /**
     * The default collation to use when a table
     * is created without specifying an explicit collation. Defaults to
     * ``utf8_general_ci``. Each character set has its own set of collations, so
     * changing the character set requires also changing the collation.
     */
    readonly defaultCollation?: pulumi.Input<string>;
    /**
     * The name of the database. This must be unique within
     * a given MySQL server and may or may not be case-sensitive depending on
     * the operating system on which the MySQL server is running.
     */
    readonly name?: pulumi.Input<string>;
}
