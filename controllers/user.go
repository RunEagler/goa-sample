package controllers

import (
	"GolandProject/goa-sample/app"
	"github.com/goadesign/goa"
	"io/ioutil"
	"encoding/json"
)

var newUserID = 5

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	// UserController_Create: start_implement
	println(ctx.Request.Header.Get("TOKEN"))
	newUserID += 1 // new user id
	// Put your logic here

	var response = app.UserResponse{
		UserID:            newUserID,
		Name:              ctx.Payload.Name,
		Age:               ctx.Payload.Age,
		ProgrammingSkills: ctx.Payload.ProgrammingSkills,
	}
	return ctx.Created(&response)
	// UserController_Create: end_implement
}

// Delete runs the delete action.
func (c *UserController) Delete(ctx *app.DeleteUserContext) error {
	// UserController_Delete: start_implement

	// Put your logic here
	println(ctx.Request.Header.Get("TOKEN"))
	return ctx.NoContent()
	// UserController_Delete: end_implement
}

// List runs the list action.
func (c *UserController) List(ctx *app.ListUserContext) error {
	// UserController_List: start_implement

	// Put your logic here
	println(ctx.Request.Header.Get("TOKEN"))

	users, err := readUsersJSON()
	if err != nil {
		return ctx.InternalServerError()
	}

	var response app.UserResponseCollection

	for _,user := range users{
		if user.Age < ctx.MaxAge{
			response = append(response,user)
		}
	}

	return ctx.OK(response)
	// UserController_List: end_implement
}

// Retrieve runs the retrieve action.
func (c *UserController) Retrieve(ctx *app.RetrieveUserContext) error {
	// UserController_Retrieve: start_implement

	// Put your logic here
	println(ctx.Request.Header.Get("TOKEN"))
	users, err := readUsersJSON()
	if err != nil {
		return ctx.InternalServerError()
	}

	user := getUser(users,ctx.UserID)

	if user == nil{
		return ctx.NotFound()
	}else{
		return ctx.OK(user)
	}
	// UserController_Retrieve: end_implement
}

// Update runs the update action.
func (c *UserController) Update(ctx *app.UpdateUserContext) error {
	// UserController_Update: start_implement

	// Put your logic here
	println(ctx.Request.Header.Get("TOKEN"))
	return ctx.NoContent()
	// UserController_Update: end_implement
}

func readUsersJSON() (app.UserResponseCollection, error) {

	var users app.UserResponseCollection

	raw, err := ioutil.ReadFile("./public/json/users.json")
	if err != nil {
		return nil, err
	}
	json.Unmarshal(raw, &users)

	return users, nil
}


func getUser(users app.UserResponseCollection,userID int) *app.UserResponse{

	for _, user := range users {
		if user.UserID == userID {
			return user
		}
	}
	return nil
}