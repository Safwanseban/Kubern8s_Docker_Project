package configs

import (
	"fmt"
	"log"

	"github.com/Safwanseban/Kubern8s_Docker_Project/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var env = Envload()
var Db *gorm.DB

func ConnectDB() {

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		"postgress-dp", env.Db_User, env.Db_Password, env.Db_Name, env.Db_Port)
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting db", err)
	}
	Db.AutoMigrate(&models.User{})
}
