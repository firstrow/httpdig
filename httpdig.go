package httpdig

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var apiURL = "https://dns.google.com/resolve"
var ednsSubnet = "0.0.0.0/0"

type Response struct {
	Status int  `json:"Status"`
	TC     bool `json:"TC"`
	RD     bool `json:"RD"`
	RA     bool `json:"RA"`
	AD     bool `json:"AD"`
	CD     bool `json:"CD"`

	Question []struct {
		Name string `json:"name"`
		Type int    `json:"type"`
	} `json:"Question"`

	Answer []struct {
		Name string `json:"name"`
		Type int    `json:"type"`
		TTL  int    `json:"TTL"`
		Data string `json:"data"`
	} `json:"Answer"`

	Authority []struct {
		Name string `json:"name"`
		Type int    `json:"type"`
		TTL  int    `json:"TTL"`
		Data string `json:"data"`
	} `json:"Authority"`

	Additional       []interface{} `json:"Additional"`
	EdnsClientSubnet string        `json:"edns_client_subnet"`
	Comment          string        `json:"Comment"`
}

func dig(host, recordType string) ([]byte, error) {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", apiURL, nil)

	query := req.URL.Query()
	query.Add("name", host)
	query.Add("type", recordType)
	query.Add("edns_client_subnet", ednsSubnet)

	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("Unable to resolve host")
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return body, nil
}

// Query sends request to Google dns service and parses response.
// e.g: httpdig.Query("google.com", "NS")
func Query(host string, t string) (Response, error) {
	resp, err := dig(host, t)
	if err != nil {
		return Response{}, err
	}

	response := Response{}
	err = json.Unmarshal(resp, &response)

	if err != nil {
		return Response{}, err
	}

	return response, nil
}
