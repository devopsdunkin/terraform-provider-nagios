# Terraform Provider: Nagios

## Installing

The provider can be downloaded from the [GitHub releases page](https://github.com/devopsdunkin/terraform-provider-nagios/releases)

Refer to Hashicorp's [documentation](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins) to install third-party providers

## Supported Nagios Versions

This provider was tested on Nagios XI 5.6.7. It may work with other versions, it will just depend on if there are differences in how the API processes commands.

## Provider setup

The provider requires the following attributes:

`url` (required): The URL to login to Nagios XI. This should just be the FQDN of the Nagios server, with `/nagiosxi` appended to it. It defaults to using the environment variable `NAGIOS_URL` if no value is provided in the provider definition.

`token` (required): The API token used to login to the Nagios XI API. This value can be found by logging in to the web application, clicking `Help`, and on the left hand side menu, click `Introduction`. Then it will be listed under the `Authentication` section. It defaults to using the environment variable `API_TOKEN` if no value is provided in the provider definition.

## Examples

```hcl
provider "nagios" {
    url = "http://localhost/nagiosxi"
    token = "pd994lfldfjfgGGHPDdj83iDjvdPv9033AAbwewwpPP69fmd4201qQmv0zCzxodD"
}
```

## Arguments

`url`: The URL for Nagios XI. This should be the URL to the web portal of Nagios
`token`: The API token for the user that will be logging in with the provider
