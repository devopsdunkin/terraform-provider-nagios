package nagios

import (
	"encoding/json"
	"log"
	"net/url"
	"strings"
)

func (c *Client) newAuthServer(authServer *AuthServer) ([]byte, error) {
	nagiosURL, err := c.buildURL("authserver", "POST", "", "", "", "")

	if err != nil {
		return nil, err
	}

	data := createAuthServerHTTPBody(authServer)

	body, err := c.post(data, nagiosURL)

	log.Printf("[DEBUG] Value of body (newAuthServer) %s", body)

	if err != nil {
		return nil, err
	}

	err = c.applyConfig()

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) getAuthServer(ID string) (*AuthServer, error) {
	// var authServerArray = []AuthServer{}
	var authServer AuthServer
	var mapAuthServer MapOfAuthServers

	log.Printf("[DEBUG] Value of ID during getAuthServer: %s", ID)

	nagiosURL, err := c.buildURL("authserver", "GET", "server_id", ID, "", "")

	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] NagiosURL = %s", nagiosURL)
	// log.Printf("[DEBUG] mapAuthServer = %s", mapAuthServer)

	data := &url.Values{}
	data.Set("server_id", ID)

	log.Printf("[DEBUG] data: %s", data.Encode())

	body, err := c.get(data.Encode(), nagiosURL)

	if err != nil {
		log.Printf("[DEBUG] Error during get(). Returning error")
		return nil, err
	}

	log.Printf("[DEBUG] body: %s", body)

	log.Printf("[DEBUG] Right before Unmarshal")

	err = json.Unmarshal(body, &mapAuthServer)
	// err = json.Unmarshal(body, &authServer)

	if err != nil {
		return nil, err
	}

	if mapAuthServer.Records > 0 {
		authServer = mapAuthServer.AuthServerEntry[0]
		authServer.ServerID = authServer.ID
	}

	// if i > 1 {
	// break
	// }
	// }
	log.Printf("[DEBUG] Made it through getAuthServer(). Returning authServer object")
	log.Printf("[DEBUG] authServer value inside getAuthServer = %s", authServer)

	return &authServer, nil
}

func (c *Client) updateAuthServer(authServer *AuthServer, oldVal interface{}) error {
	nagiosURL, err := c.buildURL("authserver", "PUT", "server_id", authServer.ID, oldVal.(string), "")

	if err != nil {
		return err
	}

	nagiosURL = setUpdateURLAuthServerParams(nagiosURL, authServer)

	_, err = c.put(nagiosURL)

	if err != nil {
		// If the error is this specific message, we want to "catch" it
		// and create a new host, then we can proceed on. Otherwise, we
		// can return the error and exit
		if strings.Contains(err.Error(), "Does the authentication server exist?") {
			c.newAuthServer(authServer)
		} else {
			return err
		}
	}

	err = c.applyConfig()

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) deleteAuthServer(name string) ([]byte, error) {
	nagiosURL, err := c.buildURL("authserver", "DELETE", "server_id", name, "", "")

	if err != nil {
		return nil, err
	}

	data := &url.Values{}
	data.Set("contact_name", name)

	body, err := c.delete(data, nagiosURL)

	if err != nil {
		return nil, err
	}

	err = c.applyConfig()

	if err != nil {
		return nil, err
	}

	return body, nil
}

func createAuthServerHTTPBody(authServer *AuthServer) *url.Values {
	data := &url.Values{}

	data.Set("enabled", authServer.Enabled)
	data.Set("conn_method", authServer.ConnectionMethod)

	// Optional attributes
	if authServer.ADAccountSuffix != "" {
		data.Set("ad_account_suffix", authServer.ADAccountSuffix)
	}

	if authServer.ADDomainControllers != "" {
		data.Set("ad_domain_controllers", authServer.ADDomainControllers)
	}

	if authServer.BaseDN != "" {
		data.Set("base_dn", authServer.BaseDN)
	}

	if authServer.SecurityLevel != "" {
		data.Set("security_level", authServer.SecurityLevel)
	}

	if authServer.LDAPPort != "" {
		data.Set("ldap_port", authServer.LDAPPort)
	}

	if authServer.LDAPHost != "" {
		data.Set("ldap_host", authServer.LDAPHost)
	}

	return data
}

func setUpdateURLAuthServerParams(originalURL string, authServer *AuthServer) string {
	var urlParams = &url.Values{}

	urlParams.Add("server_id", authServer.ID)
	urlParams.Add("enabled", authServer.Enabled)
	urlParams.Add("conn_method", authServer.ConnectionMethod)

	if authServer.ADAccountSuffix != "" {
		urlParams.Add("ad_account_suffix", authServer.ADAccountSuffix)
	}

	if authServer.ADDomainControllers != "" {
		urlParams.Add("ad_domain_controllers", authServer.ADDomainControllers)
	}

	if authServer.BaseDN != "" {
		urlParams.Add("base_dn", authServer.BaseDN)
	}

	if authServer.SecurityLevel != "" {
		urlParams.Add("security_level", authServer.SecurityLevel)
	}

	if authServer.LDAPPort != "" {
		urlParams.Add("ldap_port", authServer.LDAPPort)
	}

	if authServer.LDAPHost != "" {
		urlParams.Add("ldap_host", authServer.LDAPHost)
	}

	return urlParams.Encode()
}
