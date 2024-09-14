package main // Entry point of the program

import (
	"fmt"

	"github.com/AnishDhomase/packageLearning/auth"
	"github.com/AnishDhomase/packageLearning/user"
	"github.com/fatih/color"
)

func main() {
	// Packages is a way to organize code in Go, Reuse code
	// Compilation time is faster because of packages
	// If no changes in the package, no need to recompile the package

	// Packages are of two types:
	// 1. Predefined packages: fmt, os, bufio, etc.
	// 2. User-defined packages: Create your own packages

	// Exported names: Capitalized names are exported from the package and can be used outside the package
	// Unexported names: Lowercase names are unexported and can't be used outside the package

	auth.LoginToApp("anish", "password123")
	sessionToken := auth.GetSession()
	fmt.Println("Session token:", sessionToken)

	user1 := user.User{
		Username: "anish",
		Password: "password123",	
	}
	fmt.Println("User created:", user1.Username, user1.Password)


	// Third-party packages: Packages that are not part of the standard library
	// Go get command is used to download third-party packages
	// go get github.com/fatih/color
	// go mod tidy: To download all the dependencies
	color.Red("This is red text")
	color.Blue("This is blue text")
	color.Yellow("This is yellow text")


}