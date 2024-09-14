package auth

import "fmt"

// Capitalized function name means it is exported and can be used outside the package
func LoginToApp(username, password string) {
	// Login logic
	fmt.Println("Logging in", username, "with password", password)
}
