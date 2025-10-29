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
	// that supports a browser triggered Modal Dialog, prompting the Guest to 
	// confirm an action
	app.Get("/promptModal", func(c *fiber.Ctx) error {
		// return c.SendString("Generate Browser Modal!")
 		log.Println("Prompt Modal Dialog: Route to prompt the Guest to confirm an action")
        return c.Render("modalDialog", fiber.Map {
            "title": "Action Required",
            "body": "Please confirm you want to delete this record, or exit the dialog.",
            "action_route": "promptModalConfirm",
            "action_label": "Confirm",  
			"action_class": "btn  btn-danger",  

            })
    })

	// Route to render the Guest Confirming the prompt by clicking the Confirm button
	app.Get("/promptModalConfirm", func(c *fiber.Ctx) error {
		message := "Prompt Modal Dialog: Route to handle the Guest confirming prompt by clicking the Confirm button on the Modal Dialog"
		log.Println(message)
		return c.Status(fiber.StatusOK).SendString(message)
	})

	// Route to handle the Client request for the server to execute the logic
	// that supports a server triggered Modal Dialog, informing the Guest about
	// a server state
	app.Get("/informModal", func(c *fiber.Ctx) error {
		// return c.SendString("Generate Server Modal!")
		log.Println("Modal Dialog: Route to handle the server informing the Guest about a server state")
        data := fiber.Map {
            "title": "Authorization Error",
            "body": "You are not authorized to perform this action.",
            "action_route": "", //
            "action_label": "",       
 			"action_class": "",  
           }

        // Trigger a dialog_event in the server!
        c.Set("HX-Trigger", "dialog_event")
        return c.Render("modalDialog", data)
	})

	// Route to handle the Client request for the server to execute the logic
	// that supports collecting data, the guest's age in this case;
	app.Get("/collectModal", func(c *fiber.Ctx) error {
		// return c.SendString("Generate Server Modal!")
		log.Println("Modal Dialog: Route to handle the server requesting the Guest to provide data")
        data := fiber.Map {
            "title": "Guest Data Collection",
            "body": "Please Provide the data below",
            "action_route": "collectModalSubmit",
            "action_label": "Submit",      
 			"action_class": "btn btn-primary",  
           }

        // Trigger a dialog_event in the server!
        c.Set("HX-Trigger", "dialog_event")
        return c.Render("modalDialog", data)
	})

	// Route to render the Client Browser Modal Decline
	app.Post("/collectModalSubmit", func(c *fiber.Ctx) error {
		age := c.FormValue("age")
		message := "Collect: Route to handle age provided by guest: " + age
		log.Println(message)
		return c.Status(fiber.StatusOK).SendString(message)
	})

	// Start the server
	app.Listen(":3000")
}
 