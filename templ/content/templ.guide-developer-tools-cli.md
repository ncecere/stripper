CLI | templ docs
===============

[Skip to main content](https://templ.guide/developer-tools/cli#__docusaurus_skipToContent_fallback)

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
    
    *   [CLI](https://templ.guide/developer-tools/cli)
    *   [IDE support](https://templ.guide/developer-tools/ide-support)
    *   [Live reload](https://templ.guide/developer-tools/live-reload)
    *   [Live reload with other tools](https://templ.guide/developer-tools/live-reload-with-other-tools)
    *   [Coding assistants / LLMs](https://templ.guide/developer-tools/llm)
*   [Security](https://templ.guide/security/injection-attacks)
    
*   [Media and talks](https://templ.guide/media/)
*   [Integrations](https://templ.guide/integrations/web-frameworks)
    
*   [Experimental](https://templ.guide/experimental/overview)
    
*   [Help and community](https://templ.guide/help-and-community/)
*   [FAQ](https://templ.guide/faq/)

*   [](https://templ.guide/)
*   Developer tools
*   CLI

On this page

CLI
===

`templ` provides a command line interface. Most users will only need to run the `templ generate` command to generate Go code from `*.templ` files.

```
usage: templ <command> [<args>...]templ - build HTML UIs with GoSee docs at https://templ.guidecommands:  generate   Generates Go code from templ files  fmt        Formats templ files  lsp        Starts a language server for templ files  info       Displays information about the templ environment  version    Prints the version
```

Generating Go code from templ files[​](https://templ.guide/developer-tools/cli#generating-go-code-from-templ-files "Direct link to Generating Go code from templ files")
------------------------------------------------------------------------------------------------------------------------------------------------------------------------

The `templ generate` command generates Go code from `*.templ` files in the current directory tree.

The command provides additional options:

```
usage: templ generate [<args>...]Generates Go code from templ files.Args:  -path <path>    Generates code for all files in path. (default .)  -f <file>    Optionally generates code for a single file, e.g. -f header.templ  -source-map-visualisations    Set to true to generate HTML files to visualise the templ code and its corresponding Go code.  -include-version    Set to false to skip inclusion of the templ version in the generated code. (default true)  -include-timestamp    Set to true to include the current time in the generated code.  -watch    Set to true to watch the path for changes and regenerate code.  -cmd <cmd>    Set the command to run after generating code.  -proxy    Set the URL to proxy after generating code and executing the command.  -proxyport    The port the proxy will listen on. (default 7331)  -proxybind    The address the proxy will listen on. (default 127.0.0.1)  -w    Number of workers to use when generating code. (default runtime.NumCPUs)  -lazy    Only generate .go files if the source .templ file is newer.	  -pprof    Port to run the pprof server on.  -keep-orphaned-files    Keeps orphaned generated templ files. (default false)  -v    Set log verbosity level to "debug". (default "info")  -log-level    Set log verbosity level. (default "info", options: "debug", "info", "warn", "error")  -help    Print help and exit.
```

For example, to generate code for a single file:

```
templ generate -f header.templ
```

Formatting templ files[​](https://templ.guide/developer-tools/cli#formatting-templ-files "Direct link to Formatting templ files")
---------------------------------------------------------------------------------------------------------------------------------

The `templ fmt` command formats template files. You can use this command in different ways:

1.  Format all template files in the current directory and subdirectories:

```
templ fmt .
```

2.  Format input from stdin and output to stdout:

```
templ fmt
```

Alternatively, you can run `fmt` in CI to ensure that invalidly formatted templatess do not pass CI. This will cause the command to exit with unix error-code `1` if any templates needed to be modified.

```
templ fmt -fail .
```

Language Server for IDE integration[​](https://templ.guide/developer-tools/cli#language-server-for-ide-integration "Direct link to Language Server for IDE integration")
------------------------------------------------------------------------------------------------------------------------------------------------------------------------

`templ lsp` provides a Language Server Protocol (LSP) implementation to support IDE integrations.

This command isn't intended to be used directly by users, but is used by IDE integrations such as the VSCode extension and by Neovim support.

A number of additional options are provided to enable runtime logging and profiling tools.

```
  -goplsLog string        The file to log gopls output, or leave empty to disable logging.  -goplsRPCTrace        Set gopls to log input and output messages.  -help        Print help and exit.  -http string        Enable http debug server by setting a listen address (e.g. localhost:7474)  -log string        The file to log templ LSP output to, or leave empty to disable logging.  -pprof        Enable pprof web server (default address is localhost:9999)
```

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/09-developer-tools/01-cli.md)

[Previous Hosting using Docker](https://templ.guide/hosting-and-deployment/hosting-using-docker)[Next IDE support](https://templ.guide/developer-tools/ide-support)

*   [Generating Go code from templ files](https://templ.guide/developer-tools/cli#generating-go-code-from-templ-files)
*   [Formatting templ files](https://templ.guide/developer-tools/cli#formatting-templ-files)
*   [Language Server for IDE integration](https://templ.guide/developer-tools/cli#language-server-for-ide-integration)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.