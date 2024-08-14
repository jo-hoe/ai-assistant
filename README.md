# AI Assistant

[![Test Status](https://github.com/jo-hoe/ai-assistant/workflows/test/badge.svg)](https://github.com/jo-hoe/ai-assistant/actions?workflow=test)
[![Lint Status](https://github.com/jo-hoe/ai-assistant/workflows/lint/badge.svg)](https://github.com/jo-hoe/ai-assistant/actions?workflow=lint)
[![Go Report Card](https://goreportcard.com/badge/github.com/jo-hoe/ai-assistant)](https://goreportcard.com/report/github.com/jo-hoe/ai-assistant)
[![Coverage Status](https://coveralls.io/repos/github/jo-hoe/ai-assistant/badge.svg?branch=main)](https://coveralls.io/github/jo-hoe/ai-assistant?branch=main)

An AI assistant using LLMS to get answers quickly.
🚧 This is a work in progress.

## Interoperability

This app is intended to run one multiple OS systems (Windows, Mac, Linux).

To build it on windows you need a C complier, as Windows typically does not come with one out of the box.
You can install it for instance via [Chocolatey](https://chocolatey.org/).
Once Chocolatey is installed run the following command in Admin mode:

```PowerShell
choco install mingw
```

You may need to restart your system before you can build the UI, environment variables may not work after the compiler installation.

## Run Locally

### Pre-requisites

To build this project you will need

- [Golang](https://go.dev/dl/)
- [Templ](https://templ.guide/quick-start/installation)

### Run

The project is using `make`. `make` is typically installed by default on Linux and Mac.
`make` is not strictly required, but it helps and documents commonly used commands.

If you run on Windows, you can directly install it from [gnuwin32](https://gnuwin32.sourceforge.net/packages/make.htm) or via `winget`

```PowerShell
winget install GnuWin32.Make
```

You will also need Docker and Python.
Run `make init` to install all dependencies in a virtual Python environment.

### How to Use

You can check all `make` commands by running.

```bash
make help
```

## Technologies

- Golang [Webview](https://github.com/webview/webview_go) to create interoperable client application
- [Echo](https://echo.labstack.com/) as web framework (server)

## Linting

Project used `golangci-lint` for linting.

<https://golangci-lint.run/welcome/install/>

Run the linting locally by executing

```bash
golangci-lint run ./...
```

## ToDo

- Add a [form](https://theprimeagen.github.io/fem-htmx/lessons/htmx-basics/htmx-swap) with the content
- add containers to test the build on different platforms
