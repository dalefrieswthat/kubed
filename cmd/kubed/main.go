package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/dalefrieswthat/kubed/internal/layout"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
	cmd := os.Args[1]
	switch cmd {
	case "layout":
		if len(os.Args) < 3 {
			printUsage()
			os.Exit(1)
		}
		runLayout(os.Args[2])
	case "learned":
		if len(os.Args) < 3 {
			printUsage()
			os.Exit(1)
		}
		runLearned(os.Args[2:])
	default:
		printUsage()
		os.Exit(1)
	}
}

func runLayout(sub string) {
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

func runLearned(args []string) {
	sub := args[0]
	switch sub {
	case "show":
		cache, err := layout.LoadLearned()
		if err != nil {
			fmt.Fprintf(os.Stderr, "kubed learned show: %v\n", err)
			os.Exit(1)
		}
		data, _ := json.MarshalIndent(cache, "", "  ")
		fmt.Println(string(data))

	case "summary":
		cache, err := layout.LoadLearned()
		if err != nil {
			fmt.Fprintf(os.Stderr, "kubed learned summary: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(cache.Summary())

	case "add-fact":
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "Usage: kubed learned add-fact \"fact text\" [--category=X] [--source=Y] [--tags=a,b,c]")
			os.Exit(1)
		}
		fact := args[1]
		category, source, tags := parseFactFlags(args[2:])
		cache, err := layout.LoadLearned()
		if err != nil {
			fmt.Fprintf(os.Stderr, "kubed learned add-fact: %v\n", err)
			os.Exit(1)
		}
		cache.AddFact(fact, category, source, tags)
		if err := layout.SaveLearned(cache); err != nil {
			fmt.Fprintf(os.Stderr, "kubed learned add-fact: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Added fact: %s\n", fact)

	case "add-path":
		if len(args) < 3 {
			fmt.Fprintln(os.Stderr, "Usage: kubed learned add-path \"/path\" \"description\" [--tags=a,b,c]")
			os.Exit(1)
		}
		path, desc := args[1], args[2]
		var tags []string
		for _, a := range args[3:] {
			if strings.HasPrefix(a, "--tags=") {
				tags = strings.Split(strings.TrimPrefix(a, "--tags="), ",")
			}
		}
		cache, err := layout.LoadLearned()
		if err != nil {
			fmt.Fprintf(os.Stderr, "kubed learned add-path: %v\n", err)
			os.Exit(1)
		}
		cache.AddPath(path, desc, tags)
		if err := layout.SaveLearned(cache); err != nil {
			fmt.Fprintf(os.Stderr, "kubed learned add-path: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Added path: %s\n", path)

	case "add-dep":
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "Usage: kubed learned add-dep \"name\" [--kind=X] [--version=Y] [--used-by=Z] [--source=W]")
			os.Exit(1)
		}
		name := args[1]
		kind, version, usedBy, source := parseDepFlags(args[2:])
		cache, err := layout.LoadLearned()
		if err != nil {
			fmt.Fprintf(os.Stderr, "kubed learned add-dep: %v\n", err)
			os.Exit(1)
		}
		cache.AddDep(name, kind, version, usedBy, source)
		if err := layout.SaveLearned(cache); err != nil {
			fmt.Fprintf(os.Stderr, "kubed learned add-dep: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Added dep: %s (%s)\n", name, kind)

	default:
		printUsage()
		os.Exit(1)
	}
}

func parseFactFlags(args []string) (category, source string, tags []string) {
	for _, a := range args {
		switch {
		case strings.HasPrefix(a, "--category="):
			category = strings.TrimPrefix(a, "--category=")
		case strings.HasPrefix(a, "--source="):
			source = strings.TrimPrefix(a, "--source=")
		case strings.HasPrefix(a, "--tags="):
			tags = strings.Split(strings.TrimPrefix(a, "--tags="), ",")
		}
	}
	return
}

func parseDepFlags(args []string) (kind, version, usedBy, source string) {
	for _, a := range args {
		switch {
		case strings.HasPrefix(a, "--kind="):
			kind = strings.TrimPrefix(a, "--kind=")
		case strings.HasPrefix(a, "--version="):
			version = strings.TrimPrefix(a, "--version=")
		case strings.HasPrefix(a, "--used-by="):
			usedBy = strings.TrimPrefix(a, "--used-by=")
		case strings.HasPrefix(a, "--source="):
			source = strings.TrimPrefix(a, "--source=")
		}
	}
	return
}

func printUsage() {
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, "  kubed layout capture [--all-namespaces]   Capture infra layout to .kubed/layout.json")
	fmt.Fprintln(os.Stderr, "  kubed layout show                         Print layout.json")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "  kubed learned show                        Print learned.json (accumulated knowledge)")
	fmt.Fprintln(os.Stderr, "  kubed learned summary                     Print summary of learned cache")
	fmt.Fprintln(os.Stderr, "  kubed learned add-fact \"fact\" [flags]     Add a fact")
	fmt.Fprintln(os.Stderr, "  kubed learned add-path \"path\" \"desc\"      Add an important path")
	fmt.Fprintln(os.Stderr, "  kubed learned add-dep \"name\" [flags]      Add a dependency")
}
