package schema

import (
	"github.com/graphql-go/graphql"
)

var UserList []User

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	Admin     bool   `json:"admin"`
}

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"firstName": &graphql.Field{
			Type: graphql.String,
		},
		"admin": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type:        userType,
			Description: "Get single user",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				idQuery, isOK := params.Args["id"].(string)
				if isOK {
					// Search for el with id
					for _, todo := range UserList {
						if todo.ID == idQuery {
							return todo, nil
						}
					}
				}

				return User{}, nil
			},
		},
	},
})

var UserSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})
