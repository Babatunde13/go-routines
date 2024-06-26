package main

import (
	"errors"
	"fmt"
	"math/rand"
)


type User struct {
	ID int
	Name string
}

type DB struct {
	users []User
}

func (u User) String () string {
	return fmt.Sprintf("ID: %d, Name: %s", u.ID, u.Name)
}

func getRandomUserInternal (users []User, user chan User) {
	randIdx := rand.Intn(len(users))
	user <- users[randIdx]
}

func getUserByIdInternal (users []User, id int, user chan User) {
	user <- users[id-1]
}

func updateUserInternal (users *[]User, id int, name string, user chan User) {
	u := (*users)[id-1]
	u.Name = name
	(*users)[u.ID-1] = u
	user <- u
}

func deleteUserInternal (users *[]User, id int, user chan User) {
	u := (*users)[id-1]
	*users = append((*users)[:id-1], (*users)[id:]...)
	user <- u
}

func createUserInternal (users *[]User, newUser User, user chan User) {
	*users = append(*users, newUser)
	user <- newUser
}

func getUsersInternal (users []User, user chan *[]User) {
	user <- &users
}


func InitDB (users ...User) DB {
	if len(users) == 0 {
		users = []User{
			{ ID: 1, Name: "Babatunde"},
				{ ID: 2, Name: "Koiki"},
				{ ID: 3, Name: "Joseph"},
				{ ID: 4, Name: "Tsegen"},
				{ ID: 5, Name: "Ayo"},
		}
	}

	return DB{users: users}
}

func (db *DB) GetRandonmUser () User {
	data := make(chan User)
	go getRandomUserInternal(db.users, data)

	return <- data
}

func (db *DB) GetUserById (id int) (error, User) {
	if id >= len(db.users) || id < 1 {
		err := errors.New("User not found")
		return err, User{}
	}
	u := make(chan User)
	go getUserByIdInternal(db.users, id, u)

	return nil, <- u
}

func (db *DB) UpdateUser (id int, name string) (error, User) {
	if id >= len(db.users) || id < 1 {
		err := errors.New("User not found")
		return err, User{}
	}
	u := make(chan User)
	go updateUserInternal(&db.users, id, name, u)

	return nil, <- u
}

func (db *DB) DeleteUser (id int) (error, User) {
	if id >= len(db.users) || id < 1 {
		err := errors.New("User not found")
		return err, User{}
	}
	u := make(chan User)
	go deleteUserInternal(&db.users, id, u)

	return nil, <- u
}


func (db *DB) CreateUser (name string) User {
	user := User{
		ID: db.users[len(db.users)-1].ID + 1,
		Name: name,
	}
	data := make(chan User)
	go createUserInternal(&db.users, user, data)

	return <- data
}

func (db DB) GetUsers () *[]User {
	data := make(chan *[]User)
	go getUsersInternal(db.users, data)
	return <- data
}


func func1 (user chan int ) {
	user <- 1
}

func func2 (user chan int ) {
	user <- 2
}

func main () {
	u := []User{
		{ ID: 1, Name: "Babatunde"},
		{ ID: 2, Name: "Koiki"},
	}

	db := InitDB(u...)
	db.CreateUser("NewUser")
	db.CreateUser("AnotherUser")
	db.CreateUser("AUser")
	fmt.Println(db.users)
	fmt.Println(db.GetRandonmUser())
	fmt.Println(db.GetUserById(30))
	db.UpdateUser(1, "Babalola")
	fmt.Println(db.users)
	db.DeleteUser(2)
	fmt.Println(db.users)
	db.CreateUser("RandomUser")
	fmt.Println(db.users)
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
}
