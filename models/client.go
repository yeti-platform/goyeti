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
	ObservablesSearch(search string, searchType string) (ServerResponseObservables, error)
	Query(endpoint string, method string, data string) (map[string]interface{}, error)
	Close() error
}
