package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/simulatedsimian/cpuusage"
)

func main() {
	app := fiber.New()

	app.Get("/cpuload", func(c *fiber.Ctx) error {
		// display the cpu usage once per second
		u := cpuusage.Usage{}

		err := u.Measure()
		var s string
		if err != nil {
			s = err.Error()
		} else {
			s = fmt.Sprintf("Overall %%: %d Per Core %%: %v\n", u.Overall, u.Cores)
		}
		return c.SendString(s)
	})

	log.Fatal(app.Listen(":3000"))
}
