package ipaddress

import (
	"errors"
	"io/ioutil"
	"net/http"
)

const getIPURL = "https://api.ipify.org?format=text"

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

type Ipify struct {
	httpClient Doer
}

func NewService(doer Doer) (*Ipify, error) {
	if doer == nil {
		return nil, errors.New("please pass a valid http client")
	}
	return &Ipify{httpClient: doer}, nil
}

// LocalIP is a custom type that will be used
// to store the result of the HTTP request to the
// IP URL

// GetLocalIP makes an HTTP GET request to
// an URL defined with the url variable
// and returns to the user the Local Public IP
// under the form of *LocalIP and an error
func (s Ipify) GetLocal() (string, error) {

	req, err := http.NewRequest("GET", getIPURL, nil)
	if err != nil {
		return "", err
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}
