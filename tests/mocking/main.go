package main

import (
	"log"

	"github.com/rAndrade360/go-mock-tests/repositories"
	"github.com/rAndrade360/go-mock-tests/services"
)

func main() {
	name := "Renan"
	email := "renandotcorrea@gmail.com"

	userRepository := repositories.NewUserRepository(nil)

	userService := services.NewUserService(userRepository)

	err := userService.Create(name, email)
	if err != nil {
		log.Println("Error: ", err.Error())
	}
}
