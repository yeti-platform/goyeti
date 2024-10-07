package models

import "time"

type ClientConfig struct {
	Host    string
	Port    int
	Timeout time.Duration
	Proto   string
}

//go:generate mockgen -source=client.go -package=mocks -destination=mocks/client.gen.go
type ClientI interface {
	Init() error
	ObservablesSearch(query QuerySearch) (ServerResponseObservables, error)
	Query(endpoint string, method string, data []byte) (map[string]interface{}, error)
	Close() error
}
