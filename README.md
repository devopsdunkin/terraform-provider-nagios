# terraform-provider-nagios
Terraform provider for Nagios XI

# Supported Nagios XI versions

This provider was tested on Nagios XI 5.6.1 through 5.6.7. It may work with older versions, if the API commands are the same.

## Roadmap

The plan for this provider is to allow complete management of Nagios XI through code. The following object types will be coming soon, with the numbered list showing the approximate priority for each:

#### Resources

1. Users
2. Time periods
3. Free variables **
4. Templates **
5. Auth servers: This is currently not feasible given certain limitations within Nagios - [Issue #56](https://github.com/devopsdunkin/terraform-provider-nagios/issues/56)

#### Data sources

1. Hosts
2. Host groups
3. Services
4. Service groups
5. Free variables **
6. Contacts
7. Contact groups
8. Time period
9. Users
10. Auth servers
11. Templates **

<br />

** Nagios custom API endpoint required to implement