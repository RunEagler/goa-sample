// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "User CRUD": Application Media Types
//
// Command:
// $ goagen
// --design=GolandProject/goa-sample/src/design
// --out=$(GOPATH)/src/GolandProject/goa-sample/src
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// response body type (default view)
//
// Identifier: application/vnd.user.response+json; view=default
type UserResponse struct {
	// user age
	Age int `form:"age" json:"age" yaml:"age" xml:"age"`
	// user name
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// programming_skills
	ProgrammingSkills []*Skill `form:"programming_skills" json:"programming_skills" yaml:"programming_skills" xml:"programming_skills"`
	// user id
	UserID int `form:"user_id" json:"user_id" yaml:"user_id" xml:"user_id"`
}

// Validate validates the UserResponse media type instance.
func (mt *UserResponse) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	if mt.ProgrammingSkills == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "programming_skills"))
	}
	for _, e := range mt.ProgrammingSkills {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeUserResponse decodes the UserResponse instance encoded in resp body.
func (c *Client) DecodeUserResponse(resp *http.Response) (*UserResponse, error) {
	var decoded UserResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// UserResponseCollection is the media type for an array of UserResponse (default view)
//
// Identifier: application/vnd.user.response+json; type=collection; view=default
type UserResponseCollection []*UserResponse

// Validate validates the UserResponseCollection media type instance.
func (mt UserResponseCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeUserResponseCollection decodes the UserResponseCollection instance encoded in resp body.
func (c *Client) DecodeUserResponseCollection(resp *http.Response) (UserResponseCollection, error) {
	var decoded UserResponseCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}