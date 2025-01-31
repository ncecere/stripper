Using JavaScript with templ | templ docs
===============

[Skip to main content](https://templ.guide/syntax-and-usage/script-templates#__docusaurus_skipToContent_fallback)

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
*   Using JavaScript with templ

On this page

Using JavaScript with templ
===========================

Script tags[​](https://templ.guide/syntax-and-usage/script-templates#script-tags "Direct link to Script tags")
--------------------------------------------------------------------------------------------------------------

Use standard `<script>` tags, and standard HTML attributes to run JavaScript on the client.

```
templ body() {  <script type="text/javascript">    function handleClick(event) {      alert(event + ' clicked');    }  </script>  <button onclick="handleClick(this)">Click me</button>}
```

To pass data from the server to client-side scripts, see [Passing server-side data to scripts](https://templ.guide/syntax-and-usage/script-templates#passing-server-side-data-to-scripts).

Adding client side behaviours to components[​](https://templ.guide/syntax-and-usage/script-templates#adding-client-side-behaviours-to-components "Direct link to Adding client side behaviours to components")
--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

To ensure that a `<script>` tag within a templ component is only rendered once per HTTP response, use a [templ.OnceHandle](https://templ.guide/syntax-and-usage/render-once).

Using a `templ.OnceHandle` allows a component to define global client-side scripts that it needs to run without including the scripts multiple times in the response.

The example below also demonstrates applying behaviour that's defined in a multiline script to its sibling element.

component.templ

```
package mainimport "net/http"var helloHandle = templ.NewOnceHandle()templ hello(label, name string) {  // This script is only rendered once per HTTP request.  @helloHandle.Once() {    <script type="text/javascript">      function hello(name) {        alert('Hello, ' + name + '!');      }    </script>  }  <div>    <input type="button" value={ label } data-name={ name }/>    <script type="text/javascript">      // To prevent the variables from leaking into the global scope,      // this script is wrapped in an IIFE (Immediately Invoked Function Expression).      (() => {        let scriptElement = document.currentScript;        let parent = scriptElement.closest('div');        let nearestButtonWithName = parent.querySelector('input[data-name]');        nearestButtonWithName.addEventListener('click', function() {          let name = nearestButtonWithName.getAttribute('data-name');          hello(name);        })      })()    </script>  </div>}templ page() {  @hello("Hello User", "user")  @hello("Hello World", "world")}func main() {  http.Handle("/", templ.Handler(page()))  http.ListenAndServe("127.0.0.1:8080", nil)}
```

tip

You might find libraries like [surreal](https://github.com/gnat/surreal) useful for reducing boilerplate.

```
var helloHandle = templ.NewOnceHandle()var surrealHandle = templ.NewOnceHandle()templ hello(label, name string) {  @helloHandle.Once() {    <script type="text/javascript">      function hello(name) {        alert('Hello, ' + name + '!');      }    </script>  }  @surrealHandle.Once() {    <script src="https://cdn.jsdelivr.net/gh/gnat/surreal@3b4572dd0938ce975225ee598a1e7381cb64ffd8/surreal.js"></script>  }  <div>    <input type="button" value={ label } data-name={ name }/>    <script type="text/javascript">      // me("-") returns the previous sibling element.      me("-").addEventListener('click', function() {        let name = this.getAttribute('data-name');        hello(name);      })    </script>  </div>}
```

Importing scripts[​](https://templ.guide/syntax-and-usage/script-templates#importing-scripts "Direct link to Importing scripts")
--------------------------------------------------------------------------------------------------------------------------------

Use standard `<script>` tags to load JavaScript from a URL.

```
templ head() {	<head>		<script src="https://unpkg.com/lightweight-charts/dist/lightweight-charts.standalone.production.js"></script>	</head>}
```

And use the imported JavaScript directly in templ via `<script>` tags.

```
templ body() {	<script>		const chart = LightweightCharts.createChart(document.body, { width: 400, height: 300 });		const lineSeries = chart.addLineSeries();		lineSeries.setData([				{ time: '2019-04-11', value: 80.01 },				{ time: '2019-04-12', value: 96.63 },				{ time: '2019-04-13', value: 76.64 },				{ time: '2019-04-14', value: 81.89 },				{ time: '2019-04-15', value: 74.43 },				{ time: '2019-04-16', value: 80.01 },				{ time: '2019-04-17', value: 96.63 },				{ time: '2019-04-18', value: 76.64 },				{ time: '2019-04-19', value: 81.89 },				{ time: '2019-04-20', value: 74.43 },		]);	</script>}
```

tip

You can use a CDN to serve 3rd party scripts, or serve your own and 3rd party scripts from your server using a `http.FileServer`.

```
mux := http.NewServeMux()mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))http.ListenAndServe("localhost:8080", mux)
```

Passing server-side data to scripts[​](https://templ.guide/syntax-and-usage/script-templates#passing-server-side-data-to-scripts "Direct link to Passing server-side data to scripts")
--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Pass data from the server to the client by embedding it in the HTML as a JSON object in an attribute or script tag.

### Pass server-side data to the client in a HTML attribute[​](https://templ.guide/syntax-and-usage/script-templates#pass-server-side-data-to-the-client-in-a-html-attribute "Direct link to Pass server-side data to the client in a HTML attribute")

input.templ

```
templ body(data any) {  <button id="alerter" alert-data={ templ.JSONString(data) }>Show alert</button>}
```

output.html

```
<button id="alerter" alert-data="{&quot;msg&quot;:&quot;Hello, from the attribute data&quot;}">Show alert</button>
```

The data in the attribute can then be accessed from client-side JavaScript.

```
const button = document.getElementById('alerter');const data = JSON.parse(button.getAttribute('alert-data'));
```

### Pass server-side data to the client in a script element[​](https://templ.guide/syntax-and-usage/script-templates#pass-server-side-data-to-the-client-in-a-script-element "Direct link to Pass server-side data to the client in a script element")

input.templ

```
templ body(data any) {  @templ.JSONScript("id", data)}
```

output.html

```
<script id="id" type="application/json">{"msg":"Hello, from the script data"}</script>
```

The data in the script tag can then be accessed from client-side JavaScript.

```
const data = JSON.parse(document.getElementById('id').textContent);
```

Working with NPM projects[​](https://templ.guide/syntax-and-usage/script-templates#working-with-npm-projects "Direct link to Working with NPM projects")
--------------------------------------------------------------------------------------------------------------------------------------------------------

[https://github.com/a-h/templ/tree/main/examples/typescript](https://github.com/a-h/templ/tree/main/examples/typescript) contains a TypeScript example that uses `esbuild` to transpile TypeScript into plain JavaScript, along with any required `npm` modules.

After transpilation and bundling, the output JavaScript code can be used in a web page by including a `<script>` tag.

### Creating a TypeScript project[​](https://templ.guide/syntax-and-usage/script-templates#creating-a-typescript-project "Direct link to Creating a TypeScript project")

Create a new TypeScript project with `npm`, and install TypeScript and `esbuild` as development dependencies.

```
mkdir tscd tsnpm initnpm install --save-dev typescript esbuild
```

Create a `src` directory to hold the TypeScript code.

```
mkdir src
```

And add a TypeScript file to the `src` directory.

ts/src/index.ts

```
function hello() {  console.log('Hello, from TypeScript');}
```

### Bundling TypeScript code[​](https://templ.guide/syntax-and-usage/script-templates#bundling-typescript-code "Direct link to Bundling TypeScript code")

Add a script to build the TypeScript code in `index.ts` and copy it to an output directory (in this case `./assets/js/index.js`).

ts/package.json

```
{  "name": "ts",  "version": "1.0.0",  "scripts": {    "build": "esbuild --bundle --minify --outfile=../assets/js/index.js ./src/index.ts"  },  "devDependencies": {    "esbuild": "0.21.3",    "typescript": "^5.4.5"  }}
```

After running `npm build` in the `ts` directory, the TypeScript code is transpiled into JavaScript and copied to the output directory.

### Using the output JavaScript[​](https://templ.guide/syntax-and-usage/script-templates#using-the-output-javascript "Direct link to Using the output JavaScript")

The output file `../assets/js/index.js` can then be used in a templ project.

components/head.templ

```
templ head() {	<head>		<script src="/assets/js/index.js"></script>	</head>}
```

You will need to configure your Go web server to serve the static content.

main.go

```
func main() {	mux := http.NewServeMux()	// Serve the JS bundle.	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))	// Serve components.	data := map[string]any{"msg": "Hello, World!"}	h := templ.Handler(components.Page(data))	mux.Handle("/", h)	fmt.Println("Listening on http://localhost:8080")	http.ListenAndServe("localhost:8080", mux)}
```

Script templates[​](https://templ.guide/syntax-and-usage/script-templates#script-templates "Direct link to Script templates")
-----------------------------------------------------------------------------------------------------------------------------

warning

Script templates are a legacy feature and are not recommended for new projects. Use standard `<script>` tags to import a standalone JavaScript file, optionally created by a bundler like `esbuild`.

If you need to pass Go data to scripts, you can use a script template.

Here, the `page` HTML template includes a `script` element that loads a charting library, which is then used by the `body` element to render some data.

```
package mainscript graph(data []TimeValue) {	const chart = LightweightCharts.createChart(document.body, { width: 400, height: 300 });	const lineSeries = chart.addLineSeries();	lineSeries.setData(data);}templ page(data []TimeValue) {	<html>		<head>			<script src="https://unpkg.com/lightweight-charts/dist/lightweight-charts.standalone.production.js"></script>		</head>		<body onload={ graph(data) }></body>	</html>}
```

The data is loaded by the backend into the template. This example uses a constant, but it could easily have collected the `[]TimeValue` from a database.

main.go

```
package mainimport (	"fmt"	"log"	"net/http")type TimeValue struct {	Time  string  `json:"time"`	Value float64 `json:"value"`}func main() {	mux := http.NewServeMux()	// Handle template.	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {		data := []TimeValue{			{Time: "2019-04-11", Value: 80.01},			{Time: "2019-04-12", Value: 96.63},			{Time: "2019-04-13", Value: 76.64},			{Time: "2019-04-14", Value: 81.89},			{Time: "2019-04-15", Value: 74.43},			{Time: "2019-04-16", Value: 80.01},			{Time: "2019-04-17", Value: 96.63},			{Time: "2019-04-18", Value: 76.64},			{Time: "2019-04-19", Value: 81.89},			{Time: "2019-04-20", Value: 74.43},		}		page(data).Render(r.Context(), w)	})	// Start the server.	fmt.Println("listening on :8080")	if err := http.ListenAndServe(":8080", mux); err != nil {		log.Printf("error listening: %v", err)	}}
```

`script` elements are templ Components, so you can also directly render the Javascript function, passing in Go data, using the `@` expression:

```
package mainimport "fmt"script printToConsole(content string) {	console.log(content)}templ page(content string) {	<html>		<body>		  @printToConsole(content)		  @printToConsole(fmt.Sprintf("Again: %s", content))		</body>	</html>}
```

The data passed into the Javascript funtion will be JSON encoded, which then can be used inside the function.

main.go

```
package mainimport (	"fmt"	"log"	"net/http"	"time")func main() {	mux := http.NewServeMux()	// Handle template.	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {		// Format the current time and pass it into our template		page(time.Now().String()).Render(r.Context(), w)	})	// Start the server.	fmt.Println("listening on :8080")	if err := http.ListenAndServe(":8080", mux); err != nil {		log.Printf("error listening: %v", err)	}}
```

After building and running the executable, running `curl http://localhost:8080/` would render:

Output

```
<html>	<body>		<script type="text/javascript">function __templ_printToConsole_5a85(content){console.log(content)}</script>		<script type="text/javascript">__templ_printToConsole_5a85("2023-11-11 01:01:40.983381358 +0000 UTC")</script>		<script type="text/javascript">__templ_printToConsole_5a85("Again: 2023-11-11 01:01:40.983381358 +0000 UTC")</script>	</body></html>
```

The `JSExpression` type is used to pass arbitrary JavaScript expressions to a templ script template.

A common use case is to pass the `event` or `this` objects to an event handler.

```
package mainscript showButtonWasClicked(event templ.JSExpression) {	const originalButtonText = event.target.innerText	event.target.innerText = "I was Clicked!"	setTimeout(() => event.target.innerText = originalButtonText, 2000)}templ page() {	<html>		<body>			<button type="button" onclick={ showButtonWasClicked(templ.JSExpression("event")) }>Click Me</button>		</body>	</html>}
```

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/03-syntax-and-usage/12-script-templates.md)

[Previous CSS style management](https://templ.guide/syntax-and-usage/css-style-management)[Next Comments](https://templ.guide/syntax-and-usage/comments)

*   [Script tags](https://templ.guide/syntax-and-usage/script-templates#script-tags)
*   [Adding client side behaviours to components](https://templ.guide/syntax-and-usage/script-templates#adding-client-side-behaviours-to-components)
*   [Importing scripts](https://templ.guide/syntax-and-usage/script-templates#importing-scripts)
*   [Passing server-side data to scripts](https://templ.guide/syntax-and-usage/script-templates#passing-server-side-data-to-scripts)
    *   [Pass server-side data to the client in a HTML attribute](https://templ.guide/syntax-and-usage/script-templates#pass-server-side-data-to-the-client-in-a-html-attribute)
    *   [Pass server-side data to the client in a script element](https://templ.guide/syntax-and-usage/script-templates#pass-server-side-data-to-the-client-in-a-script-element)
*   [Working with NPM projects](https://templ.guide/syntax-and-usage/script-templates#working-with-npm-projects)
    *   [Creating a TypeScript project](https://templ.guide/syntax-and-usage/script-templates#creating-a-typescript-project)
    *   [Bundling TypeScript code](https://templ.guide/syntax-and-usage/script-templates#bundling-typescript-code)
    *   [Using the output JavaScript](https://templ.guide/syntax-and-usage/script-templates#using-the-output-javascript)
*   [Script templates](https://templ.guide/syntax-and-usage/script-templates#script-templates)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.