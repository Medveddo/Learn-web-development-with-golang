package models

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type ImageService interface {
	Create(galleryID uint, r io.ReadCloser, filename string) error
	ByGalleryID(galleryID uint) ([]string, error)
}

func NewImageService() ImageService {
	return &imageService{}

}

type imageService struct {
}

// images/galleries/:id/:filename
func (is *imageService) Create(galleryID uint, r io.ReadCloser, filename string) error {
	defer r.Close()
	path, err := is.mkImagePath(galleryID)
	if err != nil {
		return err
	}

	// Create a destination file
	dst, err := os.Create(path + filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy reader data to the destination file
	_, err = io.Copy(dst, r)
	if err != nil {
		return err
	}
	return nil
}

func (is *imageService) ByGalleryID(galleryID uint) ([]string, error) {
	path := is.imagePath(galleryID)
	stringS, err := filepath.Glob(path + "*")
	if err != nil {
		return nil, err
	}

	// separator := "" // IDK this code helps me or not
	// if runtime.GOOS == "windows" {
	// 	separator = "\\"
	// } else {
	// 	separator = "/"
	// }

	separator := "/"

	for i := range stringS {
		stringS[i] = strings.ReplaceAll(stringS[i], "\\", "/") // I think this line is only for Windows!! ????
		stringS[i] = separator + stringS[i]
	}

	return stringS, nil
}

func (is *imageService) imagePath(galleryID uint) string {
	return fmt.Sprintf("images/galleries/%v/", galleryID)
}

func (is *imageService) mkImagePath(galleryID uint) (string, error) {
	galleryPath := is.imagePath(galleryID)
	err := os.MkdirAll(galleryPath, 0755)
	if err != nil {
		return "", err
	}
	return galleryPath, nil
}
