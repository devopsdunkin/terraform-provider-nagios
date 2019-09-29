# terraform-provider-nagios
Terraform provider for Nagios XI

# Supported Nagios XI versions

This provider was tested onm Nagios XI 5.6.1. It may work with older versions, if the API commands are the same.

## Roadmap

The plan for this provider is to allow complete management of Nagios XI through code. The following object types will be coming soon, with the numbered list showing the approximate priority for each:

#### Resources

1. Services
2. Service groups
3. Free variables **
4. Contacts
5. Contact groups
6. Time periods
7. Users
8. Auth servers
9. Templates **

#### Data sources

1. Services
2. Service groups
3. Free variables **
4. Contacts
5. Contact groups
6. Time period
7. Users
8. Auth servers
9. Templates **

<br />

** Nagios custom API endpoint required to implement