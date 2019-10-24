# Resource: Hostgroup

# Overview

This resource manages Nagios host groups. Host groups are used to logically group servers together that may share a
similar function or other attribute. Refer to the object definition for [hostgroups](https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/objectdefinitions.html#hostgroup) from Nagios to learn more

Refer to the object definition for [hostgroups](https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/objectdefinitions.html#hostgroup) from Nagios to learn more

## Example

```hcl
resource "nagios_hostgroup" "hostgroup1" {
    name    = "hostgroup1"
    alias   = "This is an example hostgroup"
    members = [
        "host1.test.local",
        "host2.test.local"
    ]
}
```

## Arguments

`name`: The name of the Nagios host group
`alias`: The description or other name that the host group may be called. This field can be longer and more descriptive
`members`: A list of hosts that should be members of the host group. The members must be valid hosts within Nagios and must be active. The provider will NOT validate that the host is correctly configured in Nagios. The reason for this is, if the membership grows to hundreds or thousands of hosts, querying for each one of those would create performance issues.