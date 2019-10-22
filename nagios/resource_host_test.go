package nagios

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccHost_basic(t *testing.T) {
	hostName := "tf_" + acctest.RandString(10)
	hostAlias := "tf_" + acctest.RandString(10)
	hostAddress := "127.0.0.1"
	hostMaxCheckAttempts := "5"
	hostCheckPeriod := "24x7"
	hostNotificationInterval := "10"
	hostNotificationPeriod := "24x7"
	hostContacts := "nagiosadmin"
	hostTemplates := "generic-host"
	rName := "nagios_host.host"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckHostDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccHostResource_basic(hostName, hostAlias, hostAddress, hostMaxCheckAttempts, hostCheckPeriod, hostNotificationInterval, hostNotificationPeriod, hostContacts, hostTemplates),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostExists(rName),
				),
			},
		},
	})
}

func TestAccHost_createAfterManualDestroy(t *testing.T) {
	var host = &Host{}
	hostName := "tf_" + acctest.RandString(10)
	hostAlias := "tf_" + acctest.RandString(10)
	hostAddress := "127.0.0.1"
	hostMaxCheckAttempts := "5"
	hostCheckPeriod := "24x7"
	hostNotificationInterval := "10"
	hostNotificationPeriod := "24x7"
	hostContacts := "nagiosadmin"
	hostTemplates := "generic-host"
	rName := "nagios_host.host"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckHostDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccHostResource_basic(hostName, hostAlias, hostAddress, hostMaxCheckAttempts, hostCheckPeriod, hostNotificationInterval, hostNotificationPeriod, hostContacts, hostTemplates),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostExists(rName),
					testAccCheckHostFetch(rName, host),
				),
			},
			{
				PreConfig: func() {
					client := testAccProvider.Meta().(*Client)

					_, err := client.deleteHost(host.Name)
					if err != nil {
						t.Fatal(err)
					}
				},
				Config: testAccHostResource_basic(hostName, hostAlias, hostAddress, hostMaxCheckAttempts, hostCheckPeriod, hostNotificationInterval, hostNotificationPeriod, hostContacts, hostTemplates),
				Check:  testAccCheckHostExists(rName),
			},
		},
	})
}

func TestAccHost_updateName(t *testing.T) {
	firstHostName := "tf_" + acctest.RandString(10)
	secondHostName := "tf_" + acctest.RandString(10)
	hostAlias := "tf_" + acctest.RandString(10)
	hostAddress := "127.0.0.1"
	hostMaxCheckAttempts := "5"
	hostCheckPeriod := "24x7"
	hostNotificationInterval := "10"
	hostNotificationPeriod := "24x7"
	hostContacts := "nagiosadmin"
	hostTemplates := "generic-host"
	rName := "nagios_host.host"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckHostDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccHostResource_basic(firstHostName, hostAlias, hostAddress, hostMaxCheckAttempts, hostCheckPeriod, hostNotificationInterval, hostNotificationPeriod, hostContacts, hostTemplates),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostExists(rName),
					resource.TestCheckResourceAttr(rName, "name", firstHostName),
				),
			},
			{
				Config: testAccHostResource_basic(secondHostName, hostAlias, hostAddress, hostMaxCheckAttempts, hostCheckPeriod, hostNotificationInterval, hostNotificationPeriod, hostContacts, hostTemplates),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostExists(rName),
					resource.TestCheckResourceAttr(rName, "name", secondHostName),
				),
			},
		},
	})
}

func testAccHostResource_basic(name, alias, address, maxCheckAttempts, checkPeriod, notificationInterval, notificationPeriod, contacts, templates string) string {
	return fmt.Sprintf(`
resource "nagios_host" "host" {
	name					= "%s"
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
}
	`, name, alias, address, maxCheckAttempts, checkPeriod, notificationInterval, notificationPeriod, contacts, templates)
}

func testAccCheckHostDestroy() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "nagios_host" {
				continue
			}

			// Get the name from the state and check if it still exists
			name := rs.Primary.Attributes["name"]

			conn := testAccProvider.Meta().(*Client)

			host, _ := conn.getHost(name)
			if host.Name != "" {
				return fmt.Errorf("Host %s still exists", name)
			}
		}

		return nil
	}
}

func testAccCheckHostExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, err := getHostFromState(s, resourceName)
		if err != nil {
			return err
		}

		return nil
	}
}

func getHostFromState(s *terraform.State, rName string) (*Host, error) {
	nagiosClient := testAccProvider.Meta().(*Client)
	rs, ok := s.RootModule().Resources[rName]
	if !ok {
		return nil, fmt.Errorf("host not found: %s", rName)
	}

	name := rs.Primary.Attributes["name"]
	log.Printf("[DEBUG] Name value from state - %s", name)

	host, err := nagiosClient.getHost(name)

	if err != nil {
		return nil, fmt.Errorf("error getting host with name %s: %s", name, err)
	}

	return host, nil
}

func testAccCheckHostFetch(rName string, host *Host) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		returnedHost, err := getHostFromState(s, rName)
		if err != nil {
			return err
		}

		host.Name = returnedHost.Name
		host.Alias = returnedHost.Alias
		host.Address = returnedHost.Address
		host.MaxCheckAttempts = returnedHost.MaxCheckAttempts
		host.CheckPeriod = returnedHost.CheckPeriod
		host.NotificationInterval = returnedHost.NotificationInterval
		host.NotificationPeriod = returnedHost.NotificationPeriod
		host.Contacts = returnedHost.Contacts
		host.Templates = returnedHost.Templates

		return nil
	}
}
