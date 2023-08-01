package manager

import (
	"database/sql"
	"fmt"
	"log"
	"project/config"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type InfraManager interface {
	GetDB() *sql.DB
}

type infraManager struct {
	db  *sql.DB
	cfg config.Config
}

var onceLoadDb sync.Once

func (im *infraManager) GetDB() *sql.DB {
	onceLoadDb.Do(func() {
		psqlCon := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", im.cfg.Host, im.cfg.Port, im.cfg.User, im.cfg.Password, im.cfg.Name)
		db, err := sql.Open("postgres", psqlCon)
		if err != nil {
			log.Fatal("Cannot start app, Error when connect to DB ", err.Error())
		}
		im.db = db

		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			fmt.Println(err)
			panic("server error")
		}

		m, err := migrate.NewWithDatabaseInstance("file://database/migrations/", "postgres", driver)
		if err != nil {
			fmt.Println(err)
			panic("server error")
		}

		m.Down()
		m.Up()
	})

	return im.db
}

func (i *infraManager) DbConn() *sql.DB {
	return i.db
}

func NewInfraManager(config config.Config) InfraManager {
	infra := infraManager{
		cfg: config,
	}
	infra.GetDB()
	return &infra
}
