# Data Source: Host

## Overview

This data source retrieves a host object from Nagios. The data source will retrieve any setting that is defined in Nagios for the host.

Refer to the object definition for [hosts](https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/objectdefinitions.html#host) from Nagios to learn more

## Example

```hcl
data "nagios_host" "webserver1" {
    host_name       = "webserver1"
}
```

## Arguments

Below is a brief description of what each field is used for in Nagios. Refer to the [official Nagios documentation](https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/objectdefinitions.html) for more detail

`host_name` (required): The name of the host to retrieve from Nagios.
