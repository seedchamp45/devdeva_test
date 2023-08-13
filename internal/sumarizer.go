package internal

import (
	"devdeva_test/model"
)

func GetActivePower(query string) map[string]int {
	data := model.GetData()

	activeSum := 0
	if query == "active_power" {
		for _, entry := range data {
			activeSum += entry.ActivePower

		}
		result := map[string]int{
			"active_power_sum": activeSum,
		}

		return result
	} else if query == "input_power" {
		for _, entry := range data {
			activeSum += entry.InputPower

		}

		result := map[string]int{
			"input_power_sum": activeSum,
		}

		return result
	}

	result := map[string]int{
		"sum": activeSum,
	}

	return result
}
