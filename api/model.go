package api

type Store interface {
	GetVersion() (*DBVersion, error)
	GetActiveClient() (*DBActiveClient, error)
	GetHealth() (*DBHealth, error)
}

type DBVersion struct {
	Version string `json:"version"`
}

type DBActiveClient struct {
	ActiveClient int `json:"active_client"`
}

type DBHealth struct{}
