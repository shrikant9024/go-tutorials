package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/shrikant9024/go-tutorials/auth"
	"github.com/shrikant9024/go-tutorials/user"
)

func main() {

	auth.LoginWithCredentials("Shrikant", "124344")

	userDetails := user.User{
		Name:  "Shrikant",
		Email: "shrikant@gmail.com",
	}

	session := auth.GetUserSession()
	fmt.Println(session)

	// fmt.Println(userDetails.Email)
	// fmt.Println(userDetails.Name)
	color.Cyan(userDetails.Email)
	color.Yellow(userDetails.Name)
}
