package helpers

import "math/rand"

type User struct {
	ID int
	Name string
}

type DB struct {
	users []User
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
