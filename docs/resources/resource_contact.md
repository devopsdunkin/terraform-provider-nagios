# resource: contact

## overview

This resource manages Nagios contacts. Nagios uses contacts to send notifications about changes in state of objects that Nagios manages.

Refer to the object definition for [contacts](https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/objectdefinitions.html#contact) from Nagios to learn more

## example

```hcl
resource "nagios_contact" "contact" {
    contact_name                    = "nagiosadmin"
    host_notifications_enabled      = true
    service_notifications_enabled   = true
    host_notification_period        = "workhours"
    service_notification_period     = "workhours"
    host_notification_options       = "d,u,r,f,s"
    service_notification_options    = "w,u,c,f,s"
    host_notification_commands      = [
        "notify-host-by-email"
    ]
    service_notification_commands   = [
        "notify-service-by-email"
    ]
    alias                           = "Nagios Administrator"
    templates                       = [
        "generic-contact"
    ]
    email                           = "nagiosadmin@example.com"
    can_submit_commands             = true
}
```

## arguments

Below is a brief description of what each field is used for in Nagios. Refer to the [official Nagios documentation](https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/objectdefinitions.html) for more detail

`contact_name` (required): The name of the contact
`host_notifications_enabled` (required): Determines whether or not the contact will receive notifications about host problems and recoveries
`service_notifications_enabled`: Determines whether or not the contact will receive notifications about service problems and recoveries
`host_notification_period` (required): The short name of the time period during which the contact can be notified about host problems or recoveries
`service_notification_period` (required): The short name of the time period during which the contact can be notified about service problems or recoveries
`host_notification_options` (required): The host states for which notifications can be sent out to this contact. Valid options are a combination of one or more of the following:

```bash
d = notify on DOWN host states
u = notify on UNREACHABLE host states
r = notify on UP host states
f = notify when the host starts and stops flapping
s = notify when scheduled downtime starts and ends
n = do not notify this contact for any host events
```

`service_notification_options` (required): The service states for which notifications can be sent out to this contact. Valid options are a combination of one or more of the following:

```bash
w = notify on WARNING service states
u = notify on UNKNOWN service states
c = notify on CRITICAL service states
r = notify on UP service states
f = notify when the service starts and stops flapping
s = notify when scheduled downtime starts and ends
n = do not notify this contact for any service events
```

`host_notification_commands` (required): A list of the short names of the commands used to notify the contact of a host problem or recovery. Multiple notification commands should be separated by commas. The command object must exist in Nagios for it to validate. All notification commands are executed when the contact needs to be notified

`service_notification_commands` (required): A list of the short names of the commands used to notify the contact of a service problem or recovery. Multiple notification commands should be separated by commas. The command object must exist in Nagios for it to validate. All notification commands are executed when the contact needs to be notified

`alias`: A longer name or description for the contact

`contact_groups`: The short name(s) of the contactgroup(s) that the contact belongs to

`email`: Defines an email address for the contact

`pager`: Defines a pager number for the contact

`addresses`: Defines additional 'addresses' for the contact

`can_submit_commands`: Determines whether or not the contact can submit external commands to Nagios from the CGIs

`retain_status_information`: Determines whether or not status-related information about the contact is retained across program restarts

`retain_nonstatus_information`: Determines whether or not non-status information about the contact is retained across program restarts.
