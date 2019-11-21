# Resource: Service

## Overview

This resource manages Nagios services. Nagios monitors services based off health checks defined by an administrator.

Refer to the object definition for [services](https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/objectdefinitions.html#service) from Nagios to learn more

## Best practices

This resource allows you to configure any attribute for a service, however, we recommend to manually create a service template that contains common settings for the service, and then use this resource to create your services and set service or host specific attributes.

## Example

```hcl
resource "nagios_service" "service1" {
    service_name                    = "ping"
    host_name                       = [
        "websrv1",
        "websrv2",
        "dc1"
    ]
    description                     = "Service pings a server"
    check_command                   = "check_ping!3000.0!80%!5000.0!100%!!!!"
    max_check_attempts              = "5"
    check_interval                  = "5"
    retry_interval                  = "5"
    check_period                    = "24x7"
    notification_interval           = "10"
    notification_period             = "24x7"
    contacts                        = [
        "nagiosadmin"
    ]
    templates                       = [
        "generic-service"
    ]
    is_volatile                     = false
    initial_state                   = "s"
    active_checks_enabled           = true
    passive_checks_enabled          = true
    obsess_over_service             = false
    check_freshness                 = false
    freshness_threshold             = "10"
    event_handler                   = "xi_service_event_handler"
    event_handler_enabled           = true
    low_flap_threshold              = "20"
    high_flap_threshold             = "30"
    flap_detection_enabled          = false
    flap_detection_options          = [
        "d",
        "u"
    ]
    process_perf_data               = true
    retain_status_information       = true
    retain_nonstatus_information    = true
    first_notification_delay        = "5"
	notification_options            = [
        "d",
        "u"
    ]
    notifications_enabled           = true
    contact_groups                  = [
        "contact_group1"
    ]
    notes                           = "Some notes about the service"
    notes_url                       = "http://docs.company.com/host_alert"
    action_url                      = "http://docs.company.com/host_alert/actions"
    icon_image                      = "pingicon.jpg"
    icon_image_alt                  = "Ping check"
}
```

## Arguments

Below is a brief description of what each field is used for in Nagios. Refer to the [official Nagios documentation](https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/objectdefinitions.html) for more detail

`service_name` (required): The name of the service  
`host_name` (required): The hosts that the service should run on  
`description` (required): Defines the description of the service. It may contain spaces, dashes and colons (avoid using semicolons, apostrophes and quotation marks)  
`display_name`: Another name for the service that will be displayed in the web interface. If left blank, the value from `description` will be displayed  
`check_command` (required): The name of the command that should be used to check the status of the service  
`max_check_attempts` (required): How many times to retry the service check before alerting when the state is anything other than OK  
`templates`: A list of Nagios templates to apply to the service  
`is_volatile`: Determines if the service is "volatile". Services typically are not volatile and this should be disabled. This accepts either `true` or `false`. The default value is `false`  
`initial_state`: By default, Nagios will assume the service are in an OK state. Valid options are:  

    s = UP
    d = DOWN
    u = UNREACHABLE
  
`check_interval`: The number of minutes to wait until the next regular check of the service  
`retry_interval`: The number of minutes to wait until re-checking the service  
`active_checks_enabled`: Sets whether or not active checks are enabled for the service  
`passive_checks_enabled`: Sets whether or not passive checks are enabled for the service  
`check_period` (required): The time period during which active checks of the service can be made  
`obsess_over_service`: Sets whether or not Nagios 'obsesses' over the service using the ocsp_command  
`check_freshness`: Sets whether or not freshness checks are enabled for the service  
`freshness_threshold`: The freshness threshold used for the service  
`event_handler`: The command that should be run whenver a change in the state of the service is detected  
`event_handler_enabled`: Sets whether or not event handlers should be enabled for the service  
`low_flap_threshold`: The minimum threshold that should be used when detecting if flapping is occurring  
`high_flap_threshold`: The maximum threshold that should be used when detecting if flapping is occurring  
`flap_detection_enabled`: Sets whether or not flap detection is enabled for the service  
`flap_detection_options`: Determines what flap detection logic will be used for the service. One or more of the following valid options can be provided:  

    o = UP
    d = DOWN
    u = UNREACHABLE

`process_perf_data`: Determines if Nagios should process performance data  
`retain_status_information`: Sets whether or not status related information should be kept for the service  
`retain_nonstatus_information`: Sets whether or not non-status related information should be kept for the service  
`notification_interval` (required): How long to wait before sending another notification to a contact that the service is down  
`first_notification_delay`: The amount of time to wait to send out the first notification when a service enters a non-UP state  
`notification_period` (required): The time period during which notifications can be sent for a service alert  
`notification_options`: Determines when Nagios should alert if a host is one or more of the following options:  

    o = UP
    d = DOWN
    u = UNREACHABLE
    r = RECOVERY
    f = FLAPPING
    s = SCHEDULED DOWNTIME

`notifications_enabled`: Determines if Nagios should send notifications  
`contacts` (required): The list of users that Nagios should alert when a service is down  
`contact_groups`: A list of the contact groups that should be notified if the service goes down  
`stalking_options`: A list of options to determine which states, if any, should be stalked by Nagios. Refer to the [Nagios documentation](https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/stalking.html) for more information on stalking  
`notes`: Notes about the service that may assist with troubleshooting  
`notes_url`: URL to a third-party documentation respoitory containing more information about the service  
`action_url`: URL to a third-party documentation repository containing actions to take in the event the service goes down  
`icon_image`: The icon to display in Nagios  
`icon_image_alt`:  The text to display when hovering over the `icon_image` or the text to display if the `icon_image` is unavailable  