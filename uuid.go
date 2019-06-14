package utils

import (
	"github.com/satori/go.uuid"
)

func NewUUID() string {
	uuid := uuid.NewV4()
	return uuid.String()
}
