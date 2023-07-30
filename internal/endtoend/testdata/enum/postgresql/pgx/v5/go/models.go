// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package querytest

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type Size string

const (
	SizeXSmall Size = "x-small"
	SizeSmall  Size = "small"
	SizeMedium Size = "medium"
	SizeLarge  Size = "large"
	SizeXLarge Size = "x-large"
)

func (e *Size) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Size(s)
	case string:
		*e = Size(s)
	default:
		return fmt.Errorf("unsupported scan type for Size: %T", src)
	}
	return nil
}

type NullSize struct {
	Size  Size
	Valid bool // Valid is true if Size is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullSize) Scan(value interface{}) error {
	if value == nil {
		ns.Size, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Size.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullSize) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Size), nil
}

var EnumNames = []string{
	"size",
}

type User struct {
	ID        int32
	FirstName string
	LastName  pgtype.Text
	Age       int32
	ShoeSize  Size
	ShirtSize NullSize
}
