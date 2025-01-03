package main

import "github.com/gofiber/fiber/v2"

func main() {
	app:=fiber.New()

	// Serve static files 
	app.Static("/","./Frontend")

	// Start the server 
	app.Listen(":3000")
}