package utils

import (
	"github.com/satori/go.uuid"
	"strings"
)

func GenerateUUID() string{
	uuids, _ := uuid.NewV4()
	return strings.Replace(uuids.String(),"-","",-1)
}