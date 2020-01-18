package nagios

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Client used to store info required to communicate with Nagios
type Client struct {
	url        string
	token      string
	httpClient *http.Client
}

// NewClient creates a pointer to the client that will be used to send requests to Nagios
func NewClient(url, token string) *Client {
	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}

	nagiosClient := &Client{
		url:        url,
		token:      token,
		httpClient: httpClient,
	}

	return nagiosClient
}

func (c *Client) sendRequest(httpRequest *http.Request) ([]byte, error) {
	c.addRequestHeaders(httpRequest)

	response, err := c.httpClient.Do(httpRequest)

	// TODO: Need to validate that when Nagios is unavailable, this err check will catch it
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

// buildURL generates the appropriate URL to interact with the Nagios XI API
func (c *Client) buildURL(objectType, method, objectName, name, oldVal, objectDescription string) (string, error) {
	// TODO: This func has really become a mess...but it works. Plan is to revisit after building functionality
	// out for other objects in Nagios.
	var nagiosURL strings.Builder

	var apiURL string
	var apiType string
	if objectType == "applyconfig" {
		apiType = "system"

		if method != "POST" {
			return "", errors.New("You must use a HTTP POST when performing an applyconfig")
		}
	} else if objectType == "authserver" {
		apiType = "system"
	} else {
		apiType = "config"
	}

	apiURL = "api/v1/" + apiType + "/"

	if !strings.HasSuffix(c.url, "/") {
		apiURL = "/" + apiURL
	}

	// All of this creates the nagiosURL to get the object
	// that has been supplied through parameters
	nagiosURL.WriteString(c.url)
	nagiosURL.WriteString(apiURL)
	nagiosURL.WriteString(objectType)

	// If we are doing a GET, PUT or DELETE, we need to provide the name of the object
	// and type to filter results to only that. Otherwise, Nagios
	// will return all results for that particular object type
	// TODO: This is getting messy. Need to figure out a more streamlined way to handle all of this
	if method == "GET" {
		nagiosURL.WriteString("?apikey=")
		nagiosURL.WriteString(c.token)
		nagiosURL.WriteString("&")
		nagiosURL.WriteString(objectName)
		nagiosURL.WriteString("=")

		if name == "" {
			errMsg := "Name must be provided when using the " + method + " method"
			return "", errors.New(errMsg)
		} else {
			nagiosURL.WriteString(name)
		}

		nagiosURL.WriteString("&pretty=1")
	} else if method == "DELETE" {
		if objectType == "authserver" {
			nagiosURL.WriteString("/" + name)
		}
		nagiosURL.WriteString("?apikey=")
		nagiosURL.WriteString(c.token)
		nagiosURL.WriteString("&")
		nagiosURL.WriteString(objectName)
		nagiosURL.WriteString("=")

		if name == "" {
			errMsg := "Name must be provided when using the " + method + " method"
			return "", errors.New(errMsg)
		} else {
			nagiosURL.WriteString(name)
		}
	} else if method == "PUT" {
		nagiosURL.WriteString("/")

		if oldVal != "" {
			nagiosURL.WriteString(oldVal)
		} else {
			return "", errors.New("[ERROR] A value for oldVal must be provided when attempting a PUT")
		}

		if objectType == "service" {
			nagiosURL.WriteString("/" + objectDescription)
		}

		nagiosURL.WriteString("?apikey=")
		nagiosURL.WriteString(c.token)
		nagiosURL.WriteString("&pretty=1&force=1&")
	} else if method == "POST" {
		nagiosURL.WriteString("?apikey=")
		nagiosURL.WriteString(c.token)

		if objectType != "applyconfig" {
			nagiosURL.WriteString("&force=1&")
		}
	}

	return nagiosURL.String(), nil
}

func (c *Client) scrubToken(url string) string {
	if strings.Contains(url, c.token) {
		strings.Replace(url, c.token, "<SensitiveInfo>", 1)
	}

	return url
}

func (c *Client) addRequestHeaders(request *http.Request) {
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Accept", "/")

	return
}

func (c *Client) get(requestData, nagiosURL string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, nagiosURL, strings.NewReader(requestData))

	if err != nil {
		return nil, err
	}

	body, err := c.sendRequest(request)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) post(data *url.Values, nagiosURL string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodPost, nagiosURL, strings.NewReader(data.Encode()))

	if err != nil {
		return nil, err
	}

	body, err := c.sendRequest(request)

	if err != nil {
		return nil, err
	}

	err = c.commandResponse(body)

	if err != nil {
		return nil, err
	}

	err = c.commandResponse(body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) put(nagiosURL string) ([]byte, error) {
	if strings.Contains(nagiosURL, " ") {
		nagiosURL = strings.Replace(nagiosURL, " ", "%20", -1)
	}
	request, err := http.NewRequest(http.MethodPut, nagiosURL, nil)

	if err != nil {
		return nil, err
	}

	body, err := c.sendRequest(request)

	if err != nil {
		return nil, err
	}

	err = c.commandResponse(body)

	if err != nil {
		return nil, err
	}

	err = c.commandResponse(body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) delete(data *url.Values, nagiosURL string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodDelete, nagiosURL, strings.NewReader(data.Encode()))

	if err != nil {
		return nil, err
	}

	body, err := c.sendRequest(request)

	if err != nil {
		return nil, err
	}

	err = c.commandResponse(body)

	if err != nil {
		return nil, err
	}

	err = c.commandResponse(body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) applyConfig() error {
	nagiosURL, err := c.buildURL("applyconfig", "POST", "", "", "", "")

	if err != nil {
		return err
	}

	data := &url.Values{}

	_, err = c.post(data, nagiosURL)

	if err != nil {
		return err
	}

	return nil
}

// Function maps the elements of a string array to a single string with each value separated by commas
// Nagios expects a list of values supplied in this format via URL encoding
func mapArrayToString(sourceArray []interface{}) string {
	var destString strings.Builder

	for i, sourceObject := range sourceArray {
		// If this is the first time looping through, set the destination object euqal to the first element in array
		if i == 0 {
			destString.WriteString(sourceObject.(string))
		} else { // More than one element in array. Append a comma first before we add the next item
			destString.WriteString(",")
			destString.WriteString(sourceObject.(string))
		}
	}

	return destString.String()
}

// Function takes any boolean value, converts to integer and returns in string format
func convertBoolToIntToString(sourceVal bool) string {
	if sourceVal {
		return "1"
	}
	return "0"
}

// setURLParams loops through a struct object and returns a set of URL parameters
func setURLParams(nagiosObject interface{}) *url.Values {
	values := reflect.ValueOf(nagiosObject)
	var urlParams = &url.Values{}
	var tag string

	// If we are passing in a pointer to a struct, we need to get the actual result of what the struct is pointing to
	if values.Kind() == reflect.Ptr {
		values = values.Elem()
	}

	for i := 0; i < values.NumField(); i++ {
		var outputString strings.Builder
		curType := values.Field(i).Type().String()
		tags := strings.Split(values.Type().Field(i).Tag.Get("json"), ",")

		for k, _ := range tags {
			if tags[k] != "omitempty" {
				tag = tags[k]
				break
			}
		}

		if curType == "string" {
			if values.Field(i).Interface().(string) != "" {
				urlParams.Add(tag, values.Field(i).Interface().(string))
			}
		} else if curType == "[]interface {}" {
			if values.Field(i).Interface() != nil {
				for j, val := range values.Field(i).Interface().([]interface{}) {
					if j > 0 {
						outputString.WriteString(",")
					}

					outputString.WriteString(val.(string))
				}
				urlParams.Add(tag, outputString.String())
			}
		} else if curType == "int" {
			if strconv.Itoa(values.Field(i).Interface().(int)) != "" {
				// We need the value to be a string but first need to cast it as an integer if that is what the type is in the struct
				urlParams.Add(tag, strconv.Itoa(values.Field(i).Interface().(int)))
			}
		} else if curType == "map[string]interface {}" {
			if values.Field(i).Interface() != nil {
				// We need to loop through the map and grab the key and value for each line
				// The value is an interface, so we need to then call the Interface() method
				// and cast it as a string to get the value in string format
				mapObject := values.Field(i).MapRange()
				for mapObject.Next() {
					outputString.Reset()
					index := mapObject.Key().String()
					val := mapObject.Value()
					valString := val.Interface().(string)
					urlParams.Add(index, valString)
				}
			}
		}
	}
	return urlParams
}
