package status

type psql struct{}

func (p psql) GetVersion() (string, error) {
	return "", nil
}

func (p psql) GetActiveClient() (int, error) {
	return 0, nil
}

func (p psql) GetHealth() error {
	return nil
}

func New() Store {
	return psql{}
}
