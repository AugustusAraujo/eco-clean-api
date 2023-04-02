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
	CreatedAt  string
	UpdatadeAt string
}

func main() {
	const DSN = "host=localhost user=postgres password=123 dbname=postgres port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	db.AutoMigrate(&Point{})

	if err != nil {
		panic("Can't connect to database.")
	}
	var result int
	db.Raw("SELECT 1+7 AS result;").Scan(&result)

	fmt.Println(result)
}
