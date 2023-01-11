package database

import (
	"finance-api-v1/core/config"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"time"
)

// NewDatabase creates a new database with given config
func NewDatabase(cfg *config.Config) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)

	for i := 0; i <= 30; i++ {
		db, err = gorm.Open(sqlserver.Open(cfg.DBConfig.DataSourceName), &gorm.Config{})
		if err != nil {
			time.Sleep(500 * time.Millisecond)
		}
	}
	if err != nil {
		return nil, err
	}

	origin, err := db.DB()
	if err != nil {
		return nil, err
	}
	origin.SetMaxOpenConns(cfg.DBConfig.Pool.MaxOpen)
	origin.SetMaxIdleConns(cfg.DBConfig.Pool.MaxIdle)
	origin.SetConnMaxLifetime(time.Duration(cfg.DBConfig.Pool.MaxLifetime) * time.Second)
	return db, nil
}
