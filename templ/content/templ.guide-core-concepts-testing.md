Testing | templ docs
===============

[Skip to main content](https://templ.guide/core-concepts/testing#__docusaurus_skipToContent_fallback)

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
*   Testing

On this page

Testing
=======

To test that data is rendered as expected, there are two main ways to do it:

*   Expectation testing - testing that specific expectations are met by the output.
*   Snapshot testing - testing that outputs match a pre-written output.

Expectation testing[​](https://templ.guide/core-concepts/testing#expectation-testing "Direct link to Expectation testing")
--------------------------------------------------------------------------------------------------------------------------

Expectation testing validates that the right data appears in the output in the right format and position.

The example at [https://github.com/a-h/templ/blob/main/examples/blog/posts\_test.go](https://github.com/a-h/templ/blob/main/examples/blog/posts_test.go) shows how to test that a list of posts is rendered correctly.

These tests use the `goquery` library to parse HTML and check that expected elements are present. `goquery` is a jQuery-like library for Go, that is useful for parsing and querying HTML. You’ll need to run `go get github.com/PuerkitoBio/goquery` to add it to your `go.mod` file.

### Testing components[​](https://templ.guide/core-concepts/testing#testing-components "Direct link to Testing components")

The test sets up a pipe to write templ's HTML output to, and reads the output from the pipe, parsing it with `goquery`.

First, we test the page header. To use `goquery` to inspect the output, we’ll need to connect the header component’s `Render` method to the `goquery.NewDocumentFromReader` function with an `io.Pipe`.

```
func TestHeader(t *testing.T) {    // Pipe the rendered template into goquery.    r, w := io.Pipe()    go func () {        _ = headerTemplate("Posts").Render(context.Background(), w)        _ = w.Close()    }()    doc, err := goquery.NewDocumentFromReader(r)    if err != nil {        t.Fatalf("failed to read template: %v", err)    }    // Expect the component to be present.    if doc.Find(`[data-testid="headerTemplate"]`).Length() == 0 {        t.Error("expected data-testid attribute to be rendered, but it wasn't")    }    // Expect the page name to be set correctly.    expectedPageName := "Posts"    if actualPageName := doc.Find("h1").Text(); actualPageName != expectedPageName {        t.Errorf("expected page name %q, got %q", expectedPageName, actualPageName)    }}
```

The header template (the "subject under test") includes a placeholder for the page name, and a `data-testid` attribute that makes it easier to locate the `headerTemplate` within the HTML using a CSS selector of `[data-testid="headerTemplate"]`.

```
templ headerTemplate(name string) {    <header data-testid="headerTemplate">        <h1>{ name }</h1>    </header>}
```

We can also test that the navigation bar was rendered.

```
func TestNav(t *testing.T) {    r, w := io.Pipe()    go func() {        _ = navTemplate().Render(context.Background(), w)        _ = w.Close()    }()    doc, err := goquery.NewDocumentFromReader(r)    if err != nil {        t.Fatalf("failed to read template: %v", err)    }    // Expect the component to include a testid.    if doc.Find(`[data-testid="navTemplate"]`).Length() == 0 {        t.Error("expected data-testid attribute to be rendered, but it wasn't")    }}
```

Testing that it was rendered is useful, but it's even better to test that the navigation includes the correct `nav` items.

In this test, we find all of the `a` elements within the `nav` element, and check that they match the expected items.

```
navItems := []string{"Home", "Posts"}doc.Find("nav a").Each(func(i int, s *goquery.Selection) {    expected := navItems[i]    if actual := s.Text(); actual != expected {        t.Errorf("expected nav item %q, got %q", expected, actual)    }})
```

To test the posts, we can use the same approach. We test that the posts are rendered correctly, and that the expected data is present.

### Testing whole pages[​](https://templ.guide/core-concepts/testing#testing-whole-pages "Direct link to Testing whole pages")

Next, we may want to go a level higher and test the entire page.

Pages are also templ components, so the tests are structured in the same way.

There’s no need to test for the specifics about what gets rendered in the `navTemplate` or `homeTemplate` at the page level, because they’re already covered in other tests.

Some developers prefer to only test the external facing part of their code (e.g. at a page level), rather than testing each individual component, on the basis that it’s slower to make changes if the implementation is too tightly controlled.

For example, if a component is reused across pages, then it makes sense to test that in detail in its own test. In the pages or higher-order components that use it, there’s no point testing it again at that level, so we only check that it was rendered to the output by looking for its data-testid attribute, unless we also need to check what we're passing to it.

### Testing the HTTP handler[​](https://templ.guide/core-concepts/testing#testing-the-http-handler "Direct link to Testing the HTTP handler")

Finally, we want to test the posts HTTP handler. This requires a different approach.

We can use the `httptest` package to create a test server, and use the `net/http` package to make a request to the server and check the response.

The tests configure the `GetPosts` function on the `PostsHandler` with a mock that returns a "database error", while the other returns a list of two posts. Here's what the `PostsHandler` looks like:

```
type PostsHandler struct {    Log      *log.Logger    GetPosts func() ([]Post, error)}
```

In the error case, the test asserts that the error message was displayed, while in the success case, it checks that the `postsTemplate` is present. It does not check that the posts have actually been rendered properly or that specific fields are visible, because that’s already tested at the component level.

Testing it again here would make the code resistant to refactoring and rework, but then again, we might have missed actually passing the posts we got back from the database to the posts template, so it’s a matter of risk appetite vs refactor resistance.

Note the switch to the table-driven testing format, a popular approach in Go for testing multiple scenarios with the same test code.

```
func TestPostsHandler(t *testing.T) {    tests := []struct {        name           string        postGetter     func() (posts []Post, err error)        expectedStatus int        assert         func(doc *goquery.Document)    }{        {            name: "database errors result in a 500 error",            postGetter: func() (posts []Post, err error) {                return nil, errors.New("database error")            },            expectedStatus: http.StatusInternalServerError,            assert: func(doc *goquery.Document) {                expected := "failed to retrieve posts\n"                if actual := doc.Text(); actual != expected {                    t.Errorf("expected error message %q, got %q", expected, actual)                }            },        },        {            name: "database success renders the posts",            postGetter: func() (posts []Post, err error) {                return []Post{                    {Name: "Name1", Author: "Author1"},                    {Name: "Name2", Author: "Author2"},                }, nil            },            expectedStatus: http.StatusInternalServerError,            assert: func(doc *goquery.Document) {                if doc.Find(`[data-testid="postsTemplate"]`).Length() == 0 {                    t.Error("expected posts to be rendered, but it wasn't")                }            },        },    }    for _, test := range tests {        // Arrange.        w := httptest.NewRecorder()        r := httptest.NewRequest(http.MethodGet, "/posts", nil)        ph := NewPostsHandler()        ph.Log = log.New(io.Discard, "", 0) // Suppress logging.        ph.GetPosts = test.postGetter        // Act.        ph.ServeHTTP(w, r)        doc, err := goquery.NewDocumentFromReader(w.Result().Body)        if err != nil {            t.Fatalf("failed to read template: %v", err)        }        // Assert.        test.assert(doc)    }}
```

### Summary[​](https://templ.guide/core-concepts/testing#summary "Direct link to Summary")

*   goquery can be used effectively with templ for writing component level tests.
*   Adding `data-testid` attributes to your code simplifies the test expressions you need to write to find elements within the output and makes your tests less brittle.
*   Testing can be split between the two concerns of template rendering, and HTTP handlers.

Snapshot testing[​](https://templ.guide/core-concepts/testing#snapshot-testing "Direct link to Snapshot testing")
-----------------------------------------------------------------------------------------------------------------

Snapshot testing is a more broad check. It simply checks that the output hasn't changed since the last time you took a copy of the output.

It relies on manually checking the output to make sure it's correct, and then "locking it in" by using the snapshot.

templ uses this strategy to check for regressions in behaviour between releases, as per [https://github.com/a-h/templ/blob/main/generator/test-html-comment/render\_test.go](https://github.com/a-h/templ/blob/main/generator/test-html-comment/render_test.go)

To make it easier to compare the output against the expected HTML, templ uses a HTML formatting library before executing the diff.

```
package testcommentimport (	_ "embed"	"testing"	"github.com/a-h/templ/generator/htmldiff")//go:embed expected.htmlvar expected stringfunc Test(t *testing.T) {	component := render("sample content")	diff, err := htmldiff.Diff(component, expected)	if err != nil {		t.Fatal(err)	}	if diff != "" {		t.Error(diff)	}}
```

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/04-core-concepts/03-testing.md)

[Previous Template generation](https://templ.guide/core-concepts/template-generation)[Next View models](https://templ.guide/core-concepts/view-models)

*   [Expectation testing](https://templ.guide/core-concepts/testing#expectation-testing)
    *   [Testing components](https://templ.guide/core-concepts/testing#testing-components)
    *   [Testing whole pages](https://templ.guide/core-concepts/testing#testing-whole-pages)
    *   [Testing the HTTP handler](https://templ.guide/core-concepts/testing#testing-the-http-handler)
    *   [Summary](https://templ.guide/core-concepts/testing#summary)
*   [Snapshot testing](https://templ.guide/core-concepts/testing#snapshot-testing)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.