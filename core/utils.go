package core

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

func UUID(prefix string) string {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		log.Printf("UUID generation error: %v", err)
		return ""
	}
	return fmt.Sprintf("%s-%s", prefix, newUUID.String())
}
