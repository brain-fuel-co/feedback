package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	dotFilesDir := "./pid-dotfiles"
	outputDir := "./output/pid-svgfiles"
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(dotFilesDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".dot" {
			svgName := strings.TrimSuffix(file.Name(), ".dot") + ".svg"
			inputPath := filepath.Join(dotFilesDir, file.Name())
			outputPath := filepath.Join(outputDir, svgName)

			cmd := exec.Command("dot", "-Tsvg", inputPath, "-o", outputPath)
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error converting file:", inputPath, err)
			} else {
				fmt.Println("Converted:", outputPath)
			}
		}
	}
}
