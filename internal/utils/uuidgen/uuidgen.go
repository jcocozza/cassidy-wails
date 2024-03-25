package uuidgen

import (
	"github.com/google/uuid"
)

// Generate a uuid of the form xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
func GenerateUUID() string {
	uuidObj := uuid.New()
	return uuidObj.String()
}

