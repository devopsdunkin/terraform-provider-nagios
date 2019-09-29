package nagios

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccHostgroup_basic(t *testing.T) {
	hgName := "tf_" + acctest.RandString(10)
	hgAlias := "tf_" + acctest.RandString(10)
	rName := "nagios_hostgroup.hostgroup"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckHostgroupDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccHostgroupResource_basic(hgName, hgAlias),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostgroupExists(rName),
				),
			},
			// {
			// 	ResourceName:        rName,
			// 	ImportState:         false,
			// 	ImportStateVerify:   false,
			// 	ImportStateIdPrefix: hgName + "/",
			// },
		},
	})
}

func TestAccHostgroup_createAfterManualDestroy(t *testing.T) {
	var hostgroup = &Hostgroup{}
	hgName := "tf_" + acctest.RandString(10)
	hgAlias := "tf_" + acctest.RandString(10)
	rName := "nagios_hostgroup.hostgroup"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckHostgroupDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccHostgroupResource_basic(hgName, hgAlias),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostgroupExists(rName),
					testAccCheckHostgroupFetch(rName, hostgroup),
				),
			},
			{
				PreConfig: func() {
					client := testAccProvider.Meta().(*Client)

					_, err := client.DeleteHostgroup(hostgroup.Name)
					if err != nil {
						t.Fatal(err)
					}
				},
				Config: testAccHostgroupResource_basic(hgName, hgAlias),
				Check:  testAccCheckHostgroupExists(rName),
			},
		},
	})
}

func TestAccHostgroup_updateName(t *testing.T) {
	hgFirstName := "tf_" + acctest.RandString(10)
	hgAlias := "tf_" + acctest.RandString(10)
	hgSecondName := "tf_" + acctest.RandString(10)
	rName := "nagios_hostgroup.hostgroup"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckHostgroupDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccHostgroupResource_basic(hgFirstName, hgAlias),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostgroupExists(rName),
					resource.TestCheckResourceAttr(rName, "name", hgFirstName),
				),
			},
			{
				Config: testAccHostgroupResource_basic(hgSecondName, hgAlias),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostgroupExists(rName),
					resource.TestCheckResourceAttr(rName, "name", hgSecondName),
				),
			},
		},
	})
}

func testAccHostgroupResource_basic(name, alias string) string {
	return fmt.Sprintf(`
resource "nagios_hostgroup" "hostgroup" {
	name = "%s"
	alias = "%s"
}
	`, name, alias)
}

func testAccCheckHostgroupDestroy() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "nagios_hostgroup" {
				continue
			}

			// Get the name of the hostgroup from the state and check if it still exists
			name := rs.Primary.Attributes["name"]

			conn := testAccProvider.Meta().(*Client)

			hostgroup, _ := conn.GetHostgroup(name)
			log.Printf("[DEBUG] Hostgroup - %s", *hostgroup)
			if hostgroup.Name != "" {
				return fmt.Errorf("Hostgroup %s still exists", name)
			}
		}

		log.Printf("[DEBUG] Just seeing when we hit this in logs to deteremine if destroy is getting called early")

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
	log.Printf("[DEBUG] Right before call to s.RootModule().Resources[rName]")
	rs, ok := s.RootModule().Resources[rName]
	if !ok {
		return nil, fmt.Errorf("hostgroup not found: %s", rName)
	}

	log.Printf("[DEBUG] After if !ok check to see if hostgroup not found")

	id := rs.Primary.ID
	name := rs.Primary.Attributes["name"]

	log.Printf("[DEBUG] Name - %s", name)
	log.Printf("[DEBUG] ID - %s", id)

	hostgroup, err := nagiosClient.GetHostgroup(name)

	log.Printf("[DEBUG] Made it after calling GetHostgroup")

	if err != nil {
		return nil, fmt.Errorf("error getting hostgroup with name %s: %s", name, err)
	}

	log.Printf("Made it through gewtHostgroupFromState func. Returning hostgroup - %s", hostgroup)

	return hostgroup, nil
}

func testAccCheckHostgroupFetch(rName string, hostgroup *Hostgroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		returnedUser, err := getHostgroupFromState(s, rName)
		if err != nil {
			return err
		}

		hostgroup.Name = returnedUser.Name
		hostgroup.Alias = returnedUser.Alias

		return nil
	}
}
