package config

import (
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error cargando el archivo .env")
    }

    dsn := os.Getenv("DATABASE_URL")
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Error conectando a la base de datos:", err)
    }

    DB = database
}
