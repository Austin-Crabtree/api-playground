package models

import (
	"api-playground/utils"
	"math/rand"
)

// Data is a mockup struct of what data would look like
type Data struct {
	Id      int      `json:"id"`
	Version int      `json:"version"`
	Columns []string `json:"columns"`
	Data    [][]int  `json:"data"`
}

// Fill a Data struct with random data
func (d *Data) FillRandom() {
	d.Id = rand.Intn(10)
	d.Version = rand.Intn(20)
	numColumns := rand.Intn(25)
	// Generate random column names
	for i := 0; i < numColumns; i++ {
		lenString := rand.Intn(20)
		d.Columns = append(d.Columns, utils.RandStringBytes(lenString))
	}
	numRows := rand.Intn(25)
	// Generate random data for 2d data field
	for i := 0; i < numColumns; i++ {
		var row []int
		for j := 0; j < numRows; j++ {
			row = append(row, rand.Intn(1000))
		}
		d.Data = append(d.Data, row)
	}
}
