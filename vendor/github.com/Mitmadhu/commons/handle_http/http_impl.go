package handle_http

import (
	"bytes"
	"encoding/json"
	"errors"
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

	_, exists := ValidHTTPMethods[method]
	if !exists{
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
	if(err != nil){
		return nil, errors.New("error while creating request")
	}

	// sending the request
	resp, err := httpClient.Do(request)
	if (err != nil){
		return nil, errors.New("error while sending the request")
	}
	return ioutil.ReadAll(resp.Body)
}
