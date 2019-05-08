// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The ``mysql_user`` resource creates and manages a user on a MySQL
 * server.
 * 
 * > **Note:** The password for the user is provided in plain text, and is
 * obscured by an unsalted hash in the state
 * [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
 * Care is required when using this resource, to avoid disclosing the password.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mysql from "@pulumi/mysql";
 * 
 * const jdoe = new mysql.User("jdoe", {
 *     host: "example.com",
 *     plaintextPassword: "password",
 *     user: "jdoe",
 * });
 * ```
 * 
 * ## Example Usage with an Authentication Plugin
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mysql from "@pulumi/mysql";
 * 
 * const nologin = new mysql.User("nologin", {
 *     authPlugin: "mysql_no_login",
 *     host: "example.com",
 *     user: "nologin",
 * });
 * ```
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /**
     * Use an [authentication plugin][ref-auth-plugins] to authenticate the user instead of using password authentication.  Description of the fields allowed in the block below. Conflicts with `password` and `plaintext_password`.  
     */
    public readonly authPlugin: pulumi.Output<string | undefined>;
    /**
     * The source host of the user. Defaults to "localhost".
     */
    public readonly host: pulumi.Output<string | undefined>;
    /**
     * Deprecated alias of `plaintext_password`, whose value is *stored as plaintext in state*. Prefer to use `plaintext_password` instead, which stores the password as an unsalted hash. Conflicts with `auth_plugin`.
     */
    public readonly password: pulumi.Output<string | undefined>;
    /**
     * The password for the user. This must be provided in plain text, so the data source for it must be secured. An _unsalted_ hash of the provided password is stored in state. Conflicts with `auth_plugin`.
     */
    public readonly plaintextPassword: pulumi.Output<string | undefined>;
    /**
     * An TLS-Option for the `CREATE USER` or `ALTER USER` statement. The value is suffixed to `REQUIRE`. A value of 'SSL' will generate a `CREATE USER ... REQUIRE SSL` statement. See the [MYSQL `CREATE USER` documentation](https://dev.mysql.com/doc/refman/5.7/en/create-user.html) for more. Ignored if MySQL version is under 5.7.0.
     */
    public readonly tlsOption: pulumi.Output<string | undefined>;
    /**
     * The name of the user.
     */
    public readonly user: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: UserState = argsOrState as UserState | undefined;
            inputs["authPlugin"] = state ? state.authPlugin : undefined;
            inputs["host"] = state ? state.host : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["plaintextPassword"] = state ? state.plaintextPassword : undefined;
            inputs["tlsOption"] = state ? state.tlsOption : undefined;
            inputs["user"] = state ? state.user : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if (!args || args.user === undefined) {
                throw new Error("Missing required property 'user'");
            }
            inputs["authPlugin"] = args ? args.authPlugin : undefined;
            inputs["host"] = args ? args.host : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["plaintextPassword"] = args ? args.plaintextPassword : undefined;
            inputs["tlsOption"] = args ? args.tlsOption : undefined;
            inputs["user"] = args ? args.user : undefined;
        }
        super("mysql:index/user:User", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * Use an [authentication plugin][ref-auth-plugins] to authenticate the user instead of using password authentication.  Description of the fields allowed in the block below. Conflicts with `password` and `plaintext_password`.  
     */
    readonly authPlugin?: pulumi.Input<string>;
    /**
     * The source host of the user. Defaults to "localhost".
     */
    readonly host?: pulumi.Input<string>;
    /**
     * Deprecated alias of `plaintext_password`, whose value is *stored as plaintext in state*. Prefer to use `plaintext_password` instead, which stores the password as an unsalted hash. Conflicts with `auth_plugin`.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * The password for the user. This must be provided in plain text, so the data source for it must be secured. An _unsalted_ hash of the provided password is stored in state. Conflicts with `auth_plugin`.
     */
    readonly plaintextPassword?: pulumi.Input<string>;
    /**
     * An TLS-Option for the `CREATE USER` or `ALTER USER` statement. The value is suffixed to `REQUIRE`. A value of 'SSL' will generate a `CREATE USER ... REQUIRE SSL` statement. See the [MYSQL `CREATE USER` documentation](https://dev.mysql.com/doc/refman/5.7/en/create-user.html) for more. Ignored if MySQL version is under 5.7.0.
     */
    readonly tlsOption?: pulumi.Input<string>;
    /**
     * The name of the user.
     */
    readonly user?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * Use an [authentication plugin][ref-auth-plugins] to authenticate the user instead of using password authentication.  Description of the fields allowed in the block below. Conflicts with `password` and `plaintext_password`.  
     */
    readonly authPlugin?: pulumi.Input<string>;
    /**
     * The source host of the user. Defaults to "localhost".
     */
    readonly host?: pulumi.Input<string>;
    /**
     * Deprecated alias of `plaintext_password`, whose value is *stored as plaintext in state*. Prefer to use `plaintext_password` instead, which stores the password as an unsalted hash. Conflicts with `auth_plugin`.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * The password for the user. This must be provided in plain text, so the data source for it must be secured. An _unsalted_ hash of the provided password is stored in state. Conflicts with `auth_plugin`.
     */
    readonly plaintextPassword?: pulumi.Input<string>;
    /**
     * An TLS-Option for the `CREATE USER` or `ALTER USER` statement. The value is suffixed to `REQUIRE`. A value of 'SSL' will generate a `CREATE USER ... REQUIRE SSL` statement. See the [MYSQL `CREATE USER` documentation](https://dev.mysql.com/doc/refman/5.7/en/create-user.html) for more. Ignored if MySQL version is under 5.7.0.
     */
    readonly tlsOption?: pulumi.Input<string>;
    /**
     * The name of the user.
     */
    readonly user: pulumi.Input<string>;
}
