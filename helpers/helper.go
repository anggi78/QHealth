package helpers

import "github.com/google/uuid"

func CreateId() string {
	newUuid := uuid.New()
	return newUuid.String()
}