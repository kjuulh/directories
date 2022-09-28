# Directories

## Purpose

Directories (golang) is a library for easily getting common libraries cross
platform. This may be `.local/share`, `.config` etc. Nothing more.

## Features

This library exposes Base directories, as well as user directories. Such that
base directories expose common paths such as `.config/` and user directories
such as `$HOME/Downloads`

### Windows is currently missing (prs are welcome, I don't develop on windows myself, which is why it is missing)

## Installation

```bash
go get -u github.com/kjuulh/directories@latest
```

## Usage

For more in-depth guides see `_examples`

```go
package main

import (
	"fmt"

	"github.com/kjuulh/directories"
)

func main() {
	dirs, ok := directories.NewBaseDirs()
	if !ok {
		panic("could not get user dirs")
	}

	configDir, ok := dirs.ConfigDir()
	if !ok {
		panic("could not find config dir")
	}

	fmt.Printf("config dir: %s\n", configDir)
}
```
