# Resource: Hostgroup

## Overview

This resource manages Nagios hosts. Nagios monitors hosts based off health checks defined by an administrator.

Refer to the object definition for [hosts](https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/objectdefinitions.html#host) from Nagios to learn more

## Best practices

This resource allows you to configure any attribute on a host, however, we recommend to manually create a host template that contains shared settings for the hosts, and then use the host resource to create your hosts and set host specific attributes. 

## Example

```hcl
resource "nagios_host" "host1" {
    name                            = "host1"
    address                         = "127.0.0.1"
    display_name                    = "host1 - Test server"
    max_check_attempts              = "2"
    check_period                    = "24x7"
    notification_interval           = "5"
    notification_period             = "24x7"
    contacts                        = [
        "nagiosadmin"
    ]
    alias                           = "host1"
    templates                       = "generic-host"
    contact_groups                  = [
        "noc_staff"
    ]
    notes                           = "If this host is down for more than 20 mninutes, page out to the operations on-call"
    notes_url                       = "http://docs.company.com/host_alert"
    action_url                      = "http://docs.company.com/host_alert/actions"
    initial_state                   = "s"
    retry_interval                  = "5"
    passive_checks_enabled          = true
    active_checks_enabled           = true
    obsess_over_host                = false
    event_handler                   = 
    event_handler_enabled           = true
    flap_detection_enabled          = true
    flap_detection_options          = 
    low_flap_threshold              = "5"
    high_flap_threshold             = "5"
    process_perf_data               = true
    retain_status_information       =
    retain_nonstatus_information    =
    check_freshness                 =
    freshness_threshold             =
    first_notification_delay        =
    notification_options            =
    notifications_enabled           = true
    stalking_options                =
    icon_image                      = "windows_host.jpg"
    icon_image_alt                  = "Windows Server"
    vrml_image                      =
    statusmap_image                 =
    2d_coords                       =
    3d_coords                       =
}
```

## Arguments

`name`:
`address`:
`display_name`:
`max_check_attempts`:
`check_period`:
`notification_interval`:
`notification_period`:
`contacts`:
`alias`:
`templates`:
`contact_groups`:
`notes`:
`notes_url`:
`action_url`:
`initial_state`:
`retry_interval`:
`passive_checks_enabled`:
`active_checks_enabled`:
`obsess_over_host`:
`event_handler`:
`event_handler_enabled`:
`flap_detection_enabled`:
`flap_detection_options`:
`low_flap_threshold`:
`high_flap_threshold`:
`process_perf_data`:
`retain_status_information`:
`retain_nonstatus_information`:
`check_freshness`:
`freshness_threshold`:
`first_notification_delay`:
`notification_options`:
`notifications_enabled`:
`stalking_options`:
`icon_image`:
`icon_image_alt`: 
`vrml_image`:
`statusmap_image`:
`2d_coords`:
`3d_coords`: