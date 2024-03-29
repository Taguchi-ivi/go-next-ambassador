package main

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"math/rand"

	"github.com/bxcodec/faker/v4"
)

func main2() {

	database.Connect()
	for i := 0; i < 10; i++ {
		product := models.Product{
			Title:       faker.Username(),
			Description: faker.Username(),
			Image:       faker.URL(),
			Price:       float64(rand.Intn(90) + 10),
		}

		database.DB.Create(&product)

	}
}
