package nagios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccHostgroup_basic(t *testing.T) {
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

	rHostgroupName := "nagios_hostgroup.hostgroup"
	rHostName := "nagios_host.host"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckHostgroupDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccHostgroupResource_basic(hostName, alias, address, maxCheckAttempts, checkPeriod, notificationInterval, notificationPeriod, contacts, templates, hgName, hgAlias),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostgroupExists(rHostgroupName),
					testAccCheckHostExists(rHostName),
				),
			},
		},
	})
}

func TestAccHostgroup_createAfterManualDestroy(t *testing.T) {
	var hostgroup = &Hostgroup{}
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

	rHostgroupName := "nagios_hostgroup.hostgroup"
	rHostName := "nagios_host.host"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckHostgroupDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccHostgroupResource_basic(hostName, alias, address, maxCheckAttempts, checkPeriod, notificationInterval, notificationPeriod, contacts, templates, hgName, hgAlias),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostgroupExists(rHostgroupName),
					testAccCheckHostExists(rHostName),
					testAccCheckHostgroupFetch(rHostgroupName, hostgroup),
				),
			},
			{
				PreConfig: func() {
					client := testAccProvider.Meta().(*Client)

					_, err := client.deleteHostgroup(hostgroup.Name)
					if err != nil {
						t.Fatal(err)
					}
				},
				Config: testAccHostgroupResource_basic(hostName, alias, address, maxCheckAttempts, checkPeriod, notificationInterval, notificationPeriod, contacts, templates, hgName, hgAlias),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostgroupExists(rHostgroupName),
					testAccCheckHostExists(rHostName),
				),
			},
		},
	})
}

func TestAccHostgroup_updateName(t *testing.T) {
	// Host group info
	hgFirstName := "tf_" + acctest.RandString(10)
	hgSecondName := "tf_" + acctest.RandString(10)
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

	rHostgroupName := "nagios_hostgroup.hostgroup"
	rHostName := "nagios_host.host"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckHostgroupDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccHostgroupResource_basic(hostName, alias, address, maxCheckAttempts, checkPeriod, notificationInterval, notificationPeriod, contacts, templates, hgFirstName, hgAlias),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostgroupExists(rHostgroupName),
					testAccCheckHostExists(rHostName),
					resource.TestCheckResourceAttr(rHostgroupName, "name", hgFirstName),
				),
			},
			{
				Config: testAccHostgroupResource_basic(hostName, alias, address, maxCheckAttempts, checkPeriod, notificationInterval, notificationPeriod, contacts, templates, hgSecondName, hgAlias),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostgroupExists(rHostgroupName),
					testAccCheckHostExists(rHostName),
					resource.TestCheckResourceAttr(rHostgroupName, "name", hgSecondName),
				),
			},
		},
	})
}

func testAccHostgroupResource_basic(hostName, hostAlias, hostAddress, hostMaxCheckAttempts, hostCheckPeriod, hostNotificationInterval, hostNotificationPeriod, contacts, hostTemplates, hgName, hgAlias string) string {
	// TODO: Need to refactor to support creating N number of hosts and adding N number of hostgroup members

	return fmt.Sprintf(`
	resource "nagios_host" "host" {
		name = "%s"
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

	resource "nagios_hostgroup" "hostgroup" {
		name = "%s"
		alias = "%s"
		members = [
			"%s"
		]
	}
	`, hostName, hostAlias, hostAddress, hostMaxCheckAttempts, hostCheckPeriod, hostNotificationInterval, hostNotificationPeriod, contacts, hostTemplates, hgName, hgAlias, hostName)
}

func testAccCheckHostgroupDestroy() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "nagios_hostgroup" && rs.Type != "nagios_host" {
				continue
			}

			if rs.Type == "nagios_hostgroup" {
				// Get the name of the hostgroup from the state and check if it still exists
				name := rs.Primary.Attributes["name"]

				conn := testAccProvider.Meta().(*Client)

				hostgroup, _ := conn.getHostgroup(name)

				if hostgroup.Name != "" {
					return fmt.Errorf("Hostgroup %s still exists", name)
				}
			} else if rs.Type == "nagios_host" {
				name := rs.Primary.Attributes["name"]

				conn := testAccProvider.Meta().(*Client)

				host, _ := conn.getHost(name)
        
				if host.Name != "" {
					return fmt.Errorf("Host %s still exists", name)
				}
			}
		}

		return nil
	}
}

func testAccCheckHostgroupExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, err := getHostgroupFromState(s, resourceName)
		if err != nil {
			return err
		}

		return nil
	}
}

func getHostgroupFromState(s *terraform.State, rName string) (*Hostgroup, error) {
	nagiosClient := testAccProvider.Meta().(*Client)
	rs, ok := s.RootModule().Resources[rName]
	if !ok {
		return nil, fmt.Errorf("hostgroup not found: %s", rName)
	}

	name := rs.Primary.Attributes["name"]

	hostgroup, err := nagiosClient.getHostgroup(name)

	if err != nil {
		return nil, fmt.Errorf("error getting hostgroup with name %s: %s", name, err)
	}

	return hostgroup, nil
}

func testAccCheckHostgroupFetch(rName string, hostgroup *Hostgroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		returnedHg, err := getHostgroupFromState(s, rName)
		if err != nil {
			return err
		}

		hostgroup.Name = returnedHg.Name
		hostgroup.Alias = returnedHg.Alias
		hostgroup.Members = returnedHg.Members

		return nil
	}
}
