package main

import (
	"flag"
	"fmt"

	"github.com/darkraiden/odysseus/internal/cloudflare"
	"github.com/darkraiden/odysseus/internal/whatsmyip"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type flags struct {
	configName *string
	configPath *string
}

func init() {
	// Read flags
	var f flags
	f.configName = flag.String("config-name", "cloudflare.yml", "the name of the config file to be loaded")
	f.configPath = flag.String("config-path", ".", "the path to the config file")
	flag.Parse()

	// Initialize logrus
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	// Read configuration file
	viper.SetConfigName(*f.configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(*f.configPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	log.Info("Welcome to Odysseus")

	// Initialize Cloudflare API
	api, err := cloudflare.New(cloudflare.Config{APIKey: viper.Get("cloudflare.api_key").(string), Email: viper.Get("cloudflare.email").(string), ZoneName: viper.Get("cloudflare.zone_name").(string)})
	if err != nil {
		log.Panic(err)
	}

	// Get DNS Records
	records, err := api.GetDNSRecords(viper.Get("cloudflare.records").([]interface{}))
	if err != nil {
		log.Panic(err)
	}

	// Get Public IP Address
	ip, err := whatsmyip.GetLocalIP()
	if err != nil {
		log.Panic(err)
	}

	log.Info(fmt.Sprintf("Your local IP Address is: %s", *ip))

	log.Info(fmt.Sprintf("Your Zone ID is: %s", api.ZoneID))
	for _, r := range records {
		for _, inner := range r {
			if inner.Type == "A" && inner.Content != string(*ip) {
				err := api.UpdateDNSRecord(ip, inner.ID)
				if err != nil {
					log.Error(fmt.Sprintf("Error updating the DNS Record %s. Error: %v", inner.Name, err))
				} else {
					log.Info(fmt.Sprintf("The DNS Record %s has been updated successfully.", inner.Name))
				}
			} else {
				log.Info(fmt.Sprintf("No changes needed for the DNS record %s", inner.Name))
			}
		}
	}
}
