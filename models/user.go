package models

import (
	"github.com/beego/beego/v2/client/orm"
)

// User represents the user model
type User struct {
	Id       int64  `orm:"auto"` // Auto-incrementing ID
	Username string `orm:"size(100)"`
	Password string `orm:"size(100)"`
	Email    string `orm:"size(100)"`
}

// Register the User model
func init() {
	orm.RegisterModel(new(User))
}

// AddUser inserts a new user into the database
func AddUser(user *User) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(user)
	return id, err
}

// GetUser retrieves a user by ID
func GetUser(id int64) (*User, error) {
	o := orm.NewOrm()
	user := &User{Id: id}
	err := o.Read(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetAllUsers retrieves all users from the database
func GetAllUsers() ([]User, error) {
	o := orm.NewOrm()
	var users []User
	_, err := o.QueryTable(new(User)).All(&users)
	return users, err
}

// UpdateUser updates an existing user in the database
func UpdateUser(user *User) error {
	o := orm.NewOrm()
	_, err := o.Update(user)
	return err
}

// DeleteUser deletes a user from the database
func DeleteUser(id int64) error {
	o := orm.NewOrm()
	user := User{Id: id}
	_, err := o.Delete(&user)
	return err
}
