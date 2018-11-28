package status

type Store interface {
	GetVersion() (string, error)
	GetActiveClient() (int, error)
	GetHealth() error
}
