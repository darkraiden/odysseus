package DNS

import (
	"fmt"
)

type InvalidParamError struct {
	Param string
}

func (e InvalidParamError) Error() string {
	return fmt.Sprintf("Invalid Parameter %s", e.Param)
}

type Manager interface {
	Getter
	Updater
}

type Record struct {
	ID      string
	Name    string
	Type    string
	Content string
}

type Getter interface {
	ZoneIDByName(zoneName string) (string, error)
	GetDNSRecords(recordNames []string) ([]Record, error)
}

type Updater interface {
	UpdateDNSRecords(ipAddress string, recordID string) error
}
