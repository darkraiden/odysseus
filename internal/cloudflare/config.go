package cloudflare

import "github.com/cloudflare/cloudflare-go"

// Config holds all the basic config parameters
// used to establish a connection with Cloudflare
type Config struct {
	APIKey   string
	Email    string
	ZoneName string
}

// API is a custom type that is used as a wrapper
// for the actual Cloudflare API instance
type API struct {
	ZoneID        string
	CloudflareAPI *cloudflare.API
}
