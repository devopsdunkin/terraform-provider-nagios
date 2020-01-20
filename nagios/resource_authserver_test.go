package nagios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAuthServerBasic(t *testing.T) {
	enabled := true
	adAccountSuffix := "@test.local"
	adDomainControllers := "dc1.test.local"
	baseDN := "DC=test,DC=local"
	securityLevel := "ssl"
	ldapPort := "389"
	ldapHost := "ldap.test.local"
	resourceName := "nagios_authserver.authserver"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAuthServerDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccAuthServerResourceBasic(enabled, "ad", adAccountSuffix, adDomainControllers, baseDN, "", "", securityLevel),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthServerExists(resourceName),
				),
			},
			{
				Config: testAccAuthServerResourceBasic(enabled, "ldap", "", "", baseDN, ldapPort, ldapHost, securityLevel),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthServerExists(resourceName),
				),
			},
		},
	})
}

func TestAccAuthServerCreateAfterManualDestroy(t *testing.T) {
	var authServer = &AuthServer{}
	enabled := true
	adAccountSuffix := "@test.local"
	adDomainControllers := "dc1.test.local"
	ldapPort := "389"
	ldapHost := "ldap.test.local"
	baseDN := "DC=test,DC=local"
	securityLevel := "ssl"
	resourceName := "nagios_authserver.authserver"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthServerDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccAuthServerResourceBasic(enabled, "ad", adAccountSuffix, adDomainControllers, baseDN, "", "", securityLevel),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthServerExists(resourceName),
					testAccCheckAuthServerFetch(resourceName, authServer),
				),
			},
			{
				PreConfig: func() {
					client := testAccProvider.Meta().(*Client)

					_, err := client.deleteAuthServer(authServer.ID)
					if err != nil {
						t.Fatal(err)
					}
				},
				Config: testAccAuthServerResourceBasic(enabled, "ldap", "", "", baseDN, ldapPort, ldapHost, securityLevel),
				Check:  testAccCheckAuthServerExists(resourceName),
			},
		},
	})
}

func testAccAuthServerResourceBasic(enabled bool, connMethod, adAccountsuffix, adDomainControllers, baseDN, ldapPort, ldapHost, securityLevel string) string {
	var output string
	if connMethod == "ad" {
		output = fmt.Sprintf(`
		resource "nagios_authserver" "authserver" {
			enabled	= "%t"
			connection_method		= "%s"
			ad_account_suffix		= "%s"
			ad_domain_controllers	= "%s"
			base_dn					= "%s"
			security_level			= "%s"
		}
			`, enabled, connMethod, adAccountsuffix, adDomainControllers, baseDN, securityLevel)
	} else {
		output = fmt.Sprintf(`
		resource "nagios_authserver" "authserver" {
			enabled	= "%t"
			connection_method	= "%s"
			ldap_port			= "%s"
			ldap_host			= "%s"
			base_dn				= "%s"
			security_level		= "%s"
		}
			`, enabled, connMethod, ldapPort, ldapHost, baseDN, securityLevel)
	}

	return output
}

func testAccCheckAuthServerDestroy() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "nagios_authserver" {
				continue
			}

			// Get the ID from the state and check if it still exists
			name := rs.Primary.Attributes["server_id"]

			conn := testAccProvider.Meta().(*Client)

			authServer, _ := conn.getAuthServer(name)
			if authServer.ID != "" {
				return fmt.Errorf("Auth server %s still exists", name)
			}
		}

		return nil
	}
}

func testAccCheckAuthServerExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, err := getAuthServerFromState(s, resourceName)
		if err != nil {
			return err
		}

		return nil
	}
}

func getAuthServerFromState(s *terraform.State, resourceName string) (*AuthServer, error) {
	nagiosClient := testAccProvider.Meta().(*Client)
	rs, ok := s.RootModule().Resources[resourceName]
	if !ok {
		return nil, fmt.Errorf("auth server not found: %s", resourceName)
	}

	ID := rs.Primary.Attributes["server_id"]

	authServer, err := nagiosClient.getAuthServer(ID)

	if err != nil {
		return nil, fmt.Errorf("error getting auth server with name %s: %s", ID, err)
	}

	return authServer, nil
}

func testAccCheckAuthServerFetch(resourceName string, authServer *AuthServer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		returnedAuthServer, err := getAuthServerFromState(s, resourceName)
		if err != nil {
			return err
		}

		authServer.ID = returnedAuthServer.ID
		authServer.ServerID = returnedAuthServer.ServerID
		authServer.Enabled = returnedAuthServer.Enabled
		authServer.ConnectionMethod = returnedAuthServer.ConnectionMethod

		if authServer.ConnectionMethod == "ad" {
			authServer.ADAccountSuffix = returnedAuthServer.ADAccountSuffix
			authServer.ADDomainControllers = returnedAuthServer.ADDomainControllers
		} else {
			authServer.LDAPPort = returnedAuthServer.LDAPPort
			authServer.LDAPHost = returnedAuthServer.LDAPHost
		}

		authServer.BaseDN = returnedAuthServer.BaseDN
		authServer.SecurityLevel = returnedAuthServer.SecurityLevel

		return nil
	}
}
