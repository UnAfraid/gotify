//go:build cgo

package database

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite" // enable the sqlite3 dialect.
)

func cgoAwareDialect(dialect string) string {
	if dialect == "sqlite" {
		dialect = "sqlite3"
	}
	return dialect
}
