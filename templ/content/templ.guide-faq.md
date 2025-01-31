FAQ | templ docs
===============

[Skip to main content](https://templ.guide/faq/#__docusaurus_skipToContent_fallback)

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
*   FAQ

On this page

FAQ
===

How can I migrate from templ version 0.1.x to templ 0.2.x syntax?[​](https://templ.guide/faq/#how-can-i-migrate-from-templ-version-01x-to-templ-02x-syntax "Direct link to How can I migrate from templ version 0.1.x to templ 0.2.x syntax?")
----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Versions of templ <= v0.2.663 include a `templ migrate` command that can migrate v1 syntax to v2.

The v1 syntax used some extra characters for variable injection, e.g. `{%= name %}` whereas the latest (v2) syntax uses a single pair of braces within HTML, e.g. `{ name }`.

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/15-faq/index.md)

[Previous Help and community](https://templ.guide/help-and-community/)

*   [How can I migrate from templ version 0.1.x to templ 0.2.x syntax?](https://templ.guide/faq/#how-can-i-migrate-from-templ-version-01x-to-templ-02x-syntax)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.