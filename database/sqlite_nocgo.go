//go:build !cgo

package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "modernc.org/sqlite" // enable the sqlite dialect.
)

func init() {
	if dialect, ok := gorm.GetDialect("sqlite3"); ok {
		gorm.RegisterDialect("sqlite", dialect)
	} else {
		fmt.Println("Unable to find sqlite3 dialect")
	}
}

func cgoAwareDialect(dialect string) string {
	if dialect == "sqlite3" {
		dialect = "sqlite"
	}
	return dialect
}
