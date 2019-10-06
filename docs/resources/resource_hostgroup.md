# Resource: Hostgroup

# Overview

This resource manages Nagios hostgroups. Hostgroups are used to logically group servers together that may share a
similar function or other attribute.

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

`name`: The name of the Nagios hostgroup

`alias`: The description or other name that the hostgroup may be called. This field can be longer and more descriptive

<<<<<<< HEAD
`members`: A list of hosts that should be members of the hostgroup. The members must be valid hosts within Nagios and must be active
=======
`members`: A list of hosts that should be members of the hostgroup. The members must be valid hosts within Nagios and must be active
>>>>>>> origin/master
