package nagios

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccContactgroupBasic(t *testing.T) {
	contactgroupName := "tf_" + acctest.RandString(10)
	contactgroupAlias := "tf_" + acctest.RandString(10)
	contactgroupMembers := "nagiosadmin"
	resourceName := "nagios_contactgroup.contactgroup"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckContactgroupDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccContactgroupResourceBasic(contactgroupName, contactgroupAlias, contactgroupMembers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContactgroupExists(resourceName),
				),
			},
		},
	})
}

func TestAccContactgroupCreateAfterManualDestroy(t *testing.T) {
	var contactgroup = &Contactgroup{}
	contactgroupName := "tf_" + acctest.RandString(10)
	contactgroupAlias := "tf_" + acctest.RandString(10)
	contactgroupMembers := "nagiosadmin"
	resourceName := "nagios_contactgroup.contactgroup"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckContactgroupDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccContactgroupResourceBasic(contactgroupName, contactgroupAlias, contactgroupMembers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContactgroupExists(resourceName),
					testAccCheckContactgroupFetch(resourceName, contactgroup),
				),
			},
			{
				PreConfig: func() {
					client := testAccProvider.Meta().(*Client)

					_, err := client.deleteContactgroup(contactgroup.ContactgroupName)
					if err != nil {
						t.Fatal(err)
					}
				},
				Config: testAccContactgroupResourceBasic(contactgroupName, contactgroupAlias, contactgroupMembers),
				Check:  testAccCheckContactgroupExists(resourceName),
			},
		},
	})
}

func TestAccContactgroupUpdateName(t *testing.T) {
	firstContactgroupName := "tf_" + acctest.RandString(10)
	secondContactgroupName := "tf_" + acctest.RandString(10)
	contactgroupAlias := "tf_" + acctest.RandString(10)
	contactgroupMembers := "nagiosadmin"
	resourceName := "nagios_contactgroup.contactgroup"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckContactgroupDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccContactgroupResourceBasic(firstContactgroupName, contactgroupAlias, contactgroupMembers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContactgroupExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "contactgroup_name", firstContactgroupName),
				),
			},
			{
				Config: testAccContactgroupResourceBasic(secondContactgroupName, contactgroupAlias, contactgroupMembers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContactgroupExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "contactgroup_name", secondContactgroupName),
				),
			},
		},
	})
}

func testAccContactgroupResourceBasic(contactgroupName, alias, members string) string {
	return fmt.Sprintf(`
resource "nagios_contactgroup" "contactgroup" {
	contactgroup_name	= "%s"
	alias				= "%s"
	members				= [
		"%s"
	]
}
	`, contactgroupName, alias, members)
}

func testAccCheckContactgroupDestroy() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "nagios_contactgroup" {
				continue
			}

			// Get the name from the state and check if it still exists
			name := rs.Primary.Attributes["contactgroup_name"]

			conn := testAccProvider.Meta().(*Client)

			contactgroup, _ := conn.getContactgroup(name)
			if contactgroup.ContactgroupName != "" {
				return fmt.Errorf("Contact group %s still exists", name)
			}
		}

		return nil
	}
}

func testAccCheckContactgroupExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, err := getContactgroupFromState(s, resourceName)
		if err != nil {
			return err
		}

		return nil
	}
}

func getContactgroupFromState(s *terraform.State, resourceName string) (*Contactgroup, error) {
	nagiosClient := testAccProvider.Meta().(*Client)
	rs, ok := s.RootModule().Resources[resourceName]
	if !ok {
		return nil, fmt.Errorf("contact group not found: %s", resourceName)
	}

	name := rs.Primary.Attributes["contactgroup_name"]
	log.Printf("[DEBUG] Name value from state - %s", name)

	contactgroup, err := nagiosClient.getContactgroup(name)

	if err != nil {
		return nil, fmt.Errorf("error getting contact group with name %s: %s", name, err)
	}

	return contactgroup, nil
}

func testAccCheckContactgroupFetch(resourceName string, contactgroup *Contactgroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		returnedContactgroup, err := getContactgroupFromState(s, resourceName)
		if err != nil {
			return err
		}

		contactgroup.ContactgroupName = returnedContactgroup.ContactgroupName
		contactgroup.Alias = returnedContactgroup.Alias
		contactgroup.Members = returnedContactgroup.Members
		contactgroup.ContactgroupMembers = returnedContactgroup.ContactgroupMembers

		return nil
	}
}
