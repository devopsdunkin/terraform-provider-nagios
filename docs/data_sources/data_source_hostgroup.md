# Data Source: Hostgroup

## Overview

This data source retrieves a host group from Nagios. The data source will retrieve any setting that is defined in Nagios for the host group.

Refer to the object definition for [hostgroups](https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/objectdefinitions.html#hostgroup) from Nagios to learn more

## Example

```hcl
data "nagios_hostgroup" "hostgroup1" {
    name    = "hostgroup1"
}
```

## Arguments

`name` (required): The name of the Nagios host group  
