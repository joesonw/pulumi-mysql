// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

let __config = new pulumi.Config("mysql");

export let endpoint: string | undefined = __config.get("endpoint");
export let maxConnLifetimeSec: number | undefined = __config.getObject<number>("maxConnLifetimeSec");
export let maxOpenConns: number | undefined = __config.getObject<number>("maxOpenConns");
export let password: string | undefined = __config.get("password");
export let tls: string | undefined = __config.get("tls");
export let username: string | undefined = __config.get("username");
