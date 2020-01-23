package helper

import (
	uuid "github.com/satori/go.uuid"
)

func GenerateUUID() (string, error) {
	uid := uuid.NewV4()

	return uid.String(), nil
}
