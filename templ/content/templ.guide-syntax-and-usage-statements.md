Statements | templ docs
===============

[Skip to main content](https://templ.guide/syntax-and-usage/statements#__docusaurus_skipToContent_fallback)

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
*   Statements

On this page

Statements
==========

Control flow[​](https://templ.guide/syntax-and-usage/statements#control-flow "Direct link to Control flow")
-----------------------------------------------------------------------------------------------------------

Within a templ element, a subset of Go statements can be used directly.

These Go statements can be used to conditionally render child elements, or to iterate variables.

For individual implementation guides see:

*   [if/else](https://templ.guide/syntax-and-usage/if-else)
*   [switch](https://templ.guide/syntax-and-usage/switch)
*   [for loops](https://templ.guide/syntax-and-usage/loops)

if/switch/for within text[​](https://templ.guide/syntax-and-usage/statements#ifswitchfor-within-text "Direct link to if/switch/for within text")
------------------------------------------------------------------------------------------------------------------------------------------------

Go statements can be used without any escaping to make it simple for developers to include them.

The templ parser assumes that text that starts with `if`, `switch` or `for` denotes the start of one of those expressions as per this example.

show-hello.templ

```
package maintempl showHelloIfTrue(b bool) {	<div>		if b {			<p>Hello</p>		}	</div>}
```

If you need to start a text block with the words `if`, `switch`, or `for`:

*   Use a Go string expression.
*   Capitalise `if`, `switch`, or `for`.

paragraph.templ

```
package maintempl display(price float64, count int) {	<p>Switch to Linux</p>	<p>{ `switch to Linux` }</p>	<p>{ "for a day" }</p>	<p>{ fmt.Sprintf("%f", price) }{ "for" }{ fmt.Sprintf("%d", count) }</p>	<p>{ fmt.Sprintf("%f for %d", price, count) }</p>}
```

Design considerations[​](https://templ.guide/syntax-and-usage/statements#design-considerations "Direct link to Design considerations")
--------------------------------------------------------------------------------------------------------------------------------------

We decided to not require a special prefix for `if`, `switch` and `for` expressions on the basis that we were more likely to want to use a Go control statement than start a text run with those strings.

To reduce the risk of a broken control statement, resulting in printing out the source code of the application, templ will complain if a text run starts with `if`, `switch` or `for`, but no opening brace `{` is found.

For example, the following code causes the templ parser to return an error:

broken-if.templ

```
package maintempl showIfTrue(b bool) {	if b 	  <p>Hello</p>	}}
```

note

Note the missing `{` on line 4.

The following code also produces an error, since the text run starts with `if`, but no opening `{` is found.

paragraph.templ

```
package maintempl text(b bool) {	<p>if a tree fell in the woods</p>}
```

note

This also applies to `for` and `switch` statements.

To resolve the issue:

*   Use a Go string expression.
*   Capitalise `if`, `switch`, or `for`.

paragraph.templ

```
package maintempl display(price float64, count int) {	<p>Switch to Linux</p>	<p>{ `switch to Linux` }</p>	<p>{ "for a day" }</p>	<p>{ fmt.Sprintf("%f", price) }{ "for" }{ fmt.Sprintf("%d", count) }</p>	<p>{ fmt.Sprintf("%f for %d", price, count) }</p>}
```

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/03-syntax-and-usage/05-statements.md)

[Previous Expressions](https://templ.guide/syntax-and-usage/expressions)[Next If/else](https://templ.guide/syntax-and-usage/if-else)

*   [Control flow](https://templ.guide/syntax-and-usage/statements#control-flow)
*   [if/switch/for within text](https://templ.guide/syntax-and-usage/statements#ifswitchfor-within-text)
*   [Design considerations](https://templ.guide/syntax-and-usage/statements#design-considerations)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.