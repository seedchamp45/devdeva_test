package internal

import "errors"

func Validator(query string) error {
	paramsList := []string{"active_power", "input_power"}
	if query == "" {
		return errors.New("Please input query params")
	}

	for _, param := range paramsList {
		if query == param {
			return nil
		}
	}

	return errors.New("Please input correct params")
}
