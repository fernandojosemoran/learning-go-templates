package blog

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	config "github.com/fernandojosemoran/go-templates/pkg/config/constant" // Ajusta la ruta del paquete config
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseOptions struct {
	Migrations []func(db *gorm.DB) error
	Context    context.Context
}

var (
	DB   *gorm.DB
	once sync.Once
	err  error
)

var env = config.Env()

func InitializationDB(conf *DatabaseOptions) error {
	once.Do(func() {
		DB, err = connect(conf.Context, conf.Migrations)
	})
	return err
}

func connect(ctx context.Context, migrations []func(db *gorm.DB) error) (*gorm.DB, error) {
	select {
	case <-ctx.Done():
		return nil, errors.New("psql connection canceled.")
	default:
		db, err := gorm.Open(postgres.Open(env.POSTGRES_DNS), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})

		if err != nil {
			return nil, fmt.Errorf("failed to connect to '%s' database: %w", env.POSTGRES_DB_NAME, err)
		}

		sqlDB, err := db.DB()

		if err != nil {
			return nil, fmt.Errorf("failed to get underlying SQL database: %w", err)
		}

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)

		for _, migration := range migrations {
			if err = migration(db); err != nil {
				return nil, err
			}
		}

		return db, nil
	}
}
