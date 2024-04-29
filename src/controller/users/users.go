package users

import "time"

var usr = make(map[int64]*Users)
var usrCount int64

func GetUserById(id int64) *Users {
	return usr[id]
}

func CreateUser(name string) (int64, error) {
	u := Users{
		Id:        usrCount,
		Name:      name,
		CreatedOn: time.Now(),
	}
	usrCount++
	usr[u.Id] = &u
	return u.Id, nil
}
