package models

import "time"

type Observable struct {
	Value        string        `json:"value"`
	Type         string        `json:"type"`
	Created      string        `json:"created"`
	Context      []interface{} `json:"context"`
	LastAnalysis struct {
	} `json:"last_analysis"`
	Country     string `json:"country"`
	Description string `json:"description"`
	ID          string `json:"id"`
	Tags        struct {
		AdditionalProp1 struct {
			Source   string    `json:"source"`
			Target   string    `json:"target"`
			LastSeen time.Time `json:"last_seen"`
			Expires  time.Time `json:"expires"`
			Fresh    bool      `json:"fresh"`
			ID       string    `json:"id"`
		} `json:"additionalProp1"`
		AdditionalProp2 struct {
			Source   string    `json:"source"`
			Target   string    `json:"target"`
			LastSeen time.Time `json:"last_seen"`
			Expires  time.Time `json:"expires"`
			Fresh    bool      `json:"fresh"`
			ID       string    `json:"id"`
		} `json:"additionalProp2"`
		AdditionalProp3 struct {
			Source   string    `json:"source"`
			Target   string    `json:"target"`
			LastSeen time.Time `json:"last_seen"`
			Expires  time.Time `json:"expires"`
			Fresh    bool      `json:"fresh"`
			ID       string    `json:"id"`
		} `json:"additionalProp3"`
	} `json:"tags"`
	RootType              string    `json:"root_type"`
	LastSeen              time.Time `json:"last_seen,omitempty"`
	FirstSeen             time.Time `json:"first_seen,omitempty"`
	Issuer                string    `json:"issuer,omitempty"`
	Subject               string    `json:"subject,omitempty"`
	SerialNumber          string    `json:"serial_number,omitempty"`
	After                 time.Time `json:"after,omitempty"`
	Before                time.Time `json:"before,omitempty"`
	Fingerprint           string    `json:"fingerprint,omitempty"`
	Name                  string    `json:"name,omitempty"`
	Size                  int       `json:"size,omitempty"`
	Sha256                string    `json:"sha256,omitempty"`
	Md5                   string    `json:"md5,omitempty"`
	Sha1                  string    `json:"sha1,omitempty"`
	MimeType              string    `json:"mime_type,omitempty"`
	Key                   string    `json:"key,omitempty"`
	Data                  string    `json:"data,omitempty"`
	Hive                  string    `json:"hive,omitempty"`
	PathFile              string    `json:"path_file,omitempty"`
	UserID                string    `json:"user_id,omitempty"`
	Credential            string    `json:"credential,omitempty"`
	AccountLogin          string    `json:"account_login,omitempty"`
	AccountType           string    `json:"account_type,omitempty"`
	DisplayName           string    `json:"display_name,omitempty"`
	IsServiceAccount      bool      `json:"is_service_account,omitempty"`
	IsPrivileged          bool      `json:"is_privileged,omitempty"`
	CanEscalatePrivs      bool      `json:"can_escalate_privs,omitempty"`
	IsDisabled            bool      `json:"is_disabled,omitempty"`
	AccountCreated        time.Time `json:"account_created,omitempty"`
	AccountExpires        time.Time `json:"account_expires,omitempty"`
	CredentialLastChanged time.Time `json:"credential_last_changed,omitempty"`
	AccountFirstLogin     time.Time `json:"account_first_login,omitempty"`
	AccountLastLogin      time.Time `json:"account_last_login,omitempty"`
	Coin                  string    `json:"coin,omitempty"`
	Address               string    `json:"address,omitempty"`
}
