package DNS

import (
	"github.com/cloudflare/cloudflare-go"
)

// API is a custom type that is used as a wrapper
// for the actual Cloudflare API instance
type API struct {
	ZoneID        string
	CloudflareAPI *cloudflare.API
}

// New creates a new instance of the type *API
// which takes a Config type as an argument and returns
// an *API and an error
func New(apiKey, email, zoneName string) (*API, error) {

	switch {
	case apiKey == "":
		return nil, InvalidParamError{Param: "apiKey"}
	case email == "":
		return nil, InvalidParamError{Param: "email"}
	case zoneName == "":
		return nil, InvalidParamError{Param: "zoneName"}
	}

	api, err := cloudflare.New(apiKey, email)
	if err != nil {
		return nil, err
	}

	zoneID, err := api.ZoneIDByName(zoneName)
	if err != nil {
		return nil, err
	}

	return &API{ZoneID: zoneID, CloudflareAPI: api}, nil
}

func (api *API) ZoneIDByName(zoneName string) (string, error) {
	zoneID, err := api.CloudflareAPI.ZoneIDByName(zoneName)
	if err != nil {
		return "", err
	}
	return zoneID, nil
}

// GetDNSRecords pulls all DNS Records associated with each
// element of the recordNames slice passed to this method as
// a parameter and returns a [][]cloudflare.DNSRecord and an error
func (api *API) GetDNSRecords(recordNames []string) ([]Record, error) {
	var records []Record
	for _, recordName := range recordNames {
		r, err := api.CloudflareAPI.DNSRecords(api.ZoneID, cloudflare.DNSRecord{Name: recordName})
		if err != nil {
			return nil, err
		}
		records = append(records, cloudflareToDNSRecords(r)...)
	}
	return records, nil
}

func (api *API) UpdateDNSRecords(ipAddress string, recordID string) error {
	err := api.CloudflareAPI.UpdateDNSRecord(api.ZoneID, recordID, cloudflare.DNSRecord{Content: string(ipAddress)})
	if err != nil {
		return err
	}
	return nil
}

func cloudflareToDNSRecords(cf []cloudflare.DNSRecord) []Record {
	res := make([]Record, len(cf))
	for _, v := range cf {
		res = append(res,
			Record{
				ID:      v.ID,
				Name:    v.Name,
				Content: v.Content,
				Type:    v.Type,
			},
		)
	}
	return res
}
