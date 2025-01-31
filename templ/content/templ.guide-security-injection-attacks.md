Injection attacks | templ docs
===============

[Skip to main content](https://templ.guide/security/injection-attacks#__docusaurus_skipToContent_fallback)

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
*   Injection attacks

Injection attacks
=================

templ is designed to prevent user-provided data from being used to inject vulnerabilities.

`<script>` and `<style>` tags could allow user data to inject vulnerabilities, so variables are not permitted in these sections.

```
templ Example() {  <script type="text/javascript">    function showAlert() {      alert("hello");    }  </script>  <style type="text/css">    /* Only CSS is allowed */  </style>}
```

`onClick` attributes, and other `on*` attributes are used to execute JavaScript. To prevent user data from being unescaped, `on*` attributes accept a `templ.ComponentScript`.

```
script onClickHandler(msg string) {  alert(msg);}templ Example(msg string) {  <div onClick={ onClickHandler(msg) }>    { "will be HTML encoded using templ.Escape" }  </div>}
```

Style attributes cannot be expressions, only constants, to avoid escaping vulnerabilities. templ style templates (`css className()`) should be used instead.

```
templ Example() {  <div style={ "will throw an error" }></div>}
```

Class names are sanitized by default. A failed class name is replaced by `--templ-css-class-safe-name`. The sanitization can be bypassed using the `templ.SafeClass` function, but the result is still subject to escaping.

```
templ Example() {  <div class={ "unsafe</style&gt;-will-sanitized", templ.SafeClass("&sanitization bypassed") }></div>}
```

Rendered output:

```
<div class="--templ-css-class-safe-name &amp;sanitization bypassed"></div>
```

```
templ Example() {  <div>Node text is not modified at all.</div>  <div>{ "will be escaped using templ.EscapeString" }</div>}
```

`href` attributes must be a `templ.SafeURL` and are sanitized to remove JavaScript URLs unless bypassed.

```
templ Example() {  <a href="http://constants.example.com/are/not/sanitized">Text</a>  <a href={ templ.URL("will be sanitized by templ.URL to remove potential attacks") }</a>  <a href={ templ.SafeURL("will not be sanitized by templ.URL") }</a>}
```

Within css blocks, property names, and constant CSS property values are not sanitized or escaped.

```
css className() {	background-color: #ffffff;}
```

CSS property values based on expressions are passed through `templ.SanitizeCSS` to replace potentially unsafe values with placeholders.

```
css className() {	color: { red };}
```

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/10-security/01-injection-attacks.md)

[Previous Coding assistants / LLMs](https://templ.guide/developer-tools/llm)[Next Content security policy](https://templ.guide/security/content-security-policy)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.