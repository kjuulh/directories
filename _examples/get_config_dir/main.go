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

	dataDir, ok := dirs.DataDir()
	if !ok {
		panic("could not find data dir")
	}
	fmt.Printf("data dir: %s\n", dataDir)
}
