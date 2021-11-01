package main

import (
	"log"
	"startup-rest-api/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "itsme:Roots12345!@tcp(127.0.0.1:3306)/startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "sri andeani"
	userInput.Email = "sri@mail.com"
	userInput.Occupation = "ui/ux designer"
	userInput.Password = "password"

	userService.RegisterUser(userInput)

}
