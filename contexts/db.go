package contexts

import (
	"fmt"

	"github.com/Xuanwo/bard/model"
	"github.com/Xuanwo/storage"
	"github.com/Xuanwo/storage/coreutils"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	// Import gorm driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	// DB is the globally shared DB session.
	DB *gorm.DB
	// Storage is the globally shared Storage session.
	Storage storage.Storager

	// Server is the globally shared server related config.
	Server struct {
		PublicURL   string
		Listen      string
		Key         string
		MaxFileSize int64
	}
)

// Setup will setup the whole contexts
func Setup() (err error) {
	errorMessage := "contexts Setup: %w"

	// Setup server.
	Server.PublicURL = viper.GetString("public_url")
	Server.Key = viper.GetString("key")
	Server.Listen = viper.GetString("listen")
	Server.MaxFileSize = viper.GetInt64("max_file_size")

	// Setup DB.
	DB, err = gorm.Open(
		viper.GetString("database.type"),
		viper.GetString("database.connection"),
	)
	if err != nil {
		return fmt.Errorf(errorMessage, err)
	}
	DB.AutoMigrate(&model.Poem{})

	// Setup storage.
	Storage, err = coreutils.OpenStorager(viper.GetString("storage"))
	if err != nil {
		return fmt.Errorf(errorMessage, err)
	}
	return
}
