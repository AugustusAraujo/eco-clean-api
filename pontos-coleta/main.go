package main

import (
	"fmt"

	_ "github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

)

type CoordinatesOfPoints struct {
	gorm.Model
	Latitude  float64
	Longitude float64
}

type Point struct {
	gorm.Model
	Description string
}

func main() {
	const DSN = "host=localhost user=postgres password=123 dbname=postgres port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	// RUN MIGRATIONS
	db.AutoMigrate(&Point{})
	db.AutoMigrate(&CoordinatesOfPoints{})

	if err != nil {
		panic("Can't connect to database.")
	}

	db.Create(&Point{Description: "Lolzinho"})

	result := db.Take(&Point{})

	fmt.Println(result.Rows())
}
