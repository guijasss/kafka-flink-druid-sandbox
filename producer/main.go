package main

import (
	"encoding/json"
	"fmt"
	"producer/entity"
)

func main() {
	sale := entity.Sale{}
	jsonData, err := json.Marshal(sale.MakeSale())
	if err != nil {
		fmt.Println(err)
	}
	println(string(jsonData))
}
