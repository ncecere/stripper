Code signing | templ docs
===============

[Skip to main content](https://templ.guide/security/code-signing#__docusaurus_skipToContent_fallback)

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
    
    *   [Injection attacks](https://templ.guide/security/injection-attacks)
    *   [Content security policy](https://templ.guide/security/content-security-policy)
    *   [Code signing](https://templ.guide/security/code-signing)
*   [Media and talks](https://templ.guide/media/)
*   [Integrations](https://templ.guide/integrations/web-frameworks)
    
*   [Experimental](https://templ.guide/experimental/overview)
    
*   [Help and community](https://templ.guide/help-and-community/)
*   [FAQ](https://templ.guide/faq/)

*   [](https://templ.guide/)
*   Security
*   Code signing

Code signing
============

Binaries are created by the Github Actions workflow at [https://github.com/a-h/templ/blob/main/.github/workflows/release.yml](https://github.com/a-h/templ/blob/main/.github/workflows/release.yml)

Binaries are signed by cosign. The public key is stored in the repository at [https://github.com/a-h/templ/blob/main/cosign.pub](https://github.com/a-h/templ/blob/main/cosign.pub)

Instructions for key verification at [https://docs.sigstore.dev/verifying/verify/](https://docs.sigstore.dev/verifying/verify/)

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/10-security/03-code-signing.md)

[Previous Content security policy](https://templ.guide/security/content-security-policy)[Next Media and talks](https://templ.guide/media/)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.