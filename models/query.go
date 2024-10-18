package models

type QuerySearch struct {
	Query   map[string]interface{} `json:"query" mapstructure:"query"`
	Type    string                 `json:"type" mapstructure:"type"`
	Sorting []struct{}             `json:"sorting" mapstructure:"sorting"`
	Count   int                    `json:"count" mapstructure:"count"`
	Page    int                    `json:"page" mapstructure:"page"`
}
