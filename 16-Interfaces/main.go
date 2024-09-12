package main

import "fmt"

type PaymentGateway interface {
	chargeUser(amount int, user User)
}

// type User struct {
// 	id int
// 	gateway Stripe
// }

type User struct {
	id int
	gateway PaymentGateway // Allows all gateways that has chargeUser method 
	// with same signature
}
func (U User) makePayment(amount int) {
	U.gateway.chargeUser(amount, U)
}

type Stripe struct{}
func (S Stripe) chargeUser(amount int, user User) {
	fmt.Println("Charging user", user.id, "with amount", amount)
}

type Razorpay struct{}
func (R Razorpay) chargeUser(amount int, user User) {
	fmt.Println("Charging user", user.id, "with amount", amount)
}

func main() {
	user1 := User{id: 1, gateway: Stripe{}}
	// user2 := User{id: 1, gateway: Razorpay{}} 
	// To use Razorpay instead of Stripe we need to change the gateway type in User struct
	// This violates the Open-Closed Principle
	user2 := User{id: 2, gateway: Razorpay{}}
	user1.makePayment(100)
	user2.makePayment(200)
}