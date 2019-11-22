package contexts

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// Import gorm mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	// DB is the globally shared DB session.
	DB *gorm.DB
)

// Setup will setup the whole contexts
func Setup() (err error) {
	DB, err = gorm.Open("mysql", "test.db")
	if err != nil {
		return fmt.Errorf("context setup: %w", err)
	}
	return
}
