package main

import (
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

const jwtSecret = "igeargeek"

func main() {

	var err error
	dsn := "root:1234@tcp(127.0.0.1:3306)/igeargeekpracticdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Get("/profile", getAllProfiles)
	app.Post("/signup", signup)
	app.Post("/login", login)
	app.Delete("/profile/:id", deleteProfiles)
	app.Put("/profile", updateProfiles)

	if err := app.Listen(":8000"); err != nil {
		panic(err)
	}
}

func signup(c *fiber.Ctx) error {
	var request User
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	if request.Username == "" || request.Password == "" || strconv.Itoa(request.Age) == "0" || request.Name == "" {
		return fiber.ErrUnprocessableEntity
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return err
	} else {
		request.Password = string(password)
	}

	if db.Migrator().HasTable(&User{}) {
		if result := db.Create(&request); result.Error != nil {
			return result.Error
		}
	} else {
		db.Migrator().CreateTable(&User{})
		if result := db.Create(&request); result.Error != nil {
			return result.Error
		}
	}

	return c.Status(fiber.StatusCreated).SendString("Updated")
}

func login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var request LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	if request.Username == "" || request.Password == "" {
		return fiber.ErrUnprocessableEntity
	}

	var user User
	if result := db.Where("username = ?", request.Username).Find(&user); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, "Incorrect Username or Password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Incorrect Username or Password")
	}

	return c.Status(fiber.StatusOK).JSON(&user)
}

func getAllProfiles(c *fiber.Ctx) error {
	var people []User

	if err := db.Find(&people).Error; err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&people)
}

func deleteProfiles(c *fiber.Ctx) error {
	id := c.Params("id")

	if result := db.Delete(&User{}, id); result.Error != nil {
		return result.Error
	}

	return c.SendString("Deleted")
}

func updateProfiles(c *fiber.Ctx) error {
	type UpdateRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      int    `json:"age"`
	}
	var request UpdateRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	if request.Username == "" || request.Password == "" || (request.Name == "" && strconv.Itoa(request.Age) == "0"){
		return fiber.ErrUnprocessableEntity
	}

	var user User
	if result := db.Where("username = ?",request.Username).Find(&user); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound,"Username or Password not have in database")
	}
	
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(request.Password)); err != nil {
		return fiber.NewError(fiber.StatusNotFound,"Username or Password not have in database")
	}

	if request.Name != "" && strconv.Itoa(request.Age) != "0" {
		db.Model(&User{}).Where("username = ?",request.Username).Update("name",request.Name)
		db.Model(&User{}).Where("username = ?",request.Username).Update("age",request.Age)
	} else if request.Name != "" && strconv.Itoa(request.Age) == "0" {
		db.Model(&User{}).Where("username = ?",request.Username).Update("name",request.Name)
	} else if request.Name == "" && strconv.Itoa(request.Age) != "0" {
		db.Model(&User{}).Where("username = ?",request.Username).Update("age",request.Age)
	}

	return c.Status(fiber.StatusOK).SendString("Putted")
}

type User struct {
	// Id int `gorm:"primaryKey" json:"id"`
	gorm.Model
	Username string `gorm:"column:username;varchar(20)" json:"username"`
	Password string `gorm:"column:password;varchar(255)" json:"password"`
	Name     string `gorm:"columns:name;varchar(50) json:"name"`
	Age      int    `gorm:"column:age;int" json:"age"`
}
