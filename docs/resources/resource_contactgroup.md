# resource: contactgroup

## overview

This resource manages Nagios contact groups. Nagios uses contact groups to group togther contacts for alerting

Refer to the object definition for [contactgroups](https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/objectdefinitions.html#contactgroup) from Nagios to learn more

## example

```hcl
resource "nagios_contactgroup" "contactgroup" {
    contactgroup_name       = "noc_staff"
    alias                   = "NOC staff members"
    members                 = [
        "nagiosadmin",
        "jdoe"
    ]
    contactgroup_members    = [
        "system_admins"
    ]
}
```

## arguments

Below is a brief description of what each field is used for in Nagios. Refer to the [official Nagios documentation](https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/objectdefinitions.html) for more detail

`contactgroup_name` (required): The name of the contact group
`alias` (required): A longer or more descriptive name of the contact group
`members`: A list of the contacts that should be included in this group
`contactgroup_members`: Other contact groups that should be included in this group
