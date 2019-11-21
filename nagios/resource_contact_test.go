package nagios

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccContact_basic(t *testing.T) {
	contactName := "tf_" + acctest.RandString(10)
	contactHostNotificationPeriod := "24x7"
	contactServiceNotificationPeriod := "24x7"
	contactHostNotificationOptions := "d"
	contactServiceNotificationOptions := "d"
	contactHostNotificationCommands := "notify-host-by-email"
	contactServiceNotificationCommands := "notify-host-by-email"
	contactAlias := "tf_" + acctest.RandString(10)
	contactEmail := acctest.RandString(10) + "@example.com"
	contactTemplates := "generic-contact"
	rName := "nagios_contact.contact"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckContactDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccContactResource_basic(contactName, contactHostNotificationPeriod, contactServiceNotificationPeriod, contactHostNotificationOptions, contactServiceNotificationOptions, contactHostNotificationCommands, contactServiceNotificationCommands, contactAlias, contactTemplates, contactEmail),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContactExists(rName),
				),
			},
		},
	})
}

func TestAccContact_createAfterManualDestroy(t *testing.T) {
	var contact = &Contact{}
	contactName := "tf_" + acctest.RandString(10)
	contactHostNotificationPeriod := "24x7"
	contactServiceNotificationPeriod := "24x7"
	contactHostNotificationOptions := "d"
	contactServiceNotificationOptions := "d"
	contactHostNotificationCommands := "notify-host-by-email"
	contactServiceNotificationCommands := "notify-host-by-email"
	contactAlias := "tf_" + acctest.RandString(10)
	contactEmail := acctest.RandString(10) + "@example.com"
	contactTemplates := "generic-contact"
	rName := "nagios_contact.contact"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckContactDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccContactResource_basic(contactName, contactHostNotificationPeriod, contactServiceNotificationPeriod, contactHostNotificationOptions, contactServiceNotificationOptions, contactHostNotificationCommands, contactServiceNotificationCommands, contactAlias, contactTemplates, contactEmail),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContactExists(rName),
					testAccCheckContactFetch(rName, contact),
				),
			},
			{
				PreConfig: func() {
					client := testAccProvider.Meta().(*Client)

					_, err := client.deleteContact(contact.ContactName)
					if err != nil {
						t.Fatal(err)
					}
				},
				Config: testAccContactResource_basic(contactName, contactHostNotificationPeriod, contactServiceNotificationPeriod, contactHostNotificationOptions, contactServiceNotificationOptions, contactHostNotificationCommands, contactServiceNotificationCommands, contactAlias, contactTemplates, contactEmail),
				Check:  testAccCheckContactExists(rName),
			},
		},
	})
}

func TestAccContact_updateName(t *testing.T) {
	firstContactName := "tf_" + acctest.RandString(10)
	secondContactName := "tf_" + acctest.RandString(10)
	contactHostNotificationPeriod := "24x7"
	contactServiceNotificationPeriod := "24x7"
	contactHostNotificationOptions := "d"
	contactServiceNotificationOptions := "d"
	contactHostNotificationCommands := "notify-host-by-email"
	contactServiceNotificationCommands := "notify-host-by-email"
	contactAlias := "tf_" + acctest.RandString(10)
	contactEmail := acctest.RandString(10) + "@example.com"
	contactTemplates := "generic-contact"
	rName := "nagios_contact.contact"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckContactDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccContactResource_basic(firstContactName, contactHostNotificationPeriod, contactServiceNotificationPeriod, contactHostNotificationOptions, contactServiceNotificationOptions, contactHostNotificationCommands, contactServiceNotificationCommands, contactAlias, contactTemplates, contactEmail),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContactExists(rName),
					resource.TestCheckResourceAttr(rName, "contact_name", firstContactName),
				),
			},
			{
				Config: testAccContactResource_basic(secondContactName, contactHostNotificationPeriod, contactServiceNotificationPeriod, contactHostNotificationOptions, contactServiceNotificationOptions, contactHostNotificationCommands, contactServiceNotificationCommands, contactAlias, contactTemplates, contactEmail),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContactExists(rName),
					resource.TestCheckResourceAttr(rName, "contact_name", secondContactName),
				),
			},
		},
	})
}

func testAccContactResource_basic(contactName, hostNotificationPeriod, serviceNotificationPeriod, hostNotificationOptions, serviceNotificationOptions, hostNotificationCommands, serviceNotificationCommands, alias, templates, email string) string {
	return fmt.Sprintf(`
resource "nagios_contact" "contact" {
	contact_name				  = "%s"
	host_notifications_enabled	  = true
	service_notifications_enabled = true
	host_notification_period	  = "%s"
	service_notification_period	  = "%s"
	host_notification_options	  = "%s"
	service_notification_options  = "%s"
	host_notification_commands	  = [
		"%s"
	]
	service_notification_commands = [
		"%s"
	]
	alias						  = "%s"
	templates					  = [
		"%s"
	]
	email						  = "%s"
	can_submit_commands			  = true
}
	`, contactName, hostNotificationPeriod, serviceNotificationPeriod, hostNotificationOptions, serviceNotificationOptions, hostNotificationCommands, serviceNotificationCommands, alias, templates, email)
}

func testAccCheckContactDestroy() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "nagios_contact" {
				continue
			}

			// Get the name from the state and check if it still exists
			name := rs.Primary.Attributes["contact_name"]

			conn := testAccProvider.Meta().(*Client)

			contact, _ := conn.getContact(name)
			if contact.ContactName != "" {
				return fmt.Errorf("Contact %s still exists", name)
			}
		}

		return nil
	}
}

func testAccCheckContactExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, err := getContactFromState(s, resourceName)
		if err != nil {
			return err
		}

		return nil
	}
}

func getContactFromState(s *terraform.State, rName string) (*Contact, error) {
	nagiosClient := testAccProvider.Meta().(*Client)
	rs, ok := s.RootModule().Resources[rName]
	if !ok {
		return nil, fmt.Errorf("contact not found: %s", rName)
	}

	name := rs.Primary.Attributes["contact_name"]
	log.Printf("[DEBUG] Name value from state - %s", name)

	contact, err := nagiosClient.getContact(name)

	if err != nil {
		return nil, fmt.Errorf("error getting contact with name %s: %s", name, err)
	}

	return contact, nil
}

func testAccCheckContactFetch(rName string, contact *Contact) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		returnedContact, err := getContactFromState(s, rName)
		if err != nil {
			return err
		}

		contact.ContactName = returnedContact.ContactName
		contact.HostNotificationsEnabled = returnedContact.HostNotificationsEnabled
		contact.ServiceNotificationsEnabled = returnedContact.ServiceNotificationsEnabled
		contact.HostNotificationPeriod = returnedContact.HostNotificationPeriod
		contact.ServiceNotificationPeriod = returnedContact.ServiceNotificationPeriod
		contact.HostNotificationOptions = returnedContact.HostNotificationOptions
		contact.ServiceNotificationOptions = returnedContact.ServiceNotificationOptions
		contact.HostNotificationCommands = returnedContact.HostNotificationCommands
		contact.ServiceNotificationCommands = returnedContact.ServiceNotificationCommands

		// Optional attributes
		if returnedContact.Alias != "" {
			contact.Alias = returnedContact.Alias
		}

		if returnedContact.Templates != nil {
			contact.Templates = returnedContact.Templates
		}

		if returnedContact.Email != "" {
			contact.Email = returnedContact.Email
		}

		if returnedContact.CanSubmitCommands != "" {
			contact.CanSubmitCommands = returnedContact.CanSubmitCommands
		}

		return nil
	}
}
