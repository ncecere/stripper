Internationalization | templ docs
===============

[Skip to main content](https://templ.guide/integrations/internationalization#__docusaurus_skipToContent_fallback)

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
*   Internationalization

On this page

Internationalization
====================

templ can be used with 3rd party internationalization libraries.

ctxi18n[​](https://templ.guide/integrations/internationalization#ctxi18n "Direct link to ctxi18n")
--------------------------------------------------------------------------------------------------

[https://github.com/invopop/ctxi18n](https://github.com/invopop/ctxi18n) uses the context package to load strings based on the selected locale.

An example is available at [https://github.com/a-h/templ/tree/main/examples/internationalization](https://github.com/a-h/templ/tree/main/examples/internationalization)

### Storing translations[​](https://templ.guide/integrations/internationalization#storing-translations "Direct link to Storing translations")

Translations are stored in YAML files, according to the language.

locales/en/en.yaml

```
en:  hello: "Hello"  select_language: "Select Language"
```

### Selecting the language[​](https://templ.guide/integrations/internationalization#selecting-the-language "Direct link to Selecting the language")

HTTP middleware selects the language to load based on the URL path, `/en`, `/de`, etc.

main.go

```
func newLanguageMiddleware(next http.Handler) http.Handler {	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {		lang := "en" // Default language		pathSegments := strings.Split(r.URL.Path, "/")		if len(pathSegments) > 1 {			lang = pathSegments[1]		}		ctx, err := ctxi18n.WithLocale(r.Context(), lang)		if err != nil {			log.Printf("error setting locale: %v", err)			http.Error(w, "error setting locale", http.StatusBadRequest)			return		}		next.ServeHTTP(w, r.WithContext(ctx))	})}
```

### Using the middleware[​](https://templ.guide/integrations/internationalization#using-the-middleware "Direct link to Using the middleware")

The `ctxi18n.Load` function is used to load the translations, and the middleware is used to set the language.

main.go

```
func main() {	if err := ctxi18n.Load(locales.Content); err != nil {		log.Fatalf("error loading locales: %v", err)	}	mux := http.NewServeMux()	mux.Handle("/", templ.Handler(page()))	withLanguageMiddleware := newLanguageMiddleware(mux)	log.Println("listening on :8080")	if err := http.ListenAndServe("127.0.0.1:8080", withLanguageMiddleware); err != nil {		log.Printf("error listening: %v", err)	}}
```

### Fetching translations in templates[​](https://templ.guide/integrations/internationalization#fetching-translations-in-templates "Direct link to Fetching translations in templates")

Translations are fetched using the `i18n.T` function, passing the implicit context that's available in all templ components, and the key for the translation.

```
package mainimport (	"github.com/invopop/ctxi18n/i18n")templ page() {	<html>		<head>			<meta charset="UTF-8"/>			<meta name="viewport" content="width=device-width, initial-scale=1.0"/>			<title>{ i18n.T(ctx, "hello") }</title>		</head>		<body>			<h1>{ i18n.T(ctx, "hello") }</h1>			<h2>{ i18n.T(ctx, "select_language") }</h2>			<ul>				<li><a href="/en">English</a></li>				<li><a href="/de">Deutsch</a></li>				<li><a href="/zh-cn">中文</a></li>			</ul>		</body>	</html>}
```

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/12-integrations/02-internationalization.md)

[Previous Web frameworks](https://templ.guide/integrations/web-frameworks)[Next Experimental packages](https://templ.guide/experimental/overview)

*   [ctxi18n](https://templ.guide/integrations/internationalization#ctxi18n)
    *   [Storing translations](https://templ.guide/integrations/internationalization#storing-translations)
    *   [Selecting the language](https://templ.guide/integrations/internationalization#selecting-the-language)
    *   [Using the middleware](https://templ.guide/integrations/internationalization#using-the-middleware)
    *   [Fetching translations in templates](https://templ.guide/integrations/internationalization#fetching-translations-in-templates)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.