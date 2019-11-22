package nagios

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

type AuthServer struct {
	ID                  string `json:"server_id"`
	Enabled             string `json:"enabled"`
	ConnectionMethod    string `json:"conn_method"`
	ADAccountSuffix     string `json:"ad_account_suffix,omitempty"`
	ADDomainControllers string `json:"ad_domain_controllers,omitempty"`
	BaseDN              string `json:"base_dn,omitempty"`
	SecurityLevel       string `json:"security_level,omitempty"`
	LDAPPort            string `json:"ldap_port,omitempty"`
	LDAPHost            string `json:"ldap_host,omitempty"`
}

type MapOfAuthServers struct {
	AuthServerEntry []AuthServer `json:"authservers"`
}

// The Nagios REST API does not support PUT for authentication servers. So any change requires Terraform to destroy and add a new auth server
func resourceAuthServer() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"server_id": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
				Description: "The ID of the authentication server. This value is computed by Nagios",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Determines whether or not the contact will receive notifications about host problems and recoveries",
			},
			"connection_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The connection method used for authentication. This value can be either 'ad' or 'ldap'",
				ForceNew:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v != "ad" && v != "ldap" {
						errs = append(errs, fmt.Errorf("%q must be one of the following: ad or ldap, got: %s", key, v))
					}
					return
				},
			},
			"ad_account_suffix": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The account suffix that should be used. Thsi value is required when the connection method is 'ad'",
			},
			"ad_domain_controllers": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "A comma separated list of domain controllers to use for Active Directory authentication",
			},
			"base_dn": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The Base DN where the user accounts exist in AD or LDAP that will be authenticating to Nagios",
			},
			"security_level": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "none",
				Description: "The security level to be used to enerypt the connection. It can be either 'none', 'ssl  or 'tls'",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v != "ssl" && v != "tls" && v != "none" {
						errs = append(errs, fmt.Errorf("%q must be one of the following: ssl, tls or none, got: %s", key, v))
					}
					return
				},
			},
			"ldap_port": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The TCP port to use when connecting with LDAP.",
			},
			"ldap_host": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The LDAP host name or IP address to connect to",
			},
		},
		Create: resourceCreateAuthServer,
		Read:   resourceReadAuthServer,
		Update: resourceUpdateAuthServer,
		Delete: resourceDeleteAuthServer,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateAuthServer(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	authServer := setAuthServerFromSchema(d)
	// tempObj := &AuthServer{}

	body, err := nagiosClient.newAuthServer(authServer)

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &authServer)

	if err != nil {
		return err
	}

	// authServer.ServerID = tempObj.ServerID

	// log.Printf("[DEBUG] tempObj = %s", tempObj.ServerID)
	log.Printf("[DEBUG] authServer = %s", authServer)

	d.SetId(authServer.ID)

	log.Printf("[DEBUG] authServer inside CREATE = %s", authServer)

	return resourceReadAuthServer(d, m)
}

func resourceReadAuthServer(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	authServer, err := nagiosClient.getAuthServer(d.Id())

	log.Printf("[DEBUG] authServer inside resourceReadAuthServer = %s", authServer)
	if err != nil {
		return err
	}

	if authServer == nil {
		// contact not found. Let Terraform know to delete the state
		d.SetId("")
		return nil
	}

	log.Printf("[DEBUG] Right before calling setDataFromAuthServer()")

	setDataFromAuthServer(d, authServer)

	log.Printf("[DEBUG] authServer inside READ = %s", authServer)

	return nil
}

func resourceUpdateAuthServer(d *schema.ResourceData, m interface{}) error {
	return resourceCreateAuthServer(d, m)
}

func resourceDeleteAuthServer(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	_, err := nagiosClient.deleteAuthServer(d.Id())

	if err != nil {
		return err
	}

	return nil
}

func setDataFromAuthServer(d *schema.ResourceData, authServer *AuthServer) {
	d.SetId(authServer.ID)
	log.Printf("[DEBUG] set ID")
	d.Set("server_id", authServer.ID)
	log.Printf("[DEBUG] Set server_id")
	d.Set("enabled", authServer.Enabled)
	d.Set("connection_method", authServer.ConnectionMethod)

	if authServer.ConnectionMethod == "ad" {
		d.Set("ad_account_suffix", authServer.ADAccountSuffix)
		d.Set("ad_domain_controllers", authServer.ADDomainControllers)
	} else {
		d.Set("ldap_port", authServer.LDAPPort)
		d.Set("ldap_host", authServer.LDAPHost)
	}
	d.Set("base_dn", authServer.BaseDN)
	log.Printf("[DEBUG] Set Base DN = %s", authServer.BaseDN)
	d.Set("security_level", authServer.SecurityLevel)
	log.Printf("[DEBUG] Set security level and made it through all of these")
}

func setAuthServerFromSchema(d *schema.ResourceData) *AuthServer {
	authServer := &AuthServer{
		ID:                  d.Get("server_id").(string),
		Enabled:             convertBoolToIntToString(d.Get("enabled").(bool)),
		ConnectionMethod:    d.Get("connection_method").(string),
		ADAccountSuffix:     d.Get("ad_account_suffix").(string),
		ADDomainControllers: d.Get("ad_domain_controllers").(string),
		BaseDN:              d.Get("base_dn").(string),
		SecurityLevel:       d.Get("security_level").(string),
		LDAPPort:            d.Get("ldap_port").(string),
		LDAPHost:            d.Get("ldap_host").(string),
	}

	return authServer
}
