package models

import (
	"github.com/jinzhu/gorm"
)

// Gallery is our image container resources that visitors
// view
type Gallery struct {
	gorm.Model
	UserID uint   `gorm:"not_null;index"`
	Title  string `gorm:"not_null"`
}

type GalleryService interface {
	GalleryDB
}

type galleryValidator struct {
	GalleryDB
}

type galleryValFunc func(*Gallery) error

func runGalleryValFuncs(gallery *Gallery, fns ...galleryValFunc) error {
	for _, fn := range fns {
		if err := fn(gallery); err != nil {
			return err
		}
	}
	return nil
}

func (gv *galleryValidator) Create(gallery *Gallery) error {
	err := runGalleryValFuncs(gallery,
		gv.userIDRequired,
		gv.titleRequired,
	)
	if err != nil {
		return err
	}
	return gv.GalleryDB.Create(gallery)
}

func (gv *galleryValidator) Update(gallery *Gallery) error {
	err := runGalleryValFuncs(gallery,
		gv.userIDRequired,
		gv.titleRequired,
	)
	if err != nil {
		return err
	}
	return gv.GalleryDB.Update(gallery)
}

func (gv *galleryValidator) Delete(id uint) error {
	if id <= 0 {
		return ErrIDInvalid
	}
	return gv.GalleryDB.Delete(id)
}

func (gv *galleryValidator) userIDRequired(g *Gallery) error {
	if g.UserID <= 0 {
		return ErrUserIDRequired
	}
	return nil
}

func (gv *galleryValidator) titleRequired(g *Gallery) error {
	if g.Title == "" {
		return ErrTitleRequired
	}
	return nil
}

func NewGalleryService(db *gorm.DB) GalleryService {
	return &galleryService{
		GalleryDB: &galleryValidator{&galleryGorm{db}},
	}

}

type galleryService struct {
	GalleryDB
}

type GalleryDB interface {
	ByID(id uint) (*Gallery, error)
	Create(gallery *Gallery) error
	Update(gallery *Gallery) error
	Delete(id uint) error
	ByUserID(userID uint) ([]Gallery, error)
}

type galleryGorm struct {
	db *gorm.DB
}

var _ GalleryDB = &galleryGorm{}

func (gg *galleryGorm) Create(gallery *Gallery) error {
	return gg.db.Create(gallery).Error
}

func (gg *galleryGorm) Update(gallery *Gallery) error {
	return gg.db.Save(gallery).Error
}

func (gg *galleryGorm) Delete(id uint) error {
	gallery := Gallery{Model: gorm.Model{ID: id}}
	return gg.db.Delete(&gallery).Error
}

func (gg *galleryGorm) ByID(id uint) (*Gallery, error) {
	var gallery Gallery
	db := gg.db.Where("id = ?", id)
	err := first(db, &gallery)
	return &gallery, err
}

func (gg *galleryGorm) ByUserID(userID uint) ([]Gallery, error) {
	var galleries []Gallery
	gg.db.Where("user_id = ?", userID).Find(&galleries)
	return galleries, nil
}
