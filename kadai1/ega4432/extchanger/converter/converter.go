package converter

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
)

func Convert(filePath, newExt, to string) error {
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
