// cmd/main.go
package main

import (
	"devdeva_test/internal"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	internal.GenerateData(100)

	app.Get("/", func(c *fiber.Ctx) error {
		queryValue := c.Query("query")

		err := internal.Validator(queryValue)
		if err != nil {
			return err
		}

		result := internal.GetActivePower(queryValue)
		return c.JSON(result)
	})

	port := 8080
	log.Printf("Server is running on port %d...\n", port)
	app.Listen(":" + strconv.Itoa(port))
}
