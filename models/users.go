package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// ErrNotFound is returned when a recource cannot be found
	// in the database.
	ErrNotFound = errors.New("models: resource not found")
	// Returned when an invalid ID is provided
	// to a method like Delete.
	ErrInvalidID = errors.New("models: ID provided was invalid")
)

type User struct {
	gorm.Model
	Name  string
	Age   uint
	Email string `gorm:"not null;unique_index"`
}

type UserService struct {
	db *gorm.DB
}

func NewUserService(connectionInfro string) (*UserService, error) {
	// Connecting to our database
	db, err := gorm.Open("postgres", connectionInfro)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &UserService{
		db: db,
	}, nil
}

// ByID will look up the user by id provided
// 1 - user, nil
// 2 - nil, ErrNotFound
// 3 - nil, otherError
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	db := us.db.Where("id = ?", id)
	err := first(db, &user)
	return &user, err
}

// Looks up a user with the given email address
func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
	err := first(db, &user)
	return &user, err
}

func (us *UserService) ByAge(age uint) (*User, error) {
	var user User
	db := us.db.Where("age = ?", age)
	err := first(db, &user)
	return &user, err
}

// first will query using the provided gorm.DB and it will
// get the first item returned and place it into dst. If
// nothing is found in query, it will return ErrNotFound
func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}

//Update the provided user with all of the data
func (us *UserService) Update(user *User) error {
	return us.db.Save(user).Error
}

// Delete the user with provided ID
func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	user := User{Model: gorm.Model{ID: id}}
	return us.db.Delete(&user).Error
}

//Create the provided user and backfill data
//  like the ID, CreatedAt, and UpdatedAt
func (us *UserService) Create(user *User) error {
	return us.db.Create(user).Error
}

// Closes the UserService database connection
func (us *UserService) Close() error {
	return us.db.Close()
}

// Drops a users table and rebuilds it
func (us *UserService) DestructiveReset() error {
	if err := us.db.DropTableIfExists(&User{}).Error; err != nil {
		return err
	}
	return us.AutoMigrate()
}

// will attempt to automatically migrate the users table
func (us *UserService) AutoMigrate() error {
	if err := us.db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	return nil
}
