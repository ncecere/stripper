Basic syntax | templ docs
===============

[Skip to main content](https://templ.guide/syntax-and-usage/basic-syntax#__docusaurus_skipToContent_fallback)

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
*   Basic syntax

On this page

Basic syntax
============

Package name and imports[​](https://templ.guide/syntax-and-usage/basic-syntax#package-name-and-imports "Direct link to Package name and imports")
-------------------------------------------------------------------------------------------------------------------------------------------------

templ files start with a package name, followed by any required imports, just like Go.

```
package mainimport "fmt"import "time"
```

Components[​](https://templ.guide/syntax-and-usage/basic-syntax#components "Direct link to Components")
-------------------------------------------------------------------------------------------------------

templ files can also contain components. Components are markup and code that is compiled into functions that return a `templ.Component` interface by running the `templ generate` command.

Components can contain templ elements that render HTML, text, expressions that output text or include other templates, and branching statements such as `if` and `switch`, and `for` loops.

```
package maintempl headerTemplate(name string) {  <header data-testid="headerTemplate">    <h1>{ name }</h1>  </header>}
```

Go code[​](https://templ.guide/syntax-and-usage/basic-syntax#go-code "Direct link to Go code")
----------------------------------------------------------------------------------------------

Outside of templ Components, templ files are ordinary Go code.

```
package main// Ordinary Go code that we can use in our Component.var greeting = "Welcome!"// templ Componenttempl headerTemplate(name string) {  <header>    <h1>{ name }</h1>    <h2>"{ greeting }" comes from ordinary Go code</h2>  </header>}
```

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/03-syntax-and-usage/01-basic-syntax.md)

[Previous Running your first templ application](https://templ.guide/quick-start/running-your-first-templ-application)[Next Elements](https://templ.guide/syntax-and-usage/elements)

*   [Package name and imports](https://templ.guide/syntax-and-usage/basic-syntax#package-name-and-imports)
*   [Components](https://templ.guide/syntax-and-usage/basic-syntax#components)
*   [Go code](https://templ.guide/syntax-and-usage/basic-syntax#go-code)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.