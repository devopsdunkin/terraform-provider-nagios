# Resource: Servicegroup

## Overview

This resource manages Nagios servicegroups. It can be used to create, update and delete service groups, as well as manage all attributes currently supported by Nagios.

Refer to the object definition for [servicegroups](https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/objectdefinitions.html#servicegroup) from Nagios to learn more

## Example

```hcl
resource "nagios_servicegroup" "servicegroup1" {
    name    = "servicegroup1"
    alias   = "This is an example servicegroup"
    members = [
        "svcgroup1",
        "ping_check"
    ]
    notes   = "Example notes"
    notes_url   = "https://docs.example.com/servicegroup/notes"
    action_url  = "https://dovs.example.com/servicegroup/action"
}
```

## Arguments

`name`: The name of the Nagios service group  
`alias`: The description or other name that the servicegroup may be called. This field can be longer and more descriptive  
`members`: A list of services and other service groups that should be members of the service group. The members must be valid services and service groups within Nagios and must be active. The provider will NOT validate that the object is correctly configured in Nagios. The reason for this is, if the membership grows to hundreds or thousands of hosts, querying for each one of those would create performance issues.  
`notes`: Notes about the service group that may assist with troubleshooting  
`notes_url`: URL to a third-party documentation respoitory containing more information about the service group  
`action_url`: URL to a third-party documentation repository containing actions to take in the event the service group goes down
