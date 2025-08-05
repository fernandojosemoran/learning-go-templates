package blog

import (
	models "github.com/fernandojosemoran/go-templates/internal/apps/blog/application/models"
	"gorm.io/gorm"
)

func BlogMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Article{},
	)

	if err != nil {
		return err
	}

	return nil
}
