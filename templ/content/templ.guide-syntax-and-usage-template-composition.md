Template composition | templ docs
===============

[Skip to main content](https://templ.guide/syntax-and-usage/template-composition#__docusaurus_skipToContent_fallback)

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
*   Template composition

On this page

Template composition
====================

Templates can be composed using the import expression.

```
templ showAll() {	@left()	@middle()	@right()}templ left() {	<div>Left</div>}templ middle() {	<div>Middle</div>}templ right() {	<div>Right</div>}
```

Output

```
<div> Left</div><div> Middle</div><div> Right</div>
```

Children[​](https://templ.guide/syntax-and-usage/template-composition#children "Direct link to Children")
---------------------------------------------------------------------------------------------------------

Children can be passed to a component for it to wrap.

```
templ showAll() {	@wrapChildren() {		<div>Inserted from the top</div>	}}templ wrapChildren() {	<div id="wrapper">		{ children... }	</div>}
```

note

The use of the `{ children... }` expression in the child component.

output

```
<div id="wrapper"> <div>  Inserted from the top </div></div>
```

### Using children in code components[​](https://templ.guide/syntax-and-usage/template-composition#using-children-in-code-components "Direct link to Using children in code components")

Children are passed to a component using the Go context. To pass children to a component using Go code, use the `templ.WithChildren` function.

```
package mainimport (  "context"  "os"  "github.com/a-h/templ")templ wrapChildren() {	<div id="wrapper">		{ children... }	</div>}func main() {  contents := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {    _, err := io.WriteString(w, "<div>Inserted from Go code</div>")    return err  })  ctx := templ.WithChildren(context.Background(), contents)  wrapChildren().Render(ctx, os.Stdout)}
```

output

```
<div id="wrapper"> <div>  Inserted from Go code </div></div>
```

To get children from the context, use the `templ.GetChildren` function.

```
package mainimport (  "context"  "os"  "github.com/a-h/templ")func main() {  contents := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {    _, err := io.WriteString(w, "<div>Inserted from Go code</div>")    return err  })  wrapChildren := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {    children := templ.GetChildren(ctx)    ctx = templ.ClearChildren(ctx)    _, err := io.WriteString(w, "<div id=\"wrapper\">")    if err != nil {      return err    }    err = children.Render(ctx, w)    if err != nil {      return err    }    _, err = io.WriteString(w, "</div>")    return err  })
```

note

The `templ.ClearChildren` function is used to stop passing the children down the tree.

Components as parameters[​](https://templ.guide/syntax-and-usage/template-composition#components-as-parameters "Direct link to Components as parameters")
---------------------------------------------------------------------------------------------------------------------------------------------------------

Components can also be passed as parameters and rendered using the `@component` expression.

```
package maintempl heading() {    <h1>Heading</h1>}templ layout(contents templ.Component) {	<div id="heading">		@heading()	</div>	<div id="contents">		@contents	</div>}templ paragraph(contents string) {	<p>{ contents }</p>}
```

main.go

```
package mainimport (	"context"	"os")func main() {	c := paragraph("Dynamic contents")	layout(c).Render(context.Background(), os.Stdout)}
```

output

```
<div id="heading">	<h1>Heading</h1></div><div id="contents">	<p>Dynamic contents</p></div>
```

You can pass `templ` components as parameters to other components within templates using standard Go function call syntax.

```
package maintempl heading() {    <h1>Heading</h1>}templ layout(contents templ.Component) {	<div id="heading">		@heading()	</div>	<div id="contents">		@contents	</div>}templ paragraph(contents string) {	<p>{ contents }</p>}templ root() {	@layout(paragraph("Dynamic contents"))}
```

main.go

```
package mainimport (	"context"	"os")func main() {	root().Render(context.Background(), os.Stdout)}
```

output

```
<div id="heading">	<h1>Heading</h1></div><div id="contents">	<p>Dynamic contents</p></div>
```

Joining Components[​](https://templ.guide/syntax-and-usage/template-composition#joining-components "Direct link to Joining Components")
---------------------------------------------------------------------------------------------------------------------------------------

Components can be aggregated into a single Component using `templ.Join`.

```
package maintempl hello() {	<span>hello</span>}templ world() {	<span>world</span>}templ helloWorld() {	@templ.Join(hello(), world())}
```

main.go

```
package mainimport (	"context"	"os")func main() {	helloWorld().Render(context.Background(), os.Stdout)}
```

output

```
<span>hello</span><span>world</span>
```

Sharing and re-using components[​](https://templ.guide/syntax-and-usage/template-composition#sharing-and-re-using-components "Direct link to Sharing and re-using components")
------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Since templ components are compiled into Go functions by the `go generate` command, templ components follow the rules of Go, and are shared in exactly the same way as Go code.

templ files in the same directory can access each other's components. Components in different directories can be accessed by importing the package that contains the component, so long as the component is exported by capitalizing its name.

tip

In Go, a _package_ is a collection of Go source files in the same directory that are compiled together. All of the functions, types, variables, and constants defined in one source file in a package are available to all other source files in the same package.

Packages exist within a Go _module_, defined by the `go.mod` file.

note

Go is structured differently to JavaScript, but uses similar terminology. A single `.js` or `.ts` _file_ is like a Go package, and an NPM package is like a Go module.

### Exporting components[​](https://templ.guide/syntax-and-usage/template-composition#exporting-components "Direct link to Exporting components")

To make a templ component available to other packages, export it by capitalizing its name.

```
package componentstempl Hello() {	<div>Hello</div>}
```

### Importing components[​](https://templ.guide/syntax-and-usage/template-composition#importing-components "Direct link to Importing components")

To use a component in another package, import the package and use the component as you would any other Go function or type.

```
package mainimport "github.com/a-h/templ/examples/counter/components"templ Home() {	@components.Hello()}
```

tip

To import a component from another Go module, you must first import the module by using the `go get <module>` command. Then, you can import the component as you would any other Go package.

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/03-syntax-and-usage/10-template-composition.md)

[Previous Raw Go](https://templ.guide/syntax-and-usage/raw-go)[Next CSS style management](https://templ.guide/syntax-and-usage/css-style-management)

*   [Children](https://templ.guide/syntax-and-usage/template-composition#children)
    *   [Using children in code components](https://templ.guide/syntax-and-usage/template-composition#using-children-in-code-components)
*   [Components as parameters](https://templ.guide/syntax-and-usage/template-composition#components-as-parameters)
*   [Joining Components](https://templ.guide/syntax-and-usage/template-composition#joining-components)
*   [Sharing and re-using components](https://templ.guide/syntax-and-usage/template-composition#sharing-and-re-using-components)
    *   [Exporting components](https://templ.guide/syntax-and-usage/template-composition#exporting-components)
    *   [Importing components](https://templ.guide/syntax-and-usage/template-composition#importing-components)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.