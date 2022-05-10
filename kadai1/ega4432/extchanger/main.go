package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ega4432/extchanger/converter"
)

var (
	dir, to, from string
	recursive     bool
)

func init() {
	flag.StringVar(&dir, "d", "", "target directory")
	flag.BoolVar(&recursive, "r", false, "change extension recursively")
	flag.StringVar(&from, "from", "jpg", "from extension")
	flag.StringVar(&to, "to", "png", "new extension")
	flag.Parse()
}

func main() {
	if ok := isExistDir(dir); !ok || dir == "" {
		usage()
	}

	// Note: fs.FileInfo.Name() が ".png" みたいなものを返すため
	if !strings.HasPrefix(from, ".") {
		from = "." + from
	}

	if !strings.HasPrefix(to, ".") {
		to = "." + to
	}

	files := searchDir(dir, recursive)
	for _, file := range files {
		ext := filepath.Ext(file)
		if ext == to {
			continue
		}

		newFileName := getExtension(file, ext)
		err := converter.Convert(file, newFileName, to)
		if err != nil {
			log.Fatalln(err.Error())
			os.Exit(1)
		}

		fmt.Println("Successfully converted!")
	}
}

func isExistDir(dirName string) bool {
	f, err := os.Stat(dirName)
	return !os.IsNotExist(err) && f.IsDir()
}

func searchDir(dirName string, isRecursive bool) []string {
	var r []string
	fileInfos, err := ioutil.ReadDir(dirName)

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() && isRecursive {
			r = append(r, searchDir(filepath.Join(dirName, fileInfo.Name()), true)...)
		}

		if filepath.Ext(fileInfo.Name()) != from {
			continue
		}

		r = append(r, filepath.Join(dirName, fileInfo.Name()))
	}
	return r
}

func getExtension(filePath, extension string) string {
	return filePath[:len(filePath)-len(extension)] + to
}

func usage() {
	fmt.Printf("Usage:\n")
	fmt.Printf("\tAll of the args are required.\n\n")
	flag.PrintDefaults()
	os.Exit(1)
}
