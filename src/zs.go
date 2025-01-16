package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func listAllFilesAndDirs(showHidden bool, includeGit bool, showNodeModules bool) {
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "./")
	}
	for _, path := range os.Args[1:] {
		if path == "--show-hidden" || path == "--include-git" {
			continue
		}
		err := filepath.Walk(path, func(root string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			level := strings.Count(strings.TrimPrefix(root, path), string(os.PathSeparator))
			indent := strings.Repeat(" ", level+1)
			if includeGit || !strings.Contains(root, ".git") && (!strings.Contains(root, "entrypoint") || showNodeModules) {
				if info.IsDir() {
					fmt.Printf("%s%s%s/\n", indent, "\033[34m", info.Name())
				} else {
					if showHidden || !strings.HasPrefix(info.Name(), ".") {
						fmt.Printf("%s   %s%s\n", indent, "\033[32m", info.Name())
					}
				}
			}
			return nil
		})
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Print("\033[0m")
}

func cli() {
	showHidden := false
	includeGit := false
	showNodeModules := false
	for _, arg := range os.Args {
		if arg == "--show-hidden" {
			showHidden = true
		}
		if arg == "--show-git" {
			includeGit = true
		}
		if arg == "--show-node-modules" {
			showNodeModules = true
		}
	}
	listAllFilesAndDirs(showHidden, includeGit, showNodeModules)
}

func main() {
	cli()
}
