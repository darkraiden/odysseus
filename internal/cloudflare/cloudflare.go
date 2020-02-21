package cloudflare

import (
	"github.com/cloudflare/cloudflare-go"
)

func New(conf Config) (*API, error) {
	api, err := cloudflare.New(conf.APIKey, conf.Email)
	if err != nil {
		return nil, err
	}

	zoneID, err := api.ZoneIDByName(conf.zoneName)
	if err != nil {
		return nil, err
	}

	return &API{zoneID: zoneID}, nil
}

func (api *API) GetDNSRecords(recordNames []string) ([][]cloudflare.DNSRecord, error) {
	var records [][]cloudflare.DNSRecord
	for i, recordName := range recordNames {
		r, err := api.cloudflareAPI.DNSRecords(api.zoneID, cloudflare.DNSRecord{Name: recordName})
		if err != nil {
			return nil, err
		}
		records[i] = r
	}
	return records, nil
}
