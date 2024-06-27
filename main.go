package main

import (
	calculator "bkoiki950/test-go/calculator"
	"bkoiki950/test-go/helpers"
	"fmt"
)


func func1 (user chan int ) {
	user <- 1
}

func func2 (user chan int ) {
	user <- 2
}

func main () {
	u := []helpers.User{
		{ ID: 1, Name: "Babatunde"},
		{ ID: 2, Name: "Koiki"},
	}

	db := helpers.InitDB(u...)
	db.CreateUser("NewUser")
	db.CreateUser("AnotherUser")
	db.CreateUser("AUser")
	fmt.Println(db.GetUsers())
	fmt.Println(db.GetRandonmUser())
	fmt.Println(db.GetUserById(30))
	db.UpdateUser(1, "Babalola")
	fmt.Println(db.GetUsers())
	db.DeleteUser(2)
	fmt.Println(db.GetUsers())
	db.CreateUser("RandomUser")
	fmt.Println(db.GetUsers())

	user := make(chan int)
	user2 := make(chan int)

	go func1(user)
	go func2(user2)

	select {
	case msgFromfunc1 := <- user:
		fmt.Println("Received", msgFromfunc1)
	case msgFromfunc2 := <- user2:
		fmt.Println("Received", msgFromfunc2)	
	}

	fmt.Println(calculator.Add(1, 2))
	fmt.Println(calculator.Multiply(1, 2))
}
