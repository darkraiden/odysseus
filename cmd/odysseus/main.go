package main

import (
	"flag"
	"github.com/darkraiden/odysseus/internal/DNS"
	"github.com/darkraiden/odysseus/internal/odysseus"
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
		log.Fatalf("error loading viper config: %s", err.Error())
	}
}

func main() {
	log.Info("Welcome to Odysseus")

	a, err := DNS.New(viper.Get(
		"cloudflare.api_key").(string),
		viper.Get("cloudflare.email").(string),
		viper.Get("cloudflare.zone_name").(string),
	)

	if err != nil {
		log.Fatalf("fucked it: %s", err.Error())

	}

	s, err := odysseus.NewService(a)
	if err != nil {
		log.Fatalf("fucked it: %s", err.Error())
	}

	err = s.UpdateDNSWithLocalIP(viper.Get("cloudflare.records").([]string))
	if err != nil {
		log.Fatalf("fucked it: %s", err.Error())
	}

	log.Info("wahoo! did it!")
}
