package gen
//go:generate mockgen -package mockdns -destination internal/mock/mockdns/mockdns.go github.com/darkraiden/odysseus/internal/DNS Manager
//go:generate mockgen -package mockip -destination internal/mock/mockip/mockip.go github.com/darkraiden/odysseus/internal/ipaddress Getter,Doer


