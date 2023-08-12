package repositories

import (
	"fmt"
	"os"

	"repo/configs"
	"repo/loggers"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var log = loggers.Get()

// DBConnect ...
var DBConnect *gorm.DB

// Config ...
func Config() {
	// Init error codes
	DBConn()
	Ping()
}

// DBConn ...
func DBConn() error {
	host := configs.MustGetString("database.master.host")
	user := configs.MustGetString("database.master.username")
	pswd := configs.MustGetString("database.master.password")
	dbnm := configs.MustGetString("database.master.database")
	port := configs.MustGetString("database.master.port")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, user, pswd, dbnm, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Info(fmt.Sprintf("err : %v", err))
		return err
	}

	DBConnect = db
	return nil
}

// Ping ...
func Ping() {
	pingDB, _ := DBConnect.DB()
	ping := pingDB.Ping()
	if ping != nil {
		log.Info("Failed Connecting Database.")
		os.Exit(1)
	} else {
		log.Info("Success Connecting Database.")
	}
}

func GenerateUUID() uuid.UUID {
	id, _ := uuid.NewUUID()
	return id
}
