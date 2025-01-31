Introduction | templ docs
===============

[Skip to main content](https://templ.guide/#__docusaurus_skipToContent_fallback)

[![Image 1: Templ Logo](https://templ.guide/img/logo.svg)![Image 2: Templ Logo](https://templ.guide/img/logo.svg)](https://templ.guide/)[Docs](https://templ.guide/)

[GitHub](https://github.com/a-h/templ)

Search

*   [Introduction](https://templ.guide/)
*   [Quick start](https://templ.guide/quick-start/installation)
    
*   [Syntax and usage](https://templ.guide/syntax-and-usage/basic-syntax)
    
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
*   Introduction

On this page

Introduction
============

templ - build HTML with Go[​](https://templ.guide/#templ---build-html-with-go "Direct link to templ - build HTML with Go")
--------------------------------------------------------------------------------------------------------------------------

Create components that render fragments of HTML and compose them to create screens, pages, documents, or apps.

*   Server-side rendering: Deploy as a serverless function, Docker container, or standard Go program.
*   Static rendering: Create static HTML files to deploy however you choose.
*   Compiled code: Components are compiled into performant Go code.
*   Use Go: Call any Go code, and use standard `if`, `switch`, and `for` statements.
*   No JavaScript: Does not require any client or server-side JavaScript.
*   Great developer experience: Ships with IDE autocompletion.

```
package maintempl Hello(name string) {  <div>Hello, { name }</div>}templ Greeting(person Person) {  <div class="greeting">    @Hello(person.Name)  </div>}
```

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/index.md)

[Next Installation](https://templ.guide/quick-start/installation)

*   [templ - build HTML with Go](https://templ.guide/#templ---build-html-with-go)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.