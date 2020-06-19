package util

import (
	"strings"

	"github.com/google/uuid"
)

// GenerateUUID - this function generate as UUID using googles UUID
func GenerateUUID() string {
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return uuid
}
