package cloudflare

import "github.com/cloudflare/cloudflare-go"

type Config struct {
	APIKey   string
	Email    string
	ZoneName string
}

type API struct {
	ZoneID        string
	CloudflareAPI *cloudflare.API
}
