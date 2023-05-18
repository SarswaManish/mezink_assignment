package common

import (
	"awesomeProject/configs"
	"awesomeProject/internal/database/postgres"
	"awesomeProject/internal/server"
	"fmt"

	"gorm.io/gorm"
)

func MustCreateServer(config Config) *server.Server {

	resources := server.ServerResources{
		DB: MustCreatePgConnection(config),
	}
	resources.Cors = configs.Cors(config.Cors)
	resources.InternalAuth.Username = config.InternalAuth.Username
	resources.InternalAuth.Password = config.InternalAuth.Password
	return server.NewServer(resources)

}

func MustCreatePgConnection(config Config) *gorm.DB {

	db, err := postgres.NewPgClient(config.Database.Postgres.User, config.Database.Postgres.Password,
		config.Database.Postgres.Host,
		config.Database.Postgres.Db).Connect()

	if err != nil {
		fmt.Println(err)
	}

	return db
}
