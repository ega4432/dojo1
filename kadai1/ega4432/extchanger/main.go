package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	dir, to, from string
	recursive     bool
)

func init() {
	flag.StringVar(&dir, "d", "", "target directory")
	flag.BoolVar(&recursive, "r", false, "change extension recursively")
	flag.StringVar(&from, "from", "", "from extension")
	flag.StringVar(&to, "to", "", "new extension")
	flag.Parse()
}

func main() {
	if ok := isExistDir(dir); !ok || dir == "" || to == "" || from == "" {
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
		err := convert(file, newFileName)
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

func convert(filePath, newExt string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return errors.New("failed to open file")
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return errors.New("failed to decode image")
	}
	dst, err := os.Create(newExt)

	if err != nil {
		return errors.New("failed to create new file")
	}
	defer dst.Close()

	switch to {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(dst, img, nil)
	case ".png":
		err = png.Encode(dst, img)
	default:
		fmt.Println("default!")
		err = errors.New("invalid extension")
	}

	if err != nil {
		log.Println(err.Error())
		return errors.New("failed to encode image")
	}

	_, err = io.Copy(dst, file)
	if err != nil {
		return errors.New("failed copy file")
	}
	fmt.Printf("Converted:\t %s\t->\t %s\n", filePath, newExt)
	return nil
}

func usage() {
	fmt.Printf("Usage:\n")
	fmt.Printf("\tAll of the args are required.\n\n")
	flag.PrintDefaults()
	os.Exit(1)
}
