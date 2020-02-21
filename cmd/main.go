package main

import (
	"fmt"
	"os"

	"github.com/darkraiden/odysseus/internal/cloudflare"
	"github.com/darkraiden/odysseus/internal/logs"
	"github.com/darkraiden/odysseus/internal/template"
	"github.com/darkraiden/odysseus/internal/whatsmyip"
)

func main() {
	api, err := cloudflare.New(cloudflare.Config{APIKey: os.Getenv("CF_API_KEY"), Email: os.Getenv("CF_API_EMAIL"), ZoneName: os.Getenv("CF_ZONE_NAME")})
	if err != nil {
		panic(err)
	}

	records, err := api.GetDNSRecords([]string{"www.darkraiden.com"})
	if err != nil {
		panic(err)
	}

	// wire our template and logger together
	logger := logs.NewLogger()

	l := template.New(logger, "Info")

	l.Log("Welcome to Odysseus")

	ip, err := whatsmyip.GetLocalIp()
	if err != nil {
		panic(err)
	}

	l.Log(fmt.Sprintf("Your local IP Address is: %s", *ip))

	l.Log(fmt.Sprintf("Your Zone ID is: %s", api.ZoneID))
	for _, r := range records {
		for _, inner := range r {
			l.Log(fmt.Sprintf("The DNS Record Content is: %s", inner.Content))
		}
	}
}
