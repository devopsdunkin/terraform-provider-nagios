## 1.1.1 (October 31, 2019)

FEATURES:

* None

IMPROVEMENTS:

* None

BUG FIXES:

* Fixes syntax issue with adding service description in when performing update to a service ([#52](https://github.com/devopsdunkin/terraform-provider-nagios/pull/52))
* Fixes syntax issue with replacing spaces with `%20` for attributes when performing an update ([#52](https://github.com/devopsdunkin/terraform-provider-nagios/pull/52))
* Fixes issue where service description was not getting passed as a URL parameter, so it would not update ([#52](https://github.com/devopsdunkin/terraform-provider-nagios/pull/52))

## 1.1.0 (October 30, 2019)

FEATURES:

* Adds CHANGELOG ([#51](https://github.com/devopsdunkin/terraform-provider-nagios/pull/51))
* Adds test job to pipeline ([#51](https://github.com/devopsdunkin/terraform-provider-nagios/pull/51))

IMPROVEMENTS:

* Cleans up unused code ([#51](https://github.com/devopsdunkin/terraform-provider-nagios/pull/51))

BUG FIXES:

* Fixed syntax errors in documentation ([#51](https://github.com/devopsdunkin/terraform-provider-nagios/pull/51))
* Adds PR link to changes in v1.0.0 ([#51](https://github.com/devopsdunkin/terraform-provider-nagios/pull/51))

## 1.0.0 (October 24, 2019)

FEATURES:

* **New Resource:** `resource_host` ([#43](https://github.com/devopsdunkin/terraform-provider-nagios/pull/43))
* **New Resource:** `resource_hostgroup` ([#43](https://github.com/devopsdunkin/terraform-provider-nagios/pull/43))
* **New Resource:** `resource_service` ([#43](https://github.com/devopsdunkin/terraform-provider-nagios/pull/43))
* **New Resource:** `resource_servicegroup` ([#43](https://github.com/devopsdunkin/terraform-provider-nagios/pull/43))