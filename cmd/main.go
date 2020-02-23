package main

import (
	"fmt"

	"github.com/darkraiden/odysseus/internal/cloudflare"
	"github.com/darkraiden/odysseus/internal/logs"
	"github.com/darkraiden/odysseus/internal/template"
	"github.com/darkraiden/odysseus/internal/whatsmyip"
	"github.com/spf13/viper"
)

func init() {
	// Read configuration file
	viper.SetConfigName("cloudflare")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	// wire our template and logger together
	logger := logs.NewLogger()
	l := template.New(logger, "Info")

	l.Log("Welcome to Odysseus")

	// Initialize Cloudflare API
	api, err := cloudflare.New(cloudflare.Config{APIKey: viper.Get("cloudflare.api_key").(string), Email: viper.Get("cloudflare.email").(string), ZoneName: viper.Get("cloudflare.zone_name").(string)})
	if err != nil {
		panic(err)
	}

	// Get DNS Records
	records, err := api.GetDNSRecords(viper.Get("cloudflare.records").([]interface{}))
	if err != nil {
		panic(err)
	}

	// Get Public IP Address
	ip, err := whatsmyip.GetLocalIP()
	if err != nil {
		panic(err)
	}

	l.Log(fmt.Sprintf("Your local IP Address is: %s", *ip))

	l.Log(fmt.Sprintf("Your Zone ID is: %s", api.ZoneID))
	for _, r := range records {
		for _, inner := range r {
			if inner.Type == "A" && inner.Content != string(*ip) {
				err := api.UpdateDNSRecord(ip, inner.ID)
				if err != nil {
					l.Log(fmt.Sprintf("Error updating the DNS Record %s. Error: %v", inner.Name, err))
				} else {
					l.Log(fmt.Sprintf("The DNS Record %s has been updated successfully.", inner.Name))
				}
			} else {
				l.Log(fmt.Sprintf("No changes needed for the DNS record %s", inner.Name))
			}
		}
	}
}
