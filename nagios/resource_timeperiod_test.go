package nagios

// import (
// 	"fmt"
// 	"log"
// 	"testing"

// 	"github.com/hashicorp/terraform/helper/acctest"
// 	"github.com/hashicorp/terraform/helper/resource"
// 	"github.com/hashicorp/terraform/terraform"
// )

// func TestAccTimeperiod_basic(t *testing.T) {
// 	timeperiodName := "tf_" + acctest.RandString(10)
// 	timeperiodAlias := "tf_" + acctest.RandString(10)
// 	timeperiodSunday := "00:00-24:00"
// 	timeperiodMonday := "00:00-24:00"
// 	timeperiodTuesday := "00:00-24:00"
// 	timeperiodWednesday := "00:00-24:00"
// 	timeperiodThursday := "00:00-24:00"
// 	timeperiodFriday := "00:00-24:00"
// 	timeperiodSaturday := "00:00-24:00"
// 	resourceName := "nagios_timeperiod.timeperiod"

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccCheckTimeperiodDestroy(),
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccTimeperiodResource_basic(timeperiodName, timeperiodAlias, timeperiodSunday, timeperiodMonday, timeperiodTuesday, timeperiodWednesday, timeperiodThursday, timeperiodFriday, timeperiodSaturday),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckTimeperiodExists(resourceName),
// 				),
// 			},
// 		},
// 	})
// }

// func TestAccTimeperiod_createAfterManualDestroy(t *testing.T) {
// 	var timeperiod = &Timeperiod{}
// 	timeperiodName := "tf_" + acctest.RandString(10)
// 	timeperiodAlias := "tf_" + acctest.RandString(10)
// 	timeperiodSunday := "00:00-24:00"
// 	timeperiodMonday := "00:00-24:00"
// 	timeperiodTuesday := "00:00-24:00"
// 	timeperiodWednesday := "00:00-24:00"
// 	timeperiodThursday := "00:00-24:00"
// 	timeperiodFriday := "00:00-24:00"
// 	timeperiodSaturday := "00:00-24:00"
// 	resourceName := "nagios_timeperiod.timeperiod"

// 	resource.Test(t, resource.TestCase{
// 		Providers:    testAccProviders,
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		CheckDestroy: testAccCheckTimeperiodDestroy(),
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccTimeperiodResource_basic(timeperiodName, timeperiodAlias, timeperiodSunday, timeperiodMonday, timeperiodTuesday, timeperiodWednesday, timeperiodThursday, timeperiodFriday, timeperiodSaturday),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckTimeperiodExists(resourceName),
// 					testAccCheckTimeperiodFetch(resourceName, timeperiod),
// 				),
// 			},
// 			{
// 				PreConfig: func() {
// 					client := testAccProvider.Meta().(*Client)

// 					_, err := client.deleteTimeperiod(timeperiod.TimeperiodName)
// 					if err != nil {
// 						t.Fatal(err)
// 					}
// 				},
// 				Config: testAccTimeperiodResource_basic(timeperiodName, timeperiodAlias, timeperiodSunday, timeperiodMonday, timeperiodTuesday, timeperiodWednesday, timeperiodThursday, timeperiodFriday, timeperiodSaturday),
// 				Check:  testAccCheckTimeperiodExists(resourceName),
// 			},
// 		},
// 	})
// }

// func TestAccTimeperiod_updateName(t *testing.T) {
// 	firstTimeperiodName := "tf_" + acctest.RandString(10)
// 	secondTimeperiodName := "tf_" + acctest.RandString(10)
// 	timeperiodAlias := "tf_" + acctest.RandString(10)
// 	timeperiodSunday := "00:00-24:00"
// 	timeperiodMonday := "00:00-24:00"
// 	timeperiodTuesday := "00:00-24:00"
// 	timeperiodWednesday := "00:00-24:00"
// 	timeperiodThursday := "00:00-24:00"
// 	timeperiodFriday := "00:00-24:00"
// 	timeperiodSaturday := "00:00-24:00"
// 	resourceName := "nagios_timeperiod.timeperiod"

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccCheckTimeperiodDestroy(),
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccTimeperiodResource_basic(firstTimeperiodName, timeperiodAlias, timeperiodSunday, timeperiodMonday, timeperiodTuesday, timeperiodWednesday, timeperiodThursday, timeperiodFriday, timeperiodSaturday),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckTimeperiodExists(resourceName),
// 					resource.TestCheckResourceAttr(resourceName, "timeperiod_name", firstTimeperiodName),
// 				),
// 			},
// 			{
// 				Config: testAccTimeperiodResource_basic(secondTimeperiodName, timeperiodAlias, timeperiodSunday, timeperiodMonday, timeperiodTuesday, timeperiodWednesday, timeperiodThursday, timeperiodFriday, timeperiodSaturday),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckTimeperiodExists(resourceName),
// 					resource.TestCheckResourceAttr(resourceName, "timeperiod_name", secondTimeperiodName),
// 				),
// 			},
// 		},
// 	})
// }

