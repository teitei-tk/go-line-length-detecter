package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type result struct {
	Text   string
	Path   string
	Line   int
	Length int
}

const detectLength = 150

func detectAllRbFiles(root string) []string {
	var paths []string
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		ext := strings.LastIndex(path, ".")
		if ext <= 0 {
			return nil
		}

		if path[ext:] != ".rb" {
			return nil
		}

		paths = append(paths, path)
		return nil
	})

	return paths
}

func lineLenDetecter(path string, length int) []result {
	fp, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	var r []result
	var index int
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		index = index + 1
		if len(scanner.Text()) < length {
			continue
		}

		r = append(r, result{
			Text:   scanner.Text(),
			Path:   path,
			Line:   index,
			Length: len(scanner.Text()),
		})
	}

	return r
}

func printResult(results []result) {
	for _, r := range results {
		fmt.Printf(`Line: %d,
Length: %d,
Text: %s,`, r.Line, r.Length, r.Text)
	}
}

func main() {
	root := ""
	files := detectAllRbFiles(root)
	if len(files) <= 0 {
		log.Fatal("file is not detected")
	}

	for _, f := range files {
		m := lineLenDetecter(f, detectLength)
		if len(m) <= 0 {
			continue
		}

		printResult(m)
	}

	fmt.Println("finish")
	os.Exit(0)
}
