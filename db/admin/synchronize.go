package admin

import (
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/mkholjuraev/publico_engine/base/models"
	"gorm.io/gorm"
)

func synchronize(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "Reinitialize database",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.AutoMigrate(
					models.Chat{},
					models.Follower{},
					models.Images{},
					models.Notifications{},
					models.Post{},
					models.Specialization{},
					models.PostMetadata{},
					models.Tags{},
					models.User{},
					models.Writer{},
				); err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(
					models.Chat{},
					models.Follower{},
					models.Images{},
					models.Notifications{},
					models.Post{},
					models.Specialization{},
					models.PostMetadata{},
					models.Tags{},
					models.User{},
					models.Writer{},
				)
			},
		},
	})
	fmt.Println("migrated")
	return m.Migrate()
}
