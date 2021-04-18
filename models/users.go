package models

import (
	"errors"
	"learn-web-dev-with-go/hash"
	"learn-web-dev-with-go/rand"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrNotFound is returned when a recource cannot be found
	// in the database.
	ErrNotFound = errors.New("models: resource not found")
	// Returned when an invalid ID is provided
	// to a method like Delete.
	ErrInvalidID = errors.New("models: ID provided was invalid")

	// Returned when an invalid password is used
	// when attempting to authenticate a user
	ErrInvalidPassword = errors.New("models: incorrect password provided")
)

const userPwPepper = "secret-random-string"
const hmacSecretKey = "secret-hmac-key"

// User represents the user model stored in our database
// This is used for user accounts, storing both an email
// address and a password so users can log in and gain
// access to their content.
type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"not null;unique_index"`
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
	Remember     string `gorm:"-"`
	RememberHash string `gorm:"not null;unique_index"`
}

// UserDB is used to interact with the users database
//
// There are returns for pretty much all user single queries:
// Case 1: User is found - user, nil
// Case 2: User not found - nil, ErrNotFound
// Case 3: Another error - nil, OtherError
//
// For single queries, any error except ErrNotFound should
// probably result in a 500 error.
type UserDB interface {
	// Methods for querying for single users
	ByID(id uint) (*User, error)
	ByEmail(email string) (*User, error)
	ByRemember(token string) (*User, error)

	// Methods for altreing users
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error

	// Used to close a DB connection
	Close() error

	// Migration helpers
	AutoMigrate() error
	DestructiveReset() error
}

// UserService is a set of methods used to manipulate and
// work with the user model
type UserService interface {
	// Authenticate will verify the provided email address and
	// password are correct. If they are correct, the user
	// corresponding to that email will be returned. Otherwise
	// You will recieve either:
	// ErrNotFound, ErrInvalidPassword or another error if
	// something goes wrong
	Authenticate(email, password string) (*User, error)
	UserDB
}

func NewUserService(connectionInfro string) (UserService, error) {
	ug, err := newUserGorm(connectionInfro)
	if err != nil {
		return nil, err
	}
	return &userService{
		UserDB: &userValidator{
			UserDB: ug,
		},
	}, nil
}

// Checking that userService is implements all functionality
// that UserDB interface requires
var _ UserService = &userService{}

type userService struct {
	UserDB
}

// can be used to authenticate a user with the
// provided email and password.
// If the email address provided is invalid, this will return
//  nil, ErrNotFound
// If the password provided is invalid, this will return
//  nil, ErrInvalidPassword
// Otherwise if another error is encountered this will return
//  nil, error
func (us *userService) Authenticate(email, password string) (*User, error) {
	foundUser, err := us.ByEmail(email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password+userPwPepper))
	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return nil, ErrInvalidPassword
		default:
			return nil, err
		}
	}
	return foundUser, nil
}

type userValidator struct {
	UserDB
}

var _ UserDB = &userValidator{}

type userGorm struct {
	db   *gorm.DB
	hmac hash.HMAC
}

var _ UserDB = &userGorm{}

func newUserGorm(connectionInfro string) (*userGorm, error) {
	// Connecting to our database
	db, err := gorm.Open("postgres", connectionInfro)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	hmac := hash.NewHMAC(hmacSecretKey)
	return &userGorm{
		db:   db,
		hmac: hmac,
	}, nil
}

// ByID will look up the user by id provided
// 1 - user, nil
// 2 - nil, ErrNotFound
// 3 - nil, otherError
func (ug *userGorm) ByID(id uint) (*User, error) {
	var user User
	db := ug.db.Where("id = ?", id)
	err := first(db, &user)
	return &user, err
}

// Looks up a user with the given email address
func (ug *userGorm) ByEmail(email string) (*User, error) {
	var user User
	db := ug.db.Where("email = ?", email)
	err := first(db, &user)
	return &user, err
}

// ByRemember looks up a user with the given remember token
// and returns that user. This method will handle hashing
// the token for us
// Errors are the same as ByEmail and ByID
func (ug *userGorm) ByRemember(token string) (*User, error) {
	var user User
	rememberHash := ug.hmac.Hash(token)
	err := first(ug.db.Where("remember_hash = ?", rememberHash), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

//Update the provided user with all of the data
func (ug *userGorm) Update(user *User) error {
	if user.Remember != "" {
		user.RememberHash = ug.hmac.Hash(user.Remember)
	}
	return ug.db.Save(user).Error
}

// Delete the user with provided ID
func (ug *userGorm) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	user := User{Model: gorm.Model{ID: id}}
	return ug.db.Delete(&user).Error
}

//	Create the provided user and backfill data
//  like the ID, CreatedAt, and UpdatedAt
func (ug *userGorm) Create(user *User) error {
	// pass, salt
	// hash(pass+salt)
	// salt = per user (bcrypt does salt for us (?))
	// pepper = per application
	pwBytes := []byte(user.Password + userPwPepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedBytes)
	user.Password = ""
	if user.Remember == "" {
		token, err := rand.RememberToken()
		if err != nil {
			return err
		}
		user.Remember = token
	}
	user.RememberHash = ug.hmac.Hash(user.Remember)
	return ug.db.Create(user).Error
}

// Closes the UserService database connection
func (ug *userGorm) Close() error {
	return ug.db.Close()
}

// Drops a users table and rebuilds it
func (ug *userGorm) DestructiveReset() error {
	if err := ug.db.DropTableIfExists(&User{}).Error; err != nil {
		return err
	}
	return ug.AutoMigrate()
}

// will attempt to automatically migrate the users table
func (ug *userGorm) AutoMigrate() error {
	if err := ug.db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	return nil
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
