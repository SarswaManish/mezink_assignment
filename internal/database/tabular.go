package database

import "gorm.io/gorm"

type PgClient interface {
	Connect() (*gorm.DB, error)
}
