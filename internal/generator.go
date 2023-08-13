// internal/generator.go
package internal

import (
	"devdeva_test/model"
	"math/rand"
)

func GenerateData(numRows int) {
	data := make([]model.Data, numRows)
	for i := 0; i < numRows; i++ {
		activePower := rand.Intn(1000) + 1
		inputPower := rand.Intn(1000) + 1
		data[i] = model.Data{ActivePower: activePower, InputPower: inputPower}
	}

	model.SetData(data)
}
