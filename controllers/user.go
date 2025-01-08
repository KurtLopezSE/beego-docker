package controllers

import (
	"encoding/json"
	"new-beego-api/models"
	"new-beego-api/services"

	beego "github.com/beego/beego/v2/server/web"
)

// UserController handles user-related requests
type UserController struct {
	beego.Controller
	userService *services.UserService
}

// Prepare initializes the UserService
func (u *UserController) Prepare() {
	u.userService = services.NewUserService()
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} map[string]interface{}
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &user); err == nil {
		id, err := u.userService.AddUser(&user)
		if err == nil {
			u.Data["json"] = map[string]interface{}{"id": id}
		} else {
			u.Data["json"] = err.Error()
		}
	} else {
		u.Data["json"] = "Invalid input"
	}
	u.ServeJSON()
}

// @Title GetUser
// @Description get user by ID
// @Param	id		path 	int	true		"The ID of the user"
// @Success 200 {object} models.User
// @Failure 404 user not found
// @router /:id [get]
func (u *UserController) Get() {
	id, _ := u.GetInt64(":id")
	user, err := u.userService.GetUser(id)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = user
	}
	u.ServeJSON()
}

// @Title GetAllUsers
// @Description get all users
// @Success 200 {array} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users, err := u.userService.GetAllUsers()
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = users
	}
	u.ServeJSON()
}

// @Title UpdateUser
// @Description update the user
// @Param	id		path 	int	true		"The ID of the user"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {string} "User updated successfully"
// @Failure 403 invalid input
// @router /:id [put]
func (u *UserController) Put() {
	id, _ := u.GetInt64(":id")
	var user models.User
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &user); err == nil {
		user.Id = id
		if err := u.userService.UpdateUser(&user); err == nil {
			u.Data["json"] = "User updated successfully"
		} else {
			u.Data["json"] = err.Error()
		}
	} else {
		u.Data["json"] = "Invalid input"
	}
	u.ServeJSON()
}

// @Title DeleteUser
// @Description delete the user
// @Param	id		path 	int	true		"The ID of the user"
// @Success 200 {string} "User deleted successfully"
// @Failure 404 user not found
// @router /:id [delete]
func (u *UserController) Delete() {
	id, _ := u.GetInt64(":id")
	if err := u.userService.DeleteUser(id); err == nil {
		u.Data["json"] = "User deleted successfully"
	} else {
		u.Data["json"] = err.Error()
	}
	u.ServeJSON()
}
