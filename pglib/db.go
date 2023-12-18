package pglib

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(cfg *Config) error {
	instance, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN: cfg.dsn(),
			// disables implicit prepared statement usage
			PreferSimpleProtocol: true,
		}),
		&gorm.Config{},
	)
	if err != nil {
		return err
	}

	sqlDb, err := instance.DB()
	if err != nil {
		return err
	}

	sqlDb.SetConnMaxLifetime(cfg.getLifeDuration())
	sqlDb.SetMaxOpenConns(cfg.MaxConn)
	sqlDb.SetMaxIdleConns(cfg.MaxIdle)

	db = instance

	return nil
}

func DB() *gorm.DB {
	return db
}
