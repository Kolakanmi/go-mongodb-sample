package uuid

import uuid "github.com/satori/go.uuid"

func New() string {
	newUUID := uuid.NewV4()
	return newUUID.String()
}
