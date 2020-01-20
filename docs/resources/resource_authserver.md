# Resource: Auth Server

## Overview

This resource manages Nagios authentication servers. Authentication servers allow for LDAP or Active Directory authentication for Nagios XI. This resource provides the ability to create and delete authentication servers. The Nagios XI API currently does not support update (HTTP PUT) operations, so any changes made to an existing resource require the resource to be destroyed and created with the new values.

## Example

```hcl
resource "nagios_authserver" "authserver1" {
    enabled                 = true
    connection_method       = "ad"
    ad_account_suffix       = "@domain.local"
    ad_domain_controllers   = "dc1.domain.local"
    base_dn                 = "OU=IT, DC=domain, DC=local"
    security_level          = "ssl"
}
```

## Arguments

Below is a brief description of what each field is used for in Nagios

`enabled`: Determines if the authentication server is enabled in Nagios XI. The default is `true`
`connection_method` (required): The connection method to use. It must be either `ad` or `ldap`
`ad_account_suffix`: The account suffix to use when `connection_method` is `ad`. It must be in the format of `@domain.local`.
`ad_domain_controllers`: A list of domain controllers used for authentication. This field is required when `connection_method` is `ad`. The list should be provided as a comma separated string
`base_dn` (required): The disntguished name within the directory or tree where the query should look for users
`security_level`: The type of encryption to use, which can be either `ssl`, `tls` or `none`
`ldap_host`: The host name of the LDAP server to connect to. This field is required when the `connection_method` is set to `ldap`
`ldap_port`: The port to use when connecting with LDAP. This field is required when the `connection_method` is set to `ldap`. The default is `389`
