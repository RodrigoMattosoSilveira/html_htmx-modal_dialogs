# Abstract
This repository introduces two different methods to display `HTML / HTMX` modal dialogs utilizing the native [HTML](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/dialog) `<dialog>` element.

In general, we use browser modal dialogs to prompt the user to confirm an action, such as deleting a record, and server modal dialogs to either inform the user about the result of an operation, like lack of authorization for access a resource-- or collected addition data, like a phone number.

We will demonstrate our approach using a simple web applications with a technology stack consisting of  [Go](https://go.dev/), [Fiber](https://gofiber.io/) and [Bootstrap 5](https://getbootstrap.com/) on the server, and [HTMX](https://htmx.org/) and [HyperScript](https://hyperscript.org/) to manipulate the DOM Tree, enabling a very limited use of Javascript. 

The first method demonstrates how to show a modal dialog triggered `browser logic`.  The second method demonstrates how to show a modal dialog triggered by `server logic`.

