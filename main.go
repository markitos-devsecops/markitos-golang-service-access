package main

import (
	"log"
	"markitos-golang-service-access/infrastructure/api"
	"markitos-golang-service-access/infrastructure/configuration"
	"markitos-golang-service-access/infrastructure/database"
	"markitos-golang-service-access/internal/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	//------------------------------------------------TODO: extract-method_LoadConfiguration
	log.Println("['.']:>")
	log.Println("['.']:>--------------------------------------------")
	log.Println("['.']:>--- <starting markitos-golang-service-access>")
	log.Println("['.']:>----- <configuration>")
	config, err := configuration.LoadConfiguration(".")
	if err != nil {
		log.Fatal("['.']:>------- unable to load configuration: ", err)
	}
	log.Println("['.']:>------- all values ready to use :)")
	log.Println("['.']:>------- serverAddress: ", config.AppAddress)
	log.Println("['.']:>----- </configuration>")
	//------------------------------------------------

	//------------------------------------------------TODO: extract-method_LoadDatabase
	log.Println("['.']:>----- <database>")
	db, err := gorm.Open(postgres.Open(config.DsnDatabase), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	//------------------------------------------------
	// Migrate the schema (migrate)
	// solo usarlo en caso de no hacer uso de migrate
	// comentar este bloque si hacemos uso de migrate
	//------------------------------------------------
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatal(err)
	}
	repository := database.NewUserPostgresRepository(db)
	log.Println("['.']:>------- Connected to database - migrations")
	log.Println("['.']:>----- </database>")
	//------------------------------------------------

	//------------------------------------------------TODO: extract-method_StartServer
	log.Println("['.']:>----- <server.api>")
	gin.SetMode(gin.ReleaseMode)
	server := api.NewServer(config.AppAddress, repository)
	log.Println("['.']:>------- New server created")
	log.Println("['.']:>----- </server.api>")
	log.Println("['.']:>--- </starting markitos-golang-service-access>")
	log.Println("['.']:>--------------------------------------------")
	log.Println("['.']:>")
	err = server.Run()
	if err != nil {
		log.Fatal("unable to start server: ", err)
	}
	//------------------------------------------------
}
