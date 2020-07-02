// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.MySql
{
    /// <summary>
    /// The ``mysql.User`` resource creates and manages a user on a MySQL
    /// server.
    /// 
    /// &gt; **Note:** The password for the user is provided in plain text, and is
    /// obscured by an unsalted hash in the state
    /// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
    /// Care is required when using this resource, to avoid disclosing the password.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using MySql = Pulumi.MySql;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var jdoe = new MySql.User("jdoe", new MySql.UserArgs
    ///         {
    ///             Host = "example.com",
    ///             PlaintextPassword = "password",
    ///             User = "jdoe",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### With An Authentication Plugin
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using MySql = Pulumi.MySql;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var nologin = new MySql.User("nologin", new MySql.UserArgs
    ///         {
    ///             AuthPlugin = "mysql_no_login",
    ///             Host = "example.com",
    ///             User = "nologin",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class User : Pulumi.CustomResource
    {
        /// <summary>
        /// Use an [authentication plugin][ref-auth-plugins] to authenticate the user instead of using password authentication.  Description of the fields allowed in the block below. Conflicts with `password` and `plaintext_password`.
        /// </summary>
        [Output("authPlugin")]
        public Output<string?> AuthPlugin { get; private set; } = null!;

        /// <summary>
        /// The source host of the user. Defaults to "localhost".
        /// </summary>
        [Output("host")]
        public Output<string?> Host { get; private set; } = null!;

        /// <summary>
        /// Deprecated alias of `plaintext_password`, whose value is *stored as plaintext in state*. Prefer to use `plaintext_password` instead, which stores the password as an unsalted hash. Conflicts with `auth_plugin`.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// The password for the user. This must be provided in plain text, so the data source for it must be secured. An _unsalted_ hash of the provided password is stored in state. Conflicts with `auth_plugin`.
        /// </summary>
        [Output("plaintextPassword")]
        public Output<string?> PlaintextPassword { get; private set; } = null!;

        /// <summary>
        /// An TLS-Option for the `CREATE USER` or `ALTER USER` statement. The value is suffixed to `REQUIRE`. A value of 'SSL' will generate a `CREATE USER ... REQUIRE SSL` statement. See the [MYSQL `CREATE USER` documentation](https://dev.mysql.com/doc/refman/5.7/en/create-user.html) for more. Ignored if MySQL version is under 5.7.0.
        /// </summary>
        [Output("tlsOption")]
        public Output<string?> TlsOption { get; private set; } = null!;

        /// <summary>
        /// The name of the user.
        /// </summary>
        [Output("user")]
        public Output<string> UserName { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("mysql:index/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("mysql:index/user:User", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Use an [authentication plugin][ref-auth-plugins] to authenticate the user instead of using password authentication.  Description of the fields allowed in the block below. Conflicts with `password` and `plaintext_password`.
        /// </summary>
        [Input("authPlugin")]
        public Input<string>? AuthPlugin { get; set; }

        /// <summary>
        /// The source host of the user. Defaults to "localhost".
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Deprecated alias of `plaintext_password`, whose value is *stored as plaintext in state*. Prefer to use `plaintext_password` instead, which stores the password as an unsalted hash. Conflicts with `auth_plugin`.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The password for the user. This must be provided in plain text, so the data source for it must be secured. An _unsalted_ hash of the provided password is stored in state. Conflicts with `auth_plugin`.
        /// </summary>
        [Input("plaintextPassword")]
        public Input<string>? PlaintextPassword { get; set; }

        /// <summary>
        /// An TLS-Option for the `CREATE USER` or `ALTER USER` statement. The value is suffixed to `REQUIRE`. A value of 'SSL' will generate a `CREATE USER ... REQUIRE SSL` statement. See the [MYSQL `CREATE USER` documentation](https://dev.mysql.com/doc/refman/5.7/en/create-user.html) for more. Ignored if MySQL version is under 5.7.0.
        /// </summary>
        [Input("tlsOption")]
        public Input<string>? TlsOption { get; set; }

        /// <summary>
        /// The name of the user.
        /// </summary>
        [Input("user", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public UserArgs()
        {
        }
    }

    public sealed class UserState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Use an [authentication plugin][ref-auth-plugins] to authenticate the user instead of using password authentication.  Description of the fields allowed in the block below. Conflicts with `password` and `plaintext_password`.
        /// </summary>
        [Input("authPlugin")]
        public Input<string>? AuthPlugin { get; set; }

        /// <summary>
        /// The source host of the user. Defaults to "localhost".
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Deprecated alias of `plaintext_password`, whose value is *stored as plaintext in state*. Prefer to use `plaintext_password` instead, which stores the password as an unsalted hash. Conflicts with `auth_plugin`.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The password for the user. This must be provided in plain text, so the data source for it must be secured. An _unsalted_ hash of the provided password is stored in state. Conflicts with `auth_plugin`.
        /// </summary>
        [Input("plaintextPassword")]
        public Input<string>? PlaintextPassword { get; set; }

        /// <summary>
        /// An TLS-Option for the `CREATE USER` or `ALTER USER` statement. The value is suffixed to `REQUIRE`. A value of 'SSL' will generate a `CREATE USER ... REQUIRE SSL` statement. See the [MYSQL `CREATE USER` documentation](https://dev.mysql.com/doc/refman/5.7/en/create-user.html) for more. Ignored if MySQL version is under 5.7.0.
        /// </summary>
        [Input("tlsOption")]
        public Input<string>? TlsOption { get; set; }

        /// <summary>
        /// The name of the user.
        /// </summary>
        [Input("user")]
        public Input<string>? UserName { get; set; }

        public UserState()
        {
        }
    }
}
