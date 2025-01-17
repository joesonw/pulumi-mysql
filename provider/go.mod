module github.com/pulumi/pulumi-mysql/provider/v3

go 1.16

require (
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.0.0
	github.com/pulumi/pulumi/pkg/v3 v3.0.0
	github.com/pulumi/pulumi/sdk/v3 v3.0.0
	github.com/terraform-providers/terraform-provider-mysql v1.9.0
)

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/terraform-providers/terraform-provider-mysql => github.com/joesonw/terraform-provider-mysql v1.9.1-0.20210621041132-ab77ee518821
)
