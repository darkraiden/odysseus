package whatsmyip

import (
	"io/ioutil"
	"net/http"
)

var url = "https://api.ipify.org?format=text"

type LocalIP []byte

func GetLocalIp() (*LocalIP, error) {
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
