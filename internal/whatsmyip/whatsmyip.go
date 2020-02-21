package whatsmyip

import (
	"io/ioutil"
	"net/http"
)

var url = "https://api.ipify.org?format=text"

// LocalIP is a custom type that will be used
// to store the result of the HTTP request to the
// IP URL
type LocalIP []byte

// GetLocalIP makes an HTTP GET request to
// an URL defined with the url variable
// and returns to the user the Local Public IP
// under the form of *LocalIP and an error
func GetLocalIP() (*LocalIP, error) {
	var ip LocalIP

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ip, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &ip, nil
}
