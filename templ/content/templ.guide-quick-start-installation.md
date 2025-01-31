Installation | templ docs
===============

[Skip to main content](https://templ.guide/quick-start/installation#__docusaurus_skipToContent_fallback)

[![Image 1: Templ Logo](https://templ.guide/img/logo.svg)![Image 2: Templ Logo](https://templ.guide/img/logo.svg)](https://templ.guide/)[Docs](https://templ.guide/)

[GitHub](https://github.com/a-h/templ)

Search

*   [Introduction](https://templ.guide/)
*   [Quick start](https://templ.guide/quick-start/installation)
    
    *   [Installation](https://templ.guide/quick-start/installation)
    *   [Creating a simple templ component](https://templ.guide/quick-start/creating-a-simple-templ-component)
    *   [Running your first templ application](https://templ.guide/quick-start/running-your-first-templ-application)
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
*   Quick start
*   Installation

On this page

Installation
============

go install[​](https://templ.guide/quick-start/installation#go-install "Direct link to go install")
--------------------------------------------------------------------------------------------------

With Go 1.20 or greater installed, run:

```
go install github.com/a-h/templ/cmd/templ@latest
```

Github binaries[​](https://templ.guide/quick-start/installation#github-binaries "Direct link to Github binaries")
-----------------------------------------------------------------------------------------------------------------

Download the latest release from [https://github.com/a-h/templ/releases/latest](https://github.com/a-h/templ/releases/latest)

Nix[​](https://templ.guide/quick-start/installation#nix "Direct link to Nix")
-----------------------------------------------------------------------------

templ provides a Nix flake with an exported package containing the binary at [https://github.com/a-h/templ/blob/main/flake.nix](https://github.com/a-h/templ/blob/main/flake.nix)

```
nix run github:a-h/templ
```

templ also provides a development shell which includes all of the tools required to build templ, e.g. go, gopls etc. but not templ itself.

```
nix develop github:a-h/templ
```

To install in your Nix Flake:

This flake exposes an overlay, so you can add it to your own Flake and/or NixOS system.

```
{  inputs = {    ...    templ.url = "github:a-h/templ";    ...  };  outputs = inputs@{    ...  }:  # For NixOS configuration:  {    # Add the overlay,    nixpkgs.overlays = [      inputs.templ.overlays.default    ];    # and install the package    environment.systemPackages = with pkgs; [      templ    ];  };  # For a flake project:  let    forAllSystems = f: nixpkgs.lib.genAttrs allSystems (system: f {      inherit system;      pkgs = import nixpkgs { inherit system; };    });    templ = system: inputs.templ.packages.${system}.templ;  in {    packages = forAllSystems ({ pkgs, system }: {      myNewPackage = pkgs.buildGoModule {        ...        preBuild = ''          ${templ system}/bin/templ generate        '';      };    });    devShell = forAllSystems ({ pkgs, system }:      pkgs.mkShell {        buildInputs = with pkgs; [          go          (templ system)        ];      };  });}
```

Docker[​](https://templ.guide/quick-start/installation#docker "Direct link to Docker")
--------------------------------------------------------------------------------------

A Docker container is pushed on each release to [https://github.com/a-h/templ/pkgs/container/templ](https://github.com/a-h/templ/pkgs/container/templ)

Pull the latest version with:

```
docker pull ghcr.io/a-h/templ:latest
```

To use the container, mount the source code of your application into the `/app` directory, set the working directory to the same directory and run `templ generate`, e.g. in a Linux or Mac shell, you can generate code for the current directory with:

```
docker run -v `pwd`:/app -w=/app ghcr.io/a-h/templ:latest generate
```

If you want to build templates using a multi-stage Docker build, you can use the `templ` image as a base image.

Here's an example multi-stage Dockerfile. Note that in the `generate-stage` the source code is copied into the container, and the `templ generate` command is run. The `build-stage` then copies the generated code into the container and builds the application.

The permissions of the source code are set to a user with a UID of 65532, which is the UID of the `nonroot` user in the `ghcr.io/a-h/templ:latest` image.

Note also the use of the `RUN ["templ", "generate"]` command instead of the common `RUN templ generate` command. This is because the templ Docker container does not contain a shell environment to keep its size minimal, so the command must be ran in the ["exec" form](https://docs.docker.com/reference/dockerfile/#shell-and-exec-form).

```
# FetchFROM golang:latest AS fetch-stageCOPY go.mod go.sum /appWORKDIR /appRUN go mod download# GenerateFROM ghcr.io/a-h/templ:latest AS generate-stageCOPY --chown=65532:65532 . /appWORKDIR /appRUN ["templ", "generate"]# BuildFROM golang:latest AS build-stageCOPY --from=generate-stage /app /appWORKDIR /appRUN CGO_ENABLED=0 GOOS=linux go build -o /app/app# TestFROM build-stage AS test-stageRUN go test -v ./...# DeployFROM gcr.io/distroless/base-debian12 AS deploy-stageWORKDIR /COPY --from=build-stage /app/app /appEXPOSE 8080USER nonroot:nonrootENTRYPOINT ["/app"]
```

[Edit this page](https://github.com/a-h/templ/tree/main/docs/docs/02-quick-start/01-installation.md)

[Previous Introduction](https://templ.guide/)[Next Creating a simple templ component](https://templ.guide/quick-start/creating-a-simple-templ-component)

*   [go install](https://templ.guide/quick-start/installation#go-install)
*   [Github binaries](https://templ.guide/quick-start/installation#github-binaries)
*   [Nix](https://templ.guide/quick-start/installation#nix)
*   [Docker](https://templ.guide/quick-start/installation#docker)

Copyright © 2024 Adrian Hesketh, Built with Docusaurus.