package generator

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// ObjectID generates unique object ID.
func ObjectID() string {
	return uuid.Must(uuid.NewV4()).String()
}

// CurrentDateTime generates current datetime.
func CurrentDateTime() string {
	return time.Now().Format(time.RFC3339)
}
