urlbuilder | templ docs
===============

[Skip to main content](https://templ.guide/experimental/urlbuilder#__docusaurus_skipToContent_fallback)

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
    
    *   [Experimental packages](https://templ.guide/experimental/overview)
    *   [urlbuilder](https://templ.guide/experimental/urlbuilder)
*   [Help and community](https://templ.guide/help-and-community/)
*   [FAQ](https://templ.guide/faq/)

*   [](https://templ.guide/)
*   Experimental
*   urlbuilder

On this page

urlbuilder
==========

A simple URL builder to construct a `templ.SafeURL`.

component.templ

```
import (  "github.com/templ-go/x/urlbuilder"  "strconv"  "strings")templ component(o Order) {  <a    href={ urlbuilder.New("https", "example.com").    Path("orders").    Path(o.ID).    Path("line-items").    Query("page", strconv.Itoa(1)).    Query("limit", strconv.Itoa(10)).    Build() }  >    { strings.ToUpper(o.Name) }  </a>}
```

See [URL Attribures](https://templ.guide/syntax-and-usage/attributes#url-attributes) for more information.

Feedback[​](https://templ.guide/experimental/urlbuilder#feedback "Direct link to Feedback")
-------------------------------------------------------------------------------------------

Please leave your feedback on this feature at [https://github.com/a-h/templ/discussions/867](https://github.com/a-h/templ/discussions/867)

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/13-experimental/02-urlbuilder.md)

[Previous Experimental packages](https://templ.guide/experimental/overview)[Next Help and community](https://templ.guide/help-and-community/)

*   [Feedback](https://templ.guide/experimental/urlbuilder#feedback)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.