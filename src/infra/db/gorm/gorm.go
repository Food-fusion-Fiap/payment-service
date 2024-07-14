package gorm

import (
	"fmt"
	"github.com/CAVAh/api-tech-challenge/src/infra/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {

	//Local development
	//conectionString := "host=postgres.cao4mee9fcpi.us-east-1.rds.amazonaws.com user=postgres password=rootroot dbname=postgres port=5432 sslmode=require TimeZone=America/Fortaleza"
	//DB, err = gorm.Open(postgres.Open(conectionString))

	conectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=require TimeZone=America/Fortaleza",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	//fmt.Println(conectionString)
	DB, err = gorm.Open(postgres.Open(conectionString))

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Payment{})
}
