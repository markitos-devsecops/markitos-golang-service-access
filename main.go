package main

import (
	"log"
	"markitos-golang-service-access/infrastructure/api"
	"markitos-golang-service-access/infrastructure/configuration"
	"markitos-golang-service-access/infrastructure/database"
	"markitos-golang-service-access/infrastructure/implementations"
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/domain/dependencies"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	log.Println("['.']:>")
	log.Println("['.']:>--------------------------------------------")
	log.Println("['.']:>--- <starting markitos-golang-service-access>")
	hasher := implementations.NewHasherBCrypt()
	config := loadConfiguration()
	tokener := loadTokener(config)
	repository, err := loadDatabase(config)
	if err != nil {
		log.Fatal(err)
	}
	server := loadServer(config, repository, tokener, hasher)
	log.Println("['.']:>--- </starting markitos-golang-service-access>")
	log.Println("['.']:>--------------------------------------------")
	log.Println("['.']:>")
	err = server.Run()
	if err != nil {
		log.Fatal("unable to start server: ", err)
	}
}

func loadTokener(config configuration.MarkitosGolangServiceAccessConfig) dependencies.Tokener {
	tokener, err := implementations.NewTokenerPasseto(config.SymmetricKey)
	if err != nil {
		log.Fatal("['.']:> error unable to create tokener: ", err)
	}
	return tokener
}

func loadServer(
	config configuration.MarkitosGolangServiceAccessConfig,
	repository *database.UserPostgresRepository,
	tokener dependencies.Tokener,
	hasher dependencies.Hasher) *api.Server {
	gin.SetMode(gin.ReleaseMode)
	server := api.NewServer(config.AppAddress, repository, tokener, config.TokenDuration, hasher)
	log.Println("['.']:>------- Tokener created with duration: ", config.TokenDuration)
	log.Println("['.']:>------- New server created")
	return server
}

func loadDatabase(config configuration.MarkitosGolangServiceAccessConfig) (*database.UserPostgresRepository, error) {
	db, err := gorm.Open(postgres.Open(config.DsnDatabase), &gorm.Config{})
	if err != nil {
		log.Fatal("['.']:> error unable to connect to database:", err)
	}
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatal("['.']:> error unable to migrate database:", err)
	}
	repository := database.NewUserPostgresRepository(db)
	log.Println("['.']:>------- Connected to database - migrations")

	return repository, nil
}

func loadConfiguration() configuration.MarkitosGolangServiceAccessConfig {
	config, err := configuration.LoadConfiguration(".")
	if err != nil {
		log.Fatal("['.']:>------- unable to load configuration: ", err)
	}
	log.Println("['.']:>------- all values ready to use :)")
	log.Println("['.']:>------- serverAddress: ", config.AppAddress)

	return config
}
