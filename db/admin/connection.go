package admin

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/mkholjuraev/publico_engine/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func NewDatabaseConncetion() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")

	var enableLogging logger.Interface

	if len(dsn) == 0 {
		log.Println("can not read config from the server, reading local config")

		config, err := utils.LoadDbConfig("config")
		if err != nil {
			log.Println("Database configuration could not be read")
		}

		dsnURL := url.URL{
			User:   url.UserPassword(config.DBUsername, config.DBPassword),
			Scheme: config.DBScheme,
			Host:   fmt.Sprintf("%s:%d", config.DBHost, config.DBPort),
			Path:   config.DBDatabase,
		}
		dsn = dsnURL.String()

		if config.DB_IS_LOGGING_ENALBED {
			enableLogging = logger.Default
		}
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: enableLogging,
	})

	if err != nil {
		panic(fmt.Sprintf("database connection failed: %s", err))
	}

	synchronize(db)
	DB = db
	return DB
}
