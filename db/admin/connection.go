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
	config, err := utils.LoadDbConfig("config")
	var dsn string

	if err != nil {
		log.Println("cannot load config: ", err)
		dsn = os.Getenv("DATABASE_URL")
	} else {
		dsnURL := url.URL{
			User:   url.UserPassword(config.DBUsername, config.DBPassword),
			Scheme: config.DBScheme,
			Host:   fmt.Sprintf("%s:%d", config.DBHost, config.DBPort),
			Path:   config.DBDatabase,
		}
		dsn = dsnURL.String()
	}

	var enableLogging logger.Interface
	if config.DB_IS_LOGGING_ENALBED {
		enableLogging = logger.Default
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
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
