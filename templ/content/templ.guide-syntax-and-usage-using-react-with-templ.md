Using React with templ | templ docs
===============

[Skip to main content](https://templ.guide/syntax-and-usage/using-react-with-templ#__docusaurus_skipToContent_fallback)

[![Image 1: Templ Logo](https://templ.guide/img/logo.svg)![Image 2: Templ Logo](https://templ.guide/img/logo.svg)](https://templ.guide/)[Docs](https://templ.guide/)

[GitHub](https://github.com/a-h/templ)

Search

*   [Introduction](https://templ.guide/)
*   [Quick start](https://templ.guide/quick-start/installation)
    
*   [Syntax and usage](https://templ.guide/syntax-and-usage/basic-syntax)
    
    *   [Basic syntax](https://templ.guide/syntax-and-usage/basic-syntax)
    *   [Elements](https://templ.guide/syntax-and-usage/elements)
    *   [Attributes](https://templ.guide/syntax-and-usage/attributes)
    *   [Expressions](https://templ.guide/syntax-and-usage/expressions)
    *   [Statements](https://templ.guide/syntax-and-usage/statements)
    *   [If/else](https://templ.guide/syntax-and-usage/if-else)
    *   [Switch](https://templ.guide/syntax-and-usage/switch)
    *   [For loops](https://templ.guide/syntax-and-usage/loops)
    *   [Raw Go](https://templ.guide/syntax-and-usage/raw-go)
    *   [Template composition](https://templ.guide/syntax-and-usage/template-composition)
    *   [CSS style management](https://templ.guide/syntax-and-usage/css-style-management)
    *   [Using JavaScript with templ](https://templ.guide/syntax-and-usage/script-templates)
    *   [Comments](https://templ.guide/syntax-and-usage/comments)
    *   [Context](https://templ.guide/syntax-and-usage/context)
    *   [Using with html/template](https://templ.guide/syntax-and-usage/using-with-go-templates)
    *   [Rendering raw HTML](https://templ.guide/syntax-and-usage/rendering-raw-html)
    *   [Using React with templ](https://templ.guide/syntax-and-usage/using-react-with-templ)
    *   [Render once](https://templ.guide/syntax-and-usage/render-once)
*   [Core concepts](https://templ.guide/core-concepts/components)
    
*   [Server-side rendering](https://templ.guide/server-side-rendering/creating-an-http-server-with-templ)
    
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
*   Syntax and usage
*   Using React with templ

On this page

Using React with templ
======================

templ is great for server-side rendering. Combined with [HTMX](https://htmx.org/), it's even more powerful, since HTMX can be used to replace elements within the page with updated HTML fetched from the server, providing many of the benefits of React with reduced overall complexity. See [/server-side-rendering/htmx](https://templ.guide/server-side-rendering/htmx) for an example.

However, React has a huge ecosystem of rich interactive components, so being able to tap into the ecosystem is very useful.

With templ, it's more likely that you will use React components as [islands of interactivity](https://www.patterns.dev/vanilla/islands-architecture/) rather than taking over all aspects of displaying your app, with templ taking over server-side rendering, but using React to provide specific features on the client side.

Using React components[​](https://templ.guide/syntax-and-usage/using-react-with-templ#using-react-components "Direct link to Using React components")
-----------------------------------------------------------------------------------------------------------------------------------------------------

First, lets start by rendering simple React components.

### Create React components[​](https://templ.guide/syntax-and-usage/using-react-with-templ#create-react-components "Direct link to Create React components")

To use React components in your templ app, create your React components using TSX (TypeScript) or JSX as usual.

react/components.tsx

```
export const Header = () => (<h1>React component Header</h1>);export const Body = () => (<div>This is client-side content from React</div>);
```

### Create a templ page[​](https://templ.guide/syntax-and-usage/using-react-with-templ#create-a-templ-page "Direct link to Create a templ page")

Next, use templ to create a page containing HTML elements with specific IDs.

note

This page defines elements with ids of `react-header` and `react-content`.

A `<script>` element loads in a JavaScript bundle that we haven't created yet.

components.templ

```
package maintempl page() {	<html>		<body>			<div id="react-header"></div>			<div id="react-content"></div>			<div>This is server-side content from templ.</div>			<!-- Load the React bundle created using esbuild -->			<script src="static/index.js"></script>		</body>	</html>}
```

tip

Remember to run `templ generate` when you've finished writing your templ file.

### Render React components into the IDs[​](https://templ.guide/syntax-and-usage/using-react-with-templ#render-react-components-into-the-ids "Direct link to Render React components into the IDs")

Write TypeScript or JavaScript to render the React components into the HTML elements that are rendered by templ.

react/index.ts

```
import { createRoot } from 'react-dom/client';import { Header, Body } from './components';// Render the React component into the templ page at the react-header.const headerRoot = document.getElementById('react-header');if (!headerRoot) {	throw new Error('Could not find element with id react-header');}const headerReactRoot = createRoot(headerRoot);headerReactRoot.render(Header());// Add the body React component.const contentRoot = document.getElementById('react-content');if (!contentRoot) {	throw new Error('Could not find element with id react-content');}const contentReactRoot = createRoot(contentRoot);contentReactRoot.render(Body());
```

### Create a client-side bundle[​](https://templ.guide/syntax-and-usage/using-react-with-templ#create-a-client-side-bundle "Direct link to Create a client-side bundle")

To turn the JSX, TSX, TypeScript and JavaScript code into a bundle that can run in the browser, use a bundling tool.

[https://esbuild.github.io/](https://esbuild.github.io/) is commonly used for this task. It's fast, it's easy to use, and it's written in Go.

Executing `esbuild` with the following arguments creates an `index.js` file in the static directory.

```
esbuild --bundle index.ts --outdir=../static --minify
```

### Serve the templ component and client side bundle[​](https://templ.guide/syntax-and-usage/using-react-with-templ#serve-the-templ-component-and-client-side-bundle "Direct link to Serve the templ component and client side bundle")

To serve the server-side rendered templ template, and the client-side JavaScript bundle created in the previous step, setup a Go web server.

main.go

```
package mainimport (	"fmt"	"log"	"net/http"	"github.com/a-h/templ")func main() {	mux := http.NewServeMux()	// Serve the templ page.	mux.Handle("/", templ.Handler(page()))	// Serve static content.	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))	// Start the server.	fmt.Println("listening on localhost:8080")	if err := http.ListenAndServe("localhost:8080", mux); err != nil {		log.Printf("error listening: %v", err)	}}
```

### Results[​](https://templ.guide/syntax-and-usage/using-react-with-templ#results "Direct link to Results")

Putting this together results in a web page that renders server-side HTML using templ. The server-side HTML includes a link to the static React bundle.

Passing server-side data to React components[​](https://templ.guide/syntax-and-usage/using-react-with-templ#passing-server-side-data-to-react-components "Direct link to Passing server-side data to React components")
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Moving on from the previous example, it's possible to pass data to client-side React components.

### Add a React component that accepts data arguments[​](https://templ.guide/syntax-and-usage/using-react-with-templ#add-a-react-component-that-accepts-data-arguments "Direct link to Add a React component that accepts data arguments")

First, add a new component.

react/components.tsx

```
export const Hello = (name: string) => (  <div>Hello {name} (Client-side React, rendering server-side data)</div>);
```

### Export a JavaScript function that renders the React component to a HTML element[​](https://templ.guide/syntax-and-usage/using-react-with-templ#export-a-javascript-function-that-renders-the-react-component-to-a-html-element "Direct link to Export a JavaScript function that renders the React component to a HTML element")

react/index.ts

```
// Update the import to add the new Hello React component.import { Header, Body, Hello } from './components';// Previous script contents...  export function renderHello(e: HTMLElement) {  const name = e.getAttribute('data-name') ?? "";  createRoot(e).render(Hello(name));}
```

### Update the templ component to use the new function[​](https://templ.guide/syntax-and-usage/using-react-with-templ#update-the-templ-component-to-use-the-new-function "Direct link to Update the templ component to use the new function")

Now that we have a `renderHello` function that will render the React component to the given element, we can update the templ components to use it.

In templ, we can add a `Hello` component that does two things:

1.  Renders an element for the React component to be loaded into that sets the `data-name` attribute to the value of the server-side `name` field.
2.  Writes out JS that calls the `renderHello` function to mount the React component into the element.

note

The template renders three copies of the `Hello` React component, passing in a distinct `name` parameter ("Alice", "Bob" and "Charlie").

components.templ

```
package mainimport "fmt"templ Hello(name string) {	<div data-name={ name }>		<script type="text/javascript">			bundle.renderHello(document.currentScript.closest('div'));		</script>	</div>}templ page() {	<html>		<head>			<title>React integration</title>		</head>		<body>			<div id="react-header"></div>			<div id="react-content"></div>			<div>				This is server-side content from templ.			</div>			<!-- Load the React bundle that was created using esbuild -->			<!-- Since the bundle was coded to expect the react-header and react-content elements to exist already, in this case, the script has to be loaded after the elements are on the page -->			<script src="static/index.js"></script>			<!-- Now that the React bundle is loaded, we can use the functions that are in it -->			<!-- the renderName function in the bundle can be used, but we want to pass it some server-side data -->			for _, name := range []string{"Alice", "Bob", "Charlie"} {				@Hello(name)			}		</body>	</html>}
```

### Update the `esbuild` command[​](https://templ.guide/syntax-and-usage/using-react-with-templ#update-the-esbuild-command "Direct link to update-the-esbuild-command")

The `bundle` namespace in JavaScript is created by adding a `--global-name` argument to `esbuild`. The argument causes any exported functions in `index.ts` to be added to that namespace.

```
esbuild --bundle index.ts --outdir=../static --minify --global-name=bundle
```

### Results[​](https://templ.guide/syntax-and-usage/using-react-with-templ#results-1 "Direct link to Results")

The HTML that's rendered is:

```
<html>  <head>    <title>React integration</title>  </head>  <body>    <div id="react-header"></div>    <div id="react-content"></div>    <div>This is server-side content from templ.</div>    <script src="static/index.js"></script>    <div data-name="Alice">      <script type="text/javascript">        // Place the React component into the parent div.        bundle.renderHello(document.currentScript.closest('div'));      </script>    </div>    <div data-name="Bob">      <script type="text/javascript">        // Place the React component into the parent div.	bundle.renderHello(document.currentScript.closest('div'));      </script>    </div>    <div data-name="Charlie">      <script type="text/javascript">        // Place the React component into the parent div.	bundle.renderHello(document.currentScript.closest('div'));      </script>    </div>  </body></html>
```

And the browser shows the expected content after rendering the client side React components.

```
React component HeaderThis is client-side content from ReactThis is server-side content from templ.Hello Alice (Client-side React, rendering server-side data)Hello Bob (Client-side React, rendering server-side data)Hello Charlie (Client-side React, rendering server-side data)
```

Example code[​](https://templ.guide/syntax-and-usage/using-react-with-templ#example-code "Direct link to Example code")
-----------------------------------------------------------------------------------------------------------------------

See [https://github.com/a-h/templ/tree/main/examples/integration-react](https://github.com/a-h/templ/tree/main/examples/integration-react) for a complete example.

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/03-syntax-and-usage/17-using-react-with-templ.md)

[Previous Rendering raw HTML](https://templ.guide/syntax-and-usage/rendering-raw-html)[Next Render once](https://templ.guide/syntax-and-usage/render-once)

*   [Using React components](https://templ.guide/syntax-and-usage/using-react-with-templ#using-react-components)
    *   [Create React components](https://templ.guide/syntax-and-usage/using-react-with-templ#create-react-components)
    *   [Create a templ page](https://templ.guide/syntax-and-usage/using-react-with-templ#create-a-templ-page)
    *   [Render React components into the IDs](https://templ.guide/syntax-and-usage/using-react-with-templ#render-react-components-into-the-ids)
    *   [Create a client-side bundle](https://templ.guide/syntax-and-usage/using-react-with-templ#create-a-client-side-bundle)
    *   [Serve the templ component and client side bundle](https://templ.guide/syntax-and-usage/using-react-with-templ#serve-the-templ-component-and-client-side-bundle)
    *   [Results](https://templ.guide/syntax-and-usage/using-react-with-templ#results)
*   [Passing server-side data to React components](https://templ.guide/syntax-and-usage/using-react-with-templ#passing-server-side-data-to-react-components)
    *   [Add a React component that accepts data arguments](https://templ.guide/syntax-and-usage/using-react-with-templ#add-a-react-component-that-accepts-data-arguments)
    *   [Export a JavaScript function that renders the React component to a HTML element](https://templ.guide/syntax-and-usage/using-react-with-templ#export-a-javascript-function-that-renders-the-react-component-to-a-html-element)
    *   [Update the templ component to use the new function](https://templ.guide/syntax-and-usage/using-react-with-templ#update-the-templ-component-to-use-the-new-function)
    *   [Update the `esbuild` command](https://templ.guide/syntax-and-usage/using-react-with-templ#update-the-esbuild-command)
    *   [Results](https://templ.guide/syntax-and-usage/using-react-with-templ#results-1)
*   [Example code](https://templ.guide/syntax-and-usage/using-react-with-templ#example-code)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.