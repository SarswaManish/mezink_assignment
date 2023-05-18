package postgres

import (
	"awesomeProject/internal/database"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type pgServer struct {
	user     string
	password string
	host     string
	db       string
}

func NewPgClient(user, password, host, db string) database.PgClient {
	return &pgServer{user: user, host: host, password: password, db: db}
}

func (cs *pgServer) Connect() (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cs.host,
		cs.user,
		cs.password,
		cs.db,
		5432)

	db, err := sql.Open("postgres", dsn)

	dbs, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})

	return dbs, err
}
