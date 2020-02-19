package main

import (
	"fmt"
	"os"

	"github.com/cloudflare/cloudflare-go"
	"github.com/darkraiden/odysseus/internal/logs"
	"github.com/darkraiden/odysseus/internal/template"
	"github.com/darkraiden/odysseus/internal/whatsmyip"
)

func main() {
	api, err := cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
	if err != nil {
		panic(err)
	}

	zoneID, err := api.ZoneIDByName("darkraiden.com")
	if err != nil {
		panic(err)
	}

	record := cloudflare.DNSRecord{Name: "www.darkraiden.com"}
	recs, err := api.DNSRecords(zoneID, record)
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

	l.Log(fmt.Sprintf("Your Zone ID is: %s", zoneID))
	l.Log(fmt.Sprintf("The DNS Record Content is: %s", recs[0].Content))
}
