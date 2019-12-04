// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mysql
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-mysql/blob/master/website/docs/r/user_password.html.markdown.
    /// </summary>
    public partial class UserPassword : Pulumi.CustomResource
    {
        /// <summary>
        /// The encrypted password, base64 encoded.
        /// </summary>
        [Output("encryptedPassword")]
        public Output<string> EncryptedPassword { get; private set; } = null!;

        /// <summary>
        /// The source host of the user. Defaults to `localhost`.
        /// </summary>
        [Output("host")]
        public Output<string?> Host { get; private set; } = null!;

        /// <summary>
        /// The fingerprint of the PGP key used to encrypt the password 
        /// </summary>
        [Output("keyFingerprint")]
        public Output<string> KeyFingerprint { get; private set; } = null!;

        /// <summary>
        /// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:some_person_that_exists`.
        /// </summary>
        [Output("pgpKey")]
        public Output<string> PgpKey { get; private set; } = null!;

        /// <summary>
        /// The IAM user to associate with this access key.
        /// </summary>
        [Output("user")]
        public Output<string> User { get; private set; } = null!;


        /// <summary>
        /// Create a UserPassword resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserPassword(string name, UserPasswordArgs args, CustomResourceOptions? options = null)
            : base("mysql:index/userPassword:UserPassword", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private UserPassword(string name, Input<string> id, UserPasswordState? state = null, CustomResourceOptions? options = null)
            : base("mysql:index/userPassword:UserPassword", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserPassword resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserPassword Get(string name, Input<string> id, UserPasswordState? state = null, CustomResourceOptions? options = null)
        {
            return new UserPassword(name, id, state, options);
        }
    }

    public sealed class UserPasswordArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The source host of the user. Defaults to `localhost`.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:some_person_that_exists`.
        /// </summary>
        [Input("pgpKey", required: true)]
        public Input<string> PgpKey { get; set; } = null!;

        /// <summary>
        /// The IAM user to associate with this access key.
        /// </summary>
        [Input("user", required: true)]
        public Input<string> User { get; set; } = null!;

        public UserPasswordArgs()
        {
        }
    }

    public sealed class UserPasswordState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The encrypted password, base64 encoded.
        /// </summary>
        [Input("encryptedPassword")]
        public Input<string>? EncryptedPassword { get; set; }

        /// <summary>
        /// The source host of the user. Defaults to `localhost`.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// The fingerprint of the PGP key used to encrypt the password 
        /// </summary>
        [Input("keyFingerprint")]
        public Input<string>? KeyFingerprint { get; set; }

        /// <summary>
        /// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:some_person_that_exists`.
        /// </summary>
        [Input("pgpKey")]
        public Input<string>? PgpKey { get; set; }

        /// <summary>
        /// The IAM user to associate with this access key.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public UserPasswordState()
        {
        }
    }
}