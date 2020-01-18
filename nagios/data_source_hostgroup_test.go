package nagios

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccHostgroupDataSource_basic(t *testing.T) {
	// Host group info
	hgName := "tf_" + acctest.RandString(10)
	hgAlias := "tf_" + acctest.RandString(10)

	// Hosts to add as hostgroup members
	hostName := "test1"
	alias := "test1"
	address := "127.0.0.1"
	maxCheckAttempts := "2"
	checkPeriod := "24x7"
	notificationInterval := "2"
	notificationPeriod := "24x7"
	contacts := "nagiosadmin"
	templates := "generic-host"

	resourceName := "nagios_hostgroup.hostgroup1"
	dataSourceName := "data.nagios_hostgroup.hostgroup2"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckHostgroupDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccHostgroupDataSourceBasic(hostName, alias, address, maxCheckAttempts,
					checkPeriod, notificationInterval, notificationPeriod, contacts, templates, hgName, hgAlias),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "alias", resourceName, "alias"),
					resource.TestCheckResourceAttrPair(dataSourceName, "members", resourceName, "members"),
				),
			},
		},
	})
}

func testAccHostgroupDataSourceBasic(hostName, hostAlias, hostAddress, hostMaxCheckAttempts, hostCheckPeriod, hostNotificationInterval, hostNotificationPeriod, contacts, hostTemplates, hgName, hgAlias string) string {
	log.Printf("hostgroup name: %s", hgName)

	return fmt.Sprintf(`
	resource "nagios_host" "host" {
		host_name = "%s"
		alias = "%s"
		address = "%s"
		max_check_attempts = "%s"
		check_period = "%s"
		notification_interval = "%s"
		notification_period = "%s"
		contacts = [
			"%s"
		]
		templates = [
			"%s"
		]
	}

	resource "nagios_hostgroup" "hostgroup1" {
		name = "%s"
		alias = "%s"
		members = [
			"%s"
		]
	}

	data "nagios_hostgroup" "hostgroup2" {
		name = "${nagios_hostgroup.hostgroup1.name}"
	}
	`, hostName, hostAlias, hostAddress, hostMaxCheckAttempts, hostCheckPeriod, hostNotificationInterval, hostNotificationPeriod, contacts, hostTemplates, hgName, hgAlias, hostName)
}
