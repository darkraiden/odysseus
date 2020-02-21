package cloudflare

import (
	"github.com/cloudflare/cloudflare-go"
)

// New creates a new instance of the type *API
// which takes a Config type as an argument and returns
// an *API and an error
func New(conf Config) (*API, error) {
	api, err := cloudflare.New(conf.APIKey, conf.Email)
	if err != nil {
		return nil, err
	}

	zoneID, err := api.ZoneIDByName(conf.ZoneName)
	if err != nil {
		return nil, err
	}

	return &API{ZoneID: zoneID, CloudflareAPI: api}, nil
}

// GetDNSRecords pulls all DNS Records associated with each
// element of the recordNames slice passed to this method as
// a parameter and returns a [][]cloudflare.DNSRecord and an error
func (api *API) GetDNSRecords(recordNames []string) ([][]cloudflare.DNSRecord, error) {
	var records [][]cloudflare.DNSRecord
	for _, recordName := range recordNames {
		r, err := api.CloudflareAPI.DNSRecords(api.ZoneID, cloudflare.DNSRecord{Name: recordName})
		if err != nil {
			return nil, err
		}
		records = append(records, r)
	}
	return records, nil
}
