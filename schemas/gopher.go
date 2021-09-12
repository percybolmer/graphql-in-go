package schemas

import (
	"github.com/graphql-go/graphql"
)

// GopherType is the gopher graphQL Object that we will send on queries
// Here we define the structure of the gopher
// This has to match the STRUCT tags that are sent out later
var GopherType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Gopher",
	// Fields is the field values to declare the structure of the object
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.ID,
			Description: "The ID that is used to identify unique gophers",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "The name of the gopher",
		},
		"hired": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "True if the Gopher is employeed",
		},
		"profession": &graphql.Field{
			Type:        graphql.String,
			Description: "The gophers last/current profession",
		},
	},
})
