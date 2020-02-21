package cloudflare

import (
	"github.com/cloudflare/cloudflare-go"
)

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
