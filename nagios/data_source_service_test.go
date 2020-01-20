package nagios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccServiceDataSourceBasic(t *testing.T) {
	serviceName := "tf_" + acctest.RandString(10)
	resourceName := "nagios_service.service1"
	dataSourceName := "data.nagios_service.service2"
	hostName := "localhost"
	serviceDescription := acctest.RandString(25)
	serviceCheckCommand := "check_ping!3000.0!80%!5000.0!100%!!!!"
	serviceMaxCheckAttempts := "2"
	serviceCheckInterval := "5"
	serviceRetryInterval := "5"
	serviceCheckPeriod := "24x7"
	serviceNotificationInterval := "10"
	serviceNotificationPeriod := "24x7"
	serviceContacts := "nagiosadmin"
	serviceTemplates := "generic-service"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckServiceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccServiceDataSourceBasic(serviceName, hostName, serviceDescription, serviceCheckCommand,
					serviceMaxCheckAttempts, serviceCheckInterval, serviceRetryInterval, serviceCheckPeriod, serviceNotificationInterval,
					serviceNotificationPeriod, serviceContacts, serviceTemplates),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "service_name", resourceName, "service_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "host_name", resourceName, "host_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "check_command", resourceName, "check_command"),
					resource.TestCheckResourceAttrPair(dataSourceName, "max_check_attempts", resourceName, "max_check_attempts"),
					resource.TestCheckResourceAttrPair(dataSourceName, "check_interval", resourceName, "check_interval"),
					resource.TestCheckResourceAttrPair(dataSourceName, "retry_interval", resourceName, "retry_interval"),
					resource.TestCheckResourceAttrPair(dataSourceName, "check_period", resourceName, "check_period"),
					resource.TestCheckResourceAttrPair(dataSourceName, "notification_interval", resourceName, "notification_interval"),
					resource.TestCheckResourceAttrPair(dataSourceName, "notification_period", resourceName, "notification_period"),
					resource.TestCheckResourceAttrPair(dataSourceName, "contacts", resourceName, "contacts"),
					resource.TestCheckResourceAttrPair(dataSourceName, "templates", resourceName, "templates"),
				),
			},
		},
	})
}

func testAccServiceDataSourceBasic(serviceName, hostName, description, checkCommand, maxCheckAttempts, checkInterval, retryInterval, checkPeriod, notificationInterval, notificationPeriod, contacts, templates string) string {
	return fmt.Sprintf(`
resource "nagios_service" "service1" {
	service_name			= "%s"
	host_name				= [
		"%s"
	]
	description				= "%s"
	check_command			= "%s"
	max_check_attempts		= "%s"
	check_interval			= "%s"
	retry_interval			= "%s"
	check_period			= "%s"
	notification_interval	= "%s"
	notification_period		= "%s"
	contacts				= [
		"%s"
	]
	templates				= [
		"%s"
	]
	free_variables 			= {
		"_test" = "TestVar123"
	}
}

data "nagios_service" "service2" {
	service_name			= "${nagios_service.service1.service_name}"
	description				= "${nagios_service.service1.description}"
}
	`, serviceName, hostName, description, checkCommand, maxCheckAttempts, checkInterval, retryInterval, checkPeriod, notificationInterval, notificationPeriod, contacts, templates)
}
