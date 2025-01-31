Creating an HTTP server with templ | templ docs
===============

[Skip to main content](https://templ.guide/server-side-rendering/creating-an-http-server-with-templ#__docusaurus_skipToContent_fallback)

[![Image 1: Templ Logo](https://templ.guide/img/logo.svg)![Image 2: Templ Logo](https://templ.guide/img/logo.svg)](https://templ.guide/)[Docs](https://templ.guide/)

[GitHub](https://github.com/a-h/templ)

Search

*   [Introduction](https://templ.guide/)
*   [Quick start](https://templ.guide/quick-start/installation)
    
*   [Syntax and usage](https://templ.guide/syntax-and-usage/basic-syntax)
    
*   [Core concepts](https://templ.guide/core-concepts/components)
    
*   [Server-side rendering](https://templ.guide/server-side-rendering/creating-an-http-server-with-templ)
    
    *   [Creating an HTTP server with templ](https://templ.guide/server-side-rendering/creating-an-http-server-with-templ)
    *   [Example: Counter application](https://templ.guide/server-side-rendering/example-counter-application)
    *   [HTMX](https://templ.guide/server-side-rendering/htmx)
    *   [Datastar](https://templ.guide/server-side-rendering/datastar)
    *   [HTTP Streaming](https://templ.guide/server-side-rendering/streaming)
*   [Static rendering](https://templ.guide/static-rendering/generating-static-html-files-with-templ)
    
*   [Project structure](https://templ.guide/project-structure/project-structure)
    
*   [Hosting and deployment](https://templ.guide/hosting-and-deployment/hosting-on-aws-lambda)
    
*   [Developer tools](https://templ.guide/developer-tools/cli)
    
*   [Security](https://templ.guide/security/injection-attacks)
    
*   [Media and talks](https://templ.guide/media/)
*   [Integrations](https://templ.guide/integrations/web-frameworks)
    
*   [Experimental](https://templ.guide/experimental/overview)
    
*   [Help and community](https://templ.guide/help-and-community/)
*   [FAQ](https://templ.guide/faq/)

*   [](https://templ.guide/)
*   Server-side rendering
*   Creating an HTTP server with templ

On this page

Creating an HTTP server with templ
==================================

### Static pages[​](https://templ.guide/server-side-rendering/creating-an-http-server-with-templ#static-pages "Direct link to Static pages")

To use a templ component as a HTTP handler, the `templ.Handler` function can be used.

This is suitable for use when the component is not used to display dynamic data.

components.templ

```
package maintempl hello() {	<div>Hello</div>}
```

main.go

```
package mainimport (	"net/http"	"github.com/a-h/templ")func main() {	http.Handle("/", templ.Handler(hello()))	http.ListenAndServe(":8080", nil)}
```

### Displaying fixed data[​](https://templ.guide/server-side-rendering/creating-an-http-server-with-templ#displaying-fixed-data "Direct link to Displaying fixed data")

In the previous example, the `hello` component does not take any parameters. Let's display the time when the server was started instead.

components.templ

```
package mainimport "time"templ timeComponent(d time.Time) {	<div>{ d.String() }</div>}templ notFoundComponent() {	<div>404 - Not found</div>}
```

main.go

```
package mainimport (	"net/http"	"time"	"github.com/a-h/templ")func main() {	http.Handle("/", templ.Handler(timeComponent(time.Now())))	http.Handle("/404", templ.Handler(notFoundComponent(), templ.WithStatus(http.StatusNotFound)))	http.ListenAndServe(":8080", nil)}
```

tip

The `templ.WithStatus`, `templ.WithContentType`, and `templ.WithErrorHandler` functions can be passed as parameters to the `templ.Handler` function to control how content is rendered.

The output will always be the date and time that the web server was started up, not the current time.

```
2023-04-26 08:40:03.421358 +0100 BST m=+0.000779501
```

To display the current time, we could update the component to use the `time.Now()` function itself, but this would limit the reusability of the component. It's better when components take parameters for their display values.

tip

Good templ components are idempotent, pure functions - they don't rely on data that is not passed in through parameters. As long as the parameters are the same, they always return the same HTML - they don't rely on any network calls or disk access.

Displaying dynamic data[​](https://templ.guide/server-side-rendering/creating-an-http-server-with-templ#displaying-dynamic-data "Direct link to Displaying dynamic data")
-------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Let's update the previous example to display dynamic content.

templ components implement the `templ.Component` interface, which provides a `Render` method.

The `Render` method can be used within HTTP handlers to write HTML to the `http.ResponseWriter`.

main.go

```
package mainimport (	"net/http")func main() {	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {		hello().Render(r.Context(), w)	})	http.ListenAndServe(":8080", nil)}
```

Building on that example, we can implement the Go HTTP handler interface and use the component within our HTTP handler. In this case, displaying the latest date and time, instead of the date and time when the server started up.

main.go

```
package mainimport (	"net/http"	"time")func NewNowHandler(now func() time.Time) NowHandler {	return NowHandler{Now: now}}type NowHandler struct {	Now func() time.Time}func (nh NowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {	timeComponent(nh.Now()).Render(r.Context(), w)}func main() {	http.Handle("/", NewNowHandler(time.Now))	http.ListenAndServe(":8080", nil)}
```

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/05-server-side-rendering/01-creating-an-http-server-with-templ.md)

[Previous View models](https://templ.guide/core-concepts/view-models)[Next Example: Counter application](https://templ.guide/server-side-rendering/example-counter-application)

*   [Static pages](https://templ.guide/server-side-rendering/creating-an-http-server-with-templ#static-pages)
*   [Displaying fixed data](https://templ.guide/server-side-rendering/creating-an-http-server-with-templ#displaying-fixed-data)
*   [Displaying dynamic data](https://templ.guide/server-side-rendering/creating-an-http-server-with-templ#displaying-dynamic-data)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.