# Abstract

This repository introduces two different methods to display `HTML / HTMX` modal dialogs utilizing the native [HTML](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/dialog) `<dialog>` element.

A `modal dialog box` is a pop-up window that appears on top of the current browser page, requiring user interaction before the user can return to the main content. It is commonly used for user notifications, forms, or custom content. We use `browser modal dialogs` to prompt the user to confirm an action, such as deleting a record, and `server modal dialogs` to either `inform the user about the result of an operation`, like lack of authorization for access a resource-- or to `collected addition data`, like a phone number.

We will demonstrate our approach using a simple web applications with a technology stack consisting of  [Go](https://go.dev/), [Fiber](https://gofiber.io/), the [Fiber Template Engine](https://docs.gofiber.io/template/html_v2.x.x/html/) and [Bootstrap 5](https://getbootstrap.com/) on the server, and [HTMX](https://htmx.org/) and [HyperScript](https://hyperscript.org/) to manipulate the DOM Tree, enabling a very limited use of Javascript. 

This technology stack requires the Server to serve HTML, whereas other stacks use JSON requiring more complex technology stack on the Client. Deceivingly, both examples

The first method demonstrates how to show a modal dialog triggered `browser logic`.  The second method demonstrates how to show a modal dialog triggered by `server logic`.

Note that:
- When refering to the `Server`, I'm referring to the `Go/Fiber` HTTP server;
- When refering to the `Template`, I'm referring to the `GoFiber Template` engine;
- When refering to the `Client`, I'm referring to any **modern** browser;

# A Word About our Demo Webapp
Web applications using complex Client technology stacks, such as [React](https://react.dev/) and its vast constellation of addons, can handle the logic to triger, display, and manage `Browser Modal Dialog` boxes without any Server interaction. Their Clients' technology stacks interact with their Servers to handle `Server Modal Dialog` boxes.

Our webapp demo uses a significantly simpler Client technology stack, [HTMX](https://htmx.org/) and [HyperScript](https://hyperscript.org/) to manipulate the DOM Tree, with a limited use of Javascript. It requires the Client to request the Server to provide the `Browser Modal Dialog` box Template, but then it handles all interections with the `Browser Modal Dialog` box; like the former, its Client' technology stack interacts with its Servers to handle `Server Modal Dialog` boxes. 

In other words, whereas the former technology stack requires server interactions only for `Server Modal Dialogs`, our webapp requires two distinct types of Server interactions; one to simply serve the `Browser Modal Dialog` Template and the other to realize the requirement for a `Server Modal Dialog`, serve its Template, collect and process the Person's interaction.

Our webapp will use the same Template for its `Browser Modal Dialog` and `Server Modal Dialogs`boxes, passing arguments to the Template Engine to render them accordingly.  

# Broswer Triggered Modal Dialogs
I'll discuss the use cases, the architecture required to support them, and the implementation to support them.

## Use Cases
- **Setup Client/Server**: I, as a guest, want to launch the HTTP server, `Go_Fiber`, and the HTTP client, any `Browser` will do, I will use to experience the `browser and server triggered modal dialogs`;
- **Launch Appplication**: I, as a guest, want to type the Server listening address in the Client address Bar to launch the demo application and see a `Landing Page` consisting of a UI element named `Modal Dialog Tests`, inclulding a button, `Browser`;
- **Show Browser Modal**: I, as a guest, when at the `Modal Dialog Tests` landing page, want to clik on the `Browser` button and observe a responsde consting of a modal dialog including `Exit`, a `Decline`, and an `Accept` buttons;
- **Dismiss Browser Modal**:  I, as a guest, when at the Dialog Tests screen, want to clik the `Exit` button and dismiss the dialong without any additional system behavior;
- I, as a guest, when at the Dialog Tests screen, want to clik the `Decline` button, receive a reply from the server indicating that I clicked the Decline button, and dismiss the dialong without any additional system behavior;
- I, as a guest, when at the Dialog Tests screen, want to clik the `Accept` button, receive a reply from the server indicating I clicked the Accept button, and dismiss the dialong without any additional system behavior;

## Architecture
The best way me to describe it is thru a few architecture diagrams depicting the main technology elements supporting the use case

### Setup Client/Server
This amounts to building and launching he Go/Fiber HTTP Server, which I'll describe below, as well as launching the HTTP client--any browser will do;

### Launching the Application

![Local Image](uml/LaunchApplication.png)

<sub>Launch Sequence Diagram</sub>

- The Guest types the application URL in their Client's address bar;
- The Client issues a "/" route request;
- The Server collaborates with the Template to render the landing page, and to return it back to the client, via the Server;

![Local Image](images/LandingPage.png)

<sub>Landing Page</sub>


### Show Browser Modal
Now the Guest is ready to experiment with the browser triggered modal dialog:
![Local Image](uml/ShowBrowserModal.png)

<sub>Show Browser Modal Sequence Diagram</sub>

- The Guest clicks on the Landing Page's `Browser` button; this button includes two configuration elements to route the request and to assist HTMX to find where and then render the Browser Modal dialog:
  - **hx-get**="/browserModal": it routes the request
  - **hx-target**="#htmx-browser-dialog-container">: it points to HTMX where to render the resulting Modal Dialog
- The Client, with HTMX assistance, issues an "/browserModal" HTTP request to the Server;
- The Server fills up a map with the arguments required for Modal Dialog--note that I'm referring to a generic instead of specific Modal Dialog; this means that I'll use the same dialog for the Browser and Server triggered Modal Dialogs;
- The Server collaborates with the Template to render the Browser Modal Dialog, and to return it back to the Client, via the Server;

![Local Image](images/BrowserModal.png)

<sub>Browser Modal Dialog</sub>

In addition to the Header and Body information, notice the `Exit`, being represented by an `x`, `Decline` and `Confirm` buttons.

Also, if I attempt to click on the `Browser` and `Server` buttons in the `Landing Page`, in this case hidden, I will observe they have been disabled, and will continue to be so until I dismiss the `Browser Modal Dialog`.

## Implementatioon
### HTML Templates 
#### Landing Page

``` HTML
<!DOCTYPE html>
<html lang="en">
<head>
	<!--  See Github for details -->
</head>
<body>
    <!-- Browser initiated dialog placeholder -->
    <div id="htmx-browser-dialog-container"></div>

    <!-- Server initiated dialog placeholder -->
    <div id="htmx-server-dialog-container" _="on dialog_event from body put detail.value into me"></div>

	<!-- Markup to display the buttons to request Browser or Server Modal dialogs -->
	<!-- My styling is very poor, eventually I'll pretty it up -->
	<div class="card d-flex aligns-items-center justify-content-center text-center w-75 position-absolute top-50 start-50 translate-middle">
    	<div class="card-header">
   			<h1>Modal Dialog tests</h1>
		</div>
		<div class="card-body">
			<button type="button" class="btn btn-primary" hx-get="/browserModal" hx-target="#htmx-browser-dialog-container">Browser</button>
			<button type="button" class="btn btn-success" hx-get="/serverModal" hx-target="#htmx-server-dialog-container">Server</button>
		</div>
	</div>
	<!-- See Githiub for deetals -->
  </body>
</html>
```

This markup has the following elements:
- **head**: Where I configure the HTML and the CSS required by the webapp; I ommited them here for clarity sake;
- **#htmx-browser-dialog-container**: The location where HTMX will load the Modal Dialog Box requested by the Clilent;
- **#htmx-server-dialog-container**: The location where HTMX will load the Modal Dialog Box triggered by the Server;
- **.card**: The location of the UX element that hosts the Browser and Server buttons the I use to trigger requests for Browser or Server Modal Dialogs;
  -  **hx-get** I used it in both buttons to direct HTMX to triger the HTTP calls required to handle the Browese and Server Use Cases
  - **hx-target**: I used it in both buttons to direct where HTMX is to place the server responses;
  
#### Broswer Modal Dialog

``` HTML
<dialog class="dialog   w-75" _="on load call htmx.process(me) then call me.showModal() end
                          on dialog_close wait 10ms then remove me end
                          on keydown if the event's key is 'Escape' then log `Clicked ESC key` remove me">
	<div class="card ">
		<div class="card-header">
			<div class="row">
				<div class="col-md-10">
					<h5 class="modal-title">{{ .title }}</h5>
				</div>
				<div class="col-md-2">
					<button type="button" class="close" _="on click send dialog_close log 'Clicked Exit Button'"><span aria-hidden="true">&times;</span></button>
				</div>
			</div>
		</div>
		<div class="card-body">
			<p class="card-text">{{ .body }}.</p>
			<button type="button" class="btn btn-secondary" hx-get="/{{ .cancel_endpoint }}" hx-swap="none" _="on click send dialog_close log 'Clicked Seconday Button'">{{ .cancel_label }}</button>
			<button type="button" class="btn btn-primary" hx-get="/{{ .confirm_endpoint }}"  hx-swap="none" _="on click send dialog_close log 'Clicked Primary Button'">{{ .confirm_label }}</button>
		</div>
	</div>
</dialog>
```

Here is where things get really interesting!

### Go/Fiber logic Configuration
The logic below configures the [GoFiber Template Engine](https://docs.gofiber.io/template/html_v2.x.x/html/)  to render my HTML templates, and [Fiber](https://gofiber.io/) to support my routes.

``` go
package main

import (
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

    // ...

	app.Listen(":3000")
}
 
```

### Go/Fiber logic to handle Browser Triggered Client requests
Next I add the routes to support our `Browser Modal Dialog` use cases:
``` go
    // ...

	// Route to render the landing page
	// Route to render the Landing Page
	app.Get("/", func(c *fiber.Ctx) error {
		log.Println("Route to render the Landing Page")
		return c.Render("bsIndex", nil)
	})

	// Route to handle the Client request for the server to execute the logic
	// that supports a browser triggered Modal Dialog rendering
	app.Get("/browserModal", func(c *fiber.Ctx) error {
		// return c.SendString("Generate Browser Modal!")
 		log.Println("Route to handle the Client request for the server to execute the logic that supports a browser triggered Modal Dialog rendering")
        return c.Render("bsModalDialog", fiber.Map {
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
		message := "Route to handle the Person clicking the Decline button on the Browser Modal Dialog"
		log.Println(message)
		return c.Status(fiber.StatusOK).SendString(message)
	})

	// Route to render the Client Browser Modal Accept
	app.Get("/browserModalAccept", func(c *fiber.Ctx) error {
		message := "Route to handle the Person clicking the Accept button on the Browser Modal Dialog"
		log.Println(message)
		return c.Status(fiber.StatusOK).SendString(message)
	})
    
	// ...
```

Although one migh conjure that the logic for the `/browserModalDecline` route is superflous, there are use cases where it is important to record the Person choice to decline a choice; hence, I elected to include it here for this reason and to offer me a solid proof the the request traveled the required round trip. I also `logged` the execution of all routes on the server on the Client console.

![Local Image](images/BrowserModalDialogServerLogs.png)

<sub>Browser Modal Dialog Server Logs</sub>

Note that the third set of logs shows the Client request for a Browser Modal Dialog, but does not show the server action; this represents two uses cases, one where the Person types the `ESC` key or clicks on the  Modal Dialog box `x` button.

![Local Image](images/BrowserModalDialogConsoleLogs.png)

<sub>Browser Modal Dialog Console Logs</sub>

Note we have logs for all use cases