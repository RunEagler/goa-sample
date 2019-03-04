package design // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("User CRUD", func() {
	Title("goa sample")
	Description("goa sample")
	Scheme("http")
	Host("localhost:8080")
})

var _ = Resource("user", func() {
	BasePath("/organizations/:organizationID/users")
	Headers(func() {
		Header("TOKEN", String, "authentication token")
	})
	Params(func() {
		Param("organizationID", Integer, "organization ID")
		Param("dummy_key", String, "dummy query parameter",func(){
			MaxLength(10)
		})
	})

	Action("list", func() {
		Description("get user list")
		Routing(GET(""))
		Params(func() {
			Param("max_age", Integer, "filter for user age")
			Required("max_age")
		})

		Response(OK, UserMediaList) // Responses define the shape and status code
	})

	Action("retrieve", func() {
		Description("get user detail")
		Routing(GET("/:userID"))
		Params(func() {
			Param("userID", Integer, "userID")
		})
		Response(OK,UserMedia)       // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})

	Action("create", func() {
		Description("create user")
		Routing(POST(""))
		Payload(User)
		Response(Created, UserMedia) // Responses define the shape and status code
	})

	Action("delete", func() {
		Description("delete user")
		Routing(DELETE("/:userID"))
		Params(func() { // (shape of the request body).
			Param("userID", Integer, "userID")
		})
		Response(NoContent) // Responses define the shape and status code
		Response(NotFound)  // of HTTP responses.
	})

	Action("update", func() {
		Description("update user")
		Routing(PUT("/:userID"))
		Params(func() { // (shape of the request body).
			Param("userID", Integer, "userID")
		})
		Payload(User)
		Response(NoContent) // Responses define the shape and status code
		Response(NotFound)  // of HTTP responses.
	})
	Response(InternalServerError)
})

// User user type
var User = Type("User", func() {
	Description("request body type")
	Member("user_id", Integer, "user id")
	Member("name", String, "user name")
	Member("age", Integer, "user age")
	Member("programming_skills", Skills, "programming skills")

	Required("name", "age", "programming_skills")
})

// Skill skill type
var Skill = Type("Skill", func() {
	Attribute("item", String, "skill name", func() {
		MaxLength(20)
	})
})

// Skills skill list type
var Skills = ArrayOf(Skill)

// UserMedia response body for user resource
var UserMedia = MediaType("application/vnd.user.response+json", func() {
	Description("response body type")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("user_id", Integer, "user id")
		Attribute("name", String, "user name")
		Attribute("age", Integer, "user age")
		Attribute("programming_skills", Skills, "programming_skills")
		Required("user_id", "name", "age", "programming_skills")
	})
	View("default", func() {
		Attribute("user_id")
		Attribute("name")
		Attribute("age")
		Attribute("programming_skills")
	})
})

var UserMediaList = CollectionOf(UserMedia)
