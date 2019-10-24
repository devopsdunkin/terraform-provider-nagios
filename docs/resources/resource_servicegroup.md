# Resource: Servicegroup

# Overview

This resource manages Nagios servicegroups. Servicegroups are used to logically group services together that may share a
similar function or provide checks for a specific type of server. Refer to the object definition for [servicegroups](https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/objectdefinitions.html#servicegroup) from Nagios to learn more

## Example

```hcl
resource "nagios_servicegroup" "servicegroup1" {
    name    = "servicegroup1"
    alias   = "This is an example servicegroup"
    members = [
        "host1.test.local",
        "host2.test.local",
        "ping_check"
    ]
}
```

## Arguments

`name`: The name of the Nagios servicegroup

`alias`: The description or other name that the servicegroup may be called. This field can be longer and more descriptive

`members`: A list of hosts and/or services that should be members of the servicegroup. The members must be valid hosts and services within Nagios and must be active. The provider will NOT validate that the host is correctly configured in Nagios. The reason for this is, if the membership grows to hundreds or thousands of hosts, querying for each one of those would create performance issues.