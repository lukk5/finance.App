package database

import (
	"errors"
	"gorm.io/gorm"
)

var (
	ErrNotFound    = errors.New("record not found")
	ErrKeyConflict = errors.New("key conflict")
)

// IsRecordNotFoundErr returns true if err is gorm.ErrRecordNotFound or ErrNotFound
func IsRecordNotFoundErr(err error) bool {
	return err == gorm.ErrRecordNotFound || err == ErrNotFound
}

// IsKeyConflictErr returns true if err is ErrKeyConflict or MySQLError with 1062 code number
func IsKeyConflictErr(err error) bool {
	if err == ErrKeyConflict {
		return true
	}
	return false
}
