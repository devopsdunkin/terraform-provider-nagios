# Data Source: Service

## Overview

This data source retrieves a service object from Nagios. The data source will retrieve any setting that is defined in Nagios for the service.

Refer to the object definition for [services](https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/objectdefinitions.html#service) from Nagios to learn more

## Example

```hcl
data "nagios_service" "ping_check" {
    service_name        = "Ping"
    description         = "This is an example of a service description"
}
```

## Arguments

Below is a brief description of what each field is used for in Nagios. Refer to the [official Nagios documentation](https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/objectdefinitions.html) for more detail

`service_name` (required): The name of the service to retrieve from Nagios
`description`: (required): The description of the service
