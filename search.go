package main

import (
	"github.com/kristian-yrjola/search/index"
	"fmt"
	"flag"
	"os"
	"io/ioutil"
	"log"
	"strings"
	"bufio"
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("Specify directory\n")
		os.Exit(1)
	}
	var directory = flag.Args()[0]
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}
	var txts = FilterTxtFilenames(files)
	fmt.Printf("Found %d txt files in directory '%s'\n", len(txts), directory)
	i := index.New()
	for _, file := range txts {
		var filename = strings.TrimSuffix(directory, "/") + "/" + file
		fmt.Printf("File %s\n", filename)
		bytes, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		i.AddToIndex(filename, string(bytes))
	}
	searchMode(i)
}

func searchMode(index *index.Index) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		var text = scanner.Text()
		if strings.HasPrefix(text, ":search") {
			var fields = strings.Fields(text)[1:]
			fmt.Printf("Words: %s\n", fields)
			fmt.Printf("Result:\n")
			var result = index.Search(fields)
			for _, score := range result {
				fmt.Printf(" - %s: %d\n", score.Filename, score.Score)
			}
		}
		if strings.EqualFold(text, ":exit") || strings.EqualFold(text, ":q") {
			os.Exit(0)
		}
		fmt.Print("> ")
	}
}

func FilterTxtFilenames(target []os.FileInfo) []string {
	result := make([]string, 0)
	for _, file := range target {
		if strings.HasSuffix(file.Name(), ".txt") {
			result = append(result, file.Name())
		}
	}
	return result
}
