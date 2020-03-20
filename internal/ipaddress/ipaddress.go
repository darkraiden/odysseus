package ipaddress

type Getter interface {
	GetLocal() (string, error)
}
