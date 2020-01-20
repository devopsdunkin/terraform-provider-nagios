package nagios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccHostDataSourceBasic(t *testing.T) {
	hostName := "tf-" + acctest.RandString(10)
	resourceName := "nagios_host.host1"
	dataSourceName := "data.nagios_host.host2"
	hostAlias := "tf_" + acctest.RandString(10)
	hostAddress := "127.0.0.1"
	hostMaxCheckAttempts := "5"
	hostCheckPeriod := "24x7"
	hostNotificationInterval := "10"
	hostNotificationPeriod := "24x7"
	hostContacts := "nagiosadmin"
	hostTemplates := "generic-host"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckHostDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccHostDataSourceBasic(hostName, hostAlias, hostAddress, hostMaxCheckAttempts,
					hostCheckPeriod, hostNotificationInterval, hostNotificationPeriod, hostContacts, hostTemplates),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "host_name", resourceName, "host_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "alias", resourceName, "alias"),
					resource.TestCheckResourceAttrPair(dataSourceName, "address", resourceName, "address"),
					resource.TestCheckResourceAttrPair(dataSourceName, "max_check_attempts", resourceName, "max_check_attempts"),
					resource.TestCheckResourceAttrPair(dataSourceName, "check_command", resourceName, "check_command"),
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

func testAccHostDataSourceBasic(hostName, alias, address, maxCheckAttempts, checkPeriod, notificationInterval, notificationPeriod, contacts, templates string) string {
	return fmt.Sprintf(`
resource "nagios_host" "host1" {
	host_name				= "%s"
	alias					= "%s"
	address					= "%s"
	max_check_attempts		= "%s"
	check_command			= "check-host-alive!3000.0!80%%!5000.0!100%%!!!!"
	check_period			= "%s"
	notification_interval	= "%s"
	notification_period		= "%s"
	contacts				= [
									"%s"
							]
	templates				= [
									"%s"
							]
	notes					= "I am adding notes"
	notes_url				= "https://docs.company.local"
	action_url				= "https://docs.company.local"
	initial_state			= "o"
	retry_interval			= "10"
	passive_checks_enabled	= true
	active_checks_enabled	= true
	obsess_over_host		= false
	notification_options	= "d,u,"
	notifications_enabled	= true
	icon_image				= "icon1.jpg"
	free_variables			= {
		"_test" = "test123"
	}
}

data "nagios_host" "host2" {
	host_name = "${nagios_host.host1.host_name}"
}`, hostName, alias, address, maxCheckAttempts, checkPeriod, notificationInterval, notificationPeriod, contacts, templates)
}
