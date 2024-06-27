package helpers

import (
	"errors"
	"fmt"
)

func (u User) String () string {
	return fmt.Sprintf("ID: %d, Name: %s", u.ID, u.Name)
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

func (db *DB) GetUserById (id int) (User, error) {
	if id >= len(db.users) || id < 1 {
		err := errors.New("User not found")
		return User{}, err
	}
	u := make(chan User)
	go getUserByIdInternal(db.users, id, u)

	return <- u, nil
}

func (db *DB) UpdateUser (id int, name string) (User, error) {
	if id >= len(db.users) || id < 1 {
		err := errors.New("User not found")
		return User{}, err
	}
	u := make(chan User)
	go updateUserInternal(&db.users, id, name, u)

	return <- u, nil
}

func (db *DB) DeleteUser (id int) (User, error) {
	if id >= len(db.users) || id < 1 {
		err := errors.New("User not found")
		return User{}, err
	}
	u := make(chan User)
	go deleteUserInternal(&db.users, id, u)

	return <- u, nil
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
