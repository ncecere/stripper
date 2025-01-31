Expressions | templ docs
===============

[Skip to main content](https://templ.guide/syntax-and-usage/expressions#__docusaurus_skipToContent_fallback)

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
*   Expressions

On this page

Expressions
===========

String expressions[​](https://templ.guide/syntax-and-usage/expressions#string-expressions "Direct link to String expressions")
------------------------------------------------------------------------------------------------------------------------------

Within a templ element, expressions can be used to render strings. Content is automatically escaped using context-aware HTML encoding rules to protect against XSS and CSS injection attacks.

String literals, variables and functions that return a string can be used.

### Literals[​](https://templ.guide/syntax-and-usage/expressions#literals "Direct link to Literals")

You can use Go string literals.

component.templ

```
package maintempl component() {  <div>{ "print this" }</div>  <div>{ `and this` }</div>}
```

Output

```
<div>print this</div><div>and this</div>
```

### Variables[​](https://templ.guide/syntax-and-usage/expressions#variables "Direct link to Variables")

Any Go string variable can be used, for example:

*   A string function parameter.
*   A field on a struct.
*   A variable or constant string that is in scope.

/main.templ

```
package maintempl greet(prefix string, p Person) {  <div>{ prefix } { p.Name }{ exclamation }</div>}
```

main.go

```
package maintype Person struct {  Name string}const exclamation = "!"func main() {  p := Person{ Name: "John" }  component := greet("Hello", p)   component.Render(context.Background(), os.Stdout)}
```

Output

```
<div>Hello John!</div>
```

### Functions[​](https://templ.guide/syntax-and-usage/expressions#functions "Direct link to Functions")

Functions that return `string` or `(string, error)` can be used.

component.templ

```
package mainimport "strings"import "strconv"func getString() (string, error) {  return "DEF", nil}templ component() {  <div>{ strings.ToUpper("abc") }</div>  <div>{ getString() }</div>}
```

Output

```
<div>ABC</div><div>DEF</div>
```

If the function returns an error, the `Render` function will return an error containing the location of the error and the underlying error.

### Escaping[​](https://templ.guide/syntax-and-usage/expressions#escaping "Direct link to Escaping")

templ automatically escapes strings using HTML escaping rules.

component.templ

```
package maintempl component() {  <div>{ `</div><script>alert('hello!')</script><div>` }</div>}
```

Output

```
<div>&lt;/div&gt;&lt;script&gt;alert(&#39;hello!&#39;)&lt;/script&gt;&lt;div&gt;</div>
```

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/03-syntax-and-usage/04-expressions.md)

[Previous Attributes](https://templ.guide/syntax-and-usage/attributes)[Next Statements](https://templ.guide/syntax-and-usage/statements)

*   [String expressions](https://templ.guide/syntax-and-usage/expressions#string-expressions)
    *   [Literals](https://templ.guide/syntax-and-usage/expressions#literals)
    *   [Variables](https://templ.guide/syntax-and-usage/expressions#variables)
    *   [Functions](https://templ.guide/syntax-and-usage/expressions#functions)
    *   [Escaping](https://templ.guide/syntax-and-usage/expressions#escaping)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.