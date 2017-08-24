package gosyncengine

// Account providers information on a single account defined on the sync engine
type Account struct {
	ID               string `json:"id"`
	AccountID        string `json:"account_id"`
	Object           string `json:"object"`
	Name             string `json:"name"`
	EmailAddress     string `json:"email_address"`
	Provider         string `json:"provider"`
	OrganizationUnit string `json:"organization_unit"`
	SyncState        string `json:"sync_state"`
}

// Accounts is a typed array of Account objects
type Accounts []Account
