// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "User CRUD": CLI Commands
//
// Command:
// $ goagen
// --design=GolandProject/goa-sample/src/design
// --out=$(GOPATH)/src/GolandProject/goa-sample/src
// --version=v1.3.1

package cli

import (
	"GolandProject/goa-sample/client"
	"context"
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	// CreateUserCommand is the command line data structure for the create action of user
	CreateUserCommand struct {
		Payload     string
		ContentType string
		// organization ID
		OrganizationID int
		// dummy query parameter
		DummyKey    string
		PrettyPrint bool
	}

	// DeleteUserCommand is the command line data structure for the delete action of user
	DeleteUserCommand struct {
		// organization ID
		OrganizationID int
		// userID
		UserID int
		// dummy query parameter
		DummyKey    string
		PrettyPrint bool
	}

	// ListUserCommand is the command line data structure for the list action of user
	ListUserCommand struct {
		// organization ID
		OrganizationID int
		// dummy query parameter
		DummyKey string
		// filter for user age
		MaxAge      int
		PrettyPrint bool
	}

	// RetrieveUserCommand is the command line data structure for the retrieve action of user
	RetrieveUserCommand struct {
		// organization ID
		OrganizationID int
		// userID
		UserID int
		// dummy query parameter
		DummyKey    string
		PrettyPrint bool
	}

	// UpdateUserCommand is the command line data structure for the update action of user
	UpdateUserCommand struct {
		Payload     string
		ContentType string
		// organization ID
		OrganizationID int
		// userID
		UserID int
		// dummy query parameter
		DummyKey    string
		PrettyPrint bool
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "create",
		Short: `create user`,
	}
	tmp1 := new(CreateUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/organizations/ORGANIZATIONID/users"]`,
		Short: ``,
		Long: `

Payload example:

{
   "age": 8942812999174911700,
   "name": "Vero et sequi assumenda.",
   "programming_skills": [
      {
         "item": "bs6vrw28ln"
      },
      {
         "item": "bs6vrw28ln"
      },
      {
         "item": "bs6vrw28ln"
      }
   ],
   "user_id": 5295941926282978951
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "delete",
		Short: `delete user`,
	}
	tmp2 := new(DeleteUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/organizations/ORGANIZATIONID/users/USERID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "list",
		Short: `get user list`,
	}
	tmp3 := new(ListUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/organizations/ORGANIZATIONID/users"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "retrieve",
		Short: `get user detail`,
	}
	tmp4 := new(RetrieveUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/organizations/ORGANIZATIONID/users/USERID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "update",
		Short: `update user`,
	}
	tmp5 := new(UpdateUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/organizations/ORGANIZATIONID/users/USERID"]`,
		Short: ``,
		Long: `

Payload example:

{
   "age": 8942812999174911700,
   "name": "Vero et sequi assumenda.",
   "programming_skills": [
      {
         "item": "bs6vrw28ln"
      },
      {
         "item": "bs6vrw28ln"
      },
      {
         "item": "bs6vrw28ln"
      }
   ],
   "user_id": 5295941926282978951
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run makes the HTTP request corresponding to the CreateUserCommand command.
func (cmd *CreateUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/organizations/%v/users", cmd.OrganizationID)
	}
	var payload client.User
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateUser(ctx, path, &payload, stringFlagVal("dummy_key", cmd.DummyKey), cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var organizationID int
	cc.Flags().IntVar(&cmd.OrganizationID, "organizationID", organizationID, `organization ID`)
	var dummyKey string
	cc.Flags().StringVar(&cmd.DummyKey, "dummy_key", dummyKey, `dummy query parameter`)
}

// Run makes the HTTP request corresponding to the DeleteUserCommand command.
func (cmd *DeleteUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/organizations/%v/users/%v", cmd.OrganizationID, cmd.UserID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteUser(ctx, path, stringFlagVal("dummy_key", cmd.DummyKey))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var organizationID int
	cc.Flags().IntVar(&cmd.OrganizationID, "organizationID", organizationID, `organization ID`)
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, `userID`)
	var dummyKey string
	cc.Flags().StringVar(&cmd.DummyKey, "dummy_key", dummyKey, `dummy query parameter`)
}

// Run makes the HTTP request corresponding to the ListUserCommand command.
func (cmd *ListUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/organizations/%v/users", cmd.OrganizationID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListUser(ctx, path, cmd.MaxAge, stringFlagVal("dummy_key", cmd.DummyKey))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var organizationID int
	cc.Flags().IntVar(&cmd.OrganizationID, "organizationID", organizationID, `organization ID`)
	var dummyKey string
	cc.Flags().StringVar(&cmd.DummyKey, "dummy_key", dummyKey, `dummy query parameter`)
	var maxAge int
	cc.Flags().IntVar(&cmd.MaxAge, "max_age", maxAge, `filter for user age`)
}

// Run makes the HTTP request corresponding to the RetrieveUserCommand command.
func (cmd *RetrieveUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/organizations/%v/users/%v", cmd.OrganizationID, cmd.UserID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.RetrieveUser(ctx, path, stringFlagVal("dummy_key", cmd.DummyKey))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *RetrieveUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var organizationID int
	cc.Flags().IntVar(&cmd.OrganizationID, "organizationID", organizationID, `organization ID`)
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, `userID`)
	var dummyKey string
	cc.Flags().StringVar(&cmd.DummyKey, "dummy_key", dummyKey, `dummy query parameter`)
}

// Run makes the HTTP request corresponding to the UpdateUserCommand command.
func (cmd *UpdateUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/organizations/%v/users/%v", cmd.OrganizationID, cmd.UserID)
	}
	var payload client.User
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateUser(ctx, path, &payload, stringFlagVal("dummy_key", cmd.DummyKey), cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var organizationID int
	cc.Flags().IntVar(&cmd.OrganizationID, "organizationID", organizationID, `organization ID`)
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, `userID`)
	var dummyKey string
	cc.Flags().StringVar(&cmd.DummyKey, "dummy_key", dummyKey, `dummy query parameter`)
}
