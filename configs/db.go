package configs

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var env = Envload()
var db *gorm.DB

func ConnectDB() {

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		env.Db_Host, env.Db_User,  env.Db_Password,env.Db_Name, env.Db_Port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting db", err)
	}

}
