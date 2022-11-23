// Package silksdp is a Go library for the Silk Cloud Data Platform. It includes the ability to connect to any Silk API through the included
// Get, Post, Patch, and Delete functions as well as CRUD related opeartions against various objects on the Silk platform
package silksdp

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"reflect"
	"sort"
	"time"
)

// Type and Constants are used for escaping Get requests
type encoding int

const (
	encodePath encoding = 1 + iota
	encodePathSegment
	encodeQueryComponent
)

// Credentials contains the parameters used to authenticate against the Silk SDP server and can be consumed
// through ConnectEnv() or Connect().
type Credentials struct {
	Server   string
	Username string
	Password string
}

// Connect initializes a new API client based on manually provided Silk SDP server credentials. When possible,
// the Silk credentials should not be stored as plain text in your .go file. ConnectEnv() can be used
// as a safer alternative.
func Connect(server, username, password string) *Credentials {
	client := &Credentials{
		Server:   server,
		Username: username,
		Password: password,
	}

	return client
}

// ConnectEnv is the preferred method to initialize a new API client by attempting to read the
// following environment variables:
//
//	SILK_SDP_SERVER
//
//	SILK_SDP_USERNAME
//
//	SILK_SDP_PASSWORD
func ConnectEnv() (*Credentials, error) {

	server, ok := os.LookupEnv("SILK_SDP_SERVER")
	if ok != true {
		return nil, errors.New("The `SILK_SDP_SERVER` environment variable is not present")
	}
	username, ok := os.LookupEnv("SILK_SDP_USERNAME")
	if ok != true {
		return nil, errors.New("The `SILK_SDP_USERNAME` environment variable is not present")
	}
	password, ok := os.LookupEnv("SILK_SDP_PASSWORD")
	if ok != true {
		return nil, errors.New("The `SILK_SDP_PASSWORD` environment variable is not present")
	}

	client := &Credentials{
		Server:   server,
		Username: username,
		Password: password,
	}

	return client, nil
}

// makeHTTPCall consolidates the functionality for the GET, POST, PATCH, and DELETE functions.
func (c *Credentials) makeHTTPCall(callType, apiEndpoint string, config interface{}, timeout int) (interface{}, error) {

	if endpointValidation(apiEndpoint) == "errorStart" {
		return nil, errors.New("The API Endpoint should begin with '/' (ex: /cluster/me)")
	} else if endpointValidation(apiEndpoint) == "errorEnd" {
		return nil, errors.New("The API Endpoint should not end with '/' (ex. /cluster/me)")
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * time.Duration(timeout),
	}

	apiVersion := "v2"

	requestURL := fmt.Sprintf("https://%s/api/%s%s", c.Server, apiVersion, apiEndpoint)

	var request *http.Request
	switch callType {
	case "GET":
		request, _ = http.NewRequest(callType, getEscape(requestURL), nil)
	case "POST":
		convertedConfig, _ := json.Marshal(config)
		request, _ = http.NewRequest(callType, requestURL, bytes.NewBuffer(convertedConfig))
	case "PATCH":
		convertedConfig, _ := json.Marshal(config)
		request, _ = http.NewRequest(callType, requestURL, bytes.NewBuffer(convertedConfig))
	case "DELETE":
		request, _ = http.NewRequest(callType, requestURL, nil)
	}

	request.SetBasicAuth(c.Username, c.Password)

	request.Header.Set("Content-Type", "application/json")

	apiRequest, err := client.Do(request)
	if err, ok := err.(net.Error); ok && err.Timeout() {
		return nil, errors.New("Unable to establish a connection to the Silk SDP server")
	} else if err != nil {
		return nil, err
	}

	// Place a 1 second pause here - Post request but prior to returning the response.
	duration := time.Second // Pause for 1 second.
	time.Sleep(duration)

	body, err := ioutil.ReadAll(apiRequest.Body)

	apiResponse := []byte(body)

	var convertedAPIResponse interface{}
	if err := json.Unmarshal(apiResponse, &convertedAPIResponse); err != nil {

		// DELETE request will return a 204 No Content status
		if apiRequest.StatusCode == 204 {
			convertedAPIResponse = map[string]interface{}{}
			convertedAPIResponse.(map[string]interface{})["statusCode"] = apiRequest.StatusCode
		} else if apiRequest.StatusCode != 200 {
			return nil, fmt.Errorf("%s", apiRequest.Status)
		}

	}

	if reflect.TypeOf(convertedAPIResponse).Kind() == reflect.Slice {
		return convertedAPIResponse, nil
	}

	if _, ok := convertedAPIResponse.(map[string]interface{})["error_msg"]; ok {
		return nil, fmt.Errorf("%s", convertedAPIResponse.(map[string]interface{})["error_msg"])
	}

	return convertedAPIResponse, nil

}

// endpointValidation validates that the endpoint provided in the Base API functions starts with a / but does not end with one except if preceded by a =
func endpointValidation(apiEndpoint string) string {

	if string(apiEndpoint[0]) != "/" {
		return "errorStart"
	} else if string(apiEndpoint[len(apiEndpoint)-1]) == "/" {

		if string(apiEndpoint[len(apiEndpoint)-2]) != "=" { // accounting for exeption =/
			return "errorEnd"
		}
	}
	return "success"
}

