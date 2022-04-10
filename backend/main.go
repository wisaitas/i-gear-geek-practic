package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	// jwtmiddleware "github.com/gofiber/jwt/v3"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

const jwtSecret = "igeargeekjiw"

func main() {

	var err error
	dsn := "root:1234@tcp(127.0.0.1:3306)/igeargeekpracticdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	// app.Use("/profiles",jwtmiddleware.New(jwtmiddleware.Config{
	// 	SigningMethod: "HS256",
	// 	SigningKey: []byte(jwtSecret),
	// 	SuccessHandler: func(c *fiber.Ctx) error {
	// 		return c.Next()
	// 	},
	// 	ErrorHandler: func(c *fiber.Ctx,e error) error {
	// 		return fiber.ErrUnauthorized
	// 	},
	// }))
	
	app.Use(logger.New())

	app.Get("/profiles", getAllProfiles)
	app.Get("/profile",getProfile)
	app.Post("/signup", signup)
	app.Post("/login", login)
	app.Post("/logout",logout)
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

	if request.Username == "" || request.Password == "" || strconv.Itoa(request.UserDetail.Age) == "0" || request.UserDetail.Fname == "" || request.UserDetail.Lname == ""{
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

	var user User
	claims := jwt.StandardClaims{
		Issuer: strconv.Itoa(user.Id),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	token , err := jwtToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return c.SendString("Error token")
	}

	db.Model(&User{}).Where("username = ?",request.Username).Update("auth_jwt",token)

	return c.Status(fiber.StatusCreated).SendString("Signup Complete")
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

	return c.JSON(user)
}

func logout(c *fiber.Ctx) error {
	return c.SendString("Logout Success")
}

func getProfile(c *fiber.Ctx) error {
	var user User
	// valueFromHeader := string(request)
	valueFromHeader := c.Get("Authorization")
	splitStr := strings.Split(valueFromHeader," ")
	if result := db.Where("auth_jwt = ?",splitStr[1]).Find(&user); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, "Incorrect Token matched")
	}

	// return c.JSON(user)
	return c.Status(fiber.StatusOK).JSON(user)
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
		UserDetail struct {
			Fname string `json:"first_name"`
			Lname string `json:"last_name"`
			Image_src string `json:"image_src"`
			Age int `json:"age"`
		} `json:"userdetail"`
	}
	var request UpdateRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	if request.Username == "" || request.Password == "" || (request.UserDetail.Fname == "" && strconv.Itoa(request.UserDetail.Age) <= "0" && request.UserDetail.Lname == ""){
		return fiber.ErrUnprocessableEntity
	}

	var user User
	if result := db.Where("username = ?",request.Username).Find(&user); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound,"Username or Password not have in database")
	}
	
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(request.Password)); err != nil {
		return fiber.NewError(fiber.StatusNotFound,"Username or Password not have in database")
	}

	if request.UserDetail.Fname != "" && strconv.Itoa(request.UserDetail.Age) > "0" && request.UserDetail.Lname != "" && request.UserDetail.Image_src != "" {
		db.Model(&User{}).Where("username = ?",request.Username).Updates(map[string]interface{}{
			"first_name":request.UserDetail.Fname,
			"last_name":request.UserDetail.Lname,
			"age":request.UserDetail.Age,
			"image_src":request.UserDetail.Image_src,
		})
	} else if request.UserDetail.Fname != "" && strconv.Itoa(request.UserDetail.Age) == "0" && request.UserDetail.Lname == "" && request.UserDetail.Image_src == "" {
		db.Model(&User{}).Where("username = ?",request.Username).Update("name",request.UserDetail.Fname)
	} else if request.UserDetail.Fname == "" && strconv.Itoa(request.UserDetail.Age) != "0" && request.UserDetail.Lname == "" && request.UserDetail.Image_src == "" {
		db.Model(&User{}).Where("username = ?",request.Username).Update("age",request.UserDetail.Age)
	} else if request.UserDetail.Fname == "" && strconv.Itoa(request.UserDetail.Age) == "0" && request.UserDetail.Lname != "" && request.UserDetail.Image_src == "" {
		db.Model(&User{}).Where("username = ?",request.Username).Update("age",request.UserDetail.Lname)
	} else if request.UserDetail.Fname == "" && strconv.Itoa(request.UserDetail.Age) == "0" &&  request.UserDetail.Lname == "" && request.UserDetail.Image_src != "" {
		db.Model(&User{}).Where("username = ?",request.Username).Update("age",request.UserDetail.Image_src)
	}

	return c.Status(fiber.StatusOK).SendString("Putted")
}
type UserDetail struct {
	Fname string `gorm:"column:first_name;varchar(20)" json:"first_name"`
	Lname string `gorm:"column:last_name;varchar(20)" json:"last_name"`
	Image_src string `gorm:"column:image_src;varchar(255)" json:"image_src"`
	Age int `gorm:"column:age;int" json:"age"`
	Auth_jwt string `gorm:"column:auth_jwt;varchar(255)" json:"auth_jwt"`
}

type User struct {
	Id int `gorm:"primaryKey" json:"id"`
	Username string `gorm:"column:username;varchar(20)" json:"username"`
	Password string `gorm:"column:password;varchar(255)" json:"password"`
	UserDetail UserDetail `gorm:"embedded;varchar(11)" json:"userdetail"`
}
