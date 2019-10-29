package handler

import (
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	ID       uint
	Email    string
	Password string
	Name     string
	Address  string
	Phone    string
	Avatar   string
}

var Store = sessions.NewCookieStore([]byte("secret-key"))

var dbstring string = "root:root@tcp(192.168.9.22:3306)/demo"

// ConnectDatabase func
// func ConnectDatabase() {
// 	db, err := gorm.Open("mysql", "root:root@tcp(192.168.200.200:3306)/demo")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	defer db.Close()
// }

func InsertData(user User) {
	db, err := gorm.Open("mysql", dbstring)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.Create(&user)
}

func UpdateDataWithAvatar(user User) {
	db, err := gorm.Open("mysql", dbstring)

	mUser := GetUserByEmail(user.Email)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.Model(&mUser).Omit("email").Updates(map[string]interface{}{"password": user.Password, "name": user.Name, "address": user.Address, "phone": user.Phone, "avatar": user.Avatar})
}

func UpdateDataWithoutAvatar(user User) {
	db, err := gorm.Open("mysql", dbstring)

	mUser := GetUserByEmail(user.Email)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.Model(&mUser).Omit("email").Updates(map[string]interface{}{"password": user.Password, "name": user.Name, "address": user.Address, "phone": user.Phone})
}

// func GetUserById(id int) User {
// 	ConnectDatabase()
// 	var user User
// 	db.Where("id = ?", id).First(&user)

// 	return user
// }

func GetUserByEmail(email string) User {
	db, err := gorm.Open("mysql", dbstring)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var user User
	db.Where("email = ?", email).First(&user)

	return user
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
