package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("error getting home dir:", err)
		return
	}

	root := filepath.Join(home, "Music")

	counts := make(map[string]int)

	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		ext := filepath.Ext(path)

		fmt.Printf("file:\t\t%s\nextension:\t%s\n", relPath, ext)

		counts[ext]++

		return nil
	})

	if err != nil {
		fmt.Println("walk error:", err)
	}

	fmt.Println("\nSummary:")

	for ext, count := range counts {
		fmt.Printf("%s:\t%d\n", ext, count)
	}

	/* after creating sum of amount of files by extension could possibly create a fzf-style TUI using a library like Bubble Tea */

}
