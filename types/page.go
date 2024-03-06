package types

import "time"

// Page is a struct that represents a page in the database
type Page struct {
	Value     []byte
	ExpiresAt *time.Time
}
