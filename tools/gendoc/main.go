// Package main contains CLI documentation generator tool.
package main

import (
	_ "embed"
	"strings"

	"github.com/grafana/clireadme"
	"github.com/grafana/k6run/cmd"
)

func main() {
	root := cmd.New()
	root.Use = strings.ReplaceAll(root.Use, "run", "k6run")
	clireadme.Main(root, 1)
}
