package main

import "fmt"

type Direction int
const (
	Up Direction = iota
	Down
	Left
	Right
)

type User string
const (
	Admin User = "admin"
	Staff User = "staff"
)

func move(dir Direction) {
	switch dir {
	case Up:
		fmt.Println("Moving Up")
	case Down:
		fmt.Println("Moving Down")
	case Left:
		fmt.Println("Moving Left")
	case Right:
		fmt.Println("Moving Right")
	}
}

func main() {
	move(Up)
	fmt.Println(Up)
}