// func testAccTimeperiodResource_basic(timeperiodName, timeperiodAlias, timeperiodSunday, timeperiodMonday, timeperiodTuesday, timeperiodWednesday, timeperiodThursday, timeperiodFriday, timeperiodSaturday string) string {
// 	return fmt.Sprintf(`
// resource "nagios_timeperiod" "timeperiod" {
// 	timeperiod_name		= "%s"
// 	alias				= "%s"
// 	sunday				= "%s"
// 	monday				= "%s"
// 	tuesday				= "%s"
// 	wednesday			= "%s"
// 	thursday			= "%s"
// 	friday				= "%s"
// 	saturday			= "%s"
// }
// 	`, timeperiodName, timeperiodAlias, timeperiodSunday, timeperiodMonday, timeperiodTuesday, timeperiodWednesday, timeperiodThursday, timeperiodFriday, timeperiodSaturday)
// }

// func testAccCheckTimeperiodDestroy() resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		for _, rs := range s.RootModule().Resources {
// 			if rs.Type != "nagios_timeperiod" {
// 				continue
// 			}

// 			// Get the name from the state and check if it still exists
// 			name := rs.Primary.Attributes["timeperiod_name"]

// 			conn := testAccProvider.Meta().(*Client)

// 			timeperiod, _ := conn.getTimeperiod(name)
// 			if timeperiod.TimeperiodName != "" {
// 				return fmt.Errorf("Timeperiod %s still exists", name)
// 			}
// 		}

// 		return nil
// 	}
// }

// func testAccCheckTimeperiodExists(resourceName string) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		_, err := getTimeperiodFromState(s, resourceName)
// 		if err != nil {
// 			return err
// 		}

// 		return nil
// 	}
// }

// func getTimeperiodFromState(s *terraform.State, resourceName string) (*Timeperiod, error) {
// 	nagiosClient := testAccProvider.Meta().(*Client)
// 	rs, ok := s.RootModule().Resources[resourceName]
// 	if !ok {
// 		return nil, fmt.Errorf("timeperiod not found: %s", resourceName)
// 	}

// 	name := rs.Primary.Attributes["timeperiod_name"]
// 	log.Printf("[DEBUG] Name value from state - %s", name)

// 	contact, err := nagiosClient.getTimeperiod(name)

// 	if err != nil {
// 		return nil, fmt.Errorf("error getting timeperiod with name %s: %s", name, err)
// 	}

// 	return contact, nil
// }

// func testAccCheckTimeperiodFetch(resourceName string, timeperiod *Timeperiod) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		returnedTimeperiod, err := getTimeperiodFromState(s, resourceName)
// 		if err != nil {
// 			return err
// 		}

// 		timeperiod.TimeperiodName = returnedTimeperiod.TimeperiodName
// 		timeperiod.Alias = returnedTimeperiod.Alias

// 		if timeperiod.Sunday != "" {
// 			timeperiod.Sunday = returnedTimeperiod.Sunday
// 		}

// 		if timeperiod.Monday != "" {
// 			timeperiod.Monday = returnedTimeperiod.Monday
// 		}

// 		if timeperiod.Tuesday != "" {
// 			timeperiod.Tuesday = returnedTimeperiod.Tuesday
// 		}

// 		if timeperiod.Wednesday != "" {
// 			timeperiod.Wednesday = returnedTimeperiod.Wednesday
// 		}

// 		if timeperiod.Thursday != "" {
// 			timeperiod.Thursday = returnedTimeperiod.Thursday
// 		}

// 		if timeperiod.Friday != "" {
// 			timeperiod.Friday = returnedTimeperiod.Friday
// 		}

// 		if timeperiod.Saturday != "" {
// 			timeperiod.Saturday = returnedTimeperiod.Saturday
// 		}

// 		return nil
// 	}
// }
