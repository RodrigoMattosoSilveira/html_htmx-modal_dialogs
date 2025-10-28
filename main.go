package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Set up the HTML template engine
	engine := html.New("./views", ".html")

	// Create a new Fiber app with the template engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Route to render the Landing Page
	app.Get("/", func(c *fiber.Ctx) error {
		log.Println("Route to render the Landing Page")
		return c.Render("index", nil)
	})

	// Route to handle the Client request for the server to execute the logic
	// that supports a browser triggered Modal Dialog rendering
	app.Get("/browserModal", func(c *fiber.Ctx) error {
		// return c.SendString("Generate Browser Modal!")
 		log.Println("Route to handle the Client request for the server to execute the logic that supports a browser triggered Modal Dialog rendering")
        return c.Render("modalDialog", fiber.Map {
            "title": "Browser Triggered Modal Title",
            "body": "Browser Triggered Modal Body",
            "confirm_endpoint": "browserModalAccept",
            "confirm_label": "Accept",
            "cancel_endpoint": "browserModalDecline",
            "cancel_label": "Decline",       
            })
    })

	// Route to render the Client Browser Modal Decline
	app.Get("/browserModalDecline", func(c *fiber.Ctx) error {
		message := "Route to handle the Guest clicking the Decline button on the Browser Modal Dialog"
		log.Println(message)
		return c.Status(fiber.StatusOK).SendString(message)
	})

	// Route to render the Client Browser Modal Accept
	app.Get("/browserModalAccept", func(c *fiber.Ctx) error {
		message := "Route to handle the Guest clicking the Accept button on the Browser Modal Dialog"
		log.Println(message)
		return c.Status(fiber.StatusOK).SendString(message)
	})

	// Route to handle the Client request for the server to execute the logic
	// that supports a server triggered Modal Dialog rendering
	app.Get("/serverModal", func(c *fiber.Ctx) error {
		// return c.SendString("Generate Server Modal!")
		log.Println("Route to handle the Client request for the server to execute the logic that supports a server triggered Modal Dialog rendering")
        data := fiber.Map {
            "title": "Server Triggered Modal Title",
            "body": "Server Triggered Modal Body",
            "confirm_endpoint": "serverModalAccept",
            "confirm_label": "Accept",
            "cancel_endpoint": "serverModalDecline",
            "cancel_label": "Decline",       
            }

        // Trigger a dialog_event in the server!
        c.Set("HX-Trigger", "dialog_event")
        return c.Render("modalDialog", data)
	})

	// Route to render the Client Browser Modal Decline
	app.Get("/serverModalDecline", func(c *fiber.Ctx) error {
		message := "Route to handle the Guest clicking the Decline button on the Server Modal Dialog"
		log.Println(message)
		return c.Status(fiber.StatusOK).SendString(message)
	})

	// Route to render the main page
	app.Get("/serverModalAccept", func(c *fiber.Ctx) error {
		message := "Route to handle the Guest clicking the Accept button on the Server Modal Dialog"
		log.Println(message)
		return c.Status(fiber.StatusOK).SendString(message)
	})

	// Start the server
	app.Listen(":3000")
}
 