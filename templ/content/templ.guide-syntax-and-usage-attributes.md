Attributes | templ docs
===============

[Skip to main content](https://templ.guide/syntax-and-usage/attributes#__docusaurus_skipToContent_fallback)

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
*   Attributes

On this page

Attributes
==========

Constant attributes[​](https://templ.guide/syntax-and-usage/attributes#constant-attributes "Direct link to Constant attributes")
--------------------------------------------------------------------------------------------------------------------------------

templ elements can have HTML attributes that use the double quote character `"`.

```
templ component() {  <p data-testid="paragraph">Text</p>}
```

Output

```
<p data-testid="paragraph">Text</p>
```

String expression attributes[​](https://templ.guide/syntax-and-usage/attributes#string-expression-attributes "Direct link to String expression attributes")
-----------------------------------------------------------------------------------------------------------------------------------------------------------

Element attributes can be set to Go strings.

```
templ component(testID string) {  <p data-testid={ testID }>Text</p>}templ page() {  @component("testid-123")}
```

Rendering the `page` component results in:

Output

```
<p data-testid="testid-123">Text</p>
```

note

String values are automatically HTML attribute encoded. This is a security measure, but may make the values (especially JSON appear) look strange to you, since some characters may be converted into HTML entities. However, it is correct HTML and won't affect the behavior.

It's also possible to use function calls in string attribute expressions.

Here's a function that returns a string based on a boolean input.

```
func testID(isTrue bool) string {    if isTrue {        return "testid-123"    }    return "testid-456"}
```

```
templ component() {  <p data-testid={ testID(true) }>Text</p>}
```

The result:

Output

```
<p data-testid="testid-123">Text</p>
```

Functions in string attribute expressions can also return errors.

```
func testID(isTrue bool) (string, error) {    if isTrue {        return "testid-123", nil    }    return "", fmt.Errorf("isTrue is false")}
```

If the function returns an error, the `Render` method will return the error along with its location.

Boolean attributes[​](https://templ.guide/syntax-and-usage/attributes#boolean-attributes "Direct link to Boolean attributes")
-----------------------------------------------------------------------------------------------------------------------------

Boolean attributes (see [https://html.spec.whatwg.org/multipage/common-microsyntaxes.html#boolean-attributes](https://html.spec.whatwg.org/multipage/common-microsyntaxes.html#boolean-attributes)) where the presence of an attribute name without a value means true, and the attribute name not being present means false are supported.

```
templ component() {  <hr noshade/>}
```

Output

```
<hr noshade>
```

note

templ is aware that `<hr/>` is a void element, and renders `<hr>` instead.

To set boolean attributes using variables or template parameters, a question mark after the attribute name is used to denote that the attribute is boolean.

```
templ component() {  <hr noshade?={ false } />}
```

Output

```
<hr>
```

Conditional attributes[​](https://templ.guide/syntax-and-usage/attributes#conditional-attributes "Direct link to Conditional attributes")
-----------------------------------------------------------------------------------------------------------------------------------------

Use an `if` statement within a templ element to optionally add attributes to elements.

```
templ component() {  <hr style="padding: 10px"    if true {      class="itIsTrue"    }  />}
```

Output

```
<hr style="padding: 10px" class="itIsTrue" />
```

Spread attributes[​](https://templ.guide/syntax-and-usage/attributes#spread-attributes "Direct link to Spread attributes")
--------------------------------------------------------------------------------------------------------------------------

Use the `{ attrMap... }` syntax in the open tag of an element to append a dynamic map of attributes to the element's attributes.

It's possible to spread any variable of type `templ.Attributes`. `templ.Attributes` is a `map[string]any` type definition.

*   If the value is a `string`, the attribute is added with the string value, e.g. `<div name="value">`.
*   If the value is a `bool`, the attribute is added as a boolean attribute if the value is true, e.g. `<div name>`.
*   If the value is a `templ.KeyValue[string, bool]`, the attribute is added if the boolean is true, e.g. `<div name="value">`.
*   If the value is a `templ.KeyValue[bool, bool]`, the attribute is added if both boolean values are true, as `<div name>`.

```
templ component(shouldBeUsed bool, attrs templ.Attributes) {  <p { attrs... }>Text</p>  <hr    if shouldBeUsed {      { attrs... }    }  />}templ usage() {  @component(false, templ.Attributes{"data-testid": "paragraph"}) }
```

Output

```
<p data-testid="paragraph">Text</p><hr>
```

URL attributes[​](https://templ.guide/syntax-and-usage/attributes#url-attributes "Direct link to URL attributes")
-----------------------------------------------------------------------------------------------------------------

The `<a>` element's `href` attribute is treated differently. templ expects you to provide a `templ.SafeURL` instead of a `string`.

Typically, you would do this by using the `templ.URL` function.

The `templ.URL` function sanitizes input URLs and checks that the protocol is `http`/`https`/`mailto` rather than `javascript` or another unexpected protocol.

```
templ component(p Person) {  <a href={ templ.URL(p.URL) }>{ strings.ToUpper(p.Name) }</a>}
```

tip

In templ, all attributes are HTML-escaped. This means that:

*   `&` characters in the URL are escaped to `&amp;`.
*   `"` characters are escaped to `&quot;`.
*   `'` characters are escaped to `&#39;`.

This done to prevent XSS attacks. For example, without escaping, if a string contained `http://google.com" onclick="alert('hello')"`, the browser would interpret this as a URL followed by an `onclick` attribute, which would execute JavaScript code.

The escaping does not change the URL's functionality.

Sanitization is the process of examining the URL scheme (protocol) and structure to ensure that it's safe to use, e.g. that it doesn't contain `javascript:` or other potentially harmful schemes. If a URL is not safe, templ will replace the URL with `about:invalid#TemplFailedSanitizationURL`.

The `templ.URL` function only supports standard HTML elements and attributes (`<a href=""` and `<form action=""`).

For use on non-standard HTML elements (e.g. HTMX's `hx-*` attributes), convert the `templ.URL` to a `string` after sanitization.

```
templ component(contact model.Contact) {  <div hx-get={ string(templ.URL(fmt.Sprintf("/contacts/%s/email", contact.ID)))}>    { contact.Name }  </div>}
```

caution

If you need to bypass this sanitization, you can use `templ.SafeURL(myURL)` to mark that your string is safe to use.

This may introduce security vulnerabilities to your program.

JavaScript attributes[​](https://templ.guide/syntax-and-usage/attributes#javascript-attributes "Direct link to JavaScript attributes")
--------------------------------------------------------------------------------------------------------------------------------------

`onClick` and other `on*` handlers have special behaviour, they expect a reference to a `script` template.

info

This ensures that any client-side JavaScript that is required for a component to function is only emitted once, that script name collisions are not possible, and that script input parameters are properly sanitized.

```
script withParameters(a string, b string, c int) {	console.log(a, b, c);}script withoutParameters() {	alert("hello");}templ Button(text string) {	<button onClick={ withParameters("test", text, 123) } onMouseover={ withoutParameters() } type="button">{ text }</button>}
```

Output

```
<script type="text/javascript"> function __templ_withParameters_1056(a, b, c){console.log(a, b, c);}function __templ_withoutParameters_6bbf(){alert("hello");}</script><button onclick="__templ_withParameters_1056("test","Say hello",123)" onmouseover="__templ_withoutParameters_6bbf()" type="button"> Say hello</button>
```

CSS attributes[​](https://templ.guide/syntax-and-usage/attributes#css-attributes "Direct link to CSS attributes")
-----------------------------------------------------------------------------------------------------------------

CSS handling is discussed in detail in [CSS style management](https://templ.guide/syntax-and-usage/css-style-management).

JSON attributes[​](https://templ.guide/syntax-and-usage/attributes#json-attributes "Direct link to JSON attributes")
--------------------------------------------------------------------------------------------------------------------

To set an attribute's value to a JSON string (e.g. for HTMX's [hx-vals](https://htmx.org/attributes/hx-vals) or Alpine's [x-data](https://alpinejs.dev/directives/data)), serialize the value to a string using a function.

```
func countriesJSON() string {	countries := []string{"Czech Republic", "Slovakia", "United Kingdom", "Germany", "Austria", "Slovenia"}	bytes, _ := json.Marshal(countries)	return string(bytes)}
```

```
templ SearchBox() {	<search-webcomponent suggestions={ countriesJSON() } />}
```

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/03-syntax-and-usage/03-attributes.md)

[Previous Elements](https://templ.guide/syntax-and-usage/elements)[Next Expressions](https://templ.guide/syntax-and-usage/expressions)

*   [Constant attributes](https://templ.guide/syntax-and-usage/attributes#constant-attributes)
*   [String expression attributes](https://templ.guide/syntax-and-usage/attributes#string-expression-attributes)
*   [Boolean attributes](https://templ.guide/syntax-and-usage/attributes#boolean-attributes)
*   [Conditional attributes](https://templ.guide/syntax-and-usage/attributes#conditional-attributes)
*   [Spread attributes](https://templ.guide/syntax-and-usage/attributes#spread-attributes)
*   [URL attributes](https://templ.guide/syntax-and-usage/attributes#url-attributes)
*   [JavaScript attributes](https://templ.guide/syntax-and-usage/attributes#javascript-attributes)
*   [CSS attributes](https://templ.guide/syntax-and-usage/attributes#css-attributes)
*   [JSON attributes](https://templ.guide/syntax-and-usage/attributes#json-attributes)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.