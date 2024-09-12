package main

import (
	"fmt"
	"time"
)

// Basic
type Order struct {
	id        int
	product   string
	quantity  int
	status    string
	createdAt time.Time
}
// Constructor
func newOrder(id int, product string, quantity int, status string, createdAt time.Time) *Order {
	order := Order{
		id:        id,
		product:   product,
		quantity:  quantity,
		status:    status,
		createdAt: createdAt,
	}
	return &order
}
// Methods 
func (O Order) changeOrderStatus1(newStat string) {
	O.status = newStat
}
func (O *Order) changeOrderStatus(newStat string) {
	O.status = newStat
}
func (O Order) getOrderId() int {
	return O.id
}

// Embedding
type User struct {
	id   int
	name string
}
type Comment struct {
	id      int
	content string
	User
}


func main() {
	// Basic
	order1 := Order{
		id:        1,
		product:   "Phone",
		status:    "Pending",
		createdAt: time.Now(),
	}
	// if we don't set any feild, it will be set to default value 
	order1.quantity = 2
	order1.changeOrderStatus1("Shipped") // This will not change the status
	order1.changeOrderStatus("Confirmed") // This will change the status
	fmt.Println(order1.getOrderId())
	fmt.Println(order1)

	// Using constructor
	order2 := newOrder(2, "Laptop", 1, "Pending", time.Now())
	fmt.Println(order2, order2.getOrderId())

	// One time instance creation
	config := struct {
		version	string
		createdOn	time.Time
	}{"1.0.0", time.Now()}
	fmt.Println(config)



	// Embedding
	user1 := User{1, "Anish"}
	comment1 := Comment{1, "Nice post", user1}
	comment2 := Comment{
		id: 2,
		content: "Great post",
		User: User{
			id: 2,
			name: "Anish",
		},
	}
	comment1.User.name = "Anish Dhomase"
	fmt.Println(comment1, comment1.name)
	fmt.Println(comment2)

}