package main

import (
	"fmt"
	"os"

	"github.com/dalefrieswthat/kubed/internal/layout"
)

func main() {
	if len(os.Args) < 3 {
		printUsage()
		os.Exit(1)
	}
	cmd, sub := os.Args[1], os.Args[2]
	if cmd != "layout" {
		printUsage()
		os.Exit(1)
	}
	switch sub {
	case "capture":
		allNamespaces := false
		for i := 3; i < len(os.Args); i++ {
			if os.Args[i] == "--all-namespaces" {
				allNamespaces = true
				break
			}
		}
		if err := layout.RunCapture(allNamespaces); err != nil {
			fmt.Fprintf(os.Stderr, "kubed layout capture: %v\n", err)
			os.Exit(1)
		}
	case "show":
		if err := layout.RunShow(); err != nil {
			if !os.IsNotExist(err) {
				fmt.Fprintf(os.Stderr, "kubed layout show: %v\n", err)
			}
			os.Exit(1)
		}
	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: kubed layout capture [--all-namespaces]\n")
	fmt.Fprintf(os.Stderr, "       kubed layout show\n")
	fmt.Fprintf(os.Stderr, "  capture  Write .kubed/layout.json from current kube context (or ~/.kubed/layout.json if not in a git repo).\n")
	fmt.Fprintf(os.Stderr, "  show     Print layout.json to stdout.\n")
}
