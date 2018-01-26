## Description

Hackable Golang package installation tool which can help you install even private
repositories and local artifacts.

## Installation

go get github.com/gianebao/goget

## Instructions

1. In your working source directory, run `goget`
2. This will create a `go-package.json` file which lists down all libraries and their installation command.
  a. The script will try to install the packages using `go get`
3. You can modify the installation command inside the file by [<excutable>, <args> ... ]
  a. To use environment variables use `["sh, "-c", "<command_with_$ENV>"]`

## Expired packages

Expired packages (packages that are not used anymore) are placed in `Archive` section.
