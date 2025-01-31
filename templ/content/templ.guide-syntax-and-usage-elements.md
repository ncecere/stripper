Elements | templ docs
===============

[Skip to main content](https://templ.guide/syntax-and-usage/elements#__docusaurus_skipToContent_fallback)

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
*   Elements

On this page

Elements
========

templ elements are used to render HTML within templ components.

button.templ

```
package maintempl button(text string) {	<button class="button">{ text }</button>}
```

main.go

```
package mainimport (	"context"	"os")func main() {	button("Click me").Render(context.Background(), os.Stdout)}
```

Output

```
<button class="button"> Click me</button>
```

info

templ automatically minifies HTML responses, output is shown formatted for readability.

Tags must be closed[​](https://templ.guide/syntax-and-usage/elements#tags-must-be-closed "Direct link to Tags must be closed")
------------------------------------------------------------------------------------------------------------------------------

Unlike HTML, templ requires that all HTML elements are closed with either a closing tag (`</a>`), or by using a self-closing element (`<hr/>`).

templ is aware of which HTML elements are "void", and will not include the closing `/` in the output HTML.

button.templ

```
package maintempl component() {	<div>Test</div>	<img src="images/test.png"/>	<br/>}
```

Output

```
<div>Test</div><img src="images/test.png"><br>
```

Attributes and elements can contain expressions[​](https://templ.guide/syntax-and-usage/elements#attributes-and-elements-can-contain-expressions "Direct link to Attributes and elements can contain expressions")
------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

templ elements can contain placeholder expressions for attributes and content.

button.templ

```
package maintempl button(name string, content string) {	<button value={ name }>{ content }</button>}
```

Rendering the component to stdout, we can see the results.

main.go

```
func main() {	component := button("John", "Say Hello")	component.Render(context.Background(), os.Stdout)}
```

Output

```
<button value="John">Say Hello</button>
```

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/03-syntax-and-usage/02-elements.md)

[Previous Basic syntax](https://templ.guide/syntax-and-usage/basic-syntax)[Next Attributes](https://templ.guide/syntax-and-usage/attributes)

*   [Tags must be closed](https://templ.guide/syntax-and-usage/elements#tags-must-be-closed)
*   [Attributes and elements can contain expressions](https://templ.guide/syntax-and-usage/elements#attributes-and-elements-can-contain-expressions)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.