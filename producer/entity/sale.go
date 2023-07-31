package entity

import (
	"math/rand"
	"producer/utils"
	"time"

	"github.com/google/uuid"
	"github.com/goombaio/namegenerator"
)

type Sale struct {
	Id            string    `json:"id"`
	Salesman      string    `json:"salesman"`
	Customer      string    `json:"customer"`
	Brand         string    `json:"brand"`
	Product       string    `json:"product"`
	OriginalPrice float32   `json:"originalPrice"`
	DiscountRate  int       `json:"discountRate"`
	Timestamp     time.Time `json:"timestamp"`
}

func (sale Sale) MakeSale() Sale {
	csvSalesData := utils.ReadFromCSV("/home/caf/dev/kafka-flink-druiddb-sandbox/producer/data/products.csv")

	return Sale{
		Id:            uuid.NewString(),
		Salesman:      utils.ReadRandomLineFromTxt("/home/caf/dev/kafka-flink-druiddb-sandbox/producer/data/salesman.txt"),
		Customer:      namegenerator.NewNameGenerator(rand.Int63()).Generate(),
		Brand:         csvSalesData.Brand,
		Product:       csvSalesData.Product,
		OriginalPrice: csvSalesData.OriginalPrice,
		DiscountRate:  int(csvSalesData.DiscountRate),
		Timestamp:     getRandomTimestamp(),
	}
}

func getRandomTimestamp() time.Time {
	startTime := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2023, time.July, 30, 0, 0, 0, 0, time.UTC)

	duration := endTime.Sub(startTime)

	randomNanoseconds := rand.Int63n(duration.Nanoseconds())

	return startTime.Add(time.Duration(randomNanoseconds))
}
