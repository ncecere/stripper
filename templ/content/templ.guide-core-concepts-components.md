Components | templ docs
===============

[Skip to main content](https://templ.guide/core-concepts/components#__docusaurus_skipToContent_fallback)

[![Image 1: Templ Logo](https://templ.guide/img/logo.svg)![Image 2: Templ Logo](https://templ.guide/img/logo.svg)](https://templ.guide/)[Docs](https://templ.guide/)

[GitHub](https://github.com/a-h/templ)

Search

*   [Introduction](https://templ.guide/)
*   [Quick start](https://templ.guide/quick-start/installation)
    
*   [Syntax and usage](https://templ.guide/syntax-and-usage/basic-syntax)
    
*   [Core concepts](https://templ.guide/core-concepts/components)
    
    *   [Components](https://templ.guide/core-concepts/components)
    *   [Template generation](https://templ.guide/core-concepts/template-generation)
    *   [Testing](https://templ.guide/core-concepts/testing)
    *   [View models](https://templ.guide/core-concepts/view-models)
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
*   Core concepts
*   Components

On this page

Components
==========

templ Components are markup and code that is compiled into functions that return a `templ.Component` interface by running the `templ generate` command.

Components can contain templ elements that render HTML, text, expressions that output text or include other templates, and branching statements such as `if` and `switch`, and `for` loops.

header.templ

```
package maintempl headerTemplate(name string) {  <header data-testid="headerTemplate">    <h1>{ name }</h1>  </header>}
```

The generated code is a Go function that returns a `templ.Component`.

header\_templ.go

```
func headerTemplate(name string) templ.Component {  // Generated contents}
```

`templ.Component` is an interface that has a `Render` method on it that is used to render the component to an `io.Writer`.

```
type Component interface {	Render(ctx context.Context, w io.Writer) error}
```

tip

Since templ produces Go code, you can share templates the same way that you share Go code - by sharing your Go module.

templ follows the same rules as Go. If a `templ` block starts with an uppercase letter, then it is public, otherwise, it is private.

A `templ.Component` may write partial output to the `io.Writer` if it returns an error. If you want to ensure you only get complete output or nothing, write to a buffer first and then write the buffer to an `io.Writer`.

Code-only components[​](https://templ.guide/core-concepts/components#code-only-components "Direct link to Code-only components")
--------------------------------------------------------------------------------------------------------------------------------

Since templ Components ultimately implement the `templ.Component` interface, any code that implements the interface can be used in place of a templ component generated from a `*.templ` file.

```
package mainimport (	"context"	"io"	"os"	"github.com/a-h/templ")func button(text string) templ.Component {	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {		_, err := io.WriteString(w, "<button>"+text+"</button>")		return err	})}func main() {	button("Click me").Render(context.Background(), os.Stdout)}
```

Output

```
<button> Click me</button>
```

warning

This code is unsafe! In code-only components, you're responsible for escaping the HTML content yourself, e.g. with the `templ.EscapeString` function.

Method components[​](https://templ.guide/core-concepts/components#method-components "Direct link to Method components")
-----------------------------------------------------------------------------------------------------------------------

templ components can be returned from methods (functions attached to types).

Go code:

```
package mainimport "os"type Data struct {	message string}templ (d Data) Method() {	<div>{ d.message }</div>}func main() {	d := Data{		message: "You can implement methods on a type.",	}	d.Method().Render(context.Background(), os.Stdout)}
```

It is also possible to initialize a struct and call its component method inline.

```
package mainimport "os"type Data struct {	message string}templ (d Data) Method() {	<div>{ d.message }</div>}templ Message() {    <div>        @Data{            message: "You can implement methods on a type.",        }.Method()    </div>}func main() {	Message().Render(context.Background(), os.Stdout)}
```

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/04-core-concepts/01-components.md)

[Previous Render once](https://templ.guide/syntax-and-usage/render-once)[Next Template generation](https://templ.guide/core-concepts/template-generation)

*   [Code-only components](https://templ.guide/core-concepts/components#code-only-components)
*   [Method components](https://templ.guide/core-concepts/components#method-components)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.