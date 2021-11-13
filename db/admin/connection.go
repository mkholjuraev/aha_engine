package admin

import (
	"fmt"
	"log"
	"net/url"

	"github.com/mkholjuraev/aha_engine/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// type DatabaseConnection interface {
// 	Get() *gorm.DB
// }

// type databaseConnetion struct {
// 	DB *gorm.DB
// }

var DB *gorm.DB

func NewDatabaseConncetion() *gorm.DB {
	config, err := util.LoadDbConfig("config")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	var enableLogging logger.Interface
	if config.DB_IS_LOGGING_ENALBED {
		enableLogging = logger.Default
	}

	dsn := url.URL{
		User:     url.UserPassword(config.DBUsername, config.DBPassword),
		Scheme:   config.DBScheme,
		Host:     fmt.Sprintf("%s:%d", config.DBHost, config.DBPort),
		Path:     config.DBDatabase,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{
		Logger: enableLogging,
	})

	if err != nil {
		panic(fmt.Sprintf("database connection failed: %s", err))
	}

	if config.DBSync {
		synchronize(db)
	}
	DB = db
	return DB
}

// func (d *databaseConnetion) Get() *gorm.DB {
// 	return d.DB
// }
