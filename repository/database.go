package repository

import (
	"fmt"

	"github.com/caioeverest/transactions-api/config"
	"github.com/caioeverest/transactions-api/logger"
	"github.com/caioeverest/transactions-api/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Conn struct {
	*gorm.DB
}

var (
	conf config.DBConfig
	log  *logger.Logger
	conn Conn
)

// Open db connection
func Start() (OperationsRepo *Repository, AccountsRepo *Repository, TransactionsRepo *Repository) {
	conf = config.Get().Database
	log = logger.Get()

	log.Info("Connecting to database...")
	var (
		err        error
		addrConfig = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			conf.Host, conf.Port, conf.User, conf.DbName, conf.Password)
		db *gorm.DB
	)
	if db, err = gorm.Open("postgres", addrConfig); err != nil {
		log.Panicf("Error connecting with database %s:%d, ERRO: %+v", conf.Host, conf.Port, err)
	}

	if config.Get().ENV == config.DEV {
		db.LogMode(true)
	}

	db.AutoMigrate(&model.Account{}, &model.Operation{}, &model.Transaction{})
	conn = Conn{db}

	return &Repository{conn, model.Operation{}}, &Repository{conn, model.Account{}}, &Repository{conn, model.Transaction{}}
}

//Close db connection
func Shutdown() {
	if err := conn.Close(); err != nil {
		log.Errorf("Error closing database connection - ERROR [%+v]", err)
	}
	log.Info("Database connection closed")
}

//Database health check
func Health() bool {
	err := conn.Raw("select * from pg_database").Error
	return err == nil
}
