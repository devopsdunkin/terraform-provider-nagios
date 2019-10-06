package nagios

import (
	"encoding/json"
	"errors"
)

type ResponseStatus struct {
	StatusSuccess string `json:"success"`
	StatusError   string `json:"error"`
}

/*
	Nagios does error handling with their API in a weird way.
	We won't typically catch any errors with returning any error
	codes from a HTTP request call. We may get some errors, such as
	if it cannot connect to Nagios.

	For a POST or PUT, it will return either "success" or "error"
	but the HTTP status will always be 200. So we have to parse the
	JSON object and determine which was returned. If "error" was returned,
	we will relay that back to the calling func as an error. If no error,
	then we do not return anything and we can assume the call was good
*/

func (c *Client) commandResponse(body []byte) error {
	responseStatus := &ResponseStatus{}

	err := json.Unmarshal(body, &responseStatus)

	if err != nil {
		return err
	}

	if responseStatus.StatusSuccess != "" {
		return nil
	} else {
		return errors.New(responseStatus.StatusError)
	}
}
