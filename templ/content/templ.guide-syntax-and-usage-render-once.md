Render once | templ docs
===============

[Skip to main content](https://templ.guide/syntax-and-usage/render-once#__docusaurus_skipToContent_fallback)

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
*   Render once

On this page

Render once
===========

If you need to render something to the page once per page, you can create a `*OnceHandler` with `templ.NewOnceHandler()` and use its `Once()` method.

The `*OnceHandler.Once()` method ensures that the content is only rendered once per distinct context passed to the component's `Render` method, even if the component is rendered multiple times.

Example[​](https://templ.guide/syntax-and-usage/render-once#example "Direct link to Example")
---------------------------------------------------------------------------------------------

The `hello` JavaScript function is only rendered once, even though the `hello` component is rendered twice.

warning

Dont write `@templ.NewOnceHandle().Once()` - this creates a new `*OnceHandler` each time the `Once` method is called, and will result in the content being rendered multiple times.

component.templ

```
package oncevar helloHandle = templ.NewOnceHandle()templ hello(label, name string) {  @helloHandle.Once() {    <script type="text/javascript">      function hello(name) {        alert('Hello, ' + name + '!');      }    </script>  }  <input type="button" value={ label } data-name={ name } onclick="hello(this.getAttribute('data-name'))"/>}templ page() {  @hello("Hello User", "user")  @hello("Hello World", "world")}
```

Output

```
<script type="text/javascript">  function hello(name) {    alert('Hello, ' + name + '!');  }</script><input type="button" value="Hello User" data-name="user" onclick="hello(this.getAttribute('data-name'))"><input type="button" value="Hello World" data-name="world" onclick="hello(this.getAttribute('data-name'))">
```

tip

Note the use of the `data-name` attribute to pass the `name` value from server-side Go code to the client-side JavaScript code.

The value of `name` is collected by the `onclick` handler, and passed to the `hello` function.

To pass complex data structures, consider using a `data-` attribute to pass a JSON string using the `templ.JSONString` function, or use the `templ.JSONScript` function to create a templ component that creates a `<script>` element containing JSON data.

Common use cases[​](https://templ.guide/syntax-and-usage/render-once#common-use-cases "Direct link to Common use cases")
------------------------------------------------------------------------------------------------------------------------

*   Rendering a `<style>` tag that contains CSS classes required by a component.
*   Rendering a `<script>` tag that contains JavaScript required by a component.
*   Rendering a `<link>` tag that contains a reference to a stylesheet.

Usage across packages[​](https://templ.guide/syntax-and-usage/render-once#usage-across-packages "Direct link to Usage across packages")
---------------------------------------------------------------------------------------------------------------------------------------

Export a component that contains the `*OnceHandler` and the content to be rendered once.

For example, create a `deps` package that contains a `JQuery` component that renders a `<script>` tag that references the jQuery library.

deps/deps.templ

```
package depsvar jqueryHandle = templ.NewOnceHandle()templ JQuery() {  @jqueryHandle.Once() {    <script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>  }}
```

You can then use the `JQuery` component in other packages, and the jQuery library will only be included once in the rendered HTML.

main.templ

```
package mainimport "deps"templ page() {  <html>    <head>      @deps.JQuery()    </head>    <body>      <h1>Hello, World!</h1>      @button()    </body>  </html>}templ button() {  @deps.JQuery()  <button>Click me</button>}
```

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/03-syntax-and-usage/18-render-once.md)

[Previous Using React with templ](https://templ.guide/syntax-and-usage/using-react-with-templ)[Next Components](https://templ.guide/core-concepts/components)

*   [Example](https://templ.guide/syntax-and-usage/render-once#example)
*   [Common use cases](https://templ.guide/syntax-and-usage/render-once#common-use-cases)
*   [Usage across packages](https://templ.guide/syntax-and-usage/render-once#usage-across-packages)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.