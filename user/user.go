package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "your data base connection"

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to Database")
	}

	DB.AutoMigrate(&User{})
}

func GetUsers(c *fiber.Ctx) error {

}

func GetUser(c *fiber.Ctx) error {
	
}

func SaveUser(c *fiber.Ctx) error {
	
}

func DeleteUser(c *fiber.Ctx) error {
	
}

func UpdateUser(c *fiber.Ctx) error {
	
}