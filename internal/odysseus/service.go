package odysseus

import (
	"errors"
	"fmt"
	"github.com/darkraiden/odysseus/internal/DNS"
	"github.com/darkraiden/odysseus/internal/ipaddress"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	dnsManager DNS.Manager
	ipGetter   ipaddress.Getter
}

func NewService(manager DNS.Manager, ipGetter ipaddress.Getter) (*Service, error) {
	if manager == nil {
		return nil, DNS.InvalidParamError{Param: "manager"}
	}
	if ipGetter == nil {
		return nil, DNS.InvalidParamError{Param: "ipGetter"}
	}
	return &Service{dnsManager: manager, ipGetter: ipGetter}, nil

}

func (s Service) UpdateDNSWithLocalIP(records []string) error {
	if len(records) == 0 {
		return errors.New("no records to update")
	}

	ip, err := s.ipGetter.GetLocal()
	if err != nil {
		return err
	}
	r, err := s.dnsManager.GetDNSRecords(records)
	if err != nil {
		return err
	}
	for _, v := range r {
		if v.Type == "A" && v.Content != ip {
			err := s.dnsManager.UpdateDNSRecords(ip, v.ID)
			if err != nil {
				log.Error(fmt.Sprintf("Error updating the DNS Record %s. Error: %v", v.Name, err))
			} else {
				log.Info(fmt.Sprintf("The DNS Record %s has been updated successfully.", v.Name))
			}
		} else {
			log.Info(fmt.Sprintf("No changes needed for the DNS record %s", v.Name))
		}
	}
	return nil
}
