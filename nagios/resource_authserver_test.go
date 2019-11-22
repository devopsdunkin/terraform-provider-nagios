package nagios

// import (
// 	"fmt"
// 	"log"
// 	"testing"

// 	"github.com/hashicorp/terraform/helper/resource"
// 	"github.com/hashicorp/terraform/terraform"
// )

// func TestAccAuthServer_basic(t *testing.T) {
// 	enabled := true
// 	connMethod := "ad"
// 	adAccountSuffix := "test.local"
// 	adDomainControllers := "dc1.test.local"
// 	baseDN := "DC=test,DC=local"
// 	securityLevel := "ssl"
// 	ldapPort := "389"
// 	ldapHost := "ldap.test.local"
// 	resourceName := "nagios_authserver.authserver"

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccCheckAuthServerDestroy(),
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccAuthServerResource_basic(enabled, connMethod, adAccountSuffix, adDomainControllers, baseDN, "", "", securityLevel),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckAuthServerExists(resourceName),
// 				),
// 			},
// 			{
// 				Config: testAccAuthServerResource_basic(enabled, connMethod, "", "", baseDN, ldapPort, ldapHost, securityLevel),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckAuthServerExists(resourceName),
// 				),
// 			},
// 		},
// 	})
// }

// func TestAccAuthServer_createAfterManualDestroy(t *testing.T) {
// 	var authServer = &AuthServer{}
// 	enabled := true
// 	connMethod := "ad"
// 	adAccountSuffix := "test.local"
// 	adDomainControllers := "dc1.test.local"
// 	baseDN := "DC=test,DC=local"
// 	securityLevel := "ssl"
// 	resourceName := "nagios_authserver.authserver"

// 	resource.Test(t, resource.TestCase{
// 		Providers:    testAccProviders,
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		CheckDestroy: testAccCheckAuthServerDestroy(),
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccAuthServerResource_basic(enabled, connMethod, adAccountSuffix, adDomainControllers, baseDN, "", "", securityLevel),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckAuthServerExists(resourceName),
// 					testAccCheckAuthServerFetch(resourceName, authServer),
// 				),
// 			},
// 			{
// 				PreConfig: func() {
// 					client := testAccProvider.Meta().(*Client)

// 					_, err := client.deleteAuthServer(authServer.ID)
// 					if err != nil {
// 						t.Fatal(err)
// 					}
// 				},
// 				Config: testAccAuthServerResource_basic(enabled, connMethod, adAccountSuffix, adDomainControllers, baseDN, "", "", securityLevel),
// 				Check:  testAccCheckContactgroupExists(resourceName),
// 			},
// 		},
// 	})
// }

// func testAccAuthServerResource_basic(enabled bool, connMethod, adAccountsuffix, adDomainControllers, baseDN, ldapPort, ldapHost, securityLevel string) string {
// 	var output string
// 	if connMethod == "ad" {
// 		output = fmt.Sprintf(`
// 		resource "nagios_authserver" "authserver" {
// 			enabled	= "%t"
// 			connection_method		= "%s"
// 			ad_account_suffix		= "%s"
// 			ad_domain_controllers	= "%s"
// 			base_dn					= "%s"
// 		}
// 			`, enabled, connMethod, adAccountsuffix, adDomainControllers, baseDN)
// 	} else {
// 		output = fmt.Sprintf(`
// 		resource "nagios_authserver" "authserver" {
// 			enabled	= "%t"
// 			connection_method	= "%s"
// 			ldap_port			= "%s"
// 			ldap_host			= "%s"
// 			base_dn				= "%s"
// 		}
// 			`, enabled, connMethod, ldapPort, ldapHost, baseDN)
// 	}

// 	return output
// }

// func testAccCheckAuthServerDestroy() resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		for _, rs := range s.RootModule().Resources {
// 			if rs.Type != "nagios_authserver" {
// 				continue
// 			}

// 			// Get the name from the state and check if it still exists
// 			name := rs.Primary.Attributes["server_id"]

// 			conn := testAccProvider.Meta().(*Client)

// 			authServer, _ := conn.getAuthServer(name)
// 			if authServer.ID != "" {
// 				return fmt.Errorf("Auth server %s still exists", name)
// 			}
// 		}

// 		return nil
// 	}
// }

// func testAccCheckAuthServerExists(resourceName string) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		_, err := getAuthServerFromState(s, resourceName)
// 		if err != nil {
// 			return err
// 		}

// 		return nil
// 	}
// }

// func getAuthServerFromState(s *terraform.State, resourceName string) (*AuthServer, error) {
// 	nagiosClient := testAccProvider.Meta().(*Client)
// 	rs, ok := s.RootModule().Resources[resourceName]
// 	if !ok {
// 		return nil, fmt.Errorf("auth server not found: %s", resourceName)
// 	}

// 	name := rs.Primary.Attributes["server_id"]
// 	log.Printf("[DEBUG] Name value from state - %s", name)

// 	authServer, err := nagiosClient.getAuthServer(name)

// 	if err != nil {
// 		return nil, fmt.Errorf("error getting auth server with name %s: %s", name, err)
// 	}

// 	return authServer, nil
// }

// func testAccCheckAuthServerFetch(resourceName string, authServer *AuthServer) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		returnedAuthServer, err := getAuthServerFromState(s, resourceName)
// 		if err != nil {
// 			return err
// 		}

// 		authServer.ID = returnedAuthServer.ID
// 		authServer.Enabled = returnedAuthServer.Enabled
// 		authServer.ConnectionMethod = returnedAuthServer.ConnectionMethod
// 		authServer.ADAccountSuffix = returnedAuthServer.ADAccountSuffix
// 		authServer.ADDomainControllers = returnedAuthServer.ADDomainControllers
// 		authServer.BaseDN = returnedAuthServer.BaseDN
// 		authServer.SecurityLevel = returnedAuthServer.SecurityLevel
// 		authServer.LDAPPort = returnedAuthServer.LDAPPort
// 		authServer.LDAPHost = returnedAuthServer.LDAPHost

// 		return nil
// 	}
// }
