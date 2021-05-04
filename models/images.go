package models

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Image is NOT stored in database
type Image struct {
	GalleryID uint
	Filename  string
}

// func (i *Image) String() string {
// 	return i.Path()
// }

func (i *Image) Path() string {
	return "/" + i.RelativePath()
}

func (i *Image) RelativePath() string {
	return fmt.Sprintf("images/galleries/%v/%v", i.GalleryID, i.Filename)
}

type ImageService interface {
	Create(galleryID uint, r io.ReadCloser, filename string) error
	ByGalleryID(galleryID uint) ([]Image, error)
	Delete(i *Image) error
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

func (is *imageService) ByGalleryID(galleryID uint) ([]Image, error) {
	path := is.imagePath(galleryID)
	stringS, err := filepath.Glob(path + "*")
	if err != nil {
		return nil, err
	}

	ret := make([]Image, len(stringS))
	for i := range stringS {
		stringS[i] = strings.ReplaceAll(stringS[i], "\\", "/") // I think this line is only for Windows!! ????
		stringS[i] = strings.Replace(stringS[i], path, "", 1)
		ret[i] = Image{
			GalleryID: galleryID,
			Filename:  stringS[i],
		}
	}

	return ret, nil
}

func (is *imageService) Delete(i *Image) error {
	return os.Remove(i.RelativePath())
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
