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

	api, err := cloudflare.New(cloudflare.Config{APIKey: viper.Get("cloudflare.api_key").(string), Email: viper.Get("cloudflare.email").(string), ZoneName: viper.Get("cloudflare.zone_name").(string)})
	if err != nil {
		panic(err)
	}

	records, err := api.GetDNSRecords(viper.Get("cloudflare.records").([]interface{}))
	if err != nil {
		panic(err)
	}

	ip, err := whatsmyip.GetLocalIP()
	if err != nil {
		panic(err)
	}

	l.Log(fmt.Sprintf("Your local IP Address is: %s", *ip))

	l.Log(fmt.Sprintf("Your Zone ID is: %s", api.ZoneID))
	for _, r := range records {
		for _, inner := range r {
			l.Log(fmt.Sprintf("The DNS Record Content of '%s' is: %s", inner.Name, inner.Content))
		}
	}
}
