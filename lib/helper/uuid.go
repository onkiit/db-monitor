package helper

import (
	"github.com/satori/go.uuid"
)

func GenerateUUID() (string, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return uid.String(), nil
}
