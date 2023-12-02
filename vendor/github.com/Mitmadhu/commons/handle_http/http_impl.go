package handle_http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var ValidHTTPMethods = map[string]bool{
	http.MethodGet:     true,
	http.MethodPost:    true,
	http.MethodPut:     true,
	http.MethodDelete:  true,
	http.MethodPatch:   true,
	http.MethodHead:    true,
	http.MethodOptions: true,
}

func Call(url string, method string, body interface{}) ([]byte, error) {

	println(url)
	_, exists := ValidHTTPMethods[method]
	if !exists {
		return nil, errors.New("invalid method")
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("error during marshalling")
	}

	// create http client
	httpClient := http.Client{}
	// Creating request
	request, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error while creating request, err: %s", err.Error())
	}

	// sending the request
	resp, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}