// httpTimeout returns a default timeout value of 15 or use the value provided by the end user
func httpTimeout(timeout []int) int {
	if len(timeout) == 0 {
		return int(15) // if not timeout value is provided, set the default to 15
	}
	return int(timeout[0]) // set the timeout value to the first value in the timeout slice

}

// getEscape is a custom implementation of url.PathEscape which escapes strings for GET requests so escapes they
// can be safely placed inside a URL path segment, replacing special characters (including /) with %XX sequences as needed
func getEscape(s string) string {
	return escape(s, encodePathSegment)
}

// Called by the getEscape function
func escape(s string, mode encoding) string {
	spaceCount, hexCount := 0, 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if shouldEscape(c, mode) {
			if c == ' ' && mode == encodeQueryComponent {
				spaceCount++
			} else {
				hexCount++
			}
		}
	}

	if spaceCount == 0 && hexCount == 0 {
		return s
	}

	t := make([]byte, len(s)+2*hexCount)
	j := 0
	for i := 0; i < len(s); i++ {
		switch c := s[i]; {
		case c == ' ' && mode == encodeQueryComponent:
			t[j] = '+'
			j++
		case shouldEscape(c, mode):
			t[j] = '%'
			t[j+1] = "0123456789ABCDEF"[c>>4]
			t[j+2] = "0123456789ABCDEF"[c&15]
			j += 3
		default:
			t[j] = s[i]
			j++
		}
	}
	return string(t)
}

// Called by the escape function
func shouldEscape(c byte, mode encoding) bool {
	// §2.3 Unreserved characters (alphanum)
	if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' {
		return false
	}

	switch c {
	case '-', '_', '.', '~': // §2.3 Unreserved characters (mark)
		return false

	case '$', '&', '+', ',', '/', ':', ';', '=', '?', '@': // §2.2 Reserved characters (reserved)
		// Different sections of the URL allow a few of
		// the reserved characters to appear unescaped.

		return c == ';' || c == ','

	}

	// Everything else must be escaped.
	return true
}

// Get sends a GET request to the provided Silk SDP API endpoint and returns the full API response.
// The optional timeout value corresponds to the number of seconds to wait to establish a connection to the Silk SDP server before returning a
// timeout error. If no value is provided, a default of 15 seconds will be used.
func (c *Credentials) Get(apiEndpoint string, timeout ...int) (interface{}, error) {

	httpTimeout := httpTimeout(timeout)

	apiRequest, err := c.makeHTTPCall("GET", apiEndpoint, nil, httpTimeout)
	if err != nil {
		return nil, err
	}

	return apiRequest, nil

}

// Post sends a POST request to the provided Silk SDP API endpoint and returns the full API response.
// The optional timeout value corresponds to the number of seconds to wait to establish a connection to the Silk SDP server before returning a
// timeout error. If no value is provided, a default of 15 seconds will be used.
func (c *Credentials) Post(apiEndpoint string, config map[string]interface{}, timeout ...int) (interface{}, error) {

	httpTimeout := httpTimeout(timeout)

	apiRequest, err := c.makeHTTPCall("POST", apiEndpoint, config, httpTimeout)
	if err != nil {
		return nil, err
	}

	return apiRequest, nil

}

// Patch sends a PATCH request to the provided Silk SDP API endpoint and returns the full API response.
// The optional timeout value corresponds to the number of seconds to wait to establish a connection to the Silk SDP server before returning a
// timeout error. If no value is provided, a default of 15 seconds will be used.
func (c *Credentials) Patch(apiEndpoint string, config interface{}, timeout ...int) (interface{}, error) {

	httpTimeout := httpTimeout(timeout)

	apiRequest, err := c.makeHTTPCall("PATCH", apiEndpoint, config, httpTimeout)
	if err != nil {
		return nil, err
	}

	return apiRequest, nil

}

// Delete sends a DELETE request to the provided Silk SDP API endpoint and returns the full API response.
// The optional timeout value corresponds to the number of seconds to wait to establish a connection to the Silk SDP server before returning a
// timeout error. If no value is provided, a default of 15 seconds will be used.
func (c *Credentials) Delete(apiEndpoint string, timeout ...int) (interface{}, error) {

	httpTimeout := httpTimeout(timeout)

	apiRequest, err := c.makeHTTPCall("DELETE", apiEndpoint, nil, httpTimeout)
	if err != nil {
		return nil, err
	}

	return apiRequest, nil
}

// stringEq converts b to []string, sorts the two []string, and checks for equality
func stringEq(a []string, b []interface{}) bool {

	// Convert []interface {} to []string
	c := make([]string, len(b))
	for i, v := range b {
		c[i] = fmt.Sprint(v)
	}

	sort.Strings(a)
	sort.Strings(c)

	// If one is nil, the other must also be nil.
	if (a == nil) != (c == nil) {
		return false
	}

	if len(a) != len(c) {
		return false
	}

	for i := range a {
		if a[i] != c[i] {
			return false
		}
	}

	return true
}

// stringInSlice checks whether the e string is in the s slice
func (c *Credentials) stringInSlice(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (c *Credentials) popStringFromSlice(s []string, e string) []string {
	for i, a := range s {
		if a == e {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
