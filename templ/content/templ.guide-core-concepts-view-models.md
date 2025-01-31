View models | templ docs
===============

[Skip to main content](https://templ.guide/core-concepts/view-models#__docusaurus_skipToContent_fallback)

[![Image 1: Templ Logo](https://templ.guide/img/logo.svg)![Image 2: Templ Logo](https://templ.guide/img/logo.svg)](https://templ.guide/)[Docs](https://templ.guide/)

[GitHub](https://github.com/a-h/templ)

Search

*   [Introduction](https://templ.guide/)
*   [Quick start](https://templ.guide/quick-start/installation)
    
*   [Syntax and usage](https://templ.guide/syntax-and-usage/basic-syntax)
    
*   [Core concepts](https://templ.guide/core-concepts/components)
    
    *   [Components](https://templ.guide/core-concepts/components)
    *   [Template generation](https://templ.guide/core-concepts/template-generation)
    *   [Testing](https://templ.guide/core-concepts/testing)
    *   [View models](https://templ.guide/core-concepts/view-models)
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
*   Core concepts
*   View models

View models
===========

With templ, you can pass any Go type into your template as parameters, and you can call arbitrary functions.

However, if the parameters of your template don't closely map to what you're displaying to users, you may find yourself calling a lot of functions within your templ files to reshape or adjust data, or to carry out complex repeated string interpolation or URL constructions.

This can make template rendering hard to test, because you need to set up complex data structures in the right way in order to render the HTML. If the template calls APIs or accesses databases from within the templates, it's even harder to test, because then testing your templates becomes an integration test.

A more reliable approach can be to create a "View model" that only contains the fields that you intend to display, and where the data structure closely matches the structure of the visual layout.

```
package invitesgettype Handler struct {  Invites *InviteService}func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {  invites, err := h.Invites.Get(getUserIDFromContext(r.Context()))  if err != nil {     //TODO: Log error server side.  }  m := NewInviteComponentViewModel(invites, err)  teamInviteComponent(m).Render(r.Context(), w)}func NewInviteComponentViewModel(invites []models.Invite, err error) (m InviteComponentViewModel) {  m.InviteCount = len(invites)  if err != nil {    m.ErrorMessage = "Failed to load invites, please try again"  }  return m}type InviteComponentViewModel struct {  InviteCount int  ErrorMessage string}templ teamInviteComponent(model InviteComponentViewModel) {	if model.InviteCount > 0 {		<div>You have { fmt.Sprintf("%d", model.InviteCount) } pending invites</div>	}        if model.ErrorMessage != "" {		<div class="error">{ model.ErrorMessage }</div>        }}
```

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/04-core-concepts/04-view-models.md)

[Previous Testing](https://templ.guide/core-concepts/testing)[Next Creating an HTTP server with templ](https://templ.guide/server-side-rendering/creating-an-http-server-with-templ)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.