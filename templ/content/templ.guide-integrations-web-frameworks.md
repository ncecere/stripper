Web frameworks | templ docs
===============

[Skip to main content](https://templ.guide/integrations/web-frameworks#__docusaurus_skipToContent_fallback)

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
    
    *   [Web frameworks](https://templ.guide/integrations/web-frameworks)
    *   [Internationalization](https://templ.guide/integrations/internationalization)
*   [Experimental](https://templ.guide/experimental/overview)
    
*   [Help and community](https://templ.guide/help-and-community/)
*   [FAQ](https://templ.guide/faq/)

*   [](https://templ.guide/)
*   Integrations
*   Web frameworks

On this page

Web frameworks
==============

Templ is framework agnostic but that does not mean it can not be used with Go frameworks and other tools.

Below are some examples of how to use templ with other Go libraries, frameworks and tools, and links to systems that have built-in templ support.

### Chi[​](https://templ.guide/integrations/web-frameworks#chi "Direct link to Chi")

See an example of using [https://github.com/go-chi/chi](https://github.com/go-chi/chi) with templ at:

[https://github.com/a-h/templ/tree/main/examples/integration-chi](https://github.com/a-h/templ/tree/main/examples/integration-chi)

### Echo[​](https://templ.guide/integrations/web-frameworks#echo "Direct link to Echo")

See an example of using [https://echo.labstack.com/](https://echo.labstack.com/) with templ at:

[https://github.com/a-h/templ/tree/main/examples/integration-echo](https://github.com/a-h/templ/tree/main/examples/integration-echo)

### Gin[​](https://templ.guide/integrations/web-frameworks#gin "Direct link to Gin")

See an example of using [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin) with templ at:

[https://github.com/a-h/templ/tree/main/examples/integration-gin](https://github.com/a-h/templ/tree/main/examples/integration-gin)

### Go Fiber[​](https://templ.guide/integrations/web-frameworks#go-fiber "Direct link to Go Fiber")

See an example of using [https://github.com/gofiber/fiber](https://github.com/gofiber/fiber) with templ at:

[https://github.com/a-h/templ/tree/main/examples/integration-gofiber](https://github.com/a-h/templ/tree/main/examples/integration-gofiber)

### github.com/gorilla/csrf[​](https://templ.guide/integrations/web-frameworks#githubcomgorillacsrf "Direct link to github.com/gorilla/csrf")

`gorilla/csrf` is a HTTP middleware library that provides cross-site request forgery (CSRF) protection.

Follow the instructions at [https://github.com/gorilla/csrf](https://github.com/gorilla/csrf) to add it to your project, by using the library as HTTP middleware.

main.go

```
package mainimport (  "crypto/rand"  "fmt"  "net/http"  "github.com/gorilla/csrf")func mustGenerateCSRFKey() (key []byte) {  key = make([]byte, 32)  n, err := rand.Read(key)  if err != nil {    panic(err)  }  if n != 32 {    panic("unable to read 32 bytes for CSRF key")  }  return}func main() {  r := http.NewServeMux()  r.Handle("/", templ.Handler(Form()))  csrfMiddleware := csrf.Protect(mustGenerateCSRFKey())  withCSRFProtection := csrfMiddleware(r)  fmt.Println("Listening on localhost:8000")  http.ListenAndServe("localhost:8000", withCSRFProtection)}
```

Creating a `CSRF` templ component makes it easy to include the CSRF token in your forms.

form.templ

```
templ Form() {  <h1>CSRF Example</h1>  <form method="post" action="/">    @CSRF()    <div>      If you inspect the HTML form, you will see a hidden field with the value: { ctx.Value("gorilla.csrf.Token").(string) }    </div>    <input type="submit" value="Submit with CSRF token"/>  </form>  <form method="post" action="/">    <div>      You can also submit the form without the CSRF token to validate that the CSRF protection is working.    </div>    <input type="submit" value="Submit without CSRF token"/>  </form>}templ CSRF() {  <input type="hidden" name="gorilla.csrf.Token" value={ ctx.Value("gorilla.csrf.Token").(string) }/>}
```

Project scaffolding[​](https://templ.guide/integrations/web-frameworks#project-scaffolding "Direct link to Project scaffolding")
--------------------------------------------------------------------------------------------------------------------------------

*   Gowebly - [https://github.com/gowebly/gowebly](https://github.com/gowebly/gowebly)
*   Go-blueprint - [https://github.com/Melkeydev/go-blueprint](https://github.com/Melkeydev/go-blueprint)
*   Slick - [https://github.com/anthdm/slick](https://github.com/anthdm/slick)

Other templates[​](https://templ.guide/integrations/web-frameworks#other-templates "Direct link to Other templates")
--------------------------------------------------------------------------------------------------------------------

### `template/html`[​](https://templ.guide/integrations/web-frameworks#templatehtml "Direct link to templatehtml")

See [Using with Go templates](https://templ.guide/syntax-and-usage/using-with-go-templates)

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/12-integrations/01-web-frameworks.md)

[Previous Media and talks](https://templ.guide/media/)[Next Internationalization](https://templ.guide/integrations/internationalization)

*   [Chi](https://templ.guide/integrations/web-frameworks#chi)
*   [Echo](https://templ.guide/integrations/web-frameworks#echo)
*   [Gin](https://templ.guide/integrations/web-frameworks#gin)
*   [Go Fiber](https://templ.guide/integrations/web-frameworks#go-fiber)
*   [github.com/gorilla/csrf](https://templ.guide/integrations/web-frameworks#githubcomgorillacsrf)
*   [Project scaffolding](https://templ.guide/integrations/web-frameworks#project-scaffolding)
*   [Other templates](https://templ.guide/integrations/web-frameworks#other-templates)
    *   [`template/html`](https://templ.guide/integrations/web-frameworks#templatehtml)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.