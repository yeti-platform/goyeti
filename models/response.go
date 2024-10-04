package models

type ServerResponseAuth struct {
	AccessToken string `json:"access_token" mapstructure:"access_token"`
	TokenType   string `json:"token_type" mapstructure:"token_type"`
}

type ServerResponseObservables struct {
	Observables []Observable `json:"observables" mapstructure:"observables"`
	Total       int          `json:"total" mapstructure:"total"`
}
