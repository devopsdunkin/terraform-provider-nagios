# CHANGELOG

## 1.3.0 (December 16, 2019)

FEATURES:

* Adds automated GitHub releases through CircleCI pipeline ([#59](https://github.com/devopsdunkin/terraform-provider-nagios/pull/59))
* Adds `free_variables` field to `resource_host`, `resource_service` and `resource_contact` ([#59](https://github.com/devopsdunkin/terraform-provider-nagios/pull/59))

IMPROVEMENTS:

* Adds centralized function to create URL parameeters for all resources ([#59](https://github.com/devopsdunkin/terraform-provider-nagios/pull/59))
* Adds `omitempty` tag to all optional struct fields to prevent setting options when not specified in schema ([#59](https://github.com/devopsdunkin/terraform-provider-nagios/pull/59))
* Refactors `get` function in `client.go` to return `[]byte` to allow for more flexibility when performing an unmarshal of `[]byte` into an `interface{}` ([#59](https://github.com/devopsdunkin/terraform-provider-nagios/pull/59))
* Fixes formatting and linting issues with `docs/resources/resource_host.md` ([#59](https://github.com/devopsdunkin/terraform-provider-nagios/pull/59))

BUG FIXES:

* NA

## 1.2.0 (November 20, 2019)

FEATURES:

* **New Resource:** `resource_contact` ([#57](https://github.com/devopsdunkin/terraform-provider-nagios/pull/57))
* **New Resource:** `resource_contactgroup` ([#57](https://github.com/devopsdunkin/terraform-provider-nagios/pull/57))

IMPROVEMENTS:

* Adds link to GitHub releases page under the installing provider section of the documentation ([#57](https://github.com/devopsdunkin/terraform-provider-nagios/pull/57))
* Adds support to import state for `resource_host`, `resource_service`, `resource_hostgroup`, `resource_servicegroup`, `resource_contact` and `resource_contactgroup` ([#57](https://github.com/devopsdunkin/terraform-provider-nagios/pull/57))
* Removes unused `schema` tag on structs for `resource_host`, `resource_service`, `resource_hostgroup`, `resource_servicegroup` and `resource_contact` ([#57](https://github.com/devopsdunkin/terraform-provider-nagios/pull/57))
* Adds `register` attribute for `resource_host` and `resource_service`. It is used to set whether the object is active or not in Nagios ([#57](https://github.com/devopsdunkin/terraform-provider-nagios/pull/57))
* Removes unused files from .gitignore ([#57](https://github.com/devopsdunkin/terraform-provider-nagios/pull/57))
* Updates README with updated roadmap and supported versions ([#57](https://github.com/devopsdunkin/terraform-provider-nagios/pull/57))

BUG FIXES:

* Fixes an issue where line breaks were missing from the documentation for resource arguments ([#57](https://github.com/devopsdunkin/terraform-provider-nagios/pull/57))
* Cleans up duplicate code ([#57](https://github.com/devopsdunkin/terraform-provider-nagios/pull/57))

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
