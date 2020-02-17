package cloudflare

import "github.com/cloudflare/cloudflare-go"

type Config struct {
	apiKey   string
	email    string
	zoneName string
}

type API struct {
	zoneID        string
	cloudflareAPI *cloudflare.API
}
