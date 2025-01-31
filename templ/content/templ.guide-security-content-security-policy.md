Content security policy | templ docs
===============

[Skip to main content](https://templ.guide/security/content-security-policy#__docusaurus_skipToContent_fallback)

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
*   Content security policy

On this page

Content security policy
=======================

Nonces[​](https://templ.guide/security/content-security-policy#nonces "Direct link to Nonces")
----------------------------------------------------------------------------------------------

In templ [script templates](https://templ.guide/syntax-and-usage/script-templates#script-templates) are rendered as inline `<script>` tags.

Strict Content Security Policies (CSP) can prevent these inline scripts from executing.

By setting a nonce attribute on the `<script>` tag, and setting the same nonce in the CSP header, the browser will allow the script to execute.

info

It's your responsibility to generate a secure nonce. Nonces should be generated using a cryptographically secure random number generator.

See [https://content-security-policy.com/nonce/](https://content-security-policy.com/nonce/) for more information.

Setting a nonce[​](https://templ.guide/security/content-security-policy#setting-a-nonce "Direct link to Setting a nonce")
-------------------------------------------------------------------------------------------------------------------------

The `templ.WithNonce` function can be used to set a nonce for templ to use when rendering scripts.

It returns an updated `context.Context` with the nonce set.

In this example, the `alert` function is rendered as a script element by templ.

templates.templ

```
package mainimport "context"import "os"script onLoad() {    alert("Hello, world!")}templ template() {    @onLoad()}
```

main.go

```
package mainimport (	"fmt"	"log"	"net/http"	"time")func withNonce(next http.Handler) http.Handler {	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {		nonce := securelyGenerateRandomString()		w.Header().Add("Content-Security-Policy", fmt.Sprintf("script-src 'nonce-%s'", nonce))		// Use the context to pass the nonce to the handler.		ctx := templ.WithNonce(r.Context(), nonce)		next.ServeHTTP(w, r.WithContext(ctx))	})}func main() {	mux := http.NewServeMux()	// Handle template.	mux.HandleFunc("/", templ.Handler(template()))	// Apply middleware.	withNonceMux := withNonce(mux)	// Start the server.	fmt.Println("listening on :8080")	if err := http.ListenAndServe(":8080", withNonceMux); err != nil {		log.Printf("error listening: %v", err)	}}
```

Output

```
<script type="text/javascript" nonce="randomly generated nonce">  function __templ_onLoad_5a85() {    alert("Hello, world!")  }</script><script type="text/javascript" nonce="randomly generated nonce">  __templ_onLoad_5a85()</script>
```

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/10-security/02-content-security-policy.md)

[Previous Injection attacks](https://templ.guide/security/injection-attacks)[Next Code signing](https://templ.guide/security/code-signing)

*   [Nonces](https://templ.guide/security/content-security-policy#nonces)
*   [Setting a nonce](https://templ.guide/security/content-security-policy#setting-a-nonce)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.