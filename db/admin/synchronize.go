package admin

import (
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/mkholjuraev/aha_engine/base/models"
	"gorm.io/gorm"
)

func synchronize(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "fix type: biograph to biography on writer table",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.AutoMigrate(
					models.Writer{},
				); err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(
					models.Writer{},
				)
			},
		},
	})
	fmt.Println("migrated")
	return m.Migrate()
}
