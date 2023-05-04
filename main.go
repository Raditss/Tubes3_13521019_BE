package main

import (

	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func main(){
	var err error
	DB, err = gorm.Open(postgres.Open("postgresql://mQxYqvGvInYceiGJHDrPeHALObJDxPxU:VmaGpCqFrrSZTZwzwbuEcoBgYrJmTlhl@db.thin.dev/53a2fb09-6906-44f0-9f2e-0239d14bff07"), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	log.Println("database connection successful")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := NewAPIServer(":"+port)
	server.Start()
}
