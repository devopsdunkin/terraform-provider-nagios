<<<<<<< Updated upstream
# Resource: Hostgroup
=======
<<<<<<< HEAD
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
=======
# Resource: Hostgroup
>>>>>>> origin/master
>>>>>>> Stashed changes
