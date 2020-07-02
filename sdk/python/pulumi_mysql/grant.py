# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class Grant(pulumi.CustomResource):
    database: pulumi.Output[str]
    """
    The database to grant privileges on.
    """
    grant: pulumi.Output[bool]
    """
    Whether to also give the user privileges to grant the same privileges to other users.
    """
    host: pulumi.Output[str]
    """
    The source host of the user. Defaults to "localhost". Conflicts with `role`.
    """
    privileges: pulumi.Output[list]
    """
    A list of privileges to grant to the user. Refer to a list of privileges (such as [here](https://dev.mysql.com/doc/refman/5.5/en/grant.html)) for applicable privileges. Conflicts with `roles`.
    """
    role: pulumi.Output[str]
    """
    The role to grant `privileges` to. Conflicts with `user` and `host`.
    """
    roles: pulumi.Output[list]
    """
    A list of rols to grant to the user. Conflicts with `privileges`.
    """
    table: pulumi.Output[str]
    """
    Which table to grant `privileges` on. Defaults to `*`, which is all tables.
    """
    tls_option: pulumi.Output[str]
    """
    An TLS-Option for the `GRANT` statement. The value is suffixed to `REQUIRE`. A value of 'SSL' will generate a `GRANT ... REQUIRE SSL` statement. See the [MYSQL `GRANT` documentation](https://dev.mysql.com/doc/refman/5.7/en/grant.html) for more. Ignored if MySQL version is under 5.7.0.
    """
    user: pulumi.Output[str]
    """
    The name of the user. Conflicts with `role`.
    """
    def __init__(__self__, resource_name, opts=None, database=None, grant=None, host=None, privileges=None, role=None, roles=None, table=None, tls_option=None, user=None, __props__=None, __name__=None, __opts__=None):
        """
        The ``Grant`` resource creates and manages privileges given to
        a user on a MySQL server.

        ## Granting Privileges to a User

        ```python
        import pulumi
        import pulumi_mysql as mysql

        jdoe_user = mysql.User("jdoeUser",
            host="example.com",
            plaintext_password="password",
            user="jdoe")
        jdoe_grant = mysql.Grant("jdoeGrant",
            database="app",
            host=jdoe_user.host,
            privileges=[
                "SELECT",
                "UPDATE",
            ],
            user=jdoe_user.user)
        ```

        ## Granting Privileges to a Role

        ```python
        import pulumi
        import pulumi_mysql as mysql

        developer_role = mysql.Role("developerRole")
        developer_grant = mysql.Grant("developerGrant",
            database="app",
            privileges=[
                "SELECT",
                "UPDATE",
            ],
            role=developer_role.name)
        ```

        ## Adding a Role to a User

        ```python
        import pulumi
        import pulumi_mysql as mysql

        jdoe = mysql.User("jdoe",
            host="example.com",
            plaintext_password="password",
            user="jdoe")
        developer_role = mysql.Role("developerRole")
        developer_grant = mysql.Grant("developerGrant",
            database="app",
            host=jdoe.host,
            roles=[developer_role.name],
            user=jdoe.user)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database: The database to grant privileges on.
        :param pulumi.Input[bool] grant: Whether to also give the user privileges to grant the same privileges to other users.
        :param pulumi.Input[str] host: The source host of the user. Defaults to "localhost". Conflicts with `role`.
        :param pulumi.Input[list] privileges: A list of privileges to grant to the user. Refer to a list of privileges (such as [here](https://dev.mysql.com/doc/refman/5.5/en/grant.html)) for applicable privileges. Conflicts with `roles`.
        :param pulumi.Input[str] role: The role to grant `privileges` to. Conflicts with `user` and `host`.
        :param pulumi.Input[list] roles: A list of rols to grant to the user. Conflicts with `privileges`.
        :param pulumi.Input[str] table: Which table to grant `privileges` on. Defaults to `*`, which is all tables.
        :param pulumi.Input[str] tls_option: An TLS-Option for the `GRANT` statement. The value is suffixed to `REQUIRE`. A value of 'SSL' will generate a `GRANT ... REQUIRE SSL` statement. See the [MYSQL `GRANT` documentation](https://dev.mysql.com/doc/refman/5.7/en/grant.html) for more. Ignored if MySQL version is under 5.7.0.
        :param pulumi.Input[str] user: The name of the user. Conflicts with `role`.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if database is None:
                raise TypeError("Missing required property 'database'")
            __props__['database'] = database
            __props__['grant'] = grant
            __props__['host'] = host
            __props__['privileges'] = privileges
            __props__['role'] = role
            __props__['roles'] = roles
            __props__['table'] = table
            __props__['tls_option'] = tls_option
            __props__['user'] = user
        super(Grant, __self__).__init__(
            'mysql:index/grant:Grant',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, database=None, grant=None, host=None, privileges=None, role=None, roles=None, table=None, tls_option=None, user=None):
        """
        Get an existing Grant resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database: The database to grant privileges on.
        :param pulumi.Input[bool] grant: Whether to also give the user privileges to grant the same privileges to other users.
        :param pulumi.Input[str] host: The source host of the user. Defaults to "localhost". Conflicts with `role`.
        :param pulumi.Input[list] privileges: A list of privileges to grant to the user. Refer to a list of privileges (such as [here](https://dev.mysql.com/doc/refman/5.5/en/grant.html)) for applicable privileges. Conflicts with `roles`.
        :param pulumi.Input[str] role: The role to grant `privileges` to. Conflicts with `user` and `host`.
        :param pulumi.Input[list] roles: A list of rols to grant to the user. Conflicts with `privileges`.
        :param pulumi.Input[str] table: Which table to grant `privileges` on. Defaults to `*`, which is all tables.
        :param pulumi.Input[str] tls_option: An TLS-Option for the `GRANT` statement. The value is suffixed to `REQUIRE`. A value of 'SSL' will generate a `GRANT ... REQUIRE SSL` statement. See the [MYSQL `GRANT` documentation](https://dev.mysql.com/doc/refman/5.7/en/grant.html) for more. Ignored if MySQL version is under 5.7.0.
        :param pulumi.Input[str] user: The name of the user. Conflicts with `role`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["database"] = database
        __props__["grant"] = grant
        __props__["host"] = host
        __props__["privileges"] = privileges
        __props__["role"] = role
        __props__["roles"] = roles
        __props__["table"] = table
        __props__["tls_option"] = tls_option
        __props__["user"] = user
        return Grant(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
