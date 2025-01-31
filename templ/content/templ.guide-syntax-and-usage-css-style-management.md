CSS style management | templ docs
===============

[Skip to main content](https://templ.guide/syntax-and-usage/css-style-management#__docusaurus_skipToContent_fallback)

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
*   CSS style management

On this page

CSS style management
====================

HTML class attribute[​](https://templ.guide/syntax-and-usage/css-style-management#html-class-attribute "Direct link to HTML class attribute")
---------------------------------------------------------------------------------------------------------------------------------------------

The standard HTML `class` attribute can be added to components to set class names.

```
templ button(text string) {	<button class="button is-primary">{ text }</button>}
```

Output

```
<button class="button is-primary"> Click me</button>
```

Class expression[​](https://templ.guide/syntax-and-usage/css-style-management#class-expression "Direct link to Class expression")
---------------------------------------------------------------------------------------------------------------------------------

To use a variable as the name of a CSS class, use a CSS expression.

component.templ

```
package maintempl button(text string, className string) {	<button class={ className }>{ text }</button>}
```

The class expression can take an array of values.

component.templ

```
package maintempl button(text string, className string) {	<button class={ "button", className }>{ text }</button>}
```

### Dynamic class names[​](https://templ.guide/syntax-and-usage/css-style-management#dynamic-class-names "Direct link to Dynamic class names")

Toggle addition of CSS classes to an element based on a boolean value by passing:

*   A `templ.KV` value containing the name of the class to add to the element, and a boolean that determines whether the class is added to the attribute at render time.
    *   `templ.KV("is-primary", true)`
    *   `templ.KV("hover:red", true)`
*   A map of string class names to a boolean that determines if the class is added to the class attribute value at render time:
    *   `map[string]bool`
    *   `map[CSSClass]bool`

component.templ

```
package maincss red() {	background-color: #ff0000;}templ button(text string, isPrimary bool) {	<button class={ "button", templ.KV("is-primary", isPrimary), templ.KV(red(), isPrimary) }>{ text }</button>}
```

main.go

```
package mainimport (	"context"	"os")func main() {	button("Click me", false).Render(context.Background(), os.Stdout)}
```

Output

```
<button class="button"> Click me</button>
```

CSS elements[​](https://templ.guide/syntax-and-usage/css-style-management#css-elements "Direct link to CSS elements")
---------------------------------------------------------------------------------------------------------------------

The standard `<style>` element can be used within a template.

`<style>` element contents are rendered to the output without any changes.

```
templ page() {	<style type="text/css">		p {			font-family: sans-serif;		}		.button {			background-color: black;			foreground-color: white;		}	</style>	<p>		Paragraph contents.	</p>}
```

Output

```
<style type="text/css">	p {		font-family: sans-serif;	}	.button {		background-color: black;		foreground-color: white;	}</style><p>	Paragraph contents.</p>
```

tip

If you want to make sure that the CSS element is only output once, even if you use a template many times, use a CSS expression.

CSS components[​](https://templ.guide/syntax-and-usage/css-style-management#css-components "Direct link to CSS components")
---------------------------------------------------------------------------------------------------------------------------

When developing a component library, it may not be desirable to require that specific CSS classes are present when the HTML is rendered.

There may be CSS class name clashes, or developers may forget to include the required CSS.

To include CSS within a component library, use a CSS component.

CSS components can also be conditionally rendered.

component.templ

```
package mainvar red = "#ff0000"var blue = "#0000ff"css primaryClassName() {	background-color: #ffffff;	color: { red };}css className() {	background-color: #ffffff;	color: { blue };}templ button(text string, isPrimary bool) {	<button class={ "button", className(), templ.KV(primaryClassName(), isPrimary) }>{ text }</button>}
```

Output

```
<style type="text/css"> .className_f179{background-color:#ffffff;color:#ff0000;}</style><button class="button className_f179"> Click me</button>
```

info

The CSS class is given a unique name the first time it is used, and only rendered once per HTTP request to save bandwidth.

caution

The class name is autogenerated, don't rely on it being consistent.

### CSS component arguments[​](https://templ.guide/syntax-and-usage/css-style-management#css-component-arguments "Direct link to CSS component arguments")

CSS components can also require function arguments.

component.templ

```
package maincss loading(percent int) {	width: { fmt.Sprintf("%d%%", percent) };}templ index() {    <div class={ loading(50) }></div>    <div class={ loading(100) }></div>}
```

Output

```
<style type="text/css"> .loading_a3cc{width:50%;}</style><div class="loading_a3cc"></div><style type="text/css"> .loading_9ccc{width:100%;}</style><div class="loading_9ccc"></div>
```

### CSS Sanitization[​](https://templ.guide/syntax-and-usage/css-style-management#css-sanitization "Direct link to CSS Sanitization")

To prevent CSS injection attacks, templ automatically sanitizes dynamic CSS property names and values using the `templ.SanitizeCSS` function. Internally, this uses a lightweight fork of Google's `safehtml` package to sanitize the value.

If a property name or value has been sanitized, it will be replaced with `zTemplUnsafeCSSPropertyName` for property names, or `zTemplUnsafeCSSPropertyValue` for property values.

To bypass this sanitization, e.g. for URL values of `background-image`, you can mark the value as safe using the `templ.SafeCSSProperty` type.

```
css windVaneRotation(degrees float64) {	transform: { templ.SafeCSSProperty(fmt.Sprintf("rotate(%ddeg)", int(math.Round(degrees)))) };}templ Rotate(degrees float64) {	<div class={ windVaneRotation(degrees) }>Rotate</div>}
```

### CSS Middleware[​](https://templ.guide/syntax-and-usage/css-style-management#css-middleware "Direct link to CSS Middleware")

The use of CSS templates means that `<style>` elements containing the CSS are rendered on each HTTP request.

To save bandwidth, templ can provide a global stylesheet that includes the output of CSS templates instead of including `<style>` tags in each HTTP request.

To provide a global stylesheet, use templ's CSS middleware, and register templ classes on application startup.

The middleware adds a HTTP route to the web server (`/styles/templ.css` by default) that renders the `text/css` classes that would otherwise be added to `<style>` tags when components are rendered.

For example, to stop the `className` CSS class from being added to the output, the HTTP middleware can be used.

```
c1 := className()handler := NewCSSMiddleware(httpRoutes, c1)http.ListenAndServe(":8000", handler)
```

caution

Don't forget to add a `<link rel="stylesheet" href="/styles/templ.css">` to your HTML to include the generated CSS class names!

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/03-syntax-and-usage/11-css-style-management.md)

[Previous Template composition](https://templ.guide/syntax-and-usage/template-composition)[Next Using JavaScript with templ](https://templ.guide/syntax-and-usage/script-templates)

*   [HTML class attribute](https://templ.guide/syntax-and-usage/css-style-management#html-class-attribute)
*   [Class expression](https://templ.guide/syntax-and-usage/css-style-management#class-expression)
    *   [Dynamic class names](https://templ.guide/syntax-and-usage/css-style-management#dynamic-class-names)
*   [CSS elements](https://templ.guide/syntax-and-usage/css-style-management#css-elements)
*   [CSS components](https://templ.guide/syntax-and-usage/css-style-management#css-components)
    *   [CSS component arguments](https://templ.guide/syntax-and-usage/css-style-management#css-component-arguments)
    *   [CSS Sanitization](https://templ.guide/syntax-and-usage/css-style-management#css-sanitization)
    *   [CSS Middleware](https://templ.guide/syntax-and-usage/css-style-management#css-middleware)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.