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

func initDB () DB {
	users := []User{
		{ ID: 1, Name: "Babatunde"},
		{ ID: 2, Name: "Koiki"},
		{ ID: 3, Name: "Joseph"},
		{ ID: 4, Name: "Tsegen"},
		{ ID: 5, Name: "Ayo"},
	}

	return DB{users: users}
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

func createUserInternal (users *[]User, newUser User, user chan User) {
	*users = append(*users, newUser)
	user <- newUser
}

func (db *DB) getRandonmUser () User {
	data := make(chan User)
	go getRandomUserInternal(db.users, data)

	return <- data
}

func (db *DB) getUserById (id int) (error, User) {
	if id >= len(db.users) || id < 1 {
		err := errors.New("User not found")
		return err, User{}
	}
	u := make(chan User)
	go getUserByIdInternal(db.users, id, u)

	return nil, <- u
}

func (db *DB) updateUser (id int, name string) (error, User) {
	if id >= len(db.users) || id < 1 {
		err := errors.New("User not found")
		return err, User{}
	}
	u := make(chan User)
	go updateUserInternal(&db.users, id, name, u)

	return nil, <- u
}

func (db *DB) createUser (name string) User {
	user := User{
		ID: len(db.users)+1,
		Name: name,
	}
	data := make(chan User)
	go createUserInternal(&db.users, user, data)

	return <- data
}

func main () {
	db := initDB()
	db.createUser("NewUser")
	db.createUser("AnotherUser")
	db.createUser("AUser")
	fmt.Println(db.users)
	fmt.Println(db.getRandonmUser())
	fmt.Println(db.getUserById(30))
	db.updateUser(1, "Babalola")
	fmt.Println(db.users)
}
