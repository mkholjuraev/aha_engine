package admin

import (
	"fmt"
	"log"

	"github.com/mkholjuraev/publico_engine/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func NewDatabaseConncetion() *gorm.DB {
	config, err := utils.LoadDbConfig("config")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	var enableLogging logger.Interface
	if config.DB_IS_LOGGING_ENALBED {
		enableLogging = logger.Default
	}

	// dsn := url.URL{
	// 	User:     url.UserPassword(config.DBUsername, config.DBPassword),
	// 	Scheme:   config.DBScheme,
	// 	Host:     fmt.Sprintf("%s:%d", config.DBHost, config.DBPort),
	// 	Path:     config.DBDatabase,
	// 	RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	// }

	db, err := gorm.Open(postgres.Open("postgres://hpyzicqwgsanug:ee7101397712b5deb369e233bcfe8f75ff078549eb1e2b88a1b0635b524a5799@ec2-3-223-213-207.compute-1.amazonaws.com:5432/dh5j5virfqd8t"), &gorm.Config{
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
