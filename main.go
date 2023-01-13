// main.go
package main

import (
	"main/cmd"
)

// set by goreleaser; see https://goreleaser.com/environment/
var (
	version = "dev"
	commit  = "none"
)

func main() {
	cmd.Execute(version, commit)
}
