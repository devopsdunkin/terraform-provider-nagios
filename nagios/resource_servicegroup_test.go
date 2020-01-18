package nagios

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccservicegroupBasic(t *testing.T) {
	sgName := "tf_" + acctest.RandString(10)
	sgAlias := "tf_" + acctest.RandString(10)
	rName := "nagios_servicegroup.servicegroup"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckservicegroupDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccservicegroupResourceBasic(sgName, sgAlias),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroupExists(rName),
				),
			},
		},
	})
}

func TestAccservicegroupCreateAfterManualDestroy(t *testing.T) {
	var servicegroup = &Servicegroup{}
	sgName := "tf_" + acctest.RandString(10)
	sgAlias := "tf_" + acctest.RandString(10)
	rName := "nagios_servicegroup.servicegroup"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckservicegroupDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccservicegroupResourceBasic(sgName, sgAlias),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroupExists(rName),
					testAccCheckServicegroupFetch(rName, servicegroup),
				),
			},
			{
				PreConfig: func() {
					client := testAccProvider.Meta().(*Client)

					_, err := client.deleteServicegroup(servicegroup.Name)
					if err != nil {
						t.Fatal(err)
					}
				},
				Config: testAccservicegroupResourceBasic(sgName, sgAlias),
				Check:  testAccCheckServicegroupExists(rName),
			},
		},
	})
}

func TestAccservicegroupUpdateName(t *testing.T) {
	sgFirstName := "tf_" + acctest.RandString(10)
	sgAlias := "tf_" + acctest.RandString(10)
	sgSecondName := "tf_" + acctest.RandString(10)
	rName := "nagios_servicegroup.servicegroup"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckservicegroupDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccservicegroupResourceBasic(sgFirstName, sgAlias),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroupExists(rName),
					resource.TestCheckResourceAttr(rName, "name", sgFirstName),
				),
			},
			{
				Config: testAccservicegroupResourceBasic(sgSecondName, sgAlias),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroupExists(rName),
					resource.TestCheckResourceAttr(rName, "name", sgSecondName),
				),
			},
		},
	})
}

func testAccservicegroupResourceBasic(name, alias string) string {
	return fmt.Sprintf(`
resource "nagios_servicegroup" "servicegroup" {
	name = "%s"
	alias = "%s"
}
	`, name, alias)
}

func testAccCheckservicegroupDestroy() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "nagios_servicegroup" {
				continue
			}

			// Get the name of the servicegroup from the state and check if it still exists
			name := rs.Primary.Attributes["name"]

			conn := testAccProvider.Meta().(*Client)

			servicegroup, _ := conn.getServicegroup(name)
			if servicegroup.Name != "" {
				return fmt.Errorf("servicegroup %s still exists", name)
			}
		}

		log.Printf("[DEBUG] Just seeing when we hit this in logs to deteremine if destroy is getting called early")

		return nil
	}
}

func testAccCheckServicegroupExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, err := getServicegroupFromState(s, resourceName)
		if err != nil {
			return err
		}

		return nil
	}
}

func getServicegroupFromState(s *terraform.State, rName string) (*Servicegroup, error) {
	nagiosClient := testAccProvider.Meta().(*Client)
	rs, ok := s.RootModule().Resources[rName]
	if !ok {
		return nil, fmt.Errorf("servicegroup not found: %s", rName)
	}

	name := rs.Primary.Attributes["name"]

	servicegroup, err := nagiosClient.getServicegroup(name)

	if err != nil {
		return nil, fmt.Errorf("error getting servicegroup with name %s: %s", name, err)
	}

	return servicegroup, nil
}

func testAccCheckServicegroupFetch(rName string, servicegroup *Servicegroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		returnedSg, err := getServicegroupFromState(s, rName)
		if err != nil {
			return err
		}

		servicegroup.Name = returnedSg.Name
		servicegroup.Alias = returnedSg.Alias

		return nil
	}
}
