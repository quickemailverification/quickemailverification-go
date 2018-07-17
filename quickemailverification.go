package quickemailverification

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"errors"
)

// BaseURL is the default host to make the API requests on
const BaseURL = "http://api.quickemailverification.com/v1/verify"

// Client struct will hold the HTTP client and API key to communicate with QuickEmailVerification API
type Client struct {
	apiKey string
	http   *http.Client
}

// This accepts the api key to use for authentication. CreateClient returns a new instance of QuickEmailVerification.
func CreateClient(apiKey string) *Client {
	client := &Client{
		apiKey: apiKey,
		http:   &http.Client{},
	}

	return client
}

// Verify the given email address by using QuickEmailVerification REST API
func (c Client) Verify(address string) (*Response, error) {
	return c.getAPIResponse(BuildResponseModel{}, fmt.Sprintf(BaseURL+"?email=%s&apikey=%s", url.QueryEscape(address), url.QueryEscape(c.apiKey)))
}

// Return predefined response for predefined email address using QuickEmailVerification Sandbox mode
func (c Client) Sandbox(address string) (*Response, error) {
	return c.getAPIResponse(BuildResponseModel{}, fmt.Sprintf(BaseURL+"/sandbox?email=%s&apikey=%s", url.QueryEscape(address), url.QueryEscape(c.apiKey)))
}

// Send the API request and read the response of the API request, returning a new Response struct
func (c Client) getAPIResponse(rm ResponseModel, url string) (*Response, error) {
	// Send API request
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

    	arrStatusCode := []int{200, 400, 401, 402, 403, 404, 429}
	if(in_array(response.StatusCode, arrStatusCode) ==-1) {
      	 msg := fmt.Sprintf("API returned HTTP Status Code : %d", response.StatusCode)
		    return nil, errors.New(msg)
    	}

	// Read the response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Build Response struct
	result, err := rm.NewResponseModel(body)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func in_array(element int, data []int) (int) {
   for k, v := range data {
       if element == v {
           return k
       }
   }
   return -1
}
