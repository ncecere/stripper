Generating static HTML files with templ | templ docs
===============

[Skip to main content](https://templ.guide/static-rendering/generating-static-html-files-with-templ#__docusaurus_skipToContent_fallback)

[![Image 1: Templ Logo](https://templ.guide/img/logo.svg)![Image 2: Templ Logo](https://templ.guide/img/logo.svg)](https://templ.guide/)[Docs](https://templ.guide/)

[GitHub](https://github.com/a-h/templ)

Search

*   [Introduction](https://templ.guide/)
*   [Quick start](https://templ.guide/quick-start/installation)
    
*   [Syntax and usage](https://templ.guide/syntax-and-usage/basic-syntax)
    
*   [Core concepts](https://templ.guide/core-concepts/components)
    
*   [Server-side rendering](https://templ.guide/server-side-rendering/creating-an-http-server-with-templ)
    
*   [Static rendering](https://templ.guide/static-rendering/generating-static-html-files-with-templ)
    
    *   [Generating static HTML files with templ](https://templ.guide/static-rendering/generating-static-html-files-with-templ)
    *   [Blog example](https://templ.guide/static-rendering/blog-example)
    *   [Deploying static files](https://templ.guide/static-rendering/deploying-static-files)
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
*   Static rendering
*   Generating static HTML files with templ

On this page

Generating static HTML files with templ
=======================================

templ components implement the `templ.Component` interface.

The interface has a `Render` method which outputs HTML to an `io.Writer` that is passed in.

```
type Component interface {	// Render the template.	Render(ctx context.Context, w io.Writer) error}
```

In Go, the `io.Writer` interface is implemented by many built-in types in the standard library, including `os.File` (files), `os.Stdout`, and `http.ResponseWriter` (HTTP responses).

This makes it easy to use templ components in a variety of contexts to generate HTML.

To render static HTML files using templ component, first create a new Go project.

Setup project[​](https://templ.guide/static-rendering/generating-static-html-files-with-templ#setup-project "Direct link to Setup project")
-------------------------------------------------------------------------------------------------------------------------------------------

Create a new directory.

```
mkdir static-generator
```

Initialize a new Go project within it.

```
cd static-generatorgo mod init github.com/a-h/templ-examples/static-generator
```

Create a templ file[​](https://templ.guide/static-rendering/generating-static-html-files-with-templ#create-a-templ-file "Direct link to Create a templ file")
-------------------------------------------------------------------------------------------------------------------------------------------------------------

To use it, create a `hello.templ` file containing a component.

Components are functions that contain templ elements, markup, `if`, `switch` and `for` Go expressions.

hello.templ

```
package maintempl hello(name string) {	<div>Hello, { name }</div>}
```

Generate Go code from the templ file[​](https://templ.guide/static-rendering/generating-static-html-files-with-templ#generate-go-code-from-the-templ-file "Direct link to Generate Go code from the templ file")
----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Run the `templ generate` command.

```
templ generate
```

templ will generate a `hello_templ.go` file containing Go code.

This file will contain a function called `hello` which takes `name` as an argument, and returns a `templ.Component` that renders HTML.

```
func hello(name string) templ.Component {  // ...}
```

Write a program that renders to stdout[​](https://templ.guide/static-rendering/generating-static-html-files-with-templ#write-a-program-that-renders-to-stdout "Direct link to Write a program that renders to stdout")
----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Create a `main.go` file. The program creates a `hello.html` file and uses the component to write HTML to the file.

main.go

```
package mainimport (	"context"	"log"	"os")func main() {	f, err := os.Create("hello.html")	if err != nil {		log.Fatalf("failed to create output file: %v", err)	}	err = hello("John").Render(context.Background(), f)	if err != nil {		log.Fatalf("failed to write output file: %v", err)	}}
```

Run the program[​](https://templ.guide/static-rendering/generating-static-html-files-with-templ#run-the-program "Direct link to Run the program")
-------------------------------------------------------------------------------------------------------------------------------------------------

Running the code will create a file called `hello.html` containing the component's HTML.

```
go run *.go
```

hello.html

```
<div>Hello, John</div>
```

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/06-static-rendering/01-generating-static-html-files-with-templ.md)

[Previous HTTP Streaming](https://templ.guide/server-side-rendering/streaming)[Next Blog example](https://templ.guide/static-rendering/blog-example)

*   [Setup project](https://templ.guide/static-rendering/generating-static-html-files-with-templ#setup-project)
*   [Create a templ file](https://templ.guide/static-rendering/generating-static-html-files-with-templ#create-a-templ-file)
*   [Generate Go code from the templ file](https://templ.guide/static-rendering/generating-static-html-files-with-templ#generate-go-code-from-the-templ-file)
*   [Write a program that renders to stdout](https://templ.guide/static-rendering/generating-static-html-files-with-templ#write-a-program-that-renders-to-stdout)
*   [Run the program](https://templ.guide/static-rendering/generating-static-html-files-with-templ#run-the-program)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.