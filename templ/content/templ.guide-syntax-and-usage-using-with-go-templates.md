Using with html/template | templ docs
===============

[Skip to main content](https://templ.guide/syntax-and-usage/using-with-go-templates#__docusaurus_skipToContent_fallback)

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
*   Using with html/template

On this page

Using with `html/template`
==========================

Templ components can be used with the Go standard library [`html/template`](https://pkg.go.dev/html/template) package.

Using `html/template` in a templ component[​](https://templ.guide/syntax-and-usage/using-with-go-templates#using-htmltemplate-in-a-templ-component "Direct link to using-htmltemplate-in-a-templ-component")
------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

To use an existing `html/template` in a templ component, use the `templ.FromGoHTML` function.

component.templ

```
package testgotemplatesimport "html/template"var goTemplate = template.Must(template.New("example").Parse("<div>{{ . }}</div>"))templ Example() {	<!DOCTYPE html>	<html>		<body>			@templ.FromGoHTML(goTemplate, "Hello, World!")		</body>	</html>}
```

main.go

```
func main() {	Example.Render(context.Background(), os.Stdout)}
```

Output

```
<!DOCTYPE html><html>	<body>		<div>Hello, World!</div>	</body></html>
```

Using a templ component with `html/template`[​](https://templ.guide/syntax-and-usage/using-with-go-templates#using-a-templ-component-withhtmltemplate "Direct link to using-a-templ-component-withhtmltemplate")
----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

To use a templ component within a `html/template`, use the `templ.ToGoHTML` function to render the component into a `template.HTML value`.

component.html

```
package testgotemplatesimport "html/template"var example = template.Must(template.New("example").Parse(`<!DOCTYPE html><html>	<body>		{{ . }}	</body></html>`))templ greeting() {	<div>Hello, World!</div>}
```

main.go

```
func main() {	// Create the templ component.	templComponent := greeting()	// Render the templ component to a `template.HTML` value.	html, err := templ.ToGoHTML(context.Background(), templComponent)	if err != nil {		t.Fatalf("failed to convert to html: %v", err)	}	// Use the `template.HTML` value within the text/html template.	err = example.Execute(os.Stdout, html)	if err != nil {		t.Fatalf("failed to execute template: %v", err)	}}
```

Output

```
<!DOCTYPE html><html>	<body>		<div>Hello, World!</div>	</body></html>
```

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/03-syntax-and-usage/15-using-with-go-templates.md)

[Previous Context](https://templ.guide/syntax-and-usage/context)[Next Rendering raw HTML](https://templ.guide/syntax-and-usage/rendering-raw-html)

*   [Using `html/template` in a templ component](https://templ.guide/syntax-and-usage/using-with-go-templates#using-htmltemplate-in-a-templ-component)
*   [Using a templ component with `html/template`](https://templ.guide/syntax-and-usage/using-with-go-templates#using-a-templ-component-withhtmltemplate)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.