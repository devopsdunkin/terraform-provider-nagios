package nagios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccHostBasic(t *testing.T) {
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
				Config: testAccHostResourceBasic(hostName, hostAlias, hostAddress, hostMaxCheckAttempts, hostCheckPeriod, hostNotificationInterval, hostNotificationPeriod, hostContacts, hostTemplates),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostExists(rName),
				),
			},
		},
	})
}

func TestAccHostCreateAfterManualDestroy(t *testing.T) {
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
				Config: testAccHostResourceBasic(hostName, hostAlias, hostAddress, hostMaxCheckAttempts, hostCheckPeriod, hostNotificationInterval, hostNotificationPeriod, hostContacts, hostTemplates),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostExists(rName),
					testAccCheckHostFetch(rName, host),
				),
			},
			{
				PreConfig: func() {
					client := testAccProvider.Meta().(*Client)

					_, err := client.deleteHost(host.HostName)
					if err != nil {
						t.Fatal(err)
					}
				},
				Config: testAccHostResourceBasic(hostName, hostAlias, hostAddress, hostMaxCheckAttempts, hostCheckPeriod, hostNotificationInterval, hostNotificationPeriod, hostContacts, hostTemplates),
				Check:  testAccCheckHostExists(rName),
			},
		},
	})
}

func TestAccHostUpdateName(t *testing.T) {
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
				Config: testAccHostResourceBasic(firstHostName, hostAlias, hostAddress, hostMaxCheckAttempts, hostCheckPeriod, hostNotificationInterval, hostNotificationPeriod, hostContacts, hostTemplates),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostExists(rName),
					resource.TestCheckResourceAttr(rName, "host_name", firstHostName),
				),
			},
			{
				Config: testAccHostResourceBasic(secondHostName, hostAlias, hostAddress, hostMaxCheckAttempts, hostCheckPeriod, hostNotificationInterval, hostNotificationPeriod, hostContacts, hostTemplates),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostExists(rName),
					resource.TestCheckResourceAttr(rName, "host_name", secondHostName),
				),
			},
		},
	})
}

func testAccHostResourceBasic(name, alias, address, maxCheckAttempts, checkPeriod, notificationInterval, notificationPeriod, contacts, templates string) string {
	return fmt.Sprintf(`
resource "nagios_host" "host" {
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

	`, name, alias, address, maxCheckAttempts, checkPeriod, notificationInterval, notificationPeriod, contacts, templates)
}

func testAccCheckHostDestroy() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "nagios_host" {
				continue
			}

			// Get the name from the state and check if it still exists
			name := rs.Primary.Attributes["host_name"]

			conn := testAccProvider.Meta().(*Client)

			host, _ := conn.getHost(name)
			if host.HostName != "" {
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

	name := rs.Primary.Attributes["host_name"]

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

		host.HostName = returnedHost.HostName
		host.Alias = returnedHost.Alias
		host.Address = returnedHost.Address
		host.MaxCheckAttempts = returnedHost.MaxCheckAttempts
		host.CheckPeriod = returnedHost.CheckPeriod
		host.NotificationInterval = returnedHost.NotificationInterval
		host.NotificationPeriod = returnedHost.NotificationPeriod
		host.Contacts = returnedHost.Contacts

		// Optional attributes
		if returnedHost.Templates != nil {
			host.Templates = returnedHost.Templates
		}
		if returnedHost.CheckCommand != "" {
			host.CheckCommand = returnedHost.CheckCommand
		}

		if returnedHost.ContactGroups != nil {
			host.ContactGroups = returnedHost.ContactGroups
		}

		if returnedHost.Notes != "" {
			host.Notes = returnedHost.Notes
		}

		if returnedHost.NotesURL != "" {
			host.NotesURL = returnedHost.NotesURL
		}

		if returnedHost.ActionURL != "" {
			host.ActionURL = returnedHost.ActionURL
		}

		if returnedHost.InitialState != "" {
			host.InitialState = returnedHost.InitialState
		}

		if returnedHost.RetryInterval != "" {
			host.RetryInterval = returnedHost.RetryInterval
		}

		if returnedHost.PassiveChecksEnabled != "" {
			host.PassiveChecksEnabled = returnedHost.PassiveChecksEnabled
		}

		if returnedHost.ActiveChecksEnabled != "" {
			host.ActiveChecksEnabled = returnedHost.ActiveChecksEnabled
		}

		if returnedHost.ObsessOverHost != "" {
			host.ObsessOverHost = returnedHost.ObsessOverHost
		}

		if returnedHost.EventHandler != "" {
			host.EventHandler = returnedHost.EventHandler
		}

		if returnedHost.EventHandlerEnabled != "" {
			host.EventHandlerEnabled = returnedHost.EventHandlerEnabled
		}

		if returnedHost.FlapDetectionEnabled != "" {
			host.FlapDetectionEnabled = returnedHost.FlapDetectionEnabled
		}

		if returnedHost.FlapDetectionOptions != nil {
			host.FlapDetectionOptions = returnedHost.FlapDetectionOptions
		}

		if returnedHost.LowFlapThreshold != "" {
			host.LowFlapThreshold = returnedHost.LowFlapThreshold
		}

		if returnedHost.HighFlapThreshold != "" {
			host.HighFlapThreshold = returnedHost.HighFlapThreshold
		}

		if returnedHost.ProcessPerfData != "" {
			host.ProcessPerfData = returnedHost.ProcessPerfData
		}

		if returnedHost.RetainStatusInformation != "" {
			host.RetainStatusInformation = returnedHost.RetainStatusInformation
		}

		if returnedHost.RetainNonstatusInformation != "" {
			host.RetainNonstatusInformation = returnedHost.RetainNonstatusInformation
		}

		if returnedHost.CheckFreshness != "" {
			host.CheckFreshness = returnedHost.CheckFreshness
		}

		if returnedHost.FreshnessThreshold != "" {
			host.FreshnessThreshold = returnedHost.FreshnessThreshold
		}

		if returnedHost.FirstNotificationDelay != "" {
			host.FirstNotificationDelay = returnedHost.FirstNotificationDelay
		}

		if returnedHost.NotificationOptions != "" {
			host.NotificationOptions = returnedHost.NotificationOptions
		}

		if returnedHost.NotificationsEnabled != "" {
			host.NotificationsEnabled = returnedHost.NotificationsEnabled
		}

		if returnedHost.StalkingOptions != "" {
			host.StalkingOptions = returnedHost.StalkingOptions
		}

		if returnedHost.IconImage != "" {
			host.IconImage = returnedHost.IconImage
		}

		if returnedHost.IconImageAlt != "" {
			host.IconImageAlt = returnedHost.IconImageAlt
		}

		if returnedHost.VRMLImage != "" {
			host.VRMLImage = returnedHost.VRMLImage
		}

		if returnedHost.StatusMapImage != "" {
			host.StatusMapImage = returnedHost.StatusMapImage
		}

		if returnedHost.TwoDCoords != "" {
			host.TwoDCoords = returnedHost.TwoDCoords
		}

		if returnedHost.ThreeDCoords != "" {
			host.ThreeDCoords = returnedHost.ThreeDCoords
		}

		if returnedHost.FreeVariables != nil {
			host.FreeVariables = returnedHost.FreeVariables
		}

		return nil
	}
}